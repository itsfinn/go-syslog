go-syslog [![Build Status](https://travis-ci.org/mcuadros/go-syslog.svg?branch=master)](https://travis-ci.org/mcuadros/go-syslog) [![GoDoc](https://godoc.org/github.com/mcuadros/go-syslog?status.svg)](https://godoc.org/gopkg.in/mcuadros/go-syslog.v2) [![GitHub release](https://img.shields.io/github/release/mcuadros/go-syslog.svg)](https://github.com/mcuadros/go-syslog/releases)
==============================

Syslog server library for go, build easy your custom syslog server over UDP, TCP or Unix sockets using RFC3164, RFC6587 or RFC5424

Installation
------------

The recommended way to install go-syslog

```
go get gopkg.in/mcuadros/go-syslog.v2
```

Examples
--------

How import the package

```go
import "gopkg.in/mcuadros/go-syslog.v2"
```

Example of a basic syslog [UDP server](example/basic_udp.go):

```go
channel := make(syslog.LogPartsChannel)
handler := syslog.NewChannelHandler(channel)

server := syslog.NewServer()
server.SetFormat(syslog.RFC5424)
server.SetHandler(handler)
server.ListenUDP("0.0.0.0:514")
server.Boot()

go func(channel syslog.LogPartsChannel) {
    for logParts := range channel {
        fmt.Println(logParts)
    }
}(channel)

server.Wait()
```

Examples for custom timestamp format
--------

How import the package

```go
import (
    "gopkg.in/mcuadros/go-syslog.v2"
    "gopkg.in/mcuadros/go-syslog.v2/format"
)
```

add replace to `go.mod` file with:

```
replace (
	gopkg.in/mcuadros/go-syslog.v2 v2.3.0 => github.com/itsfinn/go-syslog v0.0.0-20230301102244-b0e9b0c49e46
	github.com/jeromer/syslogparser v1.1.0 => github.com/itsfinn/syslogparser v0.0.0-20230301101959-c5e636156131
)
```

Example of a basic syslog [UDP server](example/basic_udp.go):

```go
channel := make(syslog.LogPartsChannel)
handler := syslog.NewChannelHandler(channel)

server := syslog.NewServer()
fmt3164 = format.NewFormatRFC3164()
fmt3164.WithTimestampFormat("2006-01-02 15:04:05")
server.SetFormat(fmt3164)
server.SetHandler(handler)
server.ListenUDP("0.0.0.0:514")
server.Boot()

go func(channel syslog.LogPartsChannel) {
    for logParts := range channel {
        fmt.Println(logParts)
    }
}(channel)

server.Wait()
```
License
-------

MIT, see [LICENSE](LICENSE)
