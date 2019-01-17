package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethercflow/hookfs/hookfs"
	log "github.com/sirupsen/logrus"
)

var (
	addr       = flag.String("addr", ":65534", "The address to bind to")
	original   = flag.String("original", "", "ORIGINAL")
	mountpoint = flag.String("mountpoint", "", "MOUNTPOINT")
	logLevel   = flag.Int("log-level", 0, fmt.Sprintf("log level (%d..%d)", hookfs.LogLevelMin, hookfs.LogLevelMax))
)


func main() {
	sc := make(chan os.Signal, 1)

	signal.Notify(sc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		sig := <-sc
		fmt.Printf("\nGot signal [%v] to exit.\n", sig)

		select {
		case <-sc:
			fmt.Printf("\nGot signal [%v] again to exit.\n", sig)
			os.Exit(1)
		case <-time.After(10 * time.Second):
			fmt.Print("\nWait 10s for closed, force exit\n")
			os.Exit(1)
		}
	}()

	flag.Parse()
	if *original == "" || *mountpoint == "" {
		os.Exit(2)
	}

	hookfs.SetLogLevel(*logLevel)
	fs, err := hookfs.NewHookFs(*original, *mountpoint, &InjuredHook{addr: *addr})
	if err != nil {
		log.Fatal(err)
	}
	if err = fs.Serve(); err != nil {
		log.Fatal(err)
	}
}
