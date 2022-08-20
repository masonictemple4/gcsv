package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var DELIM string
var rootCmd = &cobra.Command{
	Use:   "gcsv",
	Short: "CLI tool to manage csvs",
	Long:  "CLI tool to manage csvs",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalf("Please provide a filepath")
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("delimiter", "d", "|", "Specify delimiter.")
	rootCmd.MarkPersistentFlagRequired("delimiter")
}
