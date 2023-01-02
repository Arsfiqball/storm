# Run this command from root project directory
# sh scripts/test_component.sh

# Start all required service with isolated environment
echo "=============== SETUP START ================="
docker-compose -f docker-compose.isolated.yml up --build -d
echo "================ SETUP END =================="
echo "Let me take a break for 10 seconds..."
sleep 10

# Run all component test
echo "================ TEST START ================="
docker-compose -f docker-compose.isolated.yml run --rm app go test ./test/component/... -v -p 1
echo "================= TEST END =================="
echo "Let me take a break for 2 seconds..."
sleep 2

# Shutdown and remove everything after it ends
echo "============== SHUTDOWN START ==============="
docker-compose -f docker-compose.isolated.yml down --volumes
echo "=============== SHUTDOWN END ================"
