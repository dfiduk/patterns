package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"

	"github.com/pierrec/lz4/v4"
)

func connect(proto string, addr string) (c net.Conn, err error) {
	if c, err = net.Dial(proto, addr); err != nil {
		return
	}
	return

}

func ddStart(r3 *io.ReadCloser, bs string, src string) (c *exec.Cmd, err error) {
	cmdName := "dd"
	// cmdName := "c:/mind/agent/agent/dd.exe"
	c3 := exec.Command(cmdName, fmt.Sprintf("if=%s", src), fmt.Sprintf("bs=%s", bs))
	c3.Env = os.Environ()

	*r3, err = c3.StdoutPipe()
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

	var ddStdout io.ReadCloser
	// var zw io.Writer
	var err error

	proto := "tcp"
	addr := "localhost:9000"
	// compressed := true

	// dst := `\\.\Volume{7CABE3D9-0000-0000-0000-501F00000000}\test.exe`
	// dst := `\\.\Volume{bdd7570a-0000-0000-0000-501f00000000}\test.txt`
	// src := "/opt/mind/hostutils/mind_discovery"
	src := "/tmp/test.data"
	bs := "16M"

	c, err := connect(proto, addr)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	_, err = ddStart(&ddStdout, bs, src)
	if err != nil {
		panic(err)
	}

	// if compressed {
	zw := lz4.NewWriter(c)
	defer zw.Close()
	// } else {
	// 	zw := c
	// }
	if _, err = io.Copy(zw, ddStdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

}
