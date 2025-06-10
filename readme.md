# Storm Framework

Storm Framework is an AI-friendly fullstack framework based on Go, designed for building modern web applications. It provides a comprehensive suite of tools for development, deployment, and management with built-in AI capabilities.

## Features

- **Fullstack Architecture**: Backend in Go with a Vue.js frontend
- **Docker-based Isolated Development**: Easy testing and development setup
- **Database Migrations**: Built-in PostgreSQL migration management
- **Event-driven Architecture**: Redis-based messaging with Watermill
- **Component Testing**: Integrated testing framework
- **Modern UI**: Responsive design with beautiful storm-themed effects

## Getting Started

Initialize a configuration file:

```sh
storm init-config
```

Start the development environment with isolated containers:

```sh
storm isolate up --expose --web
```

Expose ports to the host machine with `--expose` flag and include the web frontend with `--web`.

## Common Commands

### Serve the application

```sh
storm serve
```

Run specific components:

```sh
storm serve http     # Run only HTTP server
storm serve listener # Run only event listener
storm serve worker   # Run only worker component
```

### Database Migrations

```sh
storm migrate up     # Apply migrations
storm migrate down   # Roll back migrations
storm migrate create "create_users_table" # Create new migration files
```

### Component Testing

```sh
storm isolate test
```

## Project Structure

- storm: CLI application entry points
- internal: Internal application code
- web: Vue.js frontend application
- database: Database migrations and schemas
- build: Build configurations and scripts

## Development

The Storm framework includes built-in development tools to streamline the development process. Check system requirements with:

```sh
storm isolate check
```

## License

Built with ⚡ for AI-friendly development
