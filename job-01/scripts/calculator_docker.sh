#!/bin/bash

echo "Welcome to our Calculator App! "
echo "Your available options are (and how to call them): "
echo "1. Build service: 'build' | 'BUILD' | '1' "
echo "2. Start service: 'start' | 'run' | 'init' | 'START' | 'RUN' | 'INIT' | '2' "
echo "3. Stop service: 'stop' | 'pause' | 'STOP' | 'PAUSE' | '3' "
echo "4. Service's status: 'status' | 'info' | 'STATUS' | 'INFO' | '4' "
echo "5. Display current images: 'images' | 'IMAGES' | '5' "
echo "6. Display current running containers: 'ps' | 'containers' | 'running' | 'PS' | 'CONTAINERS' | 'RUNNING' | '6' "

read -r -p "Insert your action: " action 

build() {
    echo "We will BUILD our Calculator service here! "
    echo "The name of the resulting image will be 'calculator-app'"
    docker build -t calculator-app .
    echo "Your BUILD action was successful! "
}

start() {
    echo "We will START our Calculator service here! "
    echo "The name of the resulting running container will be 'calc'"
    docker run --name calc -d --rm -p 8080:8080 calculator-app
    echo "Your START action was successful! "
}

stop() {
    echo "We will STOP our Calculator service here! "
    docker stop calc
    echo "Your STOP action was successful! "
}

status() {
    echo "We will show our Calculator service's STATUS here! "
    CID=$(docker ps -q --filter name=calc --filter status=running)
    if [[ ! "${CID}" ]]; then
        echo "Calculator service is NOT RUNNING"
    else
        echo "Calculator service is RUNNING" 
    fi
    echo "Your STATUS action was successful! "
}

images() {
    echo "We will show our CURRENT IMAGES here! "
    docker image ls
    echo "Your CURRENT IMAGES action was successful! "
}

running_containers() {
    echo "We will show our CURRENT RUNNING CONTAINERS here! "
    docker ps -a
    echo "Your RUNNING CONTAINERS action was successful! "
}

case $action in
build|BUILD|1)
build
;;

start|run|init|START|RUN|INIT|2)
start
;;

stop|pause|STOP|PAUSE|3)
stop
;;

status|info|STATUS|INFO|4)
status
;;

images|IMAGES|5)
images
;;

ps|containers|running|PS|CONTAINERS|RUNNING|6)
running_containers
;;

*)
echo "Your chosen action was not found! Take a better look at our available options and try again next time..."
;;

esac