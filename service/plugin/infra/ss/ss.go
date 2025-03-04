package ss

import (
	"fmt"
	ss "github.com/shadowsocks/go-shadowsocks2/core"
	"github.com/v2rayA/v2rayA/plugin/infra"
	"net"
	"net/url"
)

// Shadowsocks is a base shadowsocks struct
type Shadowsocks struct {
	dialer infra.Dialer
	addr   string
	pass   string
	method string
	cipher ss.Cipher
}

// NewShadowsocks returns a shadowsocks infra.
func NewShadowsocks(s string, d infra.Dialer) (*Shadowsocks, error) {
	u, err := url.Parse(s)
	if err != nil {
		return nil, newError("[shadowsocks]").Base(err)
	}

	method := u.User.Username()
	pass, _ := u.User.Password()

	if method == "chacha20-poly1305" {
		method = "chacha20-ietf-poly1305"
	}
	cipher, err := ss.PickCipher(method, nil, pass)
	if err != nil {
		return nil, newError("[shadowsocks]").Base(err)
	}

	t := &Shadowsocks{
		dialer: d,
		addr:   u.Host,
		pass:   pass,
		method: method,
		cipher: cipher,
	}

	return t, nil
}

// Addr returns forwarder's address.
func (s *Shadowsocks) Addr() string {
	if s.addr == "" {
		return s.dialer.Addr()
	}
	return s.addr
}

// Dial connects to the address addr on the network net via the infra.
func (s *Shadowsocks) Dial(network, addr string) (net.Conn, error) {
	return s.dial(network, addr)
}

func (s *Shadowsocks) dial(network, addr string) (net.Conn, error) {
	rc, err := s.dialer.Dial("tcp", s.addr)
	if err != nil {
		return nil, newError(fmt.Sprintf("[shadowsocks]: dial to %s", s.addr)).Base(err)
	}

	return s.cipher.StreamConn(rc), nil
}

// DialUDP connects to the given address via the infra.
func (s *Shadowsocks) DialUDP(network, addr string) (net.PacketConn, net.Addr, error) {
	rc, raddr, err := s.dialer.DialUDP("udp", addr)
	if err != nil {
		return nil, nil, newError(fmt.Sprintf("[shadowsocks]: dial to %s", s.addr)).Base(err)
	}
	return s.cipher.PacketConn(rc), raddr, nil
}
