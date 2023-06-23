# Run this command from root project directory
# sh scripts/wire.sh

# Wire all packages
wire gen app/pkg/...

# Wire all dependencies for internal system
wire gen app/internal/system
