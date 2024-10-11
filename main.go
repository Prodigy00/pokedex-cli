package main

import (
	"bufio"
	"fmt"
	"github.com/Prodigy00/pokedexcli/internal/api"
	"github.com/Prodigy00/pokedexcli/internal/pokecache"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	cliName = "pokedexcli"
)

type config struct {
	NextURL       *string
	PreviousURL   *string
	cache         *pokecache.Cache
	caughtPokemon map[string]api.CatchPokemonResult
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config, args ...string) error
}

func (cmd *cliCommand) Description() string {
	return cmd.description
}

func (cmd *cliCommand) Name() string {
	return cmd.name
}

func Commands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "exit the pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "displays the names of 20 location areas in the Pokemon world at a time, each subsequent call displays 20 more locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "displays the names of 20 previous location areas, each subsequent call displays 20 more previous locations. It's a way to go back!",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "displays a list of all the Pokemon in a given area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "attempts to catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "displays the stats of a Pokemon that has been caught(seen) before",
			callback:    commandInspect,
		},
	}
}

func main() {
	cmds := Commands()
	printPrompt()

	//create cache
	cache := pokecache.NewCache(20 * time.Second)

	nextUrl := "https://pokeapi.co/api/v2/location-area"

	cfg := config{
		NextURL:       &nextUrl,
		PreviousURL:   nil,
		cache:         cache,
		caughtPokemon: make(map[string]api.CatchPokemonResult),
	}

	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		text := sanitize(reader.Text())
		parts := strings.Fields(text)

		if len(parts) > 0 {
			cmdName := parts[0]
			cmdArgs := parts[1:]

			if cmd, ok := cmds[cmdName]; ok {
				if err := cmd.callback(&cfg, cmdArgs...); err != nil {
					fmt.Fprintln(os.Stderr, err)
					return
				}
			} else {
				fmt.Fprintf(os.Stderr, "Invalid command: %s\n", text)
				return
			}

		}
		printPrompt()
	}
	// Print an additional line if we encounter an EOF character
	fmt.Println()
}

func printPrompt() {
	fmt.Print(cliName, "> ")
}

func sanitize(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}

func commandInspect(cfg *config, args ...string) error {
	name := args[0]
	cp, ok := cfg.caughtPokemon[name]
	if !ok {
		fmt.Fprintf(os.Stdout, "you have not caught that pokemon!\n")
		return nil
	}
	fmt.Printf("Name: %s\n", cp.Name)
	fmt.Printf("Height: %d\n", cp.Height)
	fmt.Printf("Weight: %d\n", cp.Weight)
	fmt.Println("Stats:")
	for _, k := range cp.Stats {
		switch k.Stat.Name {
		case "hp":
			fmt.Printf(" -hp: %d\n", k.BaseStat)
		case "attack":
			fmt.Printf(" -attack: %d\n", k.BaseStat)
		case "defense":
			fmt.Printf(" -defense: %d\n", k.BaseStat)
		case "special-attack":
			fmt.Printf(" -special-attack: %d\n", k.BaseStat)
		case "special-defense":
			fmt.Printf(" -special-defense: %d\n", k.BaseStat)
		case "speed":
			fmt.Printf(" -speed: %d\n", k.BaseStat)
		}
	}
	fmt.Println("Types:")
	for _, t := range cp.Types {
		fmt.Printf(" - %v\n", t.Type.Name)
	}

	return nil
}

func commandCatch(c *config, args ...string) error {
	newPokeApi := api.NewPokeAPI(c.cache)
	if len(args) < 1 {
		return fmt.Errorf("please provide a pokemon name")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	res, err := newPokeApi.CatchPokemon(args[0])
	if err != nil {
		return fmt.Errorf("an error occured attempting to catch pokemon with the name %s: %w", args[0], err)
	}

	chance := rand.Intn(res.BaseExperience)

	if chance < 50 {
		fmt.Printf("%s escaped!\n", args[0])
		return nil
	}

	fmt.Printf("%s was caught!\n", args[0])

	c.caughtPokemon[args[0]] = res

	keys := make([]string, 0, len(c.caughtPokemon))
	for k := range c.caughtPokemon {
		keys = append(keys, k)
	}
	fmt.Printf("caught Pokemon %s\n", keys)
	return nil
}

func commandExplore(c *config, args ...string) error {
	newPokeApi := api.NewPokeAPI(c.cache)
	if len(args) < 1 {
		return fmt.Errorf("please provide a valid area for exploration")
	}
	res, err := newPokeApi.GetLocationArea(args[0])
	if err != nil {
		return fmt.Errorf("an error occured fecthing location with the name %s: %w", args[0], err)
	}
	fmt.Printf("Exploring %s...\n", args[0])
	for _, v := range res {
		fmt.Printf("- %s\n", v)
	}
	return nil
}

func commandMapb(c *config, args ...string) error {
	newPokeApi := api.NewPokeAPI(c.cache)
	res, err := newPokeApi.GetLocationAreas(c.PreviousURL)
	if err != nil {
		return fmt.Errorf("an error occured fecthing locations: %w", err)
	}

	for _, result := range res.Results {
		fmt.Fprintf(os.Stdout, "%v\n", result.Name)
	}

	c.NextURL = res.Next
	c.PreviousURL = res.Previous
	return nil
}

func commandMap(c *config, args ...string) error {
	newPokeApi := api.NewPokeAPI(c.cache)
	res, err := newPokeApi.GetLocationAreas(c.NextURL)
	if err != nil {
		return fmt.Errorf("an error occured fecthing locations: %w", err)
	}

	for _, result := range res.Results {
		fmt.Fprintf(os.Stdout, "%v\n", result.Name)
	}

	c.NextURL = res.Next
	c.PreviousURL = res.Previous
	return nil
}

func commandHelp(c *config, args ...string) error {
	cmds := Commands()

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range cmds {
		fmt.Printf("%s:%s\n", cmd.Name(), cmd.Description())
	}
	return nil
}

func commandExit(c *config, args ...string) error {
	os.Exit(0)
	return nil
}
