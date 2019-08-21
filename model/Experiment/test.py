import numpy as np
import pandas as pd
import tushare as ts

ts.set_token('53cd3b985c649c978160c6ec04bce24f4fbd2ebcb4673e8f2fba9a43')

pro = ts.pro_api()
start_date = '20170817'
end_date = '20190817'
code_list = [5, 6, 7, 8, 9, 10]
for code in code_list:
    data = pro.daily(ts_code=str(code).zfill(6) + '.sz', start_date=start_date, end_date=end_date)
    data = data.drop(['ts_code'], axis=1)
    data['trade_date'] = pd.to_datetime(data['trade_date'])
    data = data.set_index('trade_date')
    data = data[::-1]
    data.to_csv('./data/{}.csv'.format(code))
