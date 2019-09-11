from sqlalchemy import create_engine

import tushare as ts

df = ts.get_gdp_for()
print(df)

engine = create_engine('mysql+pymysql://root:mysql@127.0.0.1:33006/citicup?charset=utf8')
df.to_sql(name='name', con=engine, if_exists='append', index=False, index_label=False)

