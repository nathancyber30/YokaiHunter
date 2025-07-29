package katanarunner

import (
    "context"
    "fmt"
    "github.com/projectdiscovery/katana/pkg/runner"
)

func RunKatana(target string, resultChan chan<- string) {
    options := runner.Options{
        Target: []string{target},
    }

    r, err := runner.New(&options)
    if err != nil {
        fmt.Println("❌ Katana error:", err)
        return
    }

    _ = r.Run(context.Background(), func(result string) {
        resultChan <- result
    })
}