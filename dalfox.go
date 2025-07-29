package dalfoxrunner

import (
    "fmt"
    "os/exec"
)

func RunDalfox(url string) {
    fmt.Println("\033[35m[DALFOX] Scanning:", url, "\033[0m")
    cmd := exec.Command("dalfox", "url", url)
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("❌ Dalfox error:", err)
        return
    }
    fmt.Println(string(output))
}