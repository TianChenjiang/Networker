# h is the host name

while getopts h: option
do
case "${option}"
in
h) BASE=${OPTARG};;
esac
done
# 复制到base dir
scp docker/Dockerfile root@"$BASE"/Dockerfile
scp docker/apigateway/Dockerfile root@"$BASE"/Dockerfile-api-gateway
scp docker/model/Dockerfile root@"$BASE"/Dockerfile-model
scp docker/dockerize root@"$BASE"/dockerize
scp bash/setup.sh root@"$BASE"/setup.sh
# api gateway
scp citicup-apigateway/target/*.jar root@"$BASE"/citicup-api-gateway-0.0.1-SNAPSHOT.jar
scp citicup-discovery-server/target/*.jar root@"$BASE"/citicup-discovery-server-0.0.1-SNAPSHOT.jar
scp citicup-plus/target/*.jar root@"$BASE"/citicup-plus-0.0.1-SNAPSHOT.jar
# 前端
scp -r citicup-frontend/* root@"$BASE"/citicup-frontend
# admin 服务器
scp -r citicup-admin/* root@"$BASE"/citicup-admin
# model
scp -r citicup-model/production/* root@"$BASE"/citicup-model/production

# docker
scp docker-compose-prod.yml root@"$BASE"/docker-compose.yml
# 数据
scp -r sql/* root@"$BASE"/sql
