package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/urfave/cli/v2"
)

var Version string = "v0.0.0-dev"

func main() {

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print version of this tool",
	}

	cli.VersionPrinter = func(cCtx *cli.Context) {
		fmt.Printf("%s\n", cCtx.App.Version)
	}

	app := &cli.App{
		Name:        "nextversion",
		Usage:       "versioning helper tool",
		Version:     Version,
		Description: "nextversion detects application version based on git tags and suggests a bumped version based on conventional commits.",
		Flags:       appFlags(),
		Action:      appAction,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func appFlags() []cli.Flag {

	return []cli.Flag{

		&cli.StringFlag{
			Name:    "repo",
			Value:   "./",
			Aliases: []string{"r"},
			Usage:   "`PATH` to a git repository",
		},

		&cli.StringFlag{
			Name:    "format",
			Value:   "shell",
			Aliases: []string{"f"},
			Usage:   "Output `FORMAT` (shell, json)",
			Action:  verifyFormat,
		},

		&cli.StringFlag{
			Name:    "default",
			Value:   "0.1.0",
			Aliases: []string{"d"},
			Usage:   "Default `VERSION` if none could be detected",
		},

		&cli.BoolFlag{
			Name:    "prestable",
			Value:   false,
			Aliases: []string{"p"},
			Usage:   "Pre-stable mode",
		},

		&cli.StringFlag{
			Name:    "missmatch",
			Value:   "skip-merge",
			Aliases: []string{"m"},
			Usage:   "Behaviour of commit message parser (skip, skip-merge, fail)",
			Action:  verifyMissmatch,
		},
	}
}

func verifyFormat(ctx *cli.Context, value string) error {

	valid := []string{"shell", "json"}
	if !slices.Contains[[]string, string](valid, value) {
		return fmt.Errorf("--format must be one of [%s]", strings.Join(valid, ", "))
	}
	return nil
}

func verifyMissmatch(ctx *cli.Context, value string) error {

	valid := []string{"skip", "skip-merge", "fail"}
	if !slices.Contains[[]string, string](valid, value) {
		return fmt.Errorf("--missmatch must be one of [%s]", strings.Join(valid, ", "))
	}
	return nil
}

func appAction(ctx *cli.Context) error {

	lines := map[string]string{
		"PREVIOUS_VERSION": "0.1.0",
		"NEXT_VERSION":     "0.2.0",
	}

	if ctx.Bool("prestable") {
		lines["NEXT_VERSION"] = "0.1.1"
	}

	switch ctx.String("format") {
	case "shell":
		for key, val := range lines {
			fmt.Printf("%s=%s%s", key, val, EOF)
		}
	case "json":
		jsonString, err := json.Marshal(lines)
		if err != nil {
			return fmt.Errorf("JSON error: %w", err)
		}

		fmt.Printf("%s%s", jsonString, EOF)
	}

	return nil

}
