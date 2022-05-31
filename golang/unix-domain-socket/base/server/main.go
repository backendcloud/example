package main

import (
	"log"
	"net"
	"os"
)

const SockAddr = "/tmp/echo.sock"

//func echoServer(c net.Conn) {
//	log.Printf("Client connected [%s]", c.RemoteAddr().Network())
//	io.Copy(c, c)
//	c.Close()
//}

func echoServer(c net.Conn) {
	for {
		buf := make([]byte, 512)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}

		data := buf[0:nr]
		println("Server got:", string(data))
		_, err = c.Write(data)
		if err != nil {
			log.Fatal("Write: ", err)
		}
	}
}

func main() {
	if err := os.RemoveAll(SockAddr); err != nil {
		log.Fatal(err)
	}
	l, err := net.Listen("unix", SockAddr)
	if err != nil {
		log.Fatal("listen error:", err)
	}

	for {
		fd, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go echoServer(fd)
	}
}
