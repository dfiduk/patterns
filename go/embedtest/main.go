package main

import "embed"

func main() {

	var f embed.FS
	data, err := f.ReadFile("discovery.py")
	if err != nil {
		panic(err)
	}

	print(string(data))
}
