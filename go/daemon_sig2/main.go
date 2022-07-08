package main

import (
	"flag"
	"os"
	"syscall"
	"time"

	daemon "github.com/sevlyar/go-daemon"
	log "github.com/sirupsen/logrus"
)

const ST_FINISHED = 5
const ST_STARTED = 1

const SIG_UPDATECFG = 1
const SIG_QUIT = 2

var (
	signal = flag.String("s", "", `Send signal to the daemon:
  quit — graceful shutdown
  stop — fast shutdown
  reload — reloading the configuration file`)

	statusChan chan int
	signalChan chan int
)

func cleanup() {

	log.Info("Cleanup")

}

func sigHandler(sig os.Signal) error {
	log.Info("Signal handler")
	if sig == syscall.SIGHUP {
		signalChan <- SIG_UPDATECFG
	} else if sig == syscall.SIGQUIT {
		log.Info("Exiting...")
		return daemon.ErrStop
		// signalChan <- SIG_QUIT
	}

	return nil
}

func updateConfig() {
	log.Info("Update config called")
}

func main() {

	statusChan = make(chan int)
	signalChan = make(chan int)

	flag.Parse()

	daemon.AddCommand(daemon.StringFlag(signal, "updatecfg"), syscall.SIGHUP, sigHandler)
	daemon.AddCommand(daemon.StringFlag(signal, "quit"), syscall.SIGQUIT, sigHandler)
	// daemon.AddCommand(daemon.StringFlag(signal, "reload"), syscall.SIGHUP, reloadHandler)

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
		log.Info("main process found: pid: %d", d.Pid)
		if err != nil {
			log.Fatalf("Unable send signal to the daemon: %s", err.Error())
		}
		daemon.SendCommands(d)
		return
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatalln(err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

	log.Println("- - - - - - - - - - - - - - -")
	log.Println("daemon started")

	go worker()

	err = daemon.ServeSignals()
	if err != nil {
		log.Printf("Error: %s", err.Error())
	}

	log.Println("daemon terminated")

	defer cleanup()

}

func worker() {

	var i int
	go initialTransfer()
	for i < 1 {

		select {
		case status := <-statusChan:
			if status == ST_FINISHED {
				i = 1
			}
		case signal := <-signalChan:
			if signal == SIG_UPDATECFG {
				updateConfig()
			}
			// } else if signal == SIG_QUIT {
			// 	log.Info("Exiting...")
			// 	return // update config or abort (call return?..)
			// }
		default:
		}

		time.Sleep(1 * time.Second)

		// check status of initialTransger channel
		// if finished - break, else - check signals and handle it, sleep and next iteration

		// check signal channel
	}

	go rollingUpdate()
	for {

		select {
		case status := <-statusChan:
			if status == ST_FINISHED {
				break
			}
		case signal := <-signalChan:
			if signal == SIG_UPDATECFG {
				updateConfig()
			} // update config or abort (call return?..)
		default:
		}

		// check status of initialTransger channel
		// if finished - break, else - check signals and handle it, sleep and next iteration

		// check signal channel
	}

	return
}

func initialTransfer() {

	statusChan <- ST_STARTED

	log.Info("Starting initial transfer")

	time.Sleep(60 * time.Second)

	log.Info("Finishing initial transfer")

	statusChan <- ST_FINISHED

}

func rollingUpdate() {

	statusChan <- ST_STARTED
	log.Info("Starting rolling update")

	time.Sleep(60 * time.Second)
	log.Info("Finishing rolling update")
	statusChan <- ST_FINISHED
}
