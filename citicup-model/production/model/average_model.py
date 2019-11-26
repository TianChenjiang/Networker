from model.files import pro, code_set, JSPATH
import datetime
from model.deeplearning_model import deep_predict
from model.lightgbm_model import lgb_predict
import json
import pandas as pd
from werkzeug.exceptions import NotFound
import bisect


def get_result(code, forecast_close_line, price_df):
    deep_prob = deep_predict(code, forecast_close_line, price_df)
    lgb_prob = lgb_predict(code, forecast_close_line, price_df)

    return 0.5 * deep_prob + 0.5 * lgb_prob


def predict(code):
    if not JSPATH.is_file():
        predict_all_init()
        raise NotFound('Risk probability file not exists! Please wait 5min and try again!')

    with open(JSPATH, 'r') as fp:
        prob_dict = json.load(fp)
        if code not in prob_dict:
            raise NotFound('Company ts_code not in our system!')

        prob_list = sorted(prob_dict.items(), key=lambda x: x[1], reverse=True)
        rank, prob = [(i, prob) for i, (curr_code, prob) in enumerate(prob_list) if curr_code == code][0]
        total = len(prob_list)

    return prob, rank, total


def predict_given_forecast_close_line(code, forecast_close_line):
    if code not in code_set:
        raise NotFound('Company not in our dataset')

    price_df = pro.daily(ts_code=code, start_date='20190101', end_date=datetime.datetime.now().strftime('%Y%m%d'))
    return get_result(code, forecast_close_line, price_df)


# Schedule task
def predict_all_init():
    today = datetime.datetime.now()
    appended_data = []
    for i in range(100):  # 不在工作日时可能没有股价
        day = today - datetime.timedelta(days=i)
        df = pro.daily(trade_date=day.strftime('%Y%m%d'))
        appended_data.append(df)
    appended_data = pd.concat(appended_data)

    prob_dict = dict()
    for code in code_set:
        price_df = appended_data[appended_data['ts_code'] == code]
        if len(price_df) < 48:
            continue
        price_df = price_df.sort_values(by='trade_date', ascending=False)
        price_df = price_df.reset_index(drop=True)
        forecast_close_line = price_df.iloc[0]['close'] * 0.7

        prob = get_result(code, forecast_close_line, price_df)
        prob_dict[code] = prob

    with open(JSPATH, 'w') as fp:
        json.dump(prob_dict, fp)


if __name__ == '__main__':
    predict_all_init()

