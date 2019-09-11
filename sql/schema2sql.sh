# 将docker内部sql存储到外部sql文件
docker exec -i citicup-mysql mysqldump -uroot -pmysql citicup > sql/framework.sql
