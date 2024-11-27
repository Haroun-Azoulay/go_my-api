#!/bin/bash

if [ "$1" == "--stop" ]; then
    docker compose down 

elif [ "$1" == "--remove" ]; then
    docker compose down -v

elif [ "$1" == "--build" ]; then
    docker compose up --build

else 
    echo "Usage: ./script.sh [OPTION]"
    echo ""
    echo "Options:"
    echo "  --stop       Stop all running containers in the current project."
    echo "  --remove     Stop and remove containers, networks, and volumes."
    echo "  --build        Build and start containers (default if no option is provided)."
fi
