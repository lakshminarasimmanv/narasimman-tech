package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"time"
)

func main() {
	for {
		conn, err := net.Dial("tcp", "192.168.64.22:5432")
		if err != nil {
			fmt.Println("Connection failed, restarting...")
			cmd := exec.Command("/bin/sh", "-c", "sudo service networking restart")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		} else {
			fmt.Println("Connection successful")
			conn.Close()
		}
		time.Sleep(time.Microsecond)
	}
}
