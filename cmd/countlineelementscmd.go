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

var countLineElementsCmd = &cobra.Command{
	Use:   "count-line-elements",
	Short: "Counts the elements on each line",
	Long:  "Counts the elements on each line",
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
	rootCmd.AddCommand(countLineElementsCmd)
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
