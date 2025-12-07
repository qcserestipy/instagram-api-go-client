# Contributing to instagram-api-go-client

Thank you for your interest in contributing to this project! This document provides guidelines and instructions for contributing.

## Getting Started

### Prerequisites

- Go 1.16 or later
- Git
- Docker (required for API client generation via Swagger)
- Make (optional, for using Makefile commands)

### Setting Up Your Development Environment

1. Fork the repository on GitHub
2. Clone your fork locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/instagram-api-go-client.git
   cd instagram-api-go-client
   ```
3. Add the upstream repository:
   ```bash
   git remote add upstream https://github.com/qcserestipy/instagram-api-go-client.git
   ```
4. Install dependencies:
   ```bash
   go mod download
   ```

## Development Workflow

### Creating a Branch

Create a feature branch for your work:
```bash
git checkout -b feature/your-feature-name
# or for bug fixes
git checkout -b fix/bug-description
```

### Making Changes

1. Make your changes following the code style conventions used in the project
2. Test your changes thoroughly
3. Keep commits logical and well-described:
   ```bash
   git commit -m "feat: add new feature" # for new features
   git commit -m "fix: resolve issue" # for bug fixes
   git commit -m "docs: update documentation" # for documentation changes
   ```

### Code Style

- Follow Go conventions and best practices
- Run `go fmt` to format your code
- Ensure code is linted with `go vet`
- Add comments for exported functions and types
- Keep functions focused and modular

### Testing

Before submitting a pull request:
1. Run all tests:
   ```bash
   go test ./...
   ```
2. Ensure all tests pass
3. Add tests for new functionality or bug fixes
4. Maintain or improve code coverage

### Building

To build the project:
```bash
make build
# or manually
go build -o bin/instagram-media-insights-go-client ./cmd/main
```

### Working with API Swagger Manifests

The API Swagger manifests are maintained in a separate repository: [instagram-api-swagger-manifests](https://github.com/qcserestipy/instagram-api-swagger-manifests)

The `api/` directory is a Git submodule that references the latest manifests. If the upstream manifests are updated, you'll need to update the submodule:

```bash
make update-submodule
```

Or manually:
```bash
cd api
git pull --rebase
cd ..
```

After updating the submodule, regenerate the API clients:
```bash
make gen-all-clients
```

This will regenerate the Go client code based on the latest Swagger specifications.

## Submitting Changes

### Preparing Your Pull Request

1. Keep your branch up to date with main:
   ```bash
   git fetch upstream
   git rebase upstream/main
   ```
2. Push your branch to your fork:
   ```bash
   git push origin feature/your-feature-name
   ```
3. Open a Pull Request on GitHub with a clear description of your changes

### Pull Request Guidelines

- Provide a clear title and description of what your PR does
- Reference any related issues (e.g., "Fixes #123")
- Include any breaking changes clearly
- Ensure all CI checks pass
- Be open to feedback and review comments

## Reporting Bugs

When reporting a bug, include:

- Clear description of the issue
- Steps to reproduce the problem
- Expected behavior
- Actual behavior
- Go version and OS information
- Relevant code or error messages

## Suggesting Enhancements

For feature requests:

- Provide a clear use case
- Explain how the feature would benefit users
- Include examples if applicable
- Be open to discussion about implementation approach

## Project Structure

Key directories:

- `cmd/` - Command-line entry points
- `pkg/` - Main application packages
  - `access/` - Token and authentication handling
  - `account/` - Account-related operations
  - `client/` - HTTP client configuration
  - `config/` - Configuration management
  - `instagram/` - Instagram data types and operations
  - `media/` - Media-related operations
- `sdk/` - Generated SDK code for different API versions
- `api/` - API definitions and schemas

## License

By contributing to this project, you agree that your contributions will be licensed under the same license as the project.

## Questions?

Feel free to open an issue for questions or discussions about the project.

Thank you for contributing!
