package cmd

import (
	"fmt"

	"github.com/derailed/k9s/internal/color"
	"github.com/spf13/cobra"
)

func versionCmd() *cobra.Command {
	var short bool

	command := cobra.Command{
		Use:   "version",
		Short: "Print version/build info",
		Long:  "Print version/build information",
		Run: func(cmd *cobra.Command, args []string) {
			printVersion(short)
		},
	}

	command.PersistentFlags().BoolVarP(&short, "short", "s", false, "Prints K9s version info in short format")

	return &command
}

func printVersion(short bool) {
	const fmat = "%-20s %s\n"
	var sectionColor color.Paint

	if short {
		sectionColor = -1
	} else {
		printLogo(color.Magenta)
		sectionColor = color.Blue
	}
	printTuple(fmat, "Version", sectionColor, version, color.Black)
	printTuple(fmat, "Commit", sectionColor, commit, color.Black)
	printTuple(fmat, "Date", sectionColor, date, color.Black)
}

func printTuple(fmat, section string, sectionColor color.Paint, value string, valueColor color.Paint) {
	if sectionColor != -1 {
		fmt.Fprintf(out, fmat, color.Colorize(section+":", sectionColor), color.Colorize(value, valueColor))
		return
	}
	fmt.Fprintf(out, fmat, section, value)
}
