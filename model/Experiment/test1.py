import pandas as pd
import numpy as np
import random
import tushare as ts
import seaborn as sns
import matplotlib.pyplot as plt


def get_stock_price_by_code(code, lookback, delay, pledge_date, close_position_date, code_price_dict):
    """
    通过股票代码，过去天数，生成三维的股票价格矩阵
    delay * lookback * 9(features)
    """
    non_close_rate = 4
    price_df = code_price_dict[code]
    price_df = price_df[pledge_date: close_position_date]
    price_values = price_df.values
    # 平仓训练集
    usable_close_len = len(price_df) - lookback
    usable_close_len = min(usable_close_len, delay + 1)
    close_list = [price_values[-lookback - i: -i] for i in range(1, usable_close_len + 1)]
    # 正常训练集
    usable_non_close_len = len(price_df) - delay - lookback
    usable_non_close_len = min(usable_non_close_len, non_close_rate * delay)
    non_close_list = [price_values[-lookback - i: -i] for i in range(delay+1, delay+1+usable_non_close_len)]
    return close_list, non_close_list


def generator(pledge, lookback, delay, min_index, max_index, code_price_dict):
    # batch_size == delay(30)
    i = min_index
    while 1:
        # loop back to min_index
        if i >= max_index:
            i = min_index
        single_pledge = pledge.loc[i]
        close_list, non_close_list = get_stock_price_by_code(code=single_pledge['code'],
                               lookback=lookback,
                               delay=delay,
                               pledge_date=single_pledge['pledge_date'],
                               close_position_date=single_pledge['close_position_date'],
                               code_price_dict=code_price_dict)
        single_pledge = single_pledge.drop(['code', 'pledge_date', 'close_position_date'])
        close_zip = [(prices, single_pledge.values, 1) for prices in close_list]
        non_close_zip = [(prices, single_pledge.values, 0) for prices in non_close_list]
        all_zip = close_zip + non_close_zip
        random.shuffle(all_zip)
        price_time = np.stack([t[0] for t in all_zip])
        pledge_info = np.stack([t[1] for t in all_zip])
        targets = np.stack([t[2] for t in all_zip])
        i += 1
        yield [price_time, pledge_info], targets


ts.set_token('53cd3b985c649c978160c6ec04bce24f4fbd2ebcb4673e8f2fba9a43')

pro = ts.pro_api()

code_df = pro.stock_basic(exchange='', list_status='L', fields='ts_code')


pledge = pd.read_csv('data/pledge_clean.csv', index_col=0)
pledge['close_position_date'] = pd.to_datetime(pledge['close_position_date'])
pledge['pledge_date'] = pd.to_datetime(pledge['pledge_date'])
pledge = pledge.sort_values(by=['close_position_date'])
pledge = pledge.reset_index(drop=True)

code_price_dict = dict()
for i in range(len(code_df)):
    code = code_df.iloc[i, 0]
    price_df = pd.read_csv('data/price/{}.csv'.format(code), index_col=0)
    price_df.index = pd.to_datetime(price_df.index)
    code_price_dict[code] = price_df

length_list = []
for i in range(len(pledge)):
    single_pledge = pledge.loc[i]
    code = single_pledge['code']
    price_df = code_price_dict[code]
    price_df = price_df[single_pledge['pledge_date']: single_pledge['close_position_date']]
    length_list.append(len(price_df))
print(min(length_list))
sns.distplot(length_list)
plt.show()
