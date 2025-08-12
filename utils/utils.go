package utils

import (
	"fmt"
	"strings"
)

func ReplaceArgument(command string, index int, value string) string {
	placeholder := fmt.Sprintf("$%d", index)
	// Ensure the placeholder is only replaced when $ is followed by a number
	command = strings.Replace(command, placeholder, value, -1)
	return command
}

func ReplaceAllArguments(command string, indexCount int, values []string) string {
	for i := 0; i < indexCount; i++ {
		command = ReplaceArgument(command, i+1, values[i])
	}
	return command
}
