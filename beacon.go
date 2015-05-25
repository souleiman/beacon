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


    c.HandleFunc(irc.PRIVMSG, func(conn *irc.Conn, line *irc.Line) {
        message := line.Args[1]
        fmt.Println(message)
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

    c.HandleFunc(irc.CONNECTED, func(conn *irc.Conn, line *irc.Line) {
        fmt.Println("Connected to Server.")

        if config.IRC.Channels == nil {
            return
        }

        for _, commands := range config.IRC.Commands {
            if commands[0] == "msg" {
                conn.Privmsg(commands[1], commands[2])
            }
        }
        for _, channel := range config.IRC.Channels {
            conn.Join(channel)
        }
    })

    c.HandleFunc(irc.INVITE, func(conn *irc.Conn, line *irc.Line) {
        message := line.Args[1]
        from := strings.Split(line.Src, "!")[0]
        if !config.IRC.StalkSet[from] {
            return
        }

        conn.Join(message)
        fmt.Printf("Invited to %s by %s\n", message, from)
    })

    quit := make(chan bool)
    c.HandleFunc(irc.DISCONNECTED, func(conn *irc.Conn, line *irc.Line) {
        quit <- true
    })

    if err := c.Connect(); err != nil {
        fmt.Printf("Connection error: %s\n", err.Error())
    }

    <-quit
}
