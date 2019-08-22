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
scp docker/dockerize root@"$BASE"/dockerize
scp bash/setup.sh root@"$BASE"/setup.sh
# api gateway
scp citicup-apigateway/target/citicup-api-gateway-0.0.1-SNAPSHOT.jar root@"$BASE"/citicup-api-gateway-0.0.1-SNAPSHOT.jar
# 前端
scp -r citicup-frontend/* root@"$BASE"/citicup-frontend
# admin 服务器
scp -r citicup-admin/app/* root@"$BASE"/citicup-admin

# docker
scp docker-compose-prod.yml root@"$BASE"/docker-compose.yml
