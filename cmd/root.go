package cmd

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

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

		var err error

		DELIM, err = cmd.Flags().GetString("delimiter")
		if err != nil {
			log.Fatalf("There was a problem parsing the flags. %v", err)
		}

		listKeys, err := cmd.Flags().GetBool("keys")
		if err != nil {
			log.Fatalf("There was a problem parsing the flags. %v", err)
		}

		counts, err := cmd.Flags().GetBool("counts")
		if err != nil {
			log.Fatalf("There was a problem parsing the flags. %v", err)
		}

		lines, err := cmd.Flags().GetBool("line")
		if err != nil {
			log.Fatalf("There was a problem parsing the flags. %v", err)
		}

		if listKeys {
			readKeys(args[0])
		}

		if counts {
			countLineElements(args[0])
		}

		if lines {
			readLines(args[0])
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
	rootCmd.PersistentFlags().BoolP("keys", "k", false, "List csv keys.")
	rootCmd.PersistentFlags().BoolP("counts", "c", false, "Count the elements on each line")
	rootCmd.PersistentFlags().BoolP("line", "l", false, "Read line values")
	rootCmd.PersistentFlags().StringP("delimiter", "d", "|", "Specify delimiter.")
	rootCmd.MarkFlagRequired("delimiter")
}

func readLines(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("File not found: %v", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	lineNum := 0
	keys := make([]string, 221)
	for {

		line, err := reader.ReadBytes(byte('\n'))
		if err == io.EOF {
			break
		}

		if lineNum == 0 {
			keys = strings.Split(string(line), DELIM)
			lineNum = lineNum + 1
			continue
		}

		elems := strings.Split(string(line), DELIM)
		fmt.Printf("LINE: %d ", lineNum)
		for i := range keys {
			fmt.Printf("%s: %s ", keys[i], elems[i])
		}

		println("-------------------------------------------------")

		lineNum = lineNum + 1

	}
}

func readKeys(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("File not found: %v", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	line, err := reader.ReadBytes(byte('\n'))
	if err != nil {
		log.Fatalf("There was a problem reading the first line.")
	}

	keys := strings.Split(string(line), DELIM)

	for k := range keys {
		println(keys[k])
	}
	fmt.Printf("The csv has: %d keys", len(keys))
}

func countLineElements(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("File not found: %v", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	lineNum := 0
	for {
		if lineNum == 0 {
			lineNum = lineNum + 1
			continue
		}

		line, err := reader.ReadBytes(byte('\n'))
		if err == io.EOF {
			break
		}

		elems := strings.Split(string(line), DELIM)

		fmt.Printf("Line: %d has Elems: %d\n", lineNum, len(elems))
		lineNum = lineNum + 1

	}

}
