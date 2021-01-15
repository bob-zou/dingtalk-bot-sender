# dingtalk-bot-sender
It's a simple golang sdk for dingtalk webhook bot

## Overview
* [Installation](#Installation)
* [Usage](#Usage)
    * [Init Bot](#Init Bot)
    * [Send Message](#Send Message)
        * [Send Text](#Send Text Message)
        * [Send Link](#Send Link)
        * [Send MarkDown](#Send MarkDown)
        * [Send ActionCard](#Send ActionCard)
        * [Send FeedCard](#Send FeedCard)
* [Reference](#Reference)
* [License](#License)

## Installation
```
go get github.com/zouyapeng/dingtalk-bot-sender
```

## Usage
### Init Bot
```go
package main
import "github.com/zouyapeng/dingtalk-bot-sender/sender"

func main() {
	bot := sender.NewBot("access_token of your bot", "secret of your bot")
}
```
### Send Message
#### Send Text Message
```go
_ = bot.SendMessage(sender.TextMessage{
    Content:   "This is a text message",
    AtAll:     true,
})
```
![example](https://cdn.jsdelivr.net/gh/zouyapeng/dingtalk-bot-sender/assets/text.png)
#### Send Link
```go
_ = bot.SendMessage(sender.LinkMessage{
    Title:      "This is a link title",
    Content:    "This is a link message",
    PicURL:     "https://cdn.jsdelivr.net/gh/zouyapeng/dingtalk-bot-sender/assets/pic-2.png",
    MessageURL: "https://github.com/zouyapeng",
})
```
![example](https://cdn.jsdelivr.net/gh/zouyapeng/dingtalk-bot-sender/assets/link.png)
#### Send MarkDown
```go
var markDownContent = `### MarkDown Title
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
```
![example](https://cdn.jsdelivr.net/gh/zouyapeng/dingtalk-bot-sender/assets/markdown.png)
#### Send ActionCard
```go
_ = bot.SendMessage(sender.ActionCardMessage{
    Title:      "This is a action card title",
    Content:    "This is a action card title",
    Horizontal: true,
    Buttons:    []sender.ActionCardButton{
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
```
![example](https://cdn.jsdelivr.net/gh/zouyapeng/dingtalk-bot-sender/assets/action-card.png)
#### Send FeedCard
```go
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
```
![example](https://cdn.jsdelivr.net/gh/zouyapeng/dingtalk-bot-sender/assets/feed-card.png)

## Reference
[https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq](https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq)

## License
[MIT](https://opensource.org/licenses/MIT)
