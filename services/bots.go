package services

import (
	"fmt"
	"os"
	// "os"
	"github.com/bwmarrin/discordgo"
)

var (
	token string
)
type Message struct{
	
}
func bot(company_name string,job_title string,job_url string,job_location string, posted_date string) {


	token := os.Getenv("DISCORD_BOT_TOKEN")
	channel_id := os.Getenv("DISCORD_CHANNEL_ID")


	dg, err := discordgo.New("Bot " + token)
	message := "Company Name : " + company_name + "\n" + "Job Location : " + job_location + "\n" + "Posted Date : " + posted_date + "\n"
	
    
	var message_embeds  = []*discordgo.MessageEmbed{&discordgo.MessageEmbed{URL:job_url,Title:job_title,Description:message}}
	dg.ChannelMessageSendEmbeds(channel_id,message_embeds)
	if err != nil {
		fmt.Println("error opening connection,", err)
	}
	error_message_2 := dg.Open()
	if error_message_2 != nil {
		fmt.Println("error opening connection,", error_message_2)
	}	
	dg.Close()

}