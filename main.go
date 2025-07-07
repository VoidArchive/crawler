package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

type pageCount struct {
	URL   string
	Count int
}

func sortPages(pages map[string]int) []pageCount {
	var sortedPages []pageCount
	for url, count := range pages {
		sortedPages = append(sortedPages, pageCount{URL: url, Count: count})
	}
	
	sort.Slice(sortedPages, func(i, j int) bool {
		if sortedPages[i].Count == sortedPages[j].Count {
			return sortedPages[i].URL < sortedPages[j].URL
		}
		return sortedPages[i].Count > sortedPages[j].Count
	})
	
	return sortedPages
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf("=============================\n")
	fmt.Printf("  REPORT for %s\n", baseURL)
	fmt.Printf("=============================\n")
	
	sortedPages := sortPages(pages)
	
	for _, page := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", page.Count, page.URL)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("not enough arguments provided")
		fmt.Println("usage: crawler URL maxConcurrency maxPages")
		os.Exit(1)
	}

	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		fmt.Println("usage: crawler URL maxConcurrency maxPages")
		os.Exit(1)
	}

	baseURL := args[0]
	
	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("invalid maxConcurrency: %v\n", err)
		os.Exit(1)
	}

	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("invalid maxPages: %v\n", err)
		os.Exit(1)
	}

	cfg, err := configure(baseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error configure :%v\n", err)
		return
	}
	fmt.Printf("starting crawl of: %s\n", baseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(baseURL)
	cfg.wg.Wait()

	printReport(cfg.pages, baseURL)
}
