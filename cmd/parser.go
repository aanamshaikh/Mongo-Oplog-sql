package cmd

import (
	"fmt"
	"os"

	"Mongo-Oplog-sql/config"
	"Mongo-Oplog-sql/parser"

	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: `A simple CLI application`,
	Long:  `A simple CLI application to parse MongoDB oplog to SQL`,
	RunE: func(cmd *cobra.Command, args []string) error {
		config.ConnectMongo()
		parser.ConvertToSql()
		return nil
	},
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "A brief description of your application",
	Long:  `A longer description of your application.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(parseCmd)
}
