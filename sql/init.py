import tushare as ts
from sqlalchemy import create_engine

pro = ts.pro_api('53cd3b985c649c978160c6ec04bce24f4fbd2ebcb4673e8f2fba9a43')
df = pro.stock_basic(
    exchange='',
    fields='ts_code,symbol,name,area,industry,fullname,enname,market,exchange,curr_type,list_status,list_date,delist_date')

engine = create_engine('mysql+pymysql://root:mysql@127.0.0.1:33006/citicup?charset=utf8')
df.to_sql(name='stock', con=engine, if_exists='append', index=True, index_label='id')
