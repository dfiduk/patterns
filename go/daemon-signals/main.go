package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"syscall"
	"time"

	daemon "github.com/sevlyar/go-daemon"
)

var (
	signal = flag.String("s", "", `Send signal to the daemon:
  quit — graceful shutdown
  stop — fast shutdown
  reload — reloading the configuration file`)
)

func main() {

	fmt.Println("VSDEBUG: main: 1")
	flag.Parse()
	daemon.AddCommand(daemon.StringFlag(signal, "quit"), syscall.SIGQUIT, termHandler)
	daemon.AddCommand(daemon.StringFlag(signal, "stop"), syscall.SIGTERM, termHandler)
	daemon.AddCommand(daemon.StringFlag(signal, "reload"), syscall.SIGHUP, reloadHandler)

	cntxt := &daemon.Context{
		PidFileName: "sample.pid",
		PidFilePerm: 0644,
		LogFileName: "sample.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"[go-daemon sample]"},
	}

	if len(daemon.ActiveFlags()) > 0 {
		d, err := cntxt.Search()
		if err != nil {
			log.Fatalf("Unable send signal to the daemon: %s", err.Error())
		}
		daemon.SendCommands(d)
		return
	}

	fmt.Println("VSDEBUG: main: 2")

	d, err := cntxt.Reborn()
	fmt.Println("VSDEBUG: main: 3")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("VSDEBUG: main: 4")
	if d != nil {
		return
	}
	fmt.Println("VSDEBUG: main: 5")
	defer cntxt.Release()

	fmt.Println("VSDEBUG: main: 6")

	log.Println("- - - - - - - - - - - - - - -")
	log.Println("daemon started")

	go worker()

	err = daemon.ServeSignals()
	if err != nil {
		log.Printf("Error: %s", err.Error())
	}

	log.Println("daemon terminated")
}

var (
	stop = make(chan bool)
	done = make(chan bool)
)

func cleanup() {
	log.Println("CALLBACK called")
}

func worker() {

	defer cleanup()

	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println("anon goroutine cycle")
		}
	}()
LOOP:
	for {
		log.Println("VSDEBUG: worker is working now")
		time.Sleep(time.Second) // this is work to be done by worker.
		select {
		case <-stop:
			log.Println("VSDEBUG: worker catched stop")
			break LOOP
		default:
		}
	}
	done <- true
}

func termHandler(sig os.Signal) error {
	log.Println("terminating...")
	stop <- true
	if sig == syscall.SIGQUIT {
		res := <-done
		log.Printf("VSDEBUG: done is %t\n", res)
	}
	log.Printf("VSDEBUG: exiting immediately\n")
	return daemon.ErrStop
}

func reloadHandler(sig os.Signal) error {
	log.Println("configuration reloaded")
	return nil
}
