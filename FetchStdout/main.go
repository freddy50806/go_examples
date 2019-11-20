package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func ListenStdout(stdRC io.ReadCloser) {
	for {
		buf := make([]byte, 500)
		if n, err := stdRC.Read(buf); err != nil {
			return
		} else {
			fmt.Println("n:", n)
			fmt.Println("str:", string(buf))
		}
	}
}
func main() {
	cmd := exec.Command("sh", "keep_stdout.sh")
	stdRC, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()
	go ListenStdout(stdRC)
	cmd.Wait()
}
