package tfcc

import (
	. "github.com/CuteReimu/mirai-sdk-http"
	"math/rand/v2"
	"slices"
	"strconv"
	"strings"
)

func init() {
	addCmdListener(&showTips{})
	addCmdListener(&ping{})
	addCmdListener(&roll{})
}

type showTips struct{}

func (t *showTips) Name() string {
	return "查看帮助"
}

func (t *showTips) ShowTips(int64, int64) string {
	return ""
}

func (t *showTips) CheckAuth(int64, int64) bool {
	return true
}

func (t *showTips) Execute(msg *GroupMessage, _ string) MessageChain {
	var ret []string
	for _, h := range cmdMap {
		if h.CheckAuth(msg.Sender.Group.Id, msg.Sender.Id) {
			if tip := h.ShowTips(msg.Sender.Group.Id, msg.Sender.Id); len(tip) > 0 {
				ret = append(ret, tip)
			}
		}
	}
	slices.Sort(ret)
	return MessageChain{&Plain{Text: "你可以使用以下功能：\n" + strings.Join(ret, "\n")}}
}

type ping struct{}

func (p *ping) Name() string {
	return "ping"
}

func (p *ping) ShowTips(int64, int64) string {
	return ""
}

func (p *ping) CheckAuth(int64, int64) bool {
	return true
}

func (p *ping) Execute(_ *GroupMessage, content string) MessageChain {
	if len(content) == 0 {
		return MessageChain{&Plain{Text: "pong"}}
	}
	return nil
}

type roll struct{}

func (r *roll) Name() string {
	return "roll"
}

func (r *roll) ShowTips(int64, int64) string {
	return ""
}

func (r *roll) CheckAuth(int64, int64) bool {
	return true
}

func (r *roll) Execute(message *GroupMessage, content string) MessageChain {
	if len(content) == 0 {
		return MessageChain{&Plain{Text: message.Sender.MemberName + " roll: " + strconv.Itoa(rand.IntN(100))}}
	}
	return nil
}
