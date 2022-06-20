package services

import (
	
	"fmt"
	// "os"

	"github.com/bwmarrin/discordgo"
)

var (
	token string
)
type Message struct{
	
}
func bot(company_name string,job_title string,job_url string,job_location string) {

    // token := os.Getenv("DCT")

	dg, err := discordgo.New("Bot " + "OTg1NTY1MDkwMDIxNTI3NjYz.G5DVtB.ZqcRPGckwaQpc26BkUQvI6qvHjiBfkxVz2l9aA")
	message := "Company Name : " + company_name + "\n" + "Job Location : " + job_location + "\n"
	fmt.Println(message)
    // dg.ChannelMessageSend("984668801188634634",})
	var message_embeds  = []*discordgo.MessageEmbed{&discordgo.MessageEmbed{URL:job_url,Title:job_title,Description:message}}
	dg.ChannelMessageSendEmbeds("984668801188634634",message_embeds)
	if err != nil {
		fmt.Println("error opening connection,", err)
	}
	error_message_2 := dg.Open()
	if error_message_2 != nil {
		fmt.Println("error opening connection,", error_message_2)
	}	
	dg.Close()

}