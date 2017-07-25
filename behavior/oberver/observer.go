package main

import (
	"fmt"
	"sync"
)

//遊戲中介面操作常需要通知其它模組
//使用觀察者模式是不錯的方案

type ui interface {
	registerNotifyChan(ch chan<- string)
	removeNotifyChan(ch chan<- string)
	notifyChan(infot string)
}

type weaponUI struct {
	notifyChans []chan<- string //write only
}

//Using chan as interface to communicate to other moudles
//Register other modules chan then write info into it to notify them
func (s *weaponUI) registerNotifyChan(ch chan<- string) {
	s.notifyChans = append(s.notifyChans, ch)
}

//remove register
func (s *weaponUI) removeNotifyChan(ch chan<- string) {

	for i, c := range s.notifyChans {
		if c == ch {
			s.notifyChans = append(s.notifyChans[:i], s.notifyChans[:i+1]...)
		}
	}

}

//Write info into chan to notify registers
func (s *weaponUI) notify(info string) {
	for _, ch := range s.notifyChans {
		ch <- info
	}
}

//另一個介面當作觀察者
//如果發生事情，觀察者的chan會收到訊息
type infoUI struct {
	infoChan chan string
	wg       *sync.WaitGroup
}

//取得infoUI的chan
func (s *infoUI) getChan() chan string {
	return s.infoChan
}

//infoUI收到訊息，後續處理
func (s *infoUI) showInfo() (resOK bool) {
	defer s.wg.Done()
	defer fmt.Println("infoUI.showInfo stop")

	active := true

	fmt.Println("infoUI.showInfo start")
	for active {
		resStr, ok := <-s.infoChan
		if ok {
			fmt.Printf("infoUI.showUI: %s\n", resStr)
			resOK = true
			active = false
		}
	}

	return resOK
}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	infoUI := &infoUI{
		infoChan: make(chan string, 100),
		wg:       &wg,
	}

	weaponUI := &weaponUI{
		notifyChans: make([]chan<- string, 0, 100),
	}

	c := infoUI.getChan()
	go infoUI.showInfo()

	weaponUI.registerNotifyChan(c)
	weaponUI.notify("got gold sword")

	wg.Wait()
}
