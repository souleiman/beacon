package main

import (
    "fmt"
    "log"
    "time"
    "strings"
    "crypto/tls"
    "github.com/souleiman/beacon/sentinel"
    irc "github.com/fluffle/goirc/client"
    "github.com/fluffle/goirc/state"
)

func main() {
    config := sentinel.Config

    cfg := irc.NewConfig(config.Username)
    cfg.NewNick = func(n string) string { return n + "_"}
    cfg.Timeout = 60 * time.Second
    cfg.SplitLen = 450
    cfg.PingFreq = 3 * time.Minute
    cfg.Recover = (*irc.Conn).LogPanic // in dispatch.go

    if config.ZNC != nil {
        znc := config.ZNC
        cfg.Server = fmt.Sprintf("%s:%d", znc.Host, znc.Port)
        cfg.Pass = znc.Password
        cfg.SSL = znc.SSL
        cfg.SSLConfig = &tls.Config{InsecureSkipVerify: znc.SSL_SKIP}
    } else {
        irc := config.IRC
        cfg.Server = fmt.Sprintf("%s:%d", irc.Host, irc.Port)
        cfg.Me = &state.Nick{Nick: irc.Nick}
    }

    config.BuildStalkSet()
    config.BuildChannelSet()

    c := irc.Client(cfg)

    c.HandleFunc("privmsg", func (conn *irc.Conn, line *irc.Line) {
        message := line.Args[1]
        if len(config.Channels) == 0 || (config.StalkSet[line.Nick] && config.ChannelSet[line.Target()]) {
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
        if config.Channels == nil {
            return
        }
        for _, channel := range config.Channels {
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
