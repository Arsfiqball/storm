# Storm Framework

The Storm Framework is a Go-based service boilerplate designed for building scalable applications of any size, from microservices to monoliths. It provides a solid foundation with best practices for modern Go service development.

[![Go Reference](https://pkg.go.dev/badge/github.com/Arsfiqball/storm.svg)](https://pkg.go.dev/github.com/Arsfiqball/storm)
[![Go Report Card](https://goreportcard.com/badge/github.com/Arsfiqball/storm)](https://goreportcard.com/report/github.com/Arsfiqball/storm)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Quick Start

```bash
# Clone the repository
git clone https://github.com/Arsfiqball/storm.git
cd storm

# Start development environment
sh ./scripts/isolate_up.sh --expose

# To stop the development environment
sh ./scripts/isolate_down.sh
```

## Core Features

- **Dependency Injection**: Uses [Wire](https://github.com/google/wire) for compile-time dependency injection
- **HTTP Server**: Built with [Fiber](https://github.com/gofiber/fiber/v2) for high-performance HTTP handling
- **Database Integration**: [GORM](https://gorm.io/) for PostgreSQL database interactions
- **Asynchronous Messaging**: [Watermill](https://github.com/ThreeDotsLabs/watermill) for event-driven architecture
- **Background Processing**: Redis-based job processing with [gocraft/work](https://github.com/gocraft/work)
- **Observability**: OpenTelemetry integration with Zipkin tracing support
- **Structured Logging**: Using Go's `slog` package with configurable outputs

## Architecture

The project follows a clean architecture pattern with:

- cmd: Application entry points
- internal: Private application code
  - `provider/`: Core service providers (Fiber, GORM, Redis, etc.)
  - `system/`: Application bootstrapping and lifecycle management
- pkg: Public packages that can be imported by other projects
- database: Database migration scripts
- scripts: Utility scripts for development workflows

## Development Workflow

The framework includes several convenience scripts:
- isolate_up.sh: Start the development environment with Docker
- isolate_down.sh: Tear down the development environment
- wire.sh: Generate dependency injection code
- test_component.sh: Run component tests in an isolated environment
- gomarkdoc.sh: Generate Markdown documentation from code comments

## Deployment Options

The framework supports multiple deployment methods:
1. **Docker containers**: Multi-stage Docker builds for minimal production images
2. **Debian packages**: Build scripts for creating Debian packages with systemd service files

## Configuration

Configuration is managed through:
- Environment variables with `APP_` prefix
- YAML configuration file (.config.yml)
- [Viper](https://github.com/spf13/viper) for configuration management

This framework is designed to be flexible enough for any Go service development while providing a solid foundation of best practices and patterns.

## Project Structure

```
storm/
├── bin/                # Binary artifacts
├── build/              # Build configurations
│   ├── docker/         # Docker build files
│   └── linux/          # Linux packaging
├── cmd/                # Application entry points
│   └── server/         # Main server executable
├── database/           # Database migrations
│   └── postgresql/     # PostgreSQL specific migrations
├── docs/               # Documentation
│   └── packages/       # Package-specific documentation
├── internal/           # Private application code
│   ├── provider/       # Service providers
│   └── system/         # Application bootstrapping
├── pkg/                # Public packages
│   ├── example/        # Example package
│   └── kernel/         # Core functionality
└── scripts/            # Utility scripts
```

## Prerequisites

- Go 1.21+
- Docker and Docker Compose
- PostgreSQL 15+
- Redis 7+

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- [Wire](https://github.com/google/wire) - Dependency injection
- [Fiber](https://github.com/gofiber/fiber) - Web framework
- [GORM](https://gorm.io/) - ORM library
- [Watermill](https://github.com/ThreeDotsLabs/watermill) - Event-driven architecture
