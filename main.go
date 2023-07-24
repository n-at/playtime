package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"playtime/web"
)

var (
	Listen = ":3000"
)

func main() {
	e := web.New()
	log.Fatal(e.Start(Listen))
}

func init() {
	verbosePtr := flag.Bool("verbose", false, "show debug output")
	templatesDebugPtr := flag.Bool("templates-debug", false, "debug page templates (do not cache)")
	listenPtr := flag.String("listen", ":3000", "address and port to listen")
	flag.Parse()

	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetOutput(os.Stdout)
	if *verbosePtr {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	web.TemplatesDebug = *templatesDebugPtr
	Listen = *listenPtr
}
