package main

import "fmt"

// 這是一個間諜聯繫系統
// 領導者跟下屬是一對一的直線聯繫，
// 屬下蒐集到情報後，往上傳遞給上級，
// 上級判斷這個請報是否屬自己可處理的權限，
// 若屬於自己權限就處理，
// 否則就往上傳遞

type spy interface {
	CollectInfo(i info)
	HandleInfo()
}

type spyer struct {
	id        int
	authority int  //權限，用來判斷情報是否屬於自己可處理範疇。情報範疇小於等於自己的權限，就可以處理，大於自己權限就往上傳。
	superior  spy  //自己的上級。若情報不屬於自己範疇，就傳給上級。
	info      info //info that spyer collected
}

func newSpy(id int, autority int, sp spy) spy {
	return &spyer{
		id:        id,
		authority: autority,
		superior:  sp,
	}
}

func (s *spyer) CollectInfo(i info) {
	fmt.Printf("id %d spy collect info level %d\n", s.id, i.level)
	s.info = i
}

func (s *spyer) HandleInfo() {

	if s.authority >= s.info.level {
		fmt.Printf("id %d spy handle info level %d\n", s.id, s.info.level)
	} else {
		if s.superior != nil {
			fmt.Printf("id %d spy pass info level %d to superior\n", s.id, s.info.level)
			//pass info to superior
			s.superior.CollectInfo(s.info)
			s.superior.HandleInfo()
		} else {
			fmt.Printf("id %d spy has no superior,hold info level %d\n", s.id, s.info.level)
		}
	}
}

//情報
type info struct {
	level int //info level
}

func main() {
	spy1 := newSpy(1, 100, nil)
	spy2 := newSpy(2, 50, spy1)
	spy3 := newSpy(3, 10, spy2)

	//spy3 collected a level 1 info
	i := info{level: 1}
	spy3.CollectInfo(i)
	spy3.HandleInfo()

	//spy3 collected a level 50 info
	i2 := info{level: 50}
	spy3.CollectInfo(i2)
	spy3.HandleInfo()

	//spy2 collected a level 80 info
	i3 := info{level: 80}
	spy2.CollectInfo(i3)
	spy2.HandleInfo()
}
