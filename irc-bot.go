package main

import (
    "github.com/thoj/go-ircevent"
    "fmt"
)

type config struct {
    room_name string
    bot_name string
    password string
    server_url string
}


func main() {
    var instance_config config
    instance_config.room_name = ""
    instance_config.bot_name = ""
    instance_config.password = ""
    instance_config.server_url = ""

    conn := irc.IRC(instance_config.bot_name, instance_config.bot_name)
    conn.Password = instance_config.password
    err := conn.Connect(instance_config.server_url)
    if err != nil {
        fmt.Println("Failed to connect")
        return
    }
    conn.AddCallback("001", func (e *irc.Event) {
        conn.Join(instance_config.room_name)
    })
    conn.AddCallback("JOIN", func (e *irc.Event) {
        conn.Privmsg(instance_config.room_name, "Hello! I am a Gobot")
    })
    conn.AddCallback("PRIVMSG", func (e *irc.Event) {
        conn.Privmsg(instance_config.room_name, e.Message())
    })
    conn.Loop()
}
