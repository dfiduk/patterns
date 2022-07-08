package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func splitter(s string) (result []string) {

	fmt.Printf("source string is '%s'\n", s)

	var opened bool
	var oneOf string
	var separator rune
	for _, char := range s {

		if separator == 0 && (char == '\'' || char == ' ') || char == separator {

			// if char == '\'' || char == ' ' {
			fmt.Println("' symbol was found")
			if !opened {
				opened = true
				separator = char
				fmt.Printf("is opened, separator is '%s'\n", string(separator))
				oneOf = ""
				continue
			} else {
				fmt.Printf("is closed, target string is '%s'\n", oneOf)
				opened = false
				separator = 0
				result = append(result, oneOf)
				continue
			}
		}
		if opened {
			// fmt.Println("adding")
			oneOf += string(char)
		}

	}
	// fmt.Println(result)
	return result
}

// func splitter(s string) (result []string) {

// 	var opened bool
// 	var oneOf string
// 	for _, char := range s {

// 		if char == '\'' {
// 			fmt.Println("' symbol was found")
// 			if !opened {
// 				fmt.Println("is opened")
// 				opened = true
// 				oneOf = ""
// 				continue
// 			} else {
// 				fmt.Printf("is closed, target string is '%s'\n", oneOf)
// 				opened = false
// 				result = append(result, oneOf)
// 				continue
// 			}
// 		}
// 		if opened {
// 			fmt.Println("adding")
// 			oneOf += string(char)
// 		}

// 	}
// 	// fmt.Println(result)
// 	return result
// }

func reverseEnum(s string) (result string) {

	// fmt.Println("TEST")

	var i int
	for i = len(s); i > 0; i-- {
		// fmt.Println(i)
		fmt.Println(string(s[i-1]))
	}
	return
}

// func CmdCall(command string, env []string) (result string, err error) {

// 	var stdout bytes.Buffer
// 	var stderr bytes.Buffer

// 	// fmt.Println(reverseEnum("12345"))
// 	fmt.Printf("%#v\n", splitter(command))
// 	return

// 	cArray := strings.Fields(command)
// 	cmd := cArray[0]
// 	args := cArray[1:]

// 	// cmd := "/usr/local/bin/winexe --user=user%c15po82vN //185.151.147.75:445 --uploadfile /opt/mind/hostutils/mind_discovery.exe 'c:/mind//mind_discovery.exe -d -g 9d55c8de-8d89-467c-9d15-e66fb203dacf -u https://dev.mindsw.io//api/v1/apitask/update -t 120'"

// 	// cmd = "/usr/local/bin/winexe"
// 	// args := []string{
// 	// 	"--user=user%c15po82vN",
// 	// 	"//185.151.147.75:445",
// 	// 	"--uploadfile=/opt/mind/hostutils/mind_discovery.exe",
// 	// 	"c:/mind//mind_discovery.exe -g 9d55c8de-8d89-467c-9d15-e66fb203dacf -u https://dev.mindsw.io//api/v1/apitask/update -t 120",
// 	// 	// "\"powershell.exe Start-Sleep 5\"",
// 	// }
// 	fmt.Printf("VSDEBUG: arg: '%#v'\n", args)
// 	c := exec.Command(cmd, args...)
// 	c.Env = os.Environ()
// 	c.Env = append(c.Env, env...)

// 	c.Stdout = &stdout

// 	c.Stderr = &stderr
// 	err = c.Run()
// 	if err != nil {
// 		err = fmt.Errorf("err: %s; cmd: '%s'; stdout: %s; stderr: %s", err, c.String(), stdout.String(), stderr.String())
// 		return stdout.String(), err
// 	}

// 	return stdout.String(), nil
// }

// // cmdCallInExecutor - execute some commandline command with specified additional environment variables
// // Returns stdout and possible error (includes stderr end error code)
// func cmdCallInExecutor(cmd string, args []string, env []string) (result string, err error) {

// 	// var stdout bytes.Buffer
// 	// var stderr bytes.Buffer

// 	c := exec.Command(cmd, args...)
// 	c.Env = os.Environ()
// 	c.Env = append(c.Env, env...)

// 	// stdout, err := c.StdoutPipe()
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// defer stdout.Close()
// 	stdin, err := c.StdinPipe()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer stdin.Close()
// 	// stderr, err := c.StderrPipe()
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// defer stderr.Close()

// 	// c.Stdout = &stdout

// 	// c.Stderr = &stderr

// 	fmt.Printf("VSDEBUG: command: '%s'\n", c.String())

// 	// out, _ := ioutil.ReadAll(stdout)
// 	// stderrData, _ := ioutil.ReadAll(stdout)
// 	// fmt.Printf("%s", b)

// 	err = c.Start()
// 	// if err != nil {
// 	// 	return stdout.String(), err
// 	// }
// 	if err != nil {
// 		// err = fmt.Errorf("1: err: %s; cmd: '%s'; stdout: %s; stderr: %s", err, c.String(), stdout.String(), stderr.String())
// 		// return stdout.String(), err
// 		// err = fmt.Errorf("1: err: %s; cmd: '%s'; stdout: %s; stderr: %s", err, c.String(), string(out), string(stderrData))
// 		// return string(out), err
// 		return "", err
// 	}

// 	// time.Sleep(time.Second)

// 	err = c.Wait()
// 	if err != nil {
// 		// err = fmt.Errorf("2: err: %s; cmd: '%s'; stdout: %s; stderr: %s", err, c.String(), stdout.String(), stderr.String())
// 		return "", err
// 	}
// 	// err = c.Run()
// 	// if err != nil {
// 	// 	err = fmt.Errorf("err: %s; cmd: '%s'; stdout: %s; stderr: %s", err.Error(), c.String(), stdout.String(), stderr.String())
// 	// 	return stdout.String(), err
// 	// }

// 	// return string(out), nil
// 	return "", err
// }

func cmdCall() {
	cmnd := exec.Command("/usr/local/bin/winexe", "--debuglevel=2", "--user=user%c15po82vN",
		"//185.151.147.75:445", "\\mind\\mind_discovery.exe > \\mind\\file.log")
	//cmnd.Run() // and wait
	stdout, err := cmnd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer stdout.Close()
	stdin, err := cmnd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer stdin.Close()
	stderr, err := cmnd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer stderr.Close()
	err = cmnd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmnd.Wait()
	if err != nil {
		log.Fatal(err)
	}

	b, _ := ioutil.ReadAll(stdout)
	fmt.Printf("%s", b)
}

// cmdCallInExecutor - execute some commandline command with specified additional environment variables
// Returns stdout and possible error (includes stderr end error code)
func cmdCallInExecutor(cmd string, args []string, env []string) (result string, err error) {

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	c := exec.Command(cmd, args...)
	c.Env = os.Environ()
	c.Env = append(c.Env, env...)

	stdin, err := c.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer stdin.Close()
	c.Stdout = &stdout
	c.Stderr = &stderr

	fmt.Printf("VSDEBUG: command: '%s'\n", c.String())

	err = c.Start()
	if err != nil {
		fmt.Println("11111")
		if len(stderr.String()) == 0 && strings.Contains(strings.ToLower(stdout.String()), "error") {
			stderr = stdout
		}
		err = fmt.Errorf("1: err: %s; cmd: '%s'; stdout: %s; stderr: %s", err, c.String(), stdout.String(), stderr.String())
		return stdout.String(), err
	}

	err = c.Wait()
	if err != nil {
		fmt.Println("21111")
		if len(stderr.String()) == 0 && strings.Contains(strings.ToLower(stdout.String()), "error") {
			stderr = stdout
		}
		err = fmt.Errorf("1: err: %s; cmd: '%s'; stdout: %s; stderr: %s", err, c.String(), stdout.String(), stderr.String())
		return stdout.String(), err
	}

	if strings.Contains(strings.ToLower(stdout.String()), "error") {
		stderr = stdout
		err = fmt.Errorf("1: err: %s; cmd: '%s'; stdout: %s; stderr: %s", "Exit Code: 999", c.String(), stdout.String(), stderr.String())
		return stdout.String(), err
	}

	return stdout.String(), err
}

func main() {

	// cmdCall()
	// os.Exit(0)

	cmd := "/usr/local/bin/winexe"
	args := []string{
		"//185.151.147.75:445",
		"--user=user%c15po82vN",
		"--debuglevel=0",
		"--interactive=0",
		// "--ostype=1",
		// "-N",
		// "--reinstall",
		// "--runas=user%c15po82vN",
		"--uploadfile=/opt/mind/hostutils/mind_discovery.exe",
		// "'powerhell.exe Start-Sleep 20'",
		// "-V",
		// "c:/mind//mind_discovery.exe -g 9d55c8de-8d89-467c-9d15-e66fb203dacf -u https://dev.mindsw.io//api/v1/apitask/update -t 120",
		"powershell.exe Start-Process -ArgumentList '-g 00b8c137-2f05-4f39-92b6-c2cf53cba06c -u https://dev.mindsw.io//api/v1/apitask/update -t 120' -FilePath c:/mind//mind_discovery.exe",
		// `"powershell.exe Start-Process -ArgumentList '-g d2b36e28-a641-4439-8590-7e2854dc749c -u https://dev.mindsw.io//api/v1/apitask/update -t 120' -FilePath c:/mind//mind_discovery.exe"`,
		// "'powershell.exe'",
	}

	// "powershell.exe Start-Process -ArgumentList '-g 00b8c137-2f05-4f39-92b6-c2cf53cba06c -u https://dev.mindsw.io//api/v1/apitask/update -t 120' -FilePath c:/mind//mind_discovery.exe"
	res, err := cmdCallInExecutor(cmd, args, []string{})
	// var stdoutPipe io.ReadCloser
	// CmdCallOutput(cmd, args, []string{}, &stdoutPipe)
	// cmd := "/usr/local/bin/winexe //185.151.147.75:445 --user=user%c15po82vN '--uploadfile=/opt/mind/hostutils/mind_discovery.exe' 'c:/mind//mind_discovery.exe -g 9d55c8de-8d89-467c-9d15-e66fb203dacf -u https://dev.mindsw.io//api/v1/apitask/update -t 120'"
	// res, err := CmdCall(cmd, []string{})
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
