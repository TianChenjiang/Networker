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
scp bash/setup.sh root@"$BASE"/setup.sh
scp -r citicup-frontend/* root@"$BASE"/citicup-frontend
scp docker-compose-prod.yml root@"$BASE"/docker-compose.yml