from model.files import price_scaler, get_info, info_scaler
import pathlib
import tensorflow as tf
from keras.backend import set_session
from keras.models import load_model

lookback = 48
PATH = pathlib.Path(__file__).parent
DEEP_MODEL_PATH = PATH.joinpath('pledge_company_model.h5')


# 加载深度学习模型

sess = tf.Session()
graph = tf.get_default_graph()

set_session(sess)
deep_model = load_model(DEEP_MODEL_PATH)

def transform_price(price_df, forecast_close_line):
    price_df = price_df.drop(['ts_code', 'trade_date', 'pre_close'], axis=1)
    price_df['delta'] = price_df.apply(lambda x: x['close'] - forecast_close_line, axis=1)
    price_values = price_scaler.transform(price_df)
    price_values = price_values[:lookback]
    price_values = price_values[::-1]
    return price_values.reshape((1, lookback, -1))


def deep_predict(code, forecast_close_line, price_df):

    price_values = transform_price(price_df, forecast_close_line)
    pledge_price = price_df.loc[0]['close']

    info_values = get_info(code, forecast_close_line, pledge_price, info_scaler)

    global graph
    global sess
    with graph.as_default():
        set_session(sess)
        prob = deep_model.predict([price_values, info_values]).ravel()
    return prob[0]

