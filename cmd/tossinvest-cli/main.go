package main

import (
	"context"
	"fmt"
	"os"

	"github.com/isul/tossinvest-cli/internal/cmd"
	"github.com/isul/tossinvest-cli/internal/output"
	"github.com/urfave/cli/v3"
)

func main() {
	app := cmd.NewApp()
	app.ExitErrHandler = func(_ context.Context, c *cli.Command, err error) {
		if err == nil {
			return
		}
		_ = output.WriteError(err, output.DefaultOptions())
		if c != nil && c.Root() != nil {
			os.Exit(1)
		}
		os.Exit(1)
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
