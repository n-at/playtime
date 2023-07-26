package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"playtime/storage"
	"playtime/web"
)

var (
	storageConfig *storage.Configuration
	webConfig     *web.Configuration
)

func main() {
	s, err := storage.New(storageConfig)
	if err != nil {
		log.Fatalf("unable to open storage db: %s", err)
	}

	w := web.New(webConfig, s)
	log.Fatal(w.Start())
}

func init() {
	verbosePtr := flag.Bool("verbose", false, "show debug output")
	templatesDebugPtr := flag.Bool("templates-debug", false, "debug page templates (do not cache)")
	listenPtr := flag.String("listen", ":3000", "address and port to listen")
	storePathPtr := flag.String("storage-path", "data/bolt.db", "storage db path")
	flag.Parse()

	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetOutput(os.Stdout)
	if *verbosePtr {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	storageConfig = &storage.Configuration{
		Path: *storePathPtr,
	}
	webConfig = &web.Configuration{
		AssetsWebRoot:      "/assets",
		AssetsRoot:         "assets",
		TemplatesDebug:     *templatesDebugPtr,
		TemplatesRoot:      "templates",
		TemplatesExtension: "twig",
		Listen:             *listenPtr,
	}
}
