package gospiderrunner

import (
    "fmt"
    "os/exec"
    "strings"
)

func RunGospider(target string, resultChan chan<- string) {
    cmd := exec.Command("gospider", "-s", target, "-d", "2", "-c", "5", "--js", "--robots", "--sitemap")
    stdout, err := cmd.Output()
    if err != nil {
        fmt.Println("❌ Gospider error:", err)
        return
    }

    lines := string(stdout)
    for _, line := range strings.Split(string(lines), "\n") {
        if line != "" {
            resultChan <- line
        }
    }
}