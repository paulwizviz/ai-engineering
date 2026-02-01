#/bin/sh

# Docker compose network
export N8N_DOCKER_NETWORK=ai-engineering_default

# The base URL for your n8n instance
# Change this to your domain if you're hosting it online
export N8N_HOST=localhost
export N8N_PORT=5678
export N8N_PROTOCOL=http

# To persist your workflow data
export DATA_FOLDER=/home/node/.n8n

# To enable community nodes as tools in AI Agents (important!)
export N8N_COMMUNITY_PACKAGES_ALLOW_TOOL_USAGE=true

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
        docker compose -f $PWD/deployments/docker-compose.yaml up -d
        ;;
    "stop")
        docker compose -f $PWD/deployments/docker-compose.yaml down
        ;;
    *)
        echo "Usage: $0 [clean|start|stop]"
        ;;
esac