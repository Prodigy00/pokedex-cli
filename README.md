# Pokedex CLI - A Command-Line Pokémon Explorer

Pokedex CLI is a command-line application built in Golang that allows users to explore Pokémon data interactively. It fetches Pokémon details from an external API and provides various commands to search, list, and view Pokémon information.

## Features

- Search for Pokémon by name or ID
- List all Pokémon in a paginated format
- View detailed stats, abilities, and types of a Pokémon
- Caching mechanism to speed up repeated queries
- Simple and intuitive command-line interface

## Tech Stack

- **Golang** - The core language for development
- **REST API** - Fetching Pokémon data


## Installation

### Prerequisites

Ensure you have the following installed:

- [Go](https://go.dev/dl/) (v1.19+ recommended)

### Clone and Build

1. Clone the repository:
   ```sh
   git clone https://github.com/Prodigy00/pokedex-cli.git
   cd pokedex-cli
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Build the CLI application:
   ```sh
   go build -o pokedex
   ```
4. Run the application:
   ```sh
   ./pokedex
   ```

## Usage

### Available Commands

| Command          | Description                             |
|-----------------|-----------------------------------------|
| `pokedex help`  | Show available commands and usage      |
| `pokedex list`  | List Pokémon with pagination          |
| `pokedex info <name>` | Get details about a specific Pokémon |
| `pokedex catch <name>` | Attempt to catch a Pokémon and save it locally |
| `pokedex cache` | View cached Pokémon data               |

### Example Usage

1. Get information about Pikachu:
   ```sh
   ./pokedex info pikachu
   ```
2. List available Pokémon:
   ```sh
   ./pokedex list
   ```
3. Catch a Pokémon:
   ```sh
   ./pokedex catch charmander
   ```

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-name`).
3. Commit your changes (`git commit -m 'Add new feature'`).
4. Push to your branch (`git push origin feature-name`).
5. Open a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

For any inquiries or suggestions, feel free to reach out:

- GitHub: [Prodigy00](https://github.com/Prodigy00)

---

Gotta catch 'em all! 🎮

