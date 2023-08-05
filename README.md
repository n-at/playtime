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
  -listen string
        address and port to listen (default ":3000")
  -templates-debug
        debug page templates (do not cache)
  -uploads-path string
        uploads path (default "uploads")
  -verbose
        show debug output
```

## Uses

* [labstack/echo](https://github.com/labstack/echo) - MIT
* [flosch/pongo2](https://github.com/flosch/pongo2) - MIT
* [sirupsen/logrus](https://github.com/sirupsen/logrus) - MIT
* [timshannon/bolthold](https://github.com/timshannon/bolthold) - MIT
* [google/uuid](https://github.com/google/uuid) - BSD-3-Clause
* [twbs/bootstrap](https://github.com/twbs/bootstrap) - MIT
* [twbs/icons](https://github.com/twbs/icons) - MIT
* [EmulatorJS](https://github.com/EmulatorJS/EmulatorJS) - GPL-3.0 (not included)
* BIOS files from respective vendors (not included)
