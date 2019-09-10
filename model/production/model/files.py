import tushare as ts
import pandas as pd
import pickle
import pathlib
import numpy as np

PATH = pathlib.Path(__file__).parent
DATA_PATH = PATH.joinpath("data").resolve()
INFO_SCALER_PATH = PATH.joinpath('info_scaler.pkl')
PRICE_SCALER_PATH = PATH.joinpath('price_scaler.pkl')
CODE_SET_PATH = DATA_PATH.joinpath('code_set.pkl')
FAM_PATH = DATA_PATH.joinpath('family_firm_clean.csv')
COM_PATH = DATA_PATH.joinpath('company_fund_latest.csv')

ts.set_token('53cd3b985c649c978160c6ec04bce24f4fbd2ebcb4673e8f2fba9a43')

pro = ts.pro_api()


# 加载info_scaler
with open(INFO_SCALER_PATH, 'rb') as f:
    info_scaler = pickle.load(f)

with open(PRICE_SCALER_PATH, 'rb') as f:
    price_scaler = pickle.load(f)

# 加载所支持的公司代码
with open(CODE_SET_PATH, 'rb') as f:
    code_set = pickle.load(f)

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


def get_info(code, forecast_close_line, pledge_price):

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
    return data_values