# playtime

Personal retro games library + [EmulatorJS](https://emulatorjs.org/).

## Building

Go 1.20+ and npm 7+ required.

```bash
./install.sh
go build -a -o app
```

## Configuration

Commandline arguments available:

```
$ ./app -help
Usage of ./app:
  -db-path string
        db path (default "data/bolt.db")
  -debug-emulator
        debug emulator (extended browser console output)
  -debug-netplay
        debug netplay (extended browser console output)
  -debug-templates
        debug page templates (do not cache)
  -listen string
        address and port to listen (default ":3000")
  -turn-server-password string
        TURN/STUN/ICE server password (if required)
  -turn-server-url string
        TURN/STUN/ICE server host, required for netplay (example: turn:turn.example.com)
  -turn-server-user string
        TURN/STUN/ICE server user name (if required)
  -uploads-path string
        uploads path (default "uploads")
  -verbose
        show debug output
```

## Netplay

TURN server is required for netplay. It can be obtained from:

* [This list](https://gist.github.com/sagivo/3a4b2f2c7ac6e1b5267c2f1f59ac6c6b)
* [Open Relay](https://www.metered.ca/tools/openrelay/)
* Hosted, for example [coturn](https://github.com/coturn/coturn)

## Uses

* [labstack/echo](https://github.com/labstack/echo) - MIT
* [flosch/pongo2](https://github.com/flosch/pongo2) - MIT
* [sirupsen/logrus](https://github.com/sirupsen/logrus) - MIT
* [timshannon/bolthold](https://github.com/timshannon/bolthold) - MIT
* [google/uuid](https://github.com/google/uuid) - BSD-3-Clause
* [twbs/bootstrap](https://github.com/twbs/bootstrap) - MIT
* [twbs/icons](https://github.com/twbs/icons) - MIT
* [sumimakito/Awesome-qr.js](https://github.com/sumimakito/Awesome-qr.js) - Apache-2.0
* [EmulatorJS](https://github.com/EmulatorJS/EmulatorJS) - GPL-3.0 (not included)
* BIOS files from respective vendors (not included)
