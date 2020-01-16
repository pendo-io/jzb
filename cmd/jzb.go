package main

import (
	"github.com/pendo-io/jzb/internal/app"
	"github.com/pendo-io/jzb/internal/cfg"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:          "jzb",
		Short:        "A tool for working with the jzb format",
		SilenceUsage: true,
	}
	config = &cfg.CommandLineArguments{}
)

func main() {
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		return app.Execute(*config)
	}
	rootCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if err := config.Validate(); err != nil {
			_ = cmd.Help()
			return err
		}
		return nil
	}
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&config.Create, "create", "c", false, "create a jzb")
	rootCmd.PersistentFlags().BoolVarP(&config.Extract, "extract", "x", false, "extract from a jzb")
	rootCmd.PersistentFlags().StringVarP(&config.InputPath, "file", "f", "-", "(Optional) path to file/jzb.")
	rootCmd.PersistentFlags().StringVarP(&config.OutputFile, "output", "o", "", "(Optional) output file path. (default stdout)")
}
