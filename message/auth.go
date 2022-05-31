package message

import (
	"errors"
	"io"
)

const SOCKS5VERSION = 0x05

type Method = byte

type ClientAuthMessage struct {
	Version  byte
	NMethods byte
	Methods  []Method
}

func NewClientAuthMessage(conn io.Reader) (*ClientAuthMessage, error) {
	// 读取 版本 和 协议
	buffer := make([]byte, 2)
	_, err := io.ReadFull(conn, buffer)
	if err != nil {
		return nil, err
	}

	// 验证版本号
	if buffer[0] != SOCKS5VERSION {
		return nil, errors.New("protocal version not supported")
	}

	// 读取 协议方法
	nmethods := buffer[1]
	buf := make([]byte, nmethods)
	_, err = io.ReadFull(conn, buf)
	if err != nil {
		return nil, err
	}

	return &ClientAuthMessage{
		Version:  SOCKS5VERSION,
		NMethods: nmethods,
		Methods:  buf,
	}, nil
}
