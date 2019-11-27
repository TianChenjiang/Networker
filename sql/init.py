import tushare as ts

if __name__ == "__main__":
    pro = ts.pro_api('53cd3b985c649c978160c6ec04bce24f4fbd2ebcb4673e8f2fba9a43')
    df = pro.query('daily', ts_code='000001.SZ', start_date='20180701', end_date='20180718')
    res = []
    for index, row in df.iterrows():
        tmp = {
            'date': row['trade_date'],
            'open': row['open'],
            'close': row['close'],
            'highest': row['high'],
            'lowest': row['low'],
            'vol': row['vol']
        }
        res.append(tmp)
    return res
