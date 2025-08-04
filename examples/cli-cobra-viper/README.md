# CLI Cobra Viper Example

A comprehensive CLI application template demonstrating how to use **Cobra** (CLI framework) and **Viper** (configuration management) together. This example shows different ways to configure your CLI application using command-line flags, environment variables, and configuration files.

## Overview

This application demonstrates:
- Building CLI applications with Cobra
- Configuration management with Viper
- Multiple configuration sources (flags, environment variables, config files)
- Configuration priority and binding
- Subcommands and nested command structures

## Project Structure

```
cli-cobra-viper/
├── main.go                 # Application entry point
├── cmd/                    # Command definitions
│   ├── root.go            # Cobra Root command and configuration setup
│   ├── command1.go        # First command with subcommand
│   ├── command2.go        # Second command
│   └── command2_sub2.go   # Subcommand of command2
├── cli-cobra-viper.yaml   # Example configuration file
├── go.mod                 # Go module file
└── README.md              # This file
```

## Available Commands

### Root Command
```bash
cli-cobra-viper [flags]
```

### Command 1 (`cmd1`)
A command that demonstrates message handling with subcommands.

**Subcommand: `sub`**
```bash
cli-cobra-viper cmd1 sub [message] [flags]
```

### Command 2 (`cmd2`)
A command that demonstrates string and number configuration.
```bash
cli-cobra-viper cmd2 [flags]
```

**Subcommand: `sub2`**
```bash
cli-cobra-viper cmd2 sub2 [flags]
```

## Configuration Methods

Viper supports multiple configuration sources with the following priority (highest to lowest):
1. **Command-line flags** (highest priority)
2. **Environment variables**
3. **Configuration file**
4. **Default values** (lowest priority)

### 1. Command-Line Flags

#### Basic Usage
```bash
# Build the application first
go build -o cli-cobra-viper

# Show help
./cli-cobra-viper --help

# Enable verbose output
./cli-cobra-viper --verbose cmd2

# Command 1 with custom message
./cli-cobra-viper cmd1 sub "Hello from CLI" --message "Config message"

# Command 2 with custom values
./cli-cobra-viper cmd2 --str "Custom string" --num 100

# Command 2 sub2 with custom values
./cli-cobra-viper cmd2 sub2 --str "Sub2 custom" --num 200

### 2. Environment Variables

Environment variables use the prefix `CLI_` and convert nested config keys using underscores.

#### Setting Environment Variables
```bash
# Global configuration
export CLI_VERBOSE=true

# MyApp configuration
export CLI_MYAPP_STR="Hello from environment"
export CLI_MYAPP_NUM=999
export CLI_MYAPP_MESSAGE="Environment message"

# Sub2 configuration
export CLI_MYAPP_SUB2_STR="Sub2 from env"
export CLI_MYAPP_SUB2_NUM=888

# API configuration
export CLI_API_ENDPOINT="https://api.staging.com"
export CLI_API_TOKEN="env-token-456"
```

#### Running with Environment Variables
```bash
# Run commands (environment variables will be automatically picked up)
./cli-cobra-viper cmd1 sub "CLI argument"
./cli-cobra-viper cmd2
./cli-cobra-viper cmd2 sub2
```

#### One-liner Examples
```bash
# Set environment and run in one command
CLI_MYAPP_STR="Inline env" CLI_MYAPP_NUM=777 ./cli-cobra-viper cmd2

# Override with multiple environment variables
CLI_MYAPP_SUB2_STR="Inline sub2" CLI_MYAPP_SUB2_NUM=555 ./cli-cobra-viper cmd2 sub2
```

### 3. Configuration File

The application looks for configuration files in the following locations:
- Current directory: `./cli-cobra-viper.yaml`
- Home directory: `$HOME/.cli-cobra-viper.yaml`
- Custom location: specified with `--config` flag

#### Example Configuration File (`cli-cobra-viper.yaml`)
```yaml
# Global settings
verbose: true

# Application-specific configuration
myapp:
  str: "Hello from config file"
  num: 333
  message: "Config file message"
  sub2:
    str: "Sub2 from config"
    num: 444

# API configuration
api:
  endpoint: "https://api.production.com"
  token: "config-token-789"
```

#### Using Configuration Files
```bash
# Use default config file location
./cli-cobra-viper cmd2

# Use custom config file
./cli-cobra-viper --config /path/to/my-config.yaml cmd2

# Create a custom config for testing
cat > test-config.yaml << EOF
verbose: false
myapp:
  str: "Test config string"
  num: 1234
  message: "Test message"
  sub2:
    str: "Test sub2 string"
    num: 5678
api:
  endpoint: "https://test-api.com"
  token: "test-token"
EOF

# Use the test config
./cli-cobra-viper --config test-config.yaml cmd2
```

## Configuration Priority Examples

### Example 1: All Sources Combined
```bash
# 1. Create config file
cat > priority-test.yaml << EOF
myapp:
  str: "From config file"
  num: 100
EOF

# 2. Set environment variable
export CLI_MYAPP_NUM=200

# 3. Run with command flag
./cli-cobra-viper --config priority-test.yaml cmd2 --str "From command line"

# Result: 
# - str: "From command line" (command flag wins)
# - num: 200 (environment variable wins over config file)
```

### Example 2: Environment vs Config File
```bash
# Set environment
export CLI_MYAPP_MESSAGE="Environment wins"

# Run with config file that also has myapp.message
./cli-cobra-viper cmd1 sub "test" 

# Environment variable will override config file value
```

## Building and Running

### Development
```bash
# Run directly with Go
go run main.go cmd2 --str "Development mode"

# Run with verbose output
go run main.go --verbose cmd1 sub "test message"
```

### Production Build
```bash
# Build binary
go build -o cli-cobra-viper

# Run binary
./cli-cobra-viper --help

# Install globally (optional)
go install
```

## VS Code Debugging Support

This project includes VS Code launch configurations in `.vscode/launch.json` that make it easy to debug and test the application directly from the editor.

### Available Debug Configurations

The launch.json file provides several pre-configured debug sessions:

1. **"Run app - no args"** - Runs the application with no arguments (shows help)
2. **"Run Command 1"** - Executes `cmd1 sub "Hello World From Args"`
3. **"Run Command 2"** - Executes `cmd2` command
4. **"Run Command 2 Sub 2"** - Executes `cmd2 sub2` subcommand

### How to Use VS Code Debugging

1. **Open the project in VS Code**
2. **Set breakpoints** in any Go file by clicking in the left margin
3. **Open the Run and Debug panel** (Ctrl+Shift+D / Cmd+Shift+D)
4. **Select a configuration** from the dropdown
5. **Press F5** or click the green play button to start debugging

### Debugging Features Demonstrated

Each configuration shows different aspects of the application:

```jsonc
{
    "name": "Run Command 1",
    "type": "go",
    "request": "launch",
    "mode": "auto",
    "program": "${workspaceFolder}/examples/cli-cobra-viper/main.go",
    "args": ["cmd1", "sub", "Hello World From Args"],
    "env": {
        // Uncomment to test environment variables:
        // "CLI_MYAPP_MESSAGE": "Hello World from Environment"
    },
    "cwd": "${workspaceFolder}/examples/cli-cobra-viper"
}
```

### Customizing Debug Configurations

You can modify the configurations to test different scenarios:

- **Change arguments**: Edit the `"args"` array to test different commands
- **Add environment variables**: Uncomment or add entries in the `"env"` object
- **Test configuration files**: The debugger will automatically pick up `cli-cobra-viper.yaml` if present

### Example Debug Scenarios

1. **Test argument parsing**: Set breakpoints in command handlers to see how arguments are processed
2. **Test configuration loading**: Set breakpoints in `root.go` to see how Viper loads config
3. **Test environment variables**: Uncomment env vars in launch.json and debug to see priority handling
4. **Test error handling**: Modify arguments to invalid values and debug error paths

This makes the development cycle much faster compared to building and running from the command line each time.

## Testing Different Configurations

### Test Script Example
```bash
#!/bin/bash
echo "=== Testing CLI Cobra Viper ==="

# Build the application
go build -o cli-cobra-viper

echo "1. Testing with defaults:"
./cli-cobra-viper cmd2

echo -e "\n2. Testing with command-line flags:"
./cli-cobra-viper cmd2 --str "CLI Flag" --num 999

echo -e "\n3. Testing with environment variables:"
CLI_MYAPP_STR="Environment" CLI_MYAPP_NUM=777 ./cli-cobra-viper cmd2

echo -e "\n4. Testing with config file:"
./cli-cobra-viper --config cli-cobra-viper.yaml cmd2

echo -e "\n5. Testing priority (config + env + flag):"
CLI_MYAPP_NUM=888 ./cli-cobra-viper --config cli-cobra-viper.yaml cmd2 --str "Flag wins"

echo -e "\n6. Testing nested subcommand:"
./cli-cobra-viper cmd2 sub2 --str "Sub2 test" --num 456
```

## Key Learning Points

1. **Cobra Framework**: Provides structure for building CLI applications with commands, subcommands, and flags.

2. **Viper Configuration**: Handles multiple configuration sources automatically with a clear priority system.

3. **Flag Binding**: Use `viper.BindPFlag()` to connect command flags to configuration keys.

4. **Environment Variables**: Automatic mapping with prefix and key transformation (dots → underscores).

5. **Configuration Hierarchy**: Command flags > Environment variables > Config files > Defaults.

6. **Flexible Config Files**: Support for YAML, JSON, TOML, and other formats.

## Best Practices Demonstrated

- **Separation of Concerns**: Each command in its own file
- **Consistent Naming**: Clear, hierarchical configuration key structure
- **Security**: Sensitive values (tokens) can be set via environment variables
- **User Experience**: Multiple ways to configure the same application
- **Documentation**: Clear help messages and examples

This template provides a solid foundation for building robust CLI applications in Go with flexible configuration management.
