package codec

import "io"

type Header struct {
	ServiceMethod string // 服务名和方法名，通常与结构体和方法相映射
	Seq           uint64 // 请求的序号，也可以认为是某个请求的 ID，用来区分不同的请求
	Error         string // 错误信息，客户端置为空
}

// Codec 对消息体进行编解码
type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

// NewCodecFunc Codec 的构造函数
type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json"
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
