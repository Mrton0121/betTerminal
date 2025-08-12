package commands

import (
	"betterminal/utils"
	"fmt"
	"slices"
)

type Commands struct {
	Commands []*Command `yaml:"commands"`
}

func (c Commands) GetCommand(name string) *Command {
	for _, command := range c.Commands {
		if command.Name == name || slices.Contains(command.Alias, name) {
			return command
		}
	}
	return nil
}

func (c Commands) ListCommands() {
	fmt.Println("List of available commands and their usage:")
	for _, command := range c.Commands {
		command.PrintHelpText()
	}
}

type Command struct {
	Name     string   `yaml:"name"`
	Alias    []string `yaml:"alias"`
	ArgCount int      `yaml:"argCount"`
	Exec     []string `yaml:"exec"`
	HelpText string   `yaml:"helpText"`
}

func (c *Command) PrintHelpText() {
	fmt.Printf("[%s] \n%s\n", c.Name, c.HelpText)
}

func (c *Command) Execute(args []string) error {
	args = args[1:]
	if len(args) < c.ArgCount {
		c.PrintHelpText()
		return fmt.Errorf("provided less arguments than needed")
	}

	if len(args) > c.ArgCount {
		c.PrintHelpText()
		return fmt.Errorf("provided more arguments than needed")
	}

	for _, command := range c.Exec {
		// Replace argument placeholders
		replacedCommand := utils.ReplaceAllArguments(command, c.ArgCount, args)

		fmt.Printf("%s\n", replacedCommand)
	}

	return nil
}
