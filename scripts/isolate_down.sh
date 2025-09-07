# Run this command from root project directory
# sh scripts/isolate_down.sh

docker-compose -f docker-compose.isolated.yml -f docker-compose.isolated.expose.yml down --volumes
