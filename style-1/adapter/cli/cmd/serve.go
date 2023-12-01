/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/booscaaa/desafio-client-server-api-go-expert-pos/style-1/adapter/http/rest"
	"github.com/booscaaa/desafio-client-server-api-go-expert-pos/style-1/adapter/httpclient"
	"github.com/booscaaa/desafio-client-server-api-go-expert-pos/style-1/adapter/sqlite"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		serverPort, _ := cmd.Flags().GetString("port")

		database := sqlite.InitializeDatabase()
		httpClient := httpclient.InitializeHttpClient()
		rest.InitializeRest(serverPort, database, httpClient)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringP("port", "p", "8080", "Server port")
}
