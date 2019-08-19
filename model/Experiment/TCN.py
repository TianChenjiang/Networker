import tushare as ts
import seaborn as sns
import keras
import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
from keras.models import Sequential, Input, Model
from tcn import TCN
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


def get_model(batch_size, timesteps, input_dim):
    i = Input(batch_shape=(batch_size, timesteps, input_dim))
    x = TCN(nb_filters=64,
            kernel_size=2,
            dilations=[1, 2, 4, 8, 16],
            nb_stacks=1,
            dropout_rate=0.3,
            kernel_initializer=Orthogonal(seed=7),
            use_skip_connections=True,
            return_sequences=False)(i)
    x = layers.Dense(16, activation='relu')(x)
    x = layers.Dropout(0.2)(x)
    y = layers.Dense(1)(x)
    model = Model(inputs=[i], outputs=[y])
    return model


start_time = time.time()
print('-----------starting-----------')

lookback = 32
step = 1
delay = 1
batch_size = 16
code_list = [5, 6, 7, 8, 9, 10]
for code in code_list:
    data = pd.read_csv('data/{}.csv'.format(code), index_col=0)
    data.index = pd.to_datetime(data.index)
    # 划分训练集
    train_mean = data[:'2019-03-17'].mean()
    train_std = data[:'2019-03-17'].std()

    data -= train_mean
    data /= train_std

    train_size = len(data[:'2019-03-17'])
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
    model = get_model(batch_size, lookback, data_values.shape[-1])
    model.compile(optimizer=RMSprop(), loss='mse')
    history = model.fit_generator(train_gen,
                                  steps_per_epoch=50,
                                  epochs=15,
                                  validation_data=test_gen,
                                  validation_steps=test_steps)
    plot_loss(history, '{}.sz train and validation loss on TCN'.format(code))

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
    plt.plot(date, pred_val, 'r', label='TCN predict')
    plt.title('TCN on {}.sz set'.format(code))
    plt.legend()

    plt.show()

end_time = time.time()
used_time = start_time - end_time
print('----------finished------------')
print('used time:%s s' % used_time)
