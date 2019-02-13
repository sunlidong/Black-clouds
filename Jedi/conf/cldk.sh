sudo docker rm -f $(sudo docker ps -aq)
sudo docker network prune
sudo docker volume prune
echo "------------clear is ok"
docker-compose up -d
echo "----------------------docker up is ok-----------------"