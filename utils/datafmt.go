package utils

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

// HexSp16 转换为带空格的16进制
func HexSp16(data []byte) string {
	hexString := ""
	for _, b := range data {
		hexString += fmt.Sprintf("%02X ", b)
	}
	return hexString
}
func ReplaceSp(hexStr string) string {
	return strings.ReplaceAll(hexStr, " ", "")
}
func HexToByte(hexString string) []byte {
	hexString = ReplaceSp(hexString)
	// 将十六进制字符串转换为 byte 数组
	byteData, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Println("转换错误:", err)
		return nil
	}
	return byteData
}
func HexToDecimal(hexString string) int64 {
	decimal, err := strconv.ParseInt(hexString, 16, 64)
	if err != nil {
		fmt.Println("转换错误:", err)
		return 0
	}
	return decimal
}

// TrimNullTerminator 丢弃无效位
func TrimNullTerminator(data []byte) []byte {
	index := bytes.IndexByte(data, 0)
	if index == -1 {
		return data
	}
	return data[:index]
}

// TrimTrailingZeros 丢弃字节数组末尾的多余零
func TrimTrailingZeros(bytes []byte) []byte {
	lastNonZeroIndex := len(bytes) - 1
	for i := len(bytes) - 1; i >= 0; i-- {
		if bytes[i] != 0 {
			break
		}
		lastNonZeroIndex = i - 1
	}
	return bytes[:lastNonZeroIndex+1]
}
