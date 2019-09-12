import tushare as ts
from sqlalchemy import create_engine

pro = ts.pro_api('53cd3b985c649c978160c6ec04bce24f4fbd2ebcb4673e8f2fba9a43')
df = pro.stock_company(fields='ts_code,exchange,chairman,manager,'
                              'secretary,reg_capital,setup_date,province,'
                              'city,introduction,website,email,office,employees,main_business,business_scope')
engine = create_engine('mysql+pymysql://root:mysql@127.0.0.1:33006/citicup?charset=utf8')
df.to_sql(name='company', con=engine, if_exists='append', index=True, index_label='id')
