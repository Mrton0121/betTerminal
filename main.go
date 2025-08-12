package main

import (
	"betterminal/commands"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	configPath := os.Getenv("BETTERMINAL_CONFIG")
	if configPath == "" {
		log.Fatalf("BETTERMINAL_CONFIG environment variable is not set")
	}

	var cmds commands.Commands
	cmds.Commands = append(cmds.Commands, &commands.Command{
		Name:     "init",
		ArgCount: 1,
		Exec:     []string{"touch ~/.betterminal/config.yaml", "curl $1 -o ~/.betterminal/config.yaml"},
		HelpText: "please provide the url of the file you want to use as config",
	})

	if len(os.Args) > 1 {
		if os.Args[1] == "init" {
			err := cmds.GetCommand(os.Args[1]).Execute(os.Args[1:])
			if err != nil {
				log.Fatalf("error executing init command: %v", err)
			}
			return
		}
	}

	data, err := ioutil.ReadFile(configPath + "/config.yaml")
	if err != nil {
		fmt.Println("Please use the init command to create your config.yaml or make it directly in your path")
		log.Fatalf("error: %v", err)
	}

	err = yaml.Unmarshal([]byte(data), &cmds)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if len(os.Args) == 1 {
		cmds.ListCommands()
		os.Exit(1)
	}

	command := cmds.GetCommand(os.Args[1])
	if command == nil {
		log.Fatalf("No such command: %s\n", os.Args[1])
	}

	err = command.Execute(os.Args[1:])
	if err != nil {
		log.Fatalf("Error executing command: %s\n\t%v\n", os.Args[1], err)
	}

	historyPath := configPath + "/history"
	historyFile, err := os.OpenFile(historyPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Could not open history file: %v", err)
	} else {
		defer historyFile.Close()
		args := ""
		if len(os.Args) > 1 {
			args = os.Args[1]
			for _, arg := range os.Args[2:] {
				args += " " + arg
			}
		}
		args += "\n"
		_, err := historyFile.WriteString(args)
		if err != nil {
			log.Printf("Could not write to history file: %v", err)
		}
	}

}
