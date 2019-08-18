import tushare as ts
import seaborn as sns
import keras
import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
from keras.models import Sequential
from keras import layers
from keras.optimizers import RMSprop
from keras.initializers import Orthogonal
import time


def plot_loss(history, title):
    loss = history.history['loss']
    val_loss = history.history['val_loss']

    epochs = range(len(loss))

    plt.figure()

    plt.plot(epochs, loss, 'r', label='Training loss')
    plt.plot(epochs, val_loss, 'b', label='Validation loss')
    plt.title(title)
    plt.legend()

    plt.show()


def generator(data, lookback, delay, min_index, max_index, shuffle=False, batch_size=128, step=6):
    if max_index is None:
        max_index = len(data) - delay - 1
    i = min_index + lookback
    while 1:
        if shuffle:
            rows = np.random.randint(
                min_index + lookback, max_index, size=batch_size)
        else:
            if i + batch_size >= max_index:
                i = min_index + lookback
            rows = np.arange(i, min(i + batch_size, max_index))
            i += len(rows)

        samples = np.zeros((len(rows),
                            lookback // step,
                            data.shape[-1]))
        targets = np.zeros((len(rows),))
        for j, row in enumerate(rows):
            indices = range(rows[j] - lookback, rows[j], step)
            samples[j] = data[indices]
            targets[j] = data[rows[j] + delay][3]
        yield samples, targets


def get_model():
    model = Sequential()
    model.add(
        layers.Bidirectional(layers.LSTM(32,
                                         dropout=0.2,
                                         recurrent_dropout=0.5,
                                         kernel_initializer=Orthogonal(seed=7)),
                             input_shape=(None, 9)
                             )
    )
    model.add(layers.Dense(16))
    model.add(layers.Dense(1))
    return model


start_time = time.time()
print('-----------starting-----------')

lookback = 100
step = 1
delay = 1
batch_size = 8
code_list = [5, 6, 7, 8, 9, 10]
for code in code_list:
    data = pd.read_csv('data/{}.csv'.format(code), index_col=0)
    data.index = pd.to_datetime(data.index)
    # 划分训练集
    train_mean = data[:'2019-01-17'].mean()
    train_std = data[:'2019-01-17'].std()

    data -= train_mean
    data /= train_std

    train_size = len(data[:'2019-01-17'])
    test_size = len(data) - train_size
    data_values = data.values
    train_gen = generator(data_values,
                          lookback=lookback,
                          delay=delay,
                          min_index=0,
                          max_index=train_size,
                          shuffle=True,
                          step=step,
                          batch_size=batch_size)
    test_gen = generator(data_values,
                         lookback=lookback,
                         delay=delay,
                         min_index=train_size,
                         max_index=None,
                         step=step,
                         batch_size=batch_size)

    test_steps = (len(data_values) - train_size - lookback) // batch_size
    model = get_model()
    model.compile(optimizer=RMSprop(), loss='mse')
    history = model.fit_generator(train_gen,
                                  steps_per_epoch=30,
                                  epochs=8,
                                  validation_data=test_gen,
                                  validation_steps=test_steps,
                                  verbose=2)
    plot_loss(history, '{}.sz train and validation loss'.format(code))

    test_set = data_values[train_size:]
    real_val = test_set[lookback:, 3]
    pred_val = []
    for i in range(test_size - lookback):
        window = test_set[i: i + lookback]
        window = window.reshape(1, len(window), -1)
        pred_val.append(model.predict(window)[0, 0])
    pred_val = np.array(pred_val)

    date = data.iloc[train_size + lookback:].index
    real_val = real_val * train_std['close'] + train_mean['close']
    pred_val = pred_val * train_std['close'] + train_mean['close']
    plt.figure()

    plt.plot(date, real_val, 'b', label='real price')
    plt.plot(date, pred_val, 'r', label='lstm predict')
    plt.title('LSTM on {}.sz set'.format(code))
    plt.legend()

    plt.show()

end_time = time.time()
used_time = start_time - end_time
print('----------finished------------')
print('used time:%s s'% used_time)