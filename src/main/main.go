package main

import (
	"fmt"
	"os"
	"timeseries/commands"
)

var m = map[string]commands.Command{
	"start": commands.Start,
}

func main() {
	if len(os.Args) <= 1 {
		panic("print help message instead of panicking")
	}

	command := os.Args[1]

	opts := commands.CommandLineOptions{
		Command: command,
	}

	_, ok := m[command]
	if !ok {
		panic("print help message instead of panicking")
	}

	opts := commands.NewCommandLineOptions(os.Args)
	fmt.Println(os.Args)
	fmt.Println(opts)
}
