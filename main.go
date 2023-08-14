package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"playtime/storage"
	"playtime/web"
	"strings"
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

	err = s.UserEnsureExists()
	if err != nil {
		log.Fatalf("unable to create default user: %s", err)
	}

	server := web.New(webConfig, s)
	log.Fatal(server.Start())
}

func init() {
	verbosePtr := flag.Bool("verbose", false, "show debug output")
	templatesDebugPtr := flag.Bool("templates-debug", false, "debug page templates (do not cache)")
	listenPtr := flag.String("listen", ":3000", "address and port to listen")
	dbPathPtr := flag.String("db-path", "data/bolt.db", "db path")
	uploadsPathPtr := flag.String("uploads-path", "uploads", "uploads path")
	turnServerUrlPtr := flag.String("turn-server-url", "", "TURN/STUN/ICE server host, required for netplay (example: turn:turn.example.com)")
	turnServerUser := flag.String("turn-server-user", "", "TURN/STUN/ICE server user name (if required)")
	turnServerPassword := flag.String("turn-server-password", "", "TURN/STUN/ICE server password (if required)")
	flag.Parse()

	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetOutput(os.Stdout)
	if *verbosePtr {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	storageConfig = &storage.Configuration{
		Path: *dbPathPtr,
	}
	webConfig = &web.Configuration{
		AssetsRoot:  "assets",
		UploadsRoot: *uploadsPathPtr,
		Listen:      *listenPtr,

		TemplatesDebug:     *templatesDebugPtr,
		TemplatesRoot:      "templates",
		TemplatesExtension: "twig",

		NetplayEnabled:     strings.TrimSpace(*turnServerUrlPtr) != "",
		TurnServerUrl:      *turnServerUrlPtr,
		TurnServerUser:     *turnServerUser,
		TurnServerPassword: *turnServerPassword,
	}
}
