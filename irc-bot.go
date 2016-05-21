package main

import (
    "github.com/thoj/go-ircevent"
    "github.com/BurntSushi/toml"
    "fmt"
    "os"
    "os/user"
    "log"
)

type Config struct {
    Room_name string
    Bot_name string
    Password string
    Server_url string
}

/* Idea for this code taken from https://blog.gopheracademy.com/advent-2014/reading-config-files-the-go-way/ */
func read_config_file (file_path string) *Config {
    var instance_config Config

    _, err := os.Stat(file_path)
    if err != nil {
	log.Fatal("Config file doesn't exist")
    }
    fmt.Println("Attempting to read config file ", file_path)
    if _, err := toml.DecodeFile(file_path, &instance_config); err != nil {
	log.Fatal(err)
    }
    fmt.Println("Config file read succesfully")
    return &instance_config
}


func main() {

    usr,_ := user.Current()
    dir :=usr.HomeDir
    instance_config := read_config_file(dir+"/.ircbot.config")

    conn := irc.IRC(instance_config.Bot_name, instance_config.Bot_name)
    conn.Password = instance_config.Password
    err := conn.Connect(instance_config.Server_url)
    if err != nil {
        fmt.Println("Failed to connect")
        return
    }
    conn.AddCallback("001", func (e *irc.Event) {
        conn.Join(instance_config.Room_name)
    })
    conn.AddCallback("JOIN", func (e *irc.Event) {
        conn.Privmsg(instance_config.Room_name, "Hello! I am a Gobot")
    })
    conn.AddCallback("PRIVMSG", func (e *irc.Event) {
        conn.Privmsg(instance_config.Room_name, e.Message())
    })
    conn.Loop()
}
