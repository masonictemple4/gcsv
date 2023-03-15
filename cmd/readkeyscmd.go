package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var readKeysCmd = &cobra.Command{
	Use:   "read-keys <filepath>",
	Short: "Get the keys from the csv.",
	Long:  "Get the keys from the csv.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalf("Please provide a filepath")
		}

		var err error
		DELIM, err = cmd.Flags().GetString("delimiter")
		if err != nil {
			log.Fatalf("There was a problem parsing the flags. %v", err)
		}
		readKeys(args[0])
	},
}

func init() {
	rootCmd.AddCommand(readKeysCmd)
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
