#!/bin/bash
if [ $# -eq 0 ]
  then
    echo "No arguments supplied. e.g: ./build.sh name host_port docker_port"
    exit 1
fi 
echo "Start - Building $1 image..."
echo "remove existing image..."
docker rm -f $1
echo "build image..."
docker build --tag $1 .
echo "Done - $1 image built successfully."
echo "Restart container with port $3 on $2 host port"
docker run --restart=unless-stopped -d -p $2:$3 --name $1 $1
docker system prune -f
echo "Finish - Running $1 container..."
exit
