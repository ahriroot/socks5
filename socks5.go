package socks5

import (
	"fmt"
	"log"
	"net"
)

type Server interface {
	Run() error
}

type SOCKS5Server struct {
	IP   string
	Port int
}

func (s *SOCKS5Server) Run() error {
	address := fmt.Sprintf("%s:%d", s.IP, s.Port)

	// 监听
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	for {
		// 接收请求
		conn, err := lis.Accept()
		if err != nil {
			log.Printf("connection failure from %s: %s", conn.RemoteAddr(), err)
			continue
		}

		go func() {
			defer conn.Close()
			err := handleConnection(conn)
			if err != nil {
				log.Printf("handle connection failure from %s: %s", conn.RemoteAddr(), err)
			}
		}()
	}
}

func handleConnection(conn net.Conn) error {
	// 协商
	if err := handleAuth(conn); err != nil {
		return nil
	}

	// 请求
	// 转发
	return nil
}

// @title		handleAuth
// @description	协商
// @auth		作者	2022/05/31 10:30
// @param		conn	net.Conn	连接
// @return		err		error		错误信息
func handleAuth(conn net.Conn) error {
	return nil
}
