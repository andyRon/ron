package codec

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
)

type GobCodec struct {
	conn io.ReadWriteCloser // 由构建函数传入，通常是通过 TCP 或者 Unix 建立 socket 时得到的链接实例
	buf  *bufio.Writer      // 为了防止阻塞而创建的带缓冲的 Writer
	dec  *gob.Decoder
	enc  *gob.Encoder
}

var _ Codec = (*GobCodec)(nil) // 类型断言，确保GobCodec类型实现了Codec接口。

func NewGobCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &GobCodec{
		conn: conn,
		buf:  buf,
		dec:  gob.NewDecoder(conn),
		enc:  gob.NewEncoder(buf),
	}
}

func (c GobCodec) Close() error {
	return c.conn.Close()
}

func (c GobCodec) ReadHeader(h *Header) error {
	return c.dec.Decode(h)
}

func (c GobCodec) ReadBody(body interface{}) error {
	return c.dec.Decode(body)
}

func (c GobCodec) Write(h *Header, body interface{}) (err error) {
	defer func() {
		_ = c.buf.Flush()
		if err != nil {
			_ = c.Close() // _用于忽略c.Close()的返回值
		}
	}()

	if err := c.enc.Encode(h); err != nil {
		log.Println("rpc codec: gob error encoding header: ", err)
	}
	if err := c.enc.Encode(body); err != nil {
		log.Println("rpc codec: gob error encoding body: ", err)
	}
	return nil
}
