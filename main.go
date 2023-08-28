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
	debugTemplatesPtr := flag.Bool("debug-templates", false, "debug page templates (do not cache)")
	debugEmulatorPtr := flag.Bool("debug-emulator", false, "debug emulator (extended browser console output)")
	debugNetplayPtr := flag.Bool("debug-netplay", false, "debug netplay (extended browser console output)")
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

	if *verbosePtr {
		log.Info("verbose output enabled")
	}
	if *debugTemplatesPtr {
		log.Info("templates debug enabled")
	}
	if *debugEmulatorPtr {
		log.Info("emulator debug enabled")
	}
	if *debugNetplayPtr {
		log.Info("netplay debug enabled")
	}

	storageConfig = &storage.Configuration{
		DatabasePath: *dbPathPtr,
		UploadsPath:  *uploadsPathPtr,
	}
	webConfig = &web.Configuration{
		AssetsRoot:  "assets",
		UploadsRoot: *uploadsPathPtr,
		Listen:      *listenPtr,

		TemplatesDebug:     *debugTemplatesPtr,
		TemplatesRoot:      "templates",
		TemplatesExtension: "twig",

		EmulatorDebug: *debugEmulatorPtr,

		NetplayEnabled:     strings.TrimSpace(*turnServerUrlPtr) != "",
		NetplayDebug:       *debugNetplayPtr,
		TurnServerUrl:      *turnServerUrlPtr,
		TurnServerUser:     *turnServerUser,
		TurnServerPassword: *turnServerPassword,
	}
}
