package roncache

// 缓存值的抽象与封装

type ByteView struct {
	b []byte // 存储真实的缓存值。byte可以支持任意的数据类型
}

func (v ByteView) Len() int {
	return len(v.b)
}

func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
