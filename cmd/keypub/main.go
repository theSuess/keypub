package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "keypub",
	Short: "Keypub is the place to discover and verify public keys",
	Long:  `A Keyserver focused on usability and identity verification`,
}

func main() {
	rootCmd.AddCommand(serveCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
