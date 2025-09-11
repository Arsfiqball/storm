# Run this command from root project directory
# sh scripts/isolate_up.sh

# Check if --expose flag is passed
if [ -n "$(echo "$*" | grep -- "--expose")" ]; then
    echo "Starting with exposed ports..."
    docker-compose -f docker-compose.isolated.yml -f docker-compose.isolated.expose.yml up --build -d
else
    echo "Starting in normal mode..."
    docker-compose -f docker-compose.isolated.yml up --build -d
fi
