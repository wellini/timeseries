package commands

import "strings"

type CommandLineOptions struct {
	Command        string
	PositionalArgs []string
	flags          map[string]string
}

func (o CommandLineOptions) HasFlag(name string) bool {
	_, ok := o.flags[name]
	return ok
}

func (o CommandLineOptions) Flag(name string) string {
	return o.flags[name]
}

func (o CommandLineOptions) addFlag(name, val string) {
	o.flags[name] = val
}

func NewCommandLineOptions(args []string) CommandLineOptions {
	o := CommandLineOptions{
		flags: make(map[string]string, 0),
	}

	argsLength := len(args)
	for i := range args {
		if strings.HasPrefix(args[i], "--") && i+1 < argsLength {
			o.addFlag(strings.TrimLeft(args[i], "-"), args[i+1])
		} else {
			o.flags[args[i]] = ""
		}
	}

	return o
}
