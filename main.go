package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/jlaffaye/ftp"
)

const (
	FTP_HOST        = "localhost:21"
	DEFAULT_TIMEOUT = 5 * time.Second
	RECONNECT_DELAY = 5 * time.Second
	MAX_JITTER      = 5
	USERNAME        = "anonymous"
	PASSWORD        = ""
)

func GetLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP
}

func addJitter() {
	jitter := rand.Intn(MAX_JITTER)
	time.Sleep(RECONNECT_DELAY + time.Duration(jitter)*time.Second)
}

func connectFTP() *ftp.ServerConn {
	c, err := ftp.Dial(FTP_HOST, ftp.DialWithTimeout(DEFAULT_TIMEOUT))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login(USERNAME, PASSWORD)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

func sendData() {
	iterations := 1

	while := true
	for while {
		now := time.Now()
		addJitter()

		c := connectFTP()

		data := bytes.NewBufferString(fmt.Sprintf("Hello, World! Iteration: %d", iterations))

		err := c.Append(fmt.Sprintf("%s.txt", GetLocalIP().String()), data)
		if err != nil {
			log.Println(err)
			c.Quit()
			while = false
		}

		log.Printf("Iteration: %d, Time: %s\n", iterations, time.Since(now))
		iterations++
		c.Quit()
	}
}

func main() {
	now := time.Now()
	time.Sleep(3 * time.Second)

	if time.Since(now) < 3*time.Second {
		return
	}

	sendData()
}
