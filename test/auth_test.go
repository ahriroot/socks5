package test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/ahriroot/socks5/message"
)

func TestNewClientAuthMessage(t *testing.T) {
	t.Run("generate message", func(t *testing.T) {
		b := []byte{message.SOCKS5VERSION, 2, 0x00, 0x01}
		r := bytes.NewReader(b)
		msg, err := message.NewClientAuthMessage(r)
		if err != nil {
			t.Fatalf("want error == nil but got %s", err)
		}

		if msg.Version != message.SOCKS5VERSION {
			t.Fatalf("want socks5version but got %d", msg.Version)
		}

		if msg.NMethods != 2 {
			t.Fatalf("want nmethods == 2 but got %d", msg.NMethods)
		}

		if !reflect.DeepEqual(msg.Methods, []byte{0x00, 0x01}) {
			t.Fatalf("want methods: %v but got %v", []byte{0x00, 0x01}, msg.Methods)
		}
	})

	t.Run("methods length is shorter than nmethods", func(t *testing.T) {
		b := []byte{message.SOCKS5VERSION, 2, 0x00}
		r := bytes.NewReader(b)
		_, err := message.NewClientAuthMessage(r)
		if err == nil {
			t.Fatalf("should get error != nil but got nil")
		}
	})
}
