/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "desafio-tecnico-stress-test",
	Short: "Teste de stress para endpoints http",
	Long: `Ferramenta para execução de testes de stress para 
endpoints http.

Nesta ferramenta, informa-se a URL, a quantidade de requisições 
e quantidade de chamadas simultâneas (vcpus) para a execução do
Teste.
Por fim tem-se um relatório do resultado do teste.`,
	Run: func(cmd *cobra.Command, args []string) {
		distribution := make(map[int]int)
		requests, err := strconv.Atoi(argRequests)
		if err != nil {
			panic("invalid requests")
		}
		concurrency, err := strconv.Atoi(argConcurrency)
		if err != nil {
			panic("invalid requests")
		}
		println(url)
		println(requests)
		println(concurrency)

		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		distribution[resp.StatusCode] = 1
		for k, v := range distribution {
			println(k)
			println(v)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var url string
var argRequests string
var argConcurrency string

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.desafio-tecnico-stress-test.yaml)")
	rootCmd.PersistentFlags().StringVar(&url, "url", "", "web service to test")
	rootCmd.PersistentFlags().StringVar(&argRequests, "requests", "1", "requests quantity")
	rootCmd.PersistentFlags().StringVar(&argConcurrency, "concurrency", "1", "simultaneous requests quantity")

	rootCmd.MarkPersistentFlagRequired("url")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
