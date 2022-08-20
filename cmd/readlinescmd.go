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

var readLinesCmd = &cobra.Command{
	Use:   "read",
	Short: "Read the data out for each line.",
	Long:  "Read the data out for each line.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalf("Please provide a filepath")
		}

		var err error
		DELIM, err = cmd.Flags().GetString("delimiter")
		if err != nil {
			log.Fatalf("There was a problem parsing the flags. %v", err)
		}
		readLines(args[0])
	},
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
