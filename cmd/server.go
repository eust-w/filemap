/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"markdownFolderMap/implement"

	"github.com/spf13/cobra"
)

var port string
var path string
var filter string

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server",
	Long:  `server`,
	Run: func(cmd *cobra.Command, args []string) {
		implement.Server(port, path, filter)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&port, "port", "o", "8080", "server port")
	serverCmd.Flags().StringVarP(&path, "path", "p", "", "mindmap path")
	serverCmd.Flags().StringVarP(&filter, "filter", "f", "", "filter ignore")
}
