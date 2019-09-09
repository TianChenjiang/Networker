from datetime import datetime
from model.files import info_scaler, price_scaler, code_set, pro, com, fam, deep_model, get_info


lookback = 48


def transform_price(price_df, forecast_close_line):
    price_df = price_df.drop(['ts_code', 'trade_date', 'pre_close'], axis=1)
    price_df['delta'] = price_df.apply(lambda x: x['close'] - forecast_close_line, axis=1)
    price_values = price_scaler.transform(price_df)
    price_values = price_values[:lookback]
    price_values = price_values[::-1]
    return price_values.reshape((1, lookback, -1))


def predict(code, forecast_close_line):

    if code not in code_set:
        return None

    # 获得股价信息
    # price_df = pd.read_csv('data/price/{}.csv'.format(code))
    price_df = pro.daily(ts_code=code, start_date='20190101', end_date=datetime.now().strftime('%Y%m%d'))
    price_values = transform_price(price_df, forecast_close_line)
    pledge_price = price_df.loc[0]['close']

    info_values = get_info(code, forecast_close_line, pledge_price)

    global graph
    with graph.as_default():
        prob = deep_model.predict([price_values, info_values]).ravel()
    return prob[0]


if __name__ == '__main__':
    prob = predict('000001.SZ', 10)
    print(prob)
