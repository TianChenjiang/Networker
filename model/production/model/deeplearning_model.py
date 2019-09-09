import pandas as pd
import numpy as np
import tushare as ts
import tensorflow as tf
from keras.models import load_model
import pickle
import pathlib
from datetime import datetime


PATH = pathlib.Path(__file__).parent
DATA_PATH = PATH.joinpath("data").resolve()
MODEL_PATH = PATH.joinpath('pledge_company_model.h5')
INFO_SCALER_PATH = PATH.joinpath('info_scaler.pkl')
PRICE_SCALER_PATH = PATH.joinpath('price_scaler.pkl')
CODE_SET_PATH = DATA_PATH.joinpath('code_set.pkl')
FAM_PATH = DATA_PATH.joinpath('family_firm_clean.csv')
COM_PATH = DATA_PATH.joinpath('company_fund_latest.csv')


ts.set_token('53cd3b985c649c978160c6ec04bce24f4fbd2ebcb4673e8f2fba9a43')

pro = ts.pro_api()

lookback = 48


# 加载训练好的模型
model = load_model(MODEL_PATH)
model._make_predict_function()
graph = tf.get_default_graph()

# 加载info_scaler
with open(INFO_SCALER_PATH, 'rb') as f:
    info_scaler = pickle.load(f)

with open(PRICE_SCALER_PATH, 'rb') as f:
    price_scaler = pickle.load(f)

# 加载所支持的公司代码
with open(CODE_SET_PATH, 'rb') as f:
    code_set = pickle.load(f)

# 加载股价字典
# price_dict = dict()
# for code in code_set:
#     price_df = pd.read_csv('data/price/{}.csv'.format(code))
#     price_df = price_df.drop(['trade_date', 'pre_close'], axis=1)
#     price_dict[code] = price_df


# 加载公司基本面字典
# company_dict = dict()
# for code in code_set:
#     company_df = pd.read_csv('data/fund/{}.csv'.format(code), index_col=0)
#     company_df = company_df.drop(['ann_date', 'end_date'], axis=1)
#     company_dict[code] = company_df

# 加载家族企业
use = ['Symbol', 'ContrshrProportion', 'IsRelatedTrading', 'ShareholderFirstProp',
       'ControlProportion', 'FamEntyp_1.0', 'FamEntyp_2.0', 'FamEntyp_3.0',
       'BoardCode_P3401', 'BoardCode_P3402', 'BoardCode_P3403', 'FamStyle_1',
       'FamStyle_2', 'ManGeneration_1.0', 'ManGeneration_2.0',
       'FamNameStatus_1', 'FamNameStatus_2', 'FamNameStatus_3',
       'FamNameStatus_4', 'FamNameStatus_5']
fam = pd.read_csv(FAM_PATH, index_col=0, usecols=use)

# 加载公司基本面
com = pd.read_csv(COM_PATH)
com = com.drop(['ann_date', 'end_date'], axis=1)


def transform_price(price_df, forecast_close_line):
    price_df = price_df.drop(['ts_code', 'trade_date', 'pre_close'], axis=1)
    price_df['delta'] = price_df.apply(lambda x: x['close'] - forecast_close_line, axis=1)
    price_values = price_scaler.transform(price_df)
    price_values = price_values[:lookback]
    price_values = price_values[::-1]
    return price_values.reshape((1, lookback, -1))


def predict(code, forecast_close_line):

    # 获得股价信息
    # price_df = pd.read_csv('data/price/{}.csv'.format(code))
    price_df = pro.daily(ts_code=code, start_date='20190101', end_date=datetime.now().strftime('%Y%m%d'))
    price_values = transform_price(price_df, forecast_close_line)

    pledge_price = price_df.loc[0]['close']
    data = pd.Series(np.array([code, pledge_price, forecast_close_line]),
                     index=['ts_code', 'pledge_price', 'forecast_close_line'])

    data = data.to_frame().T

    data = data.merge(com, how='left', left_on='ts_code', right_on='ts_code')

    data = data.merge(fam, how='left', left_on='ts_code', right_on='Symbol')
    data = data.drop(['ts_code'], axis=1)
    data = data.fillna(0)

    # 归一化数据面板
    data_values = info_scaler.transform(data)
    data_values = data_values.reshape((1, -1))

    global graph
    with graph.as_default():
        prob = model.predict([price_values, data_values]).ravel()
    return prob[0]


if __name__ == '__main__':
    prob = predict('000001.SZ', 10)
    print(prob)
