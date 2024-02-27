package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/tsdtsdtsd/nextversion/pkg/nextversion"
	"github.com/urfave/cli/v2"
)

func newApp() *cli.App {

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
		UsageText:   "nextversion [global options] [command]",
		Version:     version,
		Description: "nextversion detects application version based on git tags and suggests a bumped version based on conventional commits.",
		Flags:       appFlags(),
		Action:      appAction,
	}

	return app
}

func appAction(ctx *cli.Context) error {

	opts := &nextversion.Options{
		Repo:           ctx.String("repo"),
		Format:         ctx.String("format"),
		DefaultCurrent: ctx.String("defaultCurrent"),
		Prestable:      ctx.Bool("prestable"),
	}

	versions, err := nextversion.Versions(opts)
	if err != nil {
		return fmt.Errorf("failed to detect versions: %w", err)
	}

	return nextversion.Print(os.Stdout, versions, ctx.String("format"))

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
			Value:   "simple",
			Aliases: []string{"f"},
			Usage:   "Output `FORMAT` (simple, json)",
			Action:  verifyFormatValue,
		},

		&cli.StringFlag{
			Name:    "defaultCurrent",
			Value:   "v0.0.0",
			Aliases: []string{"d"},
			Usage:   "Fallback current `VERSION` if none could be detected",
		},

		&cli.BoolFlag{
			Name:    "prestable",
			Value:   false,
			Aliases: []string{"p"},
			Usage:   "Pre-stable mode",
		},
	}
}

func verifyFormatValue(ctx *cli.Context, value string) error {

	valid := []string{"simple", "json"}
	if !slices.Contains[[]string, string](valid, value) {
		return fmt.Errorf("--format must be one of [%s]", strings.Join(valid, ", "))
	}
	return nil
}
