package desktop

import (
	"os"
	"os/exec"
	"projecttimer/utils"
)

func LauncherFWApp(path string) {
	// 指定 Flutter 打包的可执行程序路径
	flutterAppPath := path + "/windows/runner/Release/projecttimerf"

	// 检查可执行程序文件是否存在
	if _, err := os.Stat(flutterAppPath); os.IsNotExist(err) {
		utils.Log.Error("Flutter app executable not found at %s", flutterAppPath)
	}

	// 启动 Flutter 打包的可执行程序
	cmd := exec.Command(flutterAppPath)
	err := cmd.Run()
	if err != nil {
		utils.Log.Fatalf("Failed to start Flutter app: %v", err)
	}

	utils.Log.Println("Flutter app started successfully")
}
