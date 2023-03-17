package commands

import "context"

type Command func(ctx context.Context, opts CommandLineOptions, args []string) error
