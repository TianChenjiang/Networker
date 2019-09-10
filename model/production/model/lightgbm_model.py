import lightgbm as lgb
import numpy as np
from model.files import price_scaler, pro, get_info
from datetime import datetime
import pathlib

PATH = pathlib.Path(__file__).parent
LGB_MODEL_PATH = PATH.joinpath('lgb_model_new.h5')

# 加载lightgbm模型
lgb_model = lgb.Booster(model_file=str(LGB_MODEL_PATH))

lookback = 48


def transform_price(price_df, forecast_close_line):
    price_df = price_df.drop(['ts_code', 'trade_date', 'pre_close'], axis=1)
    price_df['delta'] = price_df.apply(lambda x: x['close'] - forecast_close_line, axis=1)
    price_values = price_scaler.transform(price_df)
    price_values = price_values[:lookback, 3]
    price_values = price_values[::-1]
    return price_values


def lgb_predict(code, forecast_close_line):

    # 获得股价信息
    # price_df = pd.read_csv('data/price/{}.csv'.format(code))
    price_df = pro.daily(ts_code=code, start_date='20190101', end_date=datetime.now().strftime('%Y%m%d'))
    price_values = transform_price(price_df, forecast_close_line)
    pledge_price = price_df.loc[0]['close']

    info_values = get_info(code, forecast_close_line, pledge_price)

    data = np.append(price_values, info_values)

    return lgb_model.predict(data.reshape((1, -1)))


if __name__ == '__main__':
    print(lgb_predict('000001.SZ', 10))