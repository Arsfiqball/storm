#!/bin/bash
set -e

# Define variables
BUILD_DIR=$(dirname "$(readlink -f "$0")")
PROJECT_ROOT="$(cd "$BUILD_DIR/../../.." && pwd)"
PACKAGE_NAME="storm-app"
VERSION="1.0.0"
ARCHITECTURE="amd64"
MAINTAINER="Your Name <your.email@example.com>"
DESCRIPTION="Storm App Service"

# Create temp directory for package building
TMP_DIR=$(mktemp -d)
trap 'rm -rf $TMP_DIR' EXIT

echo "Building Debian package for $PACKAGE_NAME v$VERSION"

# Setup directory structure
INSTALL_DIR="$TMP_DIR/usr/local/bin"
SERVICE_DIR="$TMP_DIR/etc/systemd/system"
mkdir -p "$INSTALL_DIR" "$SERVICE_DIR" "$TMP_DIR/DEBIAN" "$PROJECT_ROOT/bin"

# Build the Go binary with CGO disabled for better compatibility
cd "$PROJECT_ROOT"
echo "Building Storm binary..."
CGO_ENABLED=0 go build -o "$INSTALL_DIR/storm" ./cmd/storm
chmod +x "$INSTALL_DIR/storm"

# Create systemd service file
cat > "$SERVICE_DIR/$PACKAGE_NAME.service" << EOF
[Unit]
Description=Storm App Service
After=network.target

[Service]
Type=simple
User=nobody
Group=nogroup
WorkingDirectory=/usr/local/bin
ExecStart=/usr/local/bin/storm
Restart=on-failure
RestartSec=5

[Install]
WantedBy=multi-user.target
EOF

# Create control file
cat > "$TMP_DIR/DEBIAN/control" << EOF
Package: $PACKAGE_NAME
Version: $VERSION
Section: utils
Priority: optional
Architecture: $ARCHITECTURE
Maintainer: $MAINTAINER
Description: $DESCRIPTION
EOF

# Create postinst script to enable service
cat > "$TMP_DIR/DEBIAN/postinst" << EOF
#!/bin/sh
systemctl daemon-reload
systemctl enable $PACKAGE_NAME.service
systemctl start $PACKAGE_NAME.service
exit 0
EOF
chmod +x "$TMP_DIR/DEBIAN/postinst"

# Create prerm script to stop and disable service
cat > "$TMP_DIR/DEBIAN/prerm" << EOF
#!/bin/sh
systemctl stop $PACKAGE_NAME.service || true
systemctl disable $PACKAGE_NAME.service || true
exit 0
EOF
chmod +x "$TMP_DIR/DEBIAN/prerm"

# Build the package
echo "Building Debian package..."
dpkg-deb --build "$TMP_DIR" "$PROJECT_ROOT/bin/$PACKAGE_NAME-$VERSION-$ARCHITECTURE.deb"

echo "Debian package created: $PROJECT_ROOT/bin/$PACKAGE_NAME-$VERSION-$ARCHITECTURE.deb"