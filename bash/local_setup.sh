# 用于本地一次性构建,以及文件复制
cd citicup-apigateway && mvn clean package -Pprod && cd ..
bash bash/scp.sh -h 106.14.140.93:/root/citicup

