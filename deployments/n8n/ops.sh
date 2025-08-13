#/bin/sh

if [ "$(basename $(realpath .))" != "ai-engineering" ]; then
    echo "You are outside the scope of the project"
    exit 0
fi

COMMAND=$1

case $COMMAND in
    "clean")
        rm  -rf ./n8n-data
        ;;
    "start")
        docker compose -f $PWD/deployments/n8n/docker-compose.yaml up -d
        ;;
    "stop")
        docker compose -f $PWD/deployments/n8n/docker-compose.yaml down
        ;;
    *)
        echo "Usage: $0 [clean|start|stop]"
        ;;
esac