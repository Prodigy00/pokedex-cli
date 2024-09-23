package main

import (
	"bufio"
	"fmt"
	"github.com/Prodigy00/pokedexcli/internal/api"
	"os"
	"strings"
)

const (
	cliName = "pokedexcli"
)

type config struct {
	NextURL     *string
	PreviousURL *string
}
type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
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
			callback:    commadMapb,
		},
	}
}

func main() {
	cmds := Commands()
	printPrompt()

	nextUrl := "https://pokeapi.co/api/v2/location-area"
	cfg := config{
		NextURL:     &nextUrl,
		PreviousURL: nil,
	}

	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		text := sanitize(reader.Text())
		if cmd, ok := cmds[text]; ok {
			if err := cmd.callback(&cfg); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
		} else {
			fmt.Fprintf(os.Stderr, "Invalid command: %s\n", text)
			return
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

func commadMapb(c *config) error {
	res, err := api.GetLocationAreas(c.PreviousURL)
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

func commandMap(c *config) error {
	res, err := api.GetLocationAreas(c.NextURL)
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

func commandHelp(c *config) error {
	cmds := Commands()

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range cmds {
		fmt.Printf("%s:%s\n", cmd.Name(), cmd.Description())
	}
	return nil
}

func commandExit(c *config) error {
	os.Exit(0)
	return nil
}
