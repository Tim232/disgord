// Very basic Discord Bot wrote using the discordgo package
package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/disgord/bot"
)

const (
	VERSION = "v0.0.0"
)

var (
	Bot = bot.New()
)

func init() {

	// Get Bot Token
	Bot.Token = os.Getenv("DG_TOKEN")
	if Bot.Token == "" {
		flag.StringVar(&Bot.Token, "t", "", "Discord token.")
	}

}

func main() {
	var err error

	// Parse command line arguments
	flag.Parse()

	// TODO
	// Read Configuration files

	// TODO
	// Parse Run Commands file

	// Connect to Discord if a Token exists.
	// TODO: Make Optional
	if Bot.Token != "" {
		err = Bot.Open()
		if err != nil {
			log.Println("error opening connection to Discord: ", err)
		}
	}

	log.Printf("Disgord %s is now running.  Press CTRL-C to exit.", VERSION)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
