package stream

import (
	"bufio"
	"bytes"
	"io"
)

type StringStream struct {
	read_ *bufio.Reader
}

func NewStringStream(read io.Reader) *StringStream {
	var stringStream StringStream
	stringStream.read_ = bufio.NewReader(read)
	return &stringStream
}
func (this *StringStream) ReadLine() ([]byte, error) {
	buffer := bytes.Buffer{}
	for {
		data, is, err := this.read_.ReadLine()
		if err != nil {
			return data, err
		}

		if is {
			if len(data) > 0 {
				buffer.Write(data)
			}
		} else {
			buffer.Write(data)
			return buffer.Bytes(), nil
		}
	}
	return nil, nil
}
