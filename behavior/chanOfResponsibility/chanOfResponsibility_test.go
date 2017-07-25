package main

import (
	"fmt"
	"testing"
)

func Test_main(t *testing.T) {

	spy1 := newSpy(1, 100, nil)
	spy2 := newSpy(2, 50, spy1)
	spy3 := newSpy(3, 10, spy2)

	//spy3 collected a level 1 info
	i := info{level: 1}
	spy3.CollectInfo(i)
	spy3.HandleInfo()
	fmt.Println("***************")

	//spy3 collected a level 50 info
	i2 := info{level: 50}
	spy3.CollectInfo(i2)
	spy3.HandleInfo()
	fmt.Println("***************")

	//spy3 collected a level 60 info
	i3 := info{level: 60}
	spy3.CollectInfo(i3)
	spy3.HandleInfo()
	fmt.Println("***************")

	//spy2 collected a level 80 info
	i4 := info{level: 80}
	spy2.CollectInfo(i4)
	spy2.HandleInfo()

	fmt.Println("***************")

	//spy2 collected a level 80 info
	i5 := info{level: 200}
	spy2.CollectInfo(i5)
	spy2.HandleInfo()
}
