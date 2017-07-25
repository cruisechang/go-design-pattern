package main

import (
	"fmt"
	"strconv"
	"strings"
)

//現在要開發一個外星探險遊戲，
//遊戲中有一種翻譯機，可把預先定義好的英文翻譯成「宇宙語」，
//「宇宙語」是宇宙共同語言，
//藉此跟外星人互動。

const (
	//alien's move
	alienAction0 = iota
	alienAction1
	alienAction2
	alienAction3

	//alien's thought
	alienThought0
	alienThought1
	alienThought2
	alienThought3
)

const (
	sit       = alienAction0
	stand     = alienAction1
	handShake = alienAction2
	punch     = alienAction3

	neverMind   = alienThought0
	wantToLeave = alienThought1
	angry       = alienThought2
	happy       = alienThought3
)

//alien lang to eng interpreter
func newAlienLangInterpreter() *alienLangInterpreter {
	return &alienLangInterpreter{}
}

type alienLangInterpreter struct {
}

func (s *alienLangInterpreter) interprete(alienLang []int) string {

	if len(alienLang) < 1 {
		return ""
	}

	outStr := ""
	for i := 0; i < len(alienLang); i++ {

		str := s.alienLangToEng(alienLang[i])
		if str != "" {
			outStr += str + " "
		} else {
			str = s.alienLangToEng(alienLang[i])
			if str != "" {
				outStr += str + " "
			}
		}
	}
	return outStr
}

func (s *alienLangInterpreter) alienLangToEng(alienLang int) string {
	switch alienLang {

	case alienAction0:
		return "sit"
	case alienAction1:
		return "stand"
	case alienAction2:
		return "handShake"
	case alienAction3:
		return "punch"
	case alienThought0:
		return "neverMind"
	case alienThought1:
		return "wantToLeave"
	case alienThought2:
		return "angry"
	case alienThought3:
		return "happy"
	}
	return ""
}

//eng to alien lang interpreter
func newEngInterpreter() *engInterpreter {
	return &engInterpreter{}
}

type engInterpreter struct {
}

func (s *engInterpreter) interprete(eng string) string {

	if len(eng) < 1 {
		return ""
	}

	splitedLang := strings.Split(eng, ",")

	outStr := ""
	for i := 0; i < len(splitedLang); i++ {

		interpretedLang := s.engToAlienLang(splitedLang[i])

		if interpretedLang != -1 {
			outStr += strconv.Itoa(interpretedLang) + " "
		} else {
			interpretedLang = s.engToAlienLang(splitedLang[i])
			if interpretedLang != -1 {
				outStr += strconv.Itoa(interpretedLang) + " "
			}
		}
	}
	return outStr
}

func (s *engInterpreter) engToAlienLang(eng string) int {
	switch eng {

	case "sit":
		return alienAction0
	case "stand":
		return alienAction1
	case "handShake":
		return alienAction2

	case "punch":
		return alienAction3

	case "neverMind":
		return alienThought0
	case "wantToLeave":
		return alienThought1
	case "angry":
		return alienThought2
	case "happy":
		return alienThought3
	}
	return -1
}

func main() {

	//alien language to eng
	alienLang := []int{0, 1, 3, 5, 7, 6, 9, 100, 0}
	interpreter := newAlienLangInterpreter()
	fmt.Printf("%s\n", interpreter.interprete(alienLang))

	//eng to alien languate
	eng := "sit,angry,happy"
	engInterpreter := newEngInterpreter()
	fmt.Printf("%s\n", engInterpreter.interprete(eng))
}
