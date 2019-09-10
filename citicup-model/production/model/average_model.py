from model.files import pro, code_set
from datetime import datetime
from model.deeplearning_model import deep_predict
from model.lightgbm_model import lgb_predict


def predict(code, forecast_close_line):

    if code not in code_set:
        return None

    price_df = pro.daily(ts_code=code, start_date='20190101', end_date=datetime.now().strftime('%Y%m%d'))
    deep_prob = deep_predict(code, forecast_close_line, price_df)
    lgb_prob = lgb_predict(code, forecast_close_line, price_df)

    return 0.8 * deep_prob + 0.2 * lgb_prob


if __name__ == '__main__':
    print(predict('000001.SZ', 12))