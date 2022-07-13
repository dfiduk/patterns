package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"

	"github.com/pierrec/lz4/v4"
)

func listen(proto string, addr string) (c net.Conn, err error) {
	ln, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatalln(err)
		return
	}
	c, err = ln.Accept()
	if err != nil {
		log.Fatalln(err)
		return
	}
	return
}

func ddStart(w3 *io.WriteCloser, bs string, dst string) (c *exec.Cmd, err error) {
	c3 := exec.Command("c:/mind/agent/agent/dd.exe", fmt.Sprintf("of=%s", dst), fmt.Sprintf("bs=%s", bs))
	c3.Env = os.Environ()

	*w3, err = c3.StdinPipe()
	if err != nil {
		return
	}
	e3, err := c3.StderrPipe()
	if err != nil {
		err = fmt.Errorf("stderr: %s, err: %s", e3, err.Error())
		return
	}

	c3.Start()

	return

}

func main() {

	var ddStdin io.WriteCloser
	var zr io.Reader
	var err error

	proto := "tcp"
	addr := "185.151.147.78:9000"

	dst := `\\.\Volume{7CABE3D9-0000-0000-0000-501F00000000}\test.exe`
	// dst := `\\.\Volume{bdd7570a-0000-0000-0000-501f00000000}\test.txt`
	bs := "16M"

	compressed := true

	c, err := listen(proto, addr)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	_, err = ddStart(&ddStdin, bs, dst)
	if err != nil {
		panic(err)
	}

	if compressed {
		zr = lz4.NewReader(c)
	} else {
		zr = c
	}
	if _, err = io.Copy(ddStdin, zr); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

}
