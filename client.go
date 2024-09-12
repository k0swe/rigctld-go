package rigctld

import (
	"bufio"
	"errors"
	"fmt"
	"net"
)

const localhostAddr = "127.0.0.1"
const rigctldPort = 4532

type Client struct {
	ServerAddr net.Addr
	conn       *net.TCPConn
}

// Frequency is the frequency of the radio's VFO in hertz.
type Frequency int64

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
	return Client{addr, conn}, nil
}

func (s *Client) GetFreq() (Frequency, error) {
	resp, err := s.writeRead("f\n")
	if err != nil {
		return 0, err
	}

	var freq Frequency
	_, err = fmt.Sscanf(resp, "%d\n", &freq)
	if err != nil {
		return 0, err
	}
	return freq, nil
}

func (s *Client) SetFreq(freq Frequency) error {
	resp, err := s.writeRead(fmt.Sprintf("F %d\n", freq))
	if err != nil {
		return err
	}

	var report int
	_, err = fmt.Sscanf(resp, "RPRT %d\n", &report)
	if err != nil {
		return err
	}
	if report != 0 {
		// TODO: Can this be more specific? Probably would have to delve into the rigctld source
		return errors.New(fmt.Sprintf("rigctld error %d", report))
	}
	return nil
}

func (s *Client) writeRead(send string) (string, error) {
	_, err := s.conn.Write([]byte(send))
	if err != nil {
		return "", err
	}
	reader := bufio.NewReader(s.conn)
	resp, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return resp, nil
}