This is a small IRC bot written in Go just so I can get used to the language

Config file
Currently a config file is expected and it's expected to exist in $HOME/ircbot.config. The file should be in the TOML format.
An example of the format is

Room_name = "Example"
Bot_name = "Tester"
Password = "pass"
Server_url = "test.server:6667"

Command line arguments

-config-file <string>
    A path to a user config file - default is $HOME/ircbot.config
