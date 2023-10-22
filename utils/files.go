package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
)

// FileInfo 文件信息结构体
type FileInfo struct {
	Name string
	Size int64
}

func DirExits(dir string) {
	// 检查文件夹是否存在
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		// 文件夹不存在，创建文件夹
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			fmt.Println("无法创建文件夹：", err)
			return
		}
		fmt.Println("文件夹已成功创建：", dir)
	} else {
		// 文件夹已存在
		fmt.Println("文件夹已存在：", dir)
	}
}

// GetCSVStore 递归计算文件夹的总大小
func GetCSVStore(folderPath string) (int64, []FileInfo, error) {
	var totalSize int64 = 0
	var csvFiles []FileInfo = nil

	// 遍历文件夹中的文件
	fileInfos, err := ioutil.ReadDir(folderPath)
	if err != nil {
		fmt.Println("无法读取文件夹：", err)
		return totalSize, csvFiles, err
	}

	for _, fileInfo := range fileInfos {
		filePath := filepath.Join(folderPath, fileInfo.Name())
		fmt.Println(filePath)
		if !fileInfo.IsDir() && strings.HasSuffix(strings.ToLower(fileInfo.Name()), ".csv") {
			// 累加 CSV 文件的大小
			totalSize += fileInfo.Size()
			csvFiles = append(csvFiles, FileInfo{fileInfo.Name(), fileInfo.Size()})
		}
	}
	return totalSize, csvFiles, nil
}

func StoreSizeConvert(count int64) string {
	var sizeStr string
	switch {
	case count >= PB:
		sizeStr = fmt.Sprintf("%.2f PB", float64(count)/PB)
	case count >= TB:
		sizeStr = fmt.Sprintf("%.2f TB", float64(count)/TB)
	case count >= GB:
		sizeStr = fmt.Sprintf("%.2f GB", float64(count)/GB)
	case count >= MB:
		sizeStr = fmt.Sprintf("%.2f MB", float64(count)/MB)
	case count >= KB:
		sizeStr = fmt.Sprintf("%.2f KB", float64(count)/KB)
	default:
		sizeStr = fmt.Sprintf("%.2f KB", float64(count)/KB)
	}
	return sizeStr
}
