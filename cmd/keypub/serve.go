package main

import (
	"github.com/spf13/cobra"
	"github.com/theSuess/keypub/pkg/server"
	"log"
)

var Interface string
var DBPath string

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start a server running the graphql api",
	Run: func(cmd *cobra.Command, args []string) {
		s := server.New(server.Configuration{
			DatabasePath: DBPath,
			Interface:    Interface,
		})
		log.Fatal(s.Run())
	},
}

func init() {
	serveCmd.Flags().StringVarP(&Interface, "interface", "i", ":8080", "Interface to listen on")
	serveCmd.Flags().StringVarP(&DBPath, "database", "d", "keypub.db", "Database file")
}
