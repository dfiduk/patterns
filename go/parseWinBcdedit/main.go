package main

import (
	"fmt"
	"strings"

	"github.com/mind-sw/mind-libs/winps"
)

func main() {

	type BootRecord struct {
		Identifier       string
		Description      string
		Recoverysequence string
		Resumeobject     string
		OSdevice         string
	}

	// var bootRecords []BootRecord

	res, err := winps.ExecuteRaw("bcdedit.exe /enum")
	if err != nil {
		panic(err)
	}
	// fmt.Println(res)

	var bootRecord BootRecord
	var mindRecord BootRecord
	var found bool
	data := string(res)
	for _, line := range strings.Split(data, "\r\n") {
		if strings.HasPrefix(line, "Windows Boot Loader") {
			if bootRecord.Description == "MIND" {
				mindRecord = bootRecord
				found = true
				break
			}
			// bootRecords = append(bootRecords, bootRecord)
			bootRecord = BootRecord{}
		}
		arr := strings.Split(line, " ")
		key := arr[0]
		value := strings.Trim(strings.Join(arr[1:], " "), " ")

		if key == "description" {
			bootRecord.Description = value
		} else if key == "identifier" {
			bootRecord.Identifier = value
		} else if key == "recoverysequence" {
			bootRecord.Recoverysequence = value
		} else if key == "resumeobject" {
			bootRecord.Resumeobject = value
		} else if key == "osdevice" {
			bootRecord.OSdevice = value
		}

	}

	if found {
		cmd := fmt.Sprintf(`bcdedit.exe /delete %s`, mindRecord.Resumeobject)
		winps.ExecuteRaw(cmd)
		if err != nil {
			panic(err)
		}
	} // else nothing to do
}
