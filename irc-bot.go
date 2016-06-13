package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/thoj/go-ircevent"
	"log"
	"os"
	"os/user"
)

type Config struct {
	Room_name  string
	Bot_name   string
	Password   string
	Server_url string
}

/* Idea for this code taken from https://blog.gopheracademy.com/advent-2014/reading-config-files-the-go-way/ */
func read_config_file(file_path string) (*Config, bool) {
	var instance_config Config

	_, err := os.Stat(file_path)
	if err != nil {
		return &instance_config, false
	}
	fmt.Println("Attempting to read config file ", file_path)
	if _, err := toml.DecodeFile(file_path, &instance_config); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Config file read succesfully")
	return &instance_config, true
}

func main() {

	usr, _ := user.Current()
	dir := usr.HomeDir

	config_file := flag.String("config-file", dir+"/.ircbot.config", "Path to a non-default config file")

	flag.Parse()

	instance_config, ok := read_config_file(*config_file)

	if !ok {
		log.Fatal("Config file doesn't exist, fail right now")
	}
	conn := irc.IRC(instance_config.Bot_name, instance_config.Bot_name)
	conn.Password = instance_config.Password
	err := conn.Connect(instance_config.Server_url)
	if err != nil {
		fmt.Println("Failed to connect")
		return
	}
	conn.AddCallback("001", func(e *irc.Event) {
		conn.Join(instance_config.Room_name)
	})
	conn.AddCallback("JOIN", func(e *irc.Event) {
		conn.Privmsg(instance_config.Room_name, "Hello! I am a Gobot")
	})
	conn.AddCallback("PRIVMSG", func(e *irc.Event) {
		conn.Privmsg(instance_config.Room_name, e.Message())
	})
	conn.Loop()
}
