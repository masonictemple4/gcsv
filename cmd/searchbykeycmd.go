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

var searchByKeyCmd = &cobra.Command{
	Use:   "search-by-key",
	Short: "Search for specific values for a given key.",
	Long:  "Search for specific values for a given key.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalf("Please provide a filepath")
		}

		var err error
		DELIM, err = cmd.Flags().GetString("delimiter")
		if err != nil {
			log.Fatalf("There was a problem parsing the flags. %v", err)
		}
		countLineElements(args[0])
	},
}

func init() {
	rootCmd.AddCommand(searchByKeyCmd)
}

func searchVals(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("File not found: %v", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	lineNum := 0
	for {
		if lineNum == 0 {
			// set the keys and search for valid index (corresponds to keyname).
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
