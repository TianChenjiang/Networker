import pandas as pd
import numpy as np
from sklearn.preprocessing import MinMaxScaler
from keras.models import load_model
import keras.metrics
from model.f1 import f1
import pickle

lookback = 48
keras.metrics.custom_metrics = f1


# 加载训练好的模型
model = load_model('pledge_company_model.h5')

# 加载info_scaler
with open('info_scaler.pkl', 'rb') as f:
    info_scaler = pickle.load(f)

with open('price_scaler.pkl', 'rb') as f:
    price_scaler = pickle.load(f)

# 加载所支持的公司代码
with open('data/code_set.pkl', 'rb') as f:
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
fam = pd.read_csv('data/family_firm_clean.csv', index_col=0, usecols=use)


def transform_price(price_df, forecast_close_line):
    price_df['delta'] = price_df.apply(lambda x: x['close'] - forecast_close_line, axis=1)
    price_values = price_scaler.transform(price_df)
    price_values = price_values[:lookback]
    return price_values.reshape((1, lookback, -1))


def predict(code, forecast_close_line):

    price_df = pd.read_csv('data/price/{}.csv'.format(code))
    price_df = price_df.drop(['trade_date', 'pre_close'], axis=1)
    price_values = transform_price(price_df, forecast_close_line)

    company_df = pd.read_csv('data/fund/{}.csv'.format(code), index_col=0)
    company_df = company_df.drop(['ann_date', 'end_date'], axis=1)
    company_df = company_df.iloc[0]

    pledge_price = price_df.loc[0]['close']
    data = pd.Series(np.array([code, pledge_price, forecast_close_line]),
                     index=['ts_code', 'pledge_price', 'forecast_close_line'])

    data = pd.concat([data, company_df])
    data = data.to_frame().T

    data = data.merge(fam, how='left', left_on='ts_code', right_on='Symbol')
    data = data.drop(['ts_code'], axis=1)
    data = data.fillna(0)
    #
    data_values = info_scaler.transform(data)
    data_values = data_values.reshape((1, -1))
    #
    return model.predict([price_values, data_values]).ravel()


if __name__ == '__main__':
    prob = predict('600998.SH', 12)
    print(prob)
