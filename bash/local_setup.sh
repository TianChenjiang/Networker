# 用于本地一次性构建,以及文件复制
mvn clean package -Pprod
bash bash/scp.sh -h 106.14.140.93:/root/citicup

