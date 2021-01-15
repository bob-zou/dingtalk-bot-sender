package sender

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type MessageRender interface {
	Render() ([]byte, error)
}

// TextMessage
// AtMobiles and AtAll cannot work at the same time
type TextMessage struct {
	Content   string
	AtMobiles []string
	AtAll     bool
}

func (m TextMessage) Render() (ret []byte, err error) {
	ret, _ = json.Marshal(map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": m.Content,
		},
		"at": map[string]interface{}{
			"atMobiles": m.AtMobiles,
			"isAtAll":   m.AtAll,
		},
	})
	return
}

type LinkMessage struct {
	Title      string
	Content    string
	PicURL     string
	MessageURL string
}

func (m LinkMessage) Render() (ret []byte, err error) {
	ret, _ = json.Marshal(map[string]interface{}{
		"msgtype": "link",
		"link": map[string]interface{}{
			"title":      m.Title,
			"text":       m.Content,
			"picUrl":     m.PicURL,
			"messageUrl": m.MessageURL,
		},
	})
	return
}

// MarkDownMessage
// AtMobiles and AtAll cannot work at the same time
type MarkDownMessage struct {
	Title     string
	Content   string
	AtMobiles []string
	AtAll     bool
}

func (m MarkDownMessage) Render() (ret []byte, err error) {
	ret, _ = json.Marshal(map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"title": m.Title,
			"text":  m.Content,
		},
		"at": map[string]interface{}{
			"atMobiles": m.AtMobiles,
			"isAtAll":   m.AtAll,
		},
	})
	return
}

type ActionCardButton struct {
	Title string
	URL   string
}

type ActionCardMessage struct {
	Title      string
	Content    string
	Horizontal bool
	Buttons    []ActionCardButton
}

func (m ActionCardMessage) Render() (ret []byte, err error) {
	var (
		btnOrientation = 0
	)
	if m.Horizontal {
		btnOrientation = 1
	}

	if len(m.Buttons) == 0 {
		err = fmt.Errorf("buttons cannot be nil")
		return
	}

	// single button
	if len(m.Buttons) == 1 {
		ret, _ = json.Marshal(map[string]interface{}{
			"msgtype": "actionCard",
			"actionCard": map[string]interface{}{
				"title":          m.Title,
				"text":           m.Content,
				"btnOrientation": btnOrientation,
				"singleTitle":    m.Buttons[0].Title,
				"singleURL":      m.Buttons[0].URL,
			},
		})
	}

	// multi buttons
	buttons := make([]map[string]string, 0)
	for _, button := range m.Buttons {
		buttons = append(buttons, map[string]string{
			"title":     button.Title,
			"actionURL": button.URL,
		})
	}

	ret, _ = json.Marshal(map[string]interface{}{
		"msgtype": "actionCard",
		"actionCard": map[string]interface{}{
			"title":          m.Title,
			"text":           m.Content,
			"btnOrientation": btnOrientation,
			"btns":           buttons,
		},
	})

	return
}

type FeedCardLink struct {
	Title      string
	PicURL     string
	MessageURL string
}

type FeedCardMessage struct {
	Links []FeedCardLink
}

func (m FeedCardMessage) Render() (ret []byte, err error) {
	if len(m.Links) == 0 {
		err = fmt.Errorf("links cannot be nil")
		return
	}
	links := make([]map[string]string, 0)

	for _, link := range m.Links {
		links = append(links, map[string]string{
			"title":      link.Title,
			"messageURL": link.MessageURL,
			"picURL":     link.PicURL,
		})
	}

	ret, _ = json.Marshal(map[string]interface{}{
		"msgtype": "feedCard",
		"feedCard": map[string]interface{}{
			"links": links,
		},
	})

	return
}

func NewBot(token, secret string) *Bot {
	return &Bot{
		token:  token,
		secret: secret,
	}
}

type Bot struct {
	token  string
	secret string // needed if select Additional Signature in Security Settings
}

func (s *Bot) sign(milTimestamp int64) string {
	strToHash := fmt.Sprintf("%d\n%s", milTimestamp, s.secret)
	hmac256 := hmac.New(sha256.New, []byte(s.secret))
	hmac256.Write([]byte(strToHash))
	data := hmac256.Sum(nil)
	return base64.StdEncoding.EncodeToString(data)
}

func (s *Bot) SendMessage(msg MessageRender) (err error) {
	values := url.Values{}
	values.Set("access_token", s.token)

	if s.secret != "" {
		t := time.Now().UnixNano() / 1e6
		values.Set("timestamp", fmt.Sprintf("%d", t))
		values.Set("sign", s.sign(t))
	}

	formattedMsg, err := msg.Render()
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, OAPIURL, bytes.NewBuffer(formattedMsg))
	if err != nil {
		return
	}

	req.URL.RawQuery = values.Encode()
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	res, err := (&http.Client{}).Do(req)
	if err != nil {
		return fmt.Errorf("send dingTalk message failed, error: %v", err.Error())
	}
	defer func() { _ = res.Body.Close() }()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("send dingTalk message failed, error: %v", err.Error())
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("send dingTalk message failed, http code is %d, required 200", res.StatusCode)
	}

	type response struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}

	var ret response

	if err = json.Unmarshal(result, &ret); err != nil {
		return fmt.Errorf("send dingTalk message failed, %s", err.Error())
	}

	if ret.ErrCode != 0 {
		return fmt.Errorf("send dingTalk message failed, %s", ret.ErrMsg)
	}

	return
}
