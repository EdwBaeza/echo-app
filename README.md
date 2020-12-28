# echo_app
* Docker Instructions

1.- `docker volume create mongodata`

2.- `docker network create --attachable echonet`

3.- `docker run -d --name db --mount src=mongodata,dst=/data/db mongo`

4.- `docker network connect echonet db`

5.- `docker run -d --name app -p 3000:8080 echoapp:2.0`

6.- `docker network connect echonet app`

