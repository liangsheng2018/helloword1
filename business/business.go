package business

import (
	"regexp"
)

func CharacterType(s string)(int,int,int,int){

	var Num = regexp.MustCompile(`\d`)

	var Character = regexp.MustCompile("[a-zA-Z]")

	var Blank = regexp.MustCompile(" ")

	var Special int

	num := len(Num.FindAllStringSubmatch(s,-1))

	character := len(Character.FindAllStringSubmatch(s,-1))

	blank := len(Blank.FindAllStringSubmatch(s,-1))

	Special = len(s) - num - character - blank

	return num, character, blank, Special
}


