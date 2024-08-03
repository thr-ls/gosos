# gosos (Save Our Services)

gosos is a simple CLI tool for monitoring website and API statuses.

I needed something I could use to check the multiple self-hosted services I use.

## Features

- Add and remove URLs for monitoring
- List all registered URLs
- Check the status of all URLs at once
- Real-time monitoring with customizable intervals
- User-friendly command-line interface

## Installation

To install gosos, make sure you have Go installed on your system, then run:

```
go install git.thrls.net/thrls/gosos@latest
```

## Usage

gosos provides several commands for managing and monitoring URLs:

```
gosos <command> [options]
```

### Commands

- `add <url>`: Add a URL to the monitoring list
- `remove <url>`: Remove a URL from the monitoring list
- `list`: Display all registered URLs
- `run`: Check the status of all registered URLs once
- `live [interval]`: Start monitoring all URLs in real-time
    - `[interval]`: Optional check interval in seconds (default: 30)
- `help`: Show the help message

### Examples

```
gosos add https://example.com
gosos remove https://example.com
gosos list
gosos run
gosos live 60  # Check every 60 seconds
```

## Configuration

gosos stores the list of URLs in a JSON file located at `~/.gosos-urls.json`. This file is automatically created by the tool.

## Development

To contribute to gosos or set up the development environment:

1. Clone the repository:
   ```
   git clone git.thrls.net/thrls/gosos
   ```
2. Navigate to the project directory:
   ```
   cd gosos
   ```
3. Install dependencies:
   ```
   go mod tidy
   ```
4. Build the project:
   ```
   go build
   ```

## Dependencies

- [pterm](https://github.com/pterm/pterm): For terminal output styling and live updates
- Standard Go libraries

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Support

If you encounter any problems or have any questions, please open an issue on the project repository.

## Todo
- [ ] Create tests
- [ ] Fix a few bugs with pterm terminal output