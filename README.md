# echo_app
* Docker Instructions

docker volume create mongodata
docker network create --attachable echonet
docker run -d --name db --mount src=mongodata,dst=/data/db mongo

docker network connect echonet db
docker run -d --name app -p 3000:8080 echoapp:2.0
docker network connect echonet app
