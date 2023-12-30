@echo off
chcp 65001 > nul
set GOOS=windows
set GOARCH=amd64
go build -ldflags "-H=windowsgui -s -w" -o "pt.exe" main.go
upx "pt.exe"
echo Build completed successfully.
