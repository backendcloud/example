package main

import "net"
import "time"

func main() {
	c, err := net.Dial("unix", "/tmp/echo.sock")
	if err != nil {
		panic(err)
	}
	for {
		_, err := c.Write([]byte("hi\n"))
		if err != nil {
			println(err)
		}
		time.Sleep(1e9)
	}
}
