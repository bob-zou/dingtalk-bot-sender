package main

import "github.com/bob-zou/dingtalk-bot-sender/sender"

func main() {
	// init bot
	bot := sender.NewBot("access_token of your bot", "secret of your bot")

	//text message
	_ = bot.SendMessage(sender.TextMessage{
		Content: "This is a text message",
		AtAll:   true,
	})

	// link message
	_ = bot.SendMessage(sender.LinkMessage{
		Title:      "This is a link title",
		Content:    "This is a link message",
		PicURL:     "https://cdn.jsdelivr.net/gh/bob-zou/dingtalk-bot-sender/assets/pic-2.png",
		MessageURL: "https://github.com/bob-zou",
	})

	markDownContent := `### MarkDown Title
> This is a markdown message

![screenshot](https://cdn.jsdelivr.net/gh/bob-zou/dingtalk-bot-sender/assets/pic-1.png)

 - list 1
 - list 2

[link](https://github.com/bob-zou)
`
	// markdown message
	_ = bot.SendMessage(sender.MarkDownMessage{
		Title:     "This is a markdown title",
		Content:   markDownContent,
		AtMobiles: []string{"13888888888"},
	})

	// action card message
	_ = bot.SendMessage(sender.ActionCardMessage{
		Title:      "This is a action card title",
		Content:    "This is a action card title",
		Horizontal: true,
		Buttons: []sender.ActionCardButton{
			{
				Title: "Action-1",
				URL:   "https://github.com/bob-zou",
			},
			{
				Title: "Action-2",
				URL:   "https://github.com/bob-zou",
			},
		},
	})

	// feed card message
	_ = bot.SendMessage(sender.FeedCardMessage{
		Links: []sender.FeedCardLink{
			{
				Title:      "FeedCard-1",
				PicURL:     "https://cdn.jsdelivr.net/gh/bob-zou/dingtalk-bot-sender/assets/pic-2.png",
				MessageURL: "https://github.com/bob-zou",
			},
			{
				Title:      "FeedCard-2",
				PicURL:     "https://cdn.jsdelivr.net/gh/bob-zou/dingtalk-bot-sender/assets/pic-1.png",
				MessageURL: "https://github.com/bob-zou",
			},
		},
	})

}
