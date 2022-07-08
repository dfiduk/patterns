package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

// func copier(dst io.Writer, src io.Reader) {
// 	var buf []byte
// 	for {
// 		if src != nil {
// 			n, err := src.Read(buf)
// 			if err != nil {
// 				fmt.Println(err)
// 			} else {
// 				fmt.Printf("n: %d, buf: '%s'\n", n, buf)
// 			}

// 		}
// 		_, err := io.Copy(dst, src)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}
// }

func copier(src io.Reader) {
	buf := new(strings.Builder)
	// check errors
	fmt.Println(buf.String())
	for {
		if src != nil {
			n, err := io.Copy(buf, src)
			// n, err := src.Read(buf)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("n: %d, buf: '%s'\n", n, buf)
			}

		}
		time.Sleep(1 * time.Second)
		// _, err := io.Copy(dst, src)
		// if err != nil {
		// 	panic(err)
		// }
	}
}

func main() {

	c1 := exec.Command("c:/mind/agent/agent/ncat.exe", "-l", "-p 9000")
	c1.Env = os.Environ()
	c2 := exec.Command("c:/mind/agent/agent/dd.exe", `of='\\.\Volume{bdd7570a-0000-0000-0000-501f00000000}\test.txt`, `bs=16M`, `--progress`)
	c2.Env = os.Environ()

	_, err := c1.StdinPipe()
	if err != nil {
		panic(err)
	}
	r1, err := c1.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stderr, err := c1.StderrPipe()
	if err != nil {
		fmt.Println(stderr)
		panic(err)
	}

	// w2, err := c2.StdinPipe()
	// if err != nil {
	// 	panic(err)
	// }

	err = c1.Start()
	if err != nil {
		panic(err)
	}

	// err = c2.Start()
	// if err != nil {
	// 	panic(err)
	// }

	// go copier(w2, r1)
	go copier(r1)

	// for i := 0; i < 100; i++ {
	// 	_, err = w1.Write([]byte(fmt.Sprintf("Number: %d\r\n", i)))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

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

	// err = c2.Wait()
	// if err != nil {
	// 	panic(err)
	// }
}
