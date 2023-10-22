package api

const (
	WS_TYPE_LOG          = 1  //日志
	WS_TYPE_SERIAL       = 2  //串口
	WS_TYPE_SERIAL_START = 3  //串口启动
	WS_TYPE_SERIAL_STOP  = 4  //串口停止
	WS_TYPE_SERIAL_ERROR = 5  //串口错误
	WS_TYPE_SERIAL_READ  = 6  //串口读
	WS_TYPE_SERIAL_WRITE = 7  //串口写
	WS_TYPE_SENSOR       = 8  //传感器
	WS_TYPE_SENSOR_WSD   = 9  //传感器 温湿度
	WS_TYPE_SENSOR_ZGF   = 10 //传感器 遮光阀
)

type Msg[T interface{}] struct {
	Type int8 `json:"type"`
	Data T    `json:"data"`
}

func Pack(aType int8, data interface{}) Msg[interface{}] {
	return Msg[interface{}]{Type: aType,
		Data: data}
}
func PackStr(aType int8, str string) Msg[interface{}] {
	return Msg[interface{}]{Type: aType,
		Data: str}
}
