package cmd

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"gcsv/utils"

	"github.com/spf13/cobra"
)

var readLinesCmd = &cobra.Command{
	Use:   "read-lines",
	Short: "Read the data out for each line.",
	Long:  "Read the data out for each line.",
	Run: func(cmd *cobra.Command, args []string) {
		includeKeys, _ := cmd.Flags().GetStringSlice("keys")
		if len(args) < 1 {
			log.Fatalf("Please provide a filepath")
		}

		var err error
		DELIM, err = cmd.Flags().GetString("delimiter")
		if err != nil {
			log.Fatalf("There was a problem parsing the flags. %v", err)
		}

		lc, err := cmd.Flags().GetInt("lines")
		if err != nil {
			log.Fatalf("There was a problem parsing the line flag. %v", err)
		}

		readLines(args[0], includeKeys, lc)
	},
}

func init() {
	rootCmd.AddCommand(readLinesCmd)
	readLinesCmd.PersistentFlags().StringSliceP("keys", "k", []string{}, "The key(s) that you wish to include")
	readLinesCmd.PersistentFlags().IntP("lines", "l", 0, "Specify a number of lines to read. 0 returns all.")
}

func readLines(path string, includeKeys []string, lc int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("File not found: %v", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	lineNum := 0
	keys := make([]string, 0)
	for {

		line, err := reader.ReadBytes(byte('\n'))
		if err == io.EOF {
			break
		}

		if lineNum == 0 {
			keys = strings.Split(string(line), DELIM)
			println("Key bank:")
			for k := range keys {
				println(keys[k])
			}
			lineNum = lineNum + 1
			continue
		}

		elems := strings.Split(string(line), DELIM)
		fmt.Printf("LINE: %d ", lineNum)
		for i := range keys {

			if len(includeKeys) > 0 && !utils.Contains(includeKeys, strings.TrimSpace(strings.Trim(keys[i], "\r\n"))) {
				continue
			}

			if i <= len(elems)-1 {
				fmt.Printf("%s: %s ", keys[i], elems[i])
			}

			// For csvs that don't add empty items
			// for missing fields.
			if i == len(elems)-1 {
				break
			}
		}

		println("-------------------------------------------------")

		lineNum = lineNum + 1

		// If lc specified break the loop
		if lc != 0 && lineNum > lc {
			break
		}

	}
}
