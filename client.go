package rigctld

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"time"
)

const localhostAddr = "127.0.0.1"
const rigctldPort = 4532
const defaultReadDeadline = 100 * time.Millisecond

type Client struct {
	ServerAddr   net.Addr
	conn         *net.TCPConn
	readDeadline time.Duration
}

// Connect creates a TCP connection to communicate with rigctld on the default address and port.
func Connect() (Client, error) {
	return ConnectTo(net.ParseIP(localhostAddr), rigctldPort)
}

// ConnectTo creates a TCP connection to communicate with rigctld on the given address and port.
func ConnectTo(ipAddr net.IP, port uint) (Client, error) {
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%v:%d", ipAddr, port))
	if err != nil {
		return Client{}, err
	}
	var conn *net.TCPConn
	conn, err = net.DialTCP(addr.Network(), nil, addr)
	if err != nil {
		return Client{}, err
	}
	if conn == nil {
		return Client{}, errors.New("rigctld tcp connection not opened")
	}
	return Client{addr, conn, defaultReadDeadline}, nil
}

func (s *Client) SetReadDeadline(d time.Duration) {
	s.readDeadline = d
}

func (s *Client) writeRead(send string) (string, error) {
	_, err := s.conn.Write([]byte(send))
	if err != nil {
		return "", err
	}

	reader := bufio.NewReader(s.conn)
	var resp string
	for {
		err := s.conn.SetReadDeadline(time.Now().Add(s.readDeadline))
		if err != nil {
			return "", err
		}
		line, err := reader.ReadString('\n')
		if err != nil {
			var netErr net.Error
			if errors.As(err, &netErr) && netErr.Timeout() {
				break
			}
			return "", err
		}
		resp += line
	}
	return resp, nil
}
