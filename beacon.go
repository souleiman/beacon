package main

import (
    "fmt"
    "log"
    "time"
    "strings"
    "crypto/tls"
    "github.com/souleiman/beacon/sentinel"
    irc "github.com/fluffle/goirc/client"
)

func main() {
    config := sentinel.Config

    cfg := irc.NewConfig(config.IRC.Nick)
    cfg.Server = fmt.Sprintf("%s:%d", config.IRC.Host, config.IRC.Port)
    cfg.Pass = config.IRC.Password
    cfg.SSL = config.IRC.SSL
    cfg.SSLConfig = &tls.Config{InsecureSkipVerify: config.IRC.SSL_SKIP}
    cfg.NewNick = func(n string) string { return n + "_"}
    cfg.Timeout = 60 * time.Second
    cfg.SplitLen = 450
    cfg.PingFreq = 3 * time.Minute
    cfg.Recover = (*irc.Conn).LogPanic // in dispatch.go

    c := irc.Client(cfg)


    c.HandleFunc("privmsg", func (conn *irc.Conn, line *irc.Line) {
        message := line.Args[1]
        if len(config.IRC.Channels) == 0 || (config.IRC.StalkSet[line.Nick] && config.IRC.ChannelSet[line.Target()]) {
            if !strings.HasPrefix(message, "New Torrent:") {
                return
            }

            torrent := sentinel.Fetch(message)
            if sentinel.Filter.Valid(torrent) {
                success := torrent.Download()

                if !success {
                    log.Println("Failed to log: " + message)
                }
            }
        }
    })

    c.HandleFunc("connected", func (conn *irc.Conn, line *irc.Line) {
        if config.IRC.Channels == nil {
            return
        }
        for _, channel := range config.IRC.Channels {
            conn.Join(channel)
        }
        fmt.Println("Connected to Server.")
    })

    quit := make(chan bool)
    c.HandleFunc("disconnected", func (conn *irc.Conn, line *irc.Line) {
        quit <- true
    })

    if err := c.Connect(); err != nil {
        fmt.Printf("Connection error: %s\n", err.Error())
    }

    <-quit
}
