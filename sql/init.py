from sqlalchemy import create_engine
import pandas as pd
import tushare as ts
pro = ts.pro_api('53cd3b985c649c978160c6ec04bce24f4fbd2ebcb4673e8f2fba9a43')
df = pro.stock_company()
engine = create_engine('mysql+pymysql://root:mysql@127.0.0.1:33006/citicup?charset=utf8')
df.to_sql(name='company', con=engine, if_exists='append',index=True, index_label='id')


