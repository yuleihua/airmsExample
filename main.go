package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"

	log "airman.com/airms/pkg/mslog"

	"airman.com/airmsExample/node/airmsExample"
	"airman.com/airmsExample/node/config"
	"airman.com/airmsExample/node/version"
)

var (
	configFile string
	isPProf    bool
	isVersion  bool

	gitCommit string // commit hash
	buildDate string // build datetime
)

func init() {
	flag.StringVar(&configFile, "c", "../etc/airmsExample.ini", "configure file")
	flag.BoolVar(&isPProf, "p", false, "setting of pprof")
	flag.BoolVar(&isVersion, "v", false, "version information")
}

func StartPProf(address string) {
	log.Info("Starting pprof server", "addr", fmt.Sprintf("http://%s/debug/pprof", address))
	go func() {
		if err := http.ListenAndServe(address, nil); err != nil {
			log.Error("Failure in running pprof server", "err", err)
		}
	}()
}

func main() {
	flag.Parse()

	if isVersion {
		version.Info(os.Args[0], gitCommit, buildDate)
		os.Exit(0)
	}

	// setting config
	if err := config.Setup(configFile); err != nil {
		log.Fatalf("config error:%v", err)
	}

	if isPProf {
		StartPProf("localhost:6060")
	}

	s := airms_example.NewAirmsExampleService(config.GetService().Name, 30*time.Second)
	if err := s.Run(); err != nil {
		log.Fatalf("run service  error:%v", err)
		s.Stop()
	}
	log.Warn(">> service quit")
}
