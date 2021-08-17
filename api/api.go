package api

import (
	"bytes"
	"encoding/binary"
	"github.com/pkg/errors"
	"github.com/ulikunitz/xz"
	"io/ioutil"
)

type Query struct {
	DeviceName string
}

type Response struct {
	SocksPort  uint16
	DnsPort    uint16
	DeviceName string
	Debug      bool
}

func MakeMessage(i interface{}) ([]byte, error) {
	buf := &bytes.Buffer{}
	writer, err := xz.NewWriter(buf)
	err = binary.Write(writer, binary.LittleEndian, i)
	if err != nil {
		return nil, errors.WithMessage(err, "write binary error")
	}
	err = writer.Close()
	if err != nil {
		return nil, errors.WithMessage(err, "close writer")
	}

	message, err := ioutil.ReadAll(buf)
	if err != nil {
		return nil, errors.WithMessage(err, "read buf error")
	}

	return message, nil
}

func ParseMessage(message []byte, i interface{}) (interface{}, error) {
	reader, err := xz.NewReader(bytes.NewReader(message))
	if err != nil {
		return nil, errors.WithMessage(err, "read message error")
	}
	err = binary.Read(reader, binary.LittleEndian, i)
	if err != nil {
		return nil, errors.WithMessage(err, "parse binary error")
	}
	return i, nil
}
