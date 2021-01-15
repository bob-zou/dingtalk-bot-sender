package main

import "github.com/zouyapeng/dingtalk-bot-sender/sender"

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
		PicURL:     "https://cdn.jsdelivr.net/gh/zouyapeng/dingtalk-bot-sender/assets/pic-2.png",
		MessageURL: "https://github.com/zouyapeng",
	})

	markDownContent := `### MarkDown Title
> This is a markdown message

![screenshot](https://cdn.jsdelivr.net/gh/zouyapeng/dingtalk-bot-sender/assets/pic-1.png)

 - list 1
 - list 2

[link](https://github.com/zouyapeng)
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
				URL:   "https://github.com/zouyapeng",
			},
			{
				Title: "Action-2",
				URL:   "https://github.com/zouyapeng",
			},
		},
	})

	// feed card message
	_ = bot.SendMessage(sender.FeedCardMessage{
		Links: []sender.FeedCardLink{
			{
				Title:      "FeedCard-1",
				PicURL:     "https://cdn.jsdelivr.net/gh/zouyapeng/dingtalk-bot-sender/assets/pic-2.png",
				MessageURL: "https://github.com/zouyapeng",
			},
			{
				Title:      "FeedCard-2",
				PicURL:     "https://cdn.jsdelivr.net/gh/zouyapeng/dingtalk-bot-sender/assets/pic-1.png",
				MessageURL: "https://github.com/zouyapeng",
			},
		},
	})

}
