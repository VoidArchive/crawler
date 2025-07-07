package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := args[0]

	const maxConcurrency = 3
	cfg, err := configure(baseURL, maxConcurrency)
	if err != nil {
		fmt.Printf("Error configure :%v\n", err)
		return
	}
	fmt.Printf("starting crawl of: %s\n", baseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(baseURL)
	cfg.wg.Wait()

	for normalizeURL, count := range cfg.pages {
		fmt.Printf("%d - %s\n", count, normalizeURL)
	}
}
