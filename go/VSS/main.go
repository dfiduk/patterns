package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/mind-sw/mind-libs/winps"
)

const BASEDIR = `c:/mind`

//go:embed vss.ps1
var binf embed.FS

func main() {
	// (gwmi -List Win32_ShadowCopy -ComputerName 'localhost').Create('\\?\Volume{bdd7570a-0000-0000-0000-100000000000}\', 'ClientAccessible').ShadowId

	filepath, err := writeEmbed("vss.ps1", "", true)
	if err != nil {
		panic(err)
	}
	fmt.Println(filepath)

	// args := `(gwmi -List Win32_ShadowCopy -ComputerName 'localhost').Create('\\?\Volume{bdd7570a-0000-0000-0000-100000000000}\', 'ClientAccessible').ShadowId`
	// args := `c:\mind\vss.ps1`
	res, err := winps.ExecuteRaw(filepath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("shadow copy id: %s\n", res)
}

func writeEmbed(src string, dst string, exec bool) (filepath string, err error) {

	if len(dst) == 0 {
		dst = src
	}

	filepath = fmt.Sprintf("%s/%s", BASEDIR, dst)
	data, err := binf.ReadFile(src)
	if err != nil {
		return
	}
	f, err := os.Create(filepath)
	if err != nil {
		return
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		return
	}

	if exec {
		err = os.Chmod(filepath, 0700)
		if err != nil {
			return
		}
	}
	return
}
