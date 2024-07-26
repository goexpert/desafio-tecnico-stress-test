/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

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
		// distribution := make(map[int]int)
		// var duration []int
		requests, err := strconv.Atoi(argRequests)
		if err != nil {
			panic("invalid requests")
		}
		concurrency, err := strconv.Atoi(argConcurrency)
		if err != nil {
			panic("invalid concurrency")
		}

		distOnConcurrency := distrubuteRequests(requests, concurrency)

		var wg sync.WaitGroup
		wg.Add(concurrency)
		for i := 0; i < concurrency; i++ {
			go loopRequests(&wg, url, distOnConcurrency[i])
		}

		wg.Wait()

		// for k, v := range distribution {
		// 	log.Println(k, v)
		// }
		// for k, v := range duration {
		// 	log.Println(k, v)
		// }
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

func request(url string) (code int, duration int64) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	totalTime := time.Since(start).Microseconds()
	return resp.StatusCode, totalTime

}

func loopRequests(wg *sync.WaitGroup, url string, qty int) {
	defer wg.Done()
	for i := 0; i < qty; i++ {
		code, reqDuration := request(url)
		// if _, ok := distribution[code]; !ok {
		// 	distribution[code] = code
		// }
		// distribution[code]++
		// duration = append(duration, int(reqDuration))
		log.Println(code, reqDuration)
	}
}

func distrubuteRequests(requests, concurrency int) map[int]int {
	distReqs := make(map[int]int)
	distribution := requests / concurrency
	rest := requests % concurrency
	for i := 0; i < concurrency; i++ {
		distReqs[i] = distribution
	}
	for i := 0; i < rest; i++ {
		distReqs[i]++
	}
	return distReqs
}
