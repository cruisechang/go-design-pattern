package main

import "fmt"

//對話選項系統
//玩RPG遊戲時，主角常會碰到npc，
//這時選擇對話項目會有好幾個選項
//每個選項會觸發不同狀況
//這就可以用strategy
//功能：模組化演算法
//功能：隱藏switch分支

const (
	Leave = iota
	Buy   = iota
)

//TalkingSystem is used to select talking options.
//Player select option on UI then do things here.
type TalkingSystem interface {
	SelectAction(action int)
}

type talkingSystem struct {
	leave Action
	buy   Action
}

func (t talkingSystem) SelectAction(action int) {
	switch action {
	case Leave:
		t.leave.act()
	case Buy:
		t.buy.act()
	default:
		fmt.Printf("Select action error\n")
	}
}

func NewTalkingSystem() TalkingSystem {

	return talkingSystem{
		leave: actionLeave{},
		buy:   actionBuy{},
	}
}

//Action is action for talking system.
type Action interface {
	act()
}

//actionLeave represents module of leave action
type actionLeave struct {
}

func (a actionLeave) act() {
	fmt.Printf("Select action leave\n")
}

//actionBuy represents module of buy action
type actionBuy struct {
}

func (a actionBuy) act() {
	fmt.Printf("Select action buy\n")

}

func main() {

	ts := NewTalkingSystem()

	ts.SelectAction(Leave)
	ts.SelectAction(Buy)
	ts.SelectAction(-1)

}
