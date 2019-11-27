import pandas as pd
import tushare as ts
from sqlalchemy import create_engine



if __name__ == "__main__":
    pro = ts.pro_api('53cd3b985c649c978160c6ec04bce24f4fbd2ebcb4673e8f2fba9a43')

    engine = create_engine('mysql+pymysql://root:mysql@106.14.140.93:3306/citicup?charset=utf8')
    d = pd.read_sql('select ts_code from market', con=engine)
    for index, row in d.iterrows():
        tscode = row['ts_code']
        # ple
        df_ple = pro.pledge_detail(ts_code=tscode)
        # finance
        # 财务指标
        d1 = pro.fina_indicator(ts_code=tscode)
        # d2 = pro.income_vip(ts_code=tscode)
        df_ple.to_sql(name='pledge', con=engine, if_exists='append', index=True, index_label='id')
        d1.to_sql(name='finance', con=engine, if_exists='append', index=True, index_label='id')
        print(index)