package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func copier(dst io.Writer, src io.Reader) {
	for {
		_, err := io.Copy(dst, src)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	// _, w := io.Pipe()

	c1 := exec.Command("/home/vsorokin/prog/go/pipereader/reader")
	c1.Env = os.Environ()
	c2 := exec.Command("/home/vsorokin/prog/go/pipereader2/reader")
	c2.Env = os.Environ()

	w1, err := c1.StdinPipe()
	if err != nil {
		panic(err)
	}
	r1, err := c1.StdoutPipe()
	if err != nil {
		panic(err)
	}

	w2, err := c2.StdinPipe()
	if err != nil {
		panic(err)
	}

	err = c1.Start()
	if err != nil {
		panic(err)
	}

	err = c2.Start()
	if err != nil {
		panic(err)
	}

	go copier(w2, r1)

	for i := 0; i < 100; i++ {
		_, err = w1.Write([]byte(fmt.Sprintf("Number: %d\r\n", i)))
		if err != nil {
			panic(err)
		}
	}

	// _, err = w1.Write([]byte("TEST\n"))
	// if err != nil {
	// 	panic(err)
	// }

	// _, err = w1.Write([]byte("TEST2\n"))
	// if err != nil {
	// 	panic(err)
	// }

	err = c1.Wait()
	if err != nil {
		panic(err)
	}

	err = c2.Wait()
	if err != nil {
		panic(err)
	}
}
