# dingtalk-bot-sender
It's a simple golang sdk for dingtalk webhook bot

## Overview
* [Installation](#Installation)
* [Usage](#Usage)
    * [Init Bot](#InitBot)
    * [Send Message](#SendMessage)
        * [Send Text](#SendTextMessage)
        * [Send Link](#SendLink)
        * [Send MarkDown](#SendMarkDown)
        * [Send ActionCard](#SendActionCard)
        * [Send FeedCard](#SendFeedCard)
* [Reference](#Reference)
* [License](#License)

## Installation
```
go get github.com/bob-zou/dingtalk-bot-sender
```

## Usage
### InitBot
```go
package main
import "github.com/bob-zou/dingtalk-bot-sender/sender"

func main() {
	bot := sender.NewBot("access_token of your bot", "secret of your bot")
}
```
### SendMessage
#### SendTextMessage
```go
_ = bot.SendMessage(sender.TextMessage{
    Content:   "This is a text message",
    AtAll:     true,
})
```
![example](https://cdn.jsdelivr.net/gh/bob-zou/dingtalk-bot-sender/assets/text.png)
#### SendLink
```go
_ = bot.SendMessage(sender.LinkMessage{
    Title:      "This is a link title",
    Content:    "This is a link message",
    PicURL:     "https://cdn.jsdelivr.net/gh/bob-zou/dingtalk-bot-sender/assets/pic-2.png",
    MessageURL: "https://github.com/bob-zou",
})
```
![example](https://cdn.jsdelivr.net/gh/bob-zou/dingtalk-bot-sender/assets/link.png)
#### SendMarkDown
```go
var markDownContent = `### MarkDown Title
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
```
![example](https://cdn.jsdelivr.net/gh/bob-zou/dingtalk-bot-sender/assets/markdown.png)
#### SendActionCard
```go
_ = bot.SendMessage(sender.ActionCardMessage{
    Title:      "This is a action card title",
    Content:    "This is a action card title",
    Horizontal: true,
    Buttons:    []sender.ActionCardButton{
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
```
![example](https://cdn.jsdelivr.net/gh/bob-zou/dingtalk-bot-sender/assets/action-card.png)
#### SendFeedCard
```go
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
```
![example](https://cdn.jsdelivr.net/gh/bob-zou/dingtalk-bot-sender/assets/feed-card.png)

## Reference
[https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq](https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq)

## License
[MIT](https://opensource.org/licenses/MIT)
