docker-compose down
docker rmi $(docker images | grep "citicup/" | awk '{print $3}')
docker rmi $(docker images | grep "none" | awk '{print $3}')
docker network create proxy
docker-compose up --build -d && docker-compose logs -f &
