


package main

import (

	"fmt"
	"os"
    // "os/signal"
    // "syscall"
    "github.com/joho/godotenv"
	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	token string
)

func main() {
    godotenv.Load()
    token := os.Getenv("DCT")
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
    dg.ChannelMessageSend("984668801188634634","hi")
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	dg.Close()
}
