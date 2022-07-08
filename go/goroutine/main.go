package main

import (
	"os"
	"sync"

	"github.com/mind-sw/mind-libs/executor"

	"github.com/sirupsen/logrus"
)

const BASEDIR = "/tmp/mind"

var env []string
var execLib executor.Executor
var log *logrus.Logger

func main() {

	log := logrus.New()

	execLib.Init(BASEDIR, "", 30, "", log, true)
	go execLib.Callbacker()

	var wg sync.WaitGroup
	wg.Add(1)
	go action(&wg)
	wg.Wait()

}

func cleanup() {
	defer os.RemoveAll(BASEDIR)
}

func action(wg *sync.WaitGroup) {

	logrus.Infof("Host discovery finished successfully")
	wg.Done()

}
