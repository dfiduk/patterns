package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fd, err := os.Create("/tmp/reader.out")
	if err != nil {
		panic(err)
		// Handle error.
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fd.WriteString(fmt.Sprintf("%s\n", scanner.Text()))
		fmt.Println(scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
		// Handle error.
	}
}
