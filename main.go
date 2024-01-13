package main

import (
	"fmt"
	"os"
	"projecttimer/api"
	"projecttimer/config"
	"projecttimer/db"
	"projecttimer/desktop"
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
	fmt.Println(currentDir)
	wg.Add(1)
	go func() {
		defer wg.Done()
		api.Start(conf)
	}()
	if conf.App.Platform == "windows" {
		//运行交互层
		desktop.LauncherFWApp(currentDir)
	} else {
		wg.Wait()
	}
}
