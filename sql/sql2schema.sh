# 存储到docker数据库
docker exec -i citicup-mysql mysql -uroot -pmysql citicup < sql/framework.sql
