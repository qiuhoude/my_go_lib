package ac

import (
	"bufio"
	"os"
	"testing"
)

func getDirtyword() []string {
	file, err := os.Open("dirtyword.txt")
	if err != nil {
		panic(err)
	}
	var ret []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		t := scanner.Text()
		ret = append(ret, t)
	}
	return ret
}
func TestAC_Match(t *testing.T) {

	acAuto := NewAc()
	dirtWorld := getDirtyword()

	acAuto.AddWorlds(dirtWorld)
	//t.Log(acAuto.Contains("好屌"))

	acAuto.BuildFailurePointer()

	acAuto.Remove("fuck")
	acAuto.BuildFailurePointer()

	text := `妈的-我看他说话的语气，好屌啊 fuck gangcunxiushu 心灵法门“白话佛法”系列节目 @@@69式111`
	afterText := []rune(text)
	acAuto.Match(text, func(start, end int) {
		//t.Logf("%v [%v-%v]\n", text[start:end+1], start, end)
		for i := start; i <= end; i++ {
			afterText[i] = '*'
		}
	})
	t.Log(string(afterText))
}
