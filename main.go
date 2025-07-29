package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "strings"
    "sync"

    "github.com/jojo/yokaihunter/internal/gospiderrunner"
    "github.com/jojo/yokaihunter/internal/katanarunner"
    "github.com/jojo/yokaihunter/internal/dalfoxrunner"

    "github.com/schollz/progressbar/v3"
)

func main() {
    var target string
    var tool string
    var output string

    flag.StringVar(&target, "u", "", "Target URL (eg: https://example.com)")
    flag.StringVar(&tool, "tool", "all", "Tool to run: gospider, katana, all")
    flag.StringVar(&output, "o", "output/yokaihunter.txt", "Output file path")
    flag.Parse()

    if target == "" {
        log.Fatal("Target URL required. Use -u https://example.com")
    }

    var wg sync.WaitGroup
    resultChan := make(chan string, 1000)

    bar := progressbar.Default(-1, "Scanning...")

    if tool == "gospider" || tool == "all" {
        wg.Add(1)
        go func() {
            defer wg.Done()
            gospiderrunner.RunGospider(target, resultChan)
        }()
    }

    if tool == "katana" || tool == "all" {
        wg.Add(1)
        go func() {
            defer wg.Done()
            katanarunner.RunKatana(target, resultChan)
        }()
    }

    go func() {
        wg.Wait()
        close(resultChan)
    }()

    unique := make(map[string]bool)
    var finalResult []string
    for url := range resultChan {
        bar.Add(1)
        if strings.Contains(url, "?") && !unique[url] {
            unique[url] = true
            finalResult = append(finalResult, url)
            fmt.Printf("\033[32m[FOUND] %s\033[0m\n", url)
            dalfoxrunner.RunDalfox(url)
        }
    }

    bar.Finish()

    err := os.WriteFile(output, []byte(strings.Join(finalResult, "\n")), 0644)
    if err != nil {
        log.Fatal("Error writing output:", err)
    }
    fmt.Println("✅ Output saved to:", output)
}