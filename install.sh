#!/bin/bash
echo "[*] Installing dependencies for YokaiHunter..."

go mod init github.com/jojo/yokaihunter
go get github.com/jaeles-project/gospider
go get github.com/projectdiscovery/katana/pkg/runner
go get github.com/schollz/progressbar/v3

echo "[*] Make sure dalfox is installed: https://github.com/hahwul/dalfox"
echo "[*] Building..."
go build -o yokaihunter main.go
echo "[*] Done. Run with ./yokaihunter -u https://target.com"