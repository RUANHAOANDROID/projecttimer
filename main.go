package main

import (
	"fmt"
	"os"
	"projecttimer/api"
	"projecttimer/config"
	"projecttimer/db"
	"projecttimer/desktop"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	conf, err := config.Load("config.yml")
	if err != nil {
		panic("config.yml read error")
	}
	db.Create(conf)

	currentDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("failed to get current directory: %v", err))
	}
	debug := os.Getenv("GODEBUG")
	isDebug := strings.Contains(debug, "gctrace=1")
	fmt.Println(currentDir)
	fmt.Println(isDebug)
	wg.Add(1)
	go func() {
		defer wg.Done()
		api.Start(conf)
	}()
	if !isDebug {
		//运行交互层
		desktop.LauncherFWApp(currentDir)
	} else {
		wg.Wait()
	}
}
