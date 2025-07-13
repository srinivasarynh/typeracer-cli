# TypeRacer CLI

A production-ready command-line typing speed test application built with Go.

## Installation

1. **Create a new directory for your project:**
   ```bash
   mkdir typeracer-cli
   cd typeracer-cli
   ```

2. **Initialize Go module:**
   ```bash
   go mod init typeracer-cli
   ```

3. **Install dependencies:**
   ```bash
   go mod tidy
   ```

4. **Build the application:**
   ```bash
   go build -o typeracer .
   ```

5. **Run the application:**
   ```bash
   ./typeracer
   ```

## Usage

### Basic Usage
```bash
./typeracer
```

### With Options
```bash
./typeracer --words 30 --difficulty hard --time 120
```

### View Statistics
```bash
./typeracer stats
```

### View Configuration
```bash
./typeracer config
```

## Available Flags

| Flag | Short | Type | Description | Default |
|------|-------|------|-------------|---------|
| `--words` | `-w` | int | Number of words to type | 20 |
| `--difficulty` | `-d` | string | Difficulty level: easy, medium, hard | medium |
| `--time` | `-t` | int | Time limit in seconds, 0 for no limit | 60 |
| `--help` | `-h` | - | Help for typeracer | - |

## Features

1. **Production-grade architecture** with proper folder structure
2. **CLI interface** using Cobra
3. **Configuration management** with Viper
4. **Colorized output** using fatih/color
5. **Statistics tracking** and persistence
6. **Real-time typing feedback**
7. **Goroutines** for input handling and timer
8. **Channels** for communication between goroutines
9. **Cross-platform terminal** input handling
10. **Difficulty levels** with appropriate word lists

## Architecture

```
typeracer-cli/
├── cmd/                    # CLI commands and flags
├── internal/              # Internal application logic
├── config/                # Configuration management
├── game/                  # Game engine and statistics
├── ui/                    # User interface and input handling
├── pkg/                   # Public packages (word generator)
└── main.go               # Application entry point
```

### Directory Structure Details

- **`cmd/`**: CLI commands and flags
- **`internal/`**: Internal application logic
- **`config/`**: Configuration management
- **`game/`**: Game engine and statistics
- **`ui/`**: User interface and input handling
- **`pkg/`**: Public packages (word generator)
- **`main.go`**: Application entry point

## Technical Highlights

This is a production-ready application that demonstrates:

- **Clean architecture** patterns
- **Proper error handling** throughout the codebase
- **Concurrent programming** with goroutines
- **Channel communication** between components
- **File I/O** for persistence
- **Terminal UI** programming
- **CLI best practices** and conventions

## Requirements

- Go 1.19 or higher
- Terminal with color support (recommended)

## Contributing

This application follows Go best practices and clean architecture principles, making it easy to extend and maintain.
