package main

import (
	"fmt"
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
		Version:     Version,
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
		OnMismatch:     ctx.String("on-mismatch"),
	}

	versions, err := nextversion.Versions(opts)
	if err != nil {
		return fmt.Errorf("failed to detect versions: %w", err)
	}

	return nextversion.Print(versions, ctx.String("format"))

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

		&cli.StringFlag{
			Name:    "on-mismatch",
			Value:   "skip-merge",
			Aliases: []string{"m"},
			Usage:   "Behaviour of commit message parser (skip, skip-merge, fail)",
			Action:  verifyOnMismatchValue,
		},
	}
}

func verifyFormatValue(ctx *cli.Context, value string) error {

	valid := []string{"shell", "json"}
	if !slices.Contains[[]string, string](valid, value) {
		return fmt.Errorf("--format must be one of [%s]", strings.Join(valid, ", "))
	}
	return nil
}

func verifyOnMismatchValue(ctx *cli.Context, value string) error {

	valid := []string{"skip", "skip-merge", "fail"}
	if !slices.Contains[[]string, string](valid, value) {
		return fmt.Errorf("--on-mismatch must be one of [%s]", strings.Join(valid, ", "))
	}
	return nil
}
