package sort_

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"testing"
	"time"
)

type PowerRankItem struct {
	id    PlayerId
	name  string
	power int
}

func NewPowerRankItem(id PlayerId, name string, power int) *PowerRankItem {
	return &PowerRankItem{
		id:    id,
		name:  name,
		power: power,
	}
}

func (p *PowerRankItem) GetPlayerId() PlayerId {
	return p.id
}
func (p *PowerRankItem) GetScore() int {
	return p.power
}
func (p *PowerRankItem) SetScore(newScore int) {
	p.power = newScore
}

func (p *PowerRankItem) String() string {
	return fmt.Sprintf("{%v %v %v}", p.id, p.name, p.power)
}

func buildRankItem(r *rand.Rand, cnt int) []IRankItem {
	var items []IRankItem
	for i := 0; i < cnt; i++ {
		power := r.Intn(cnt)
		item := NewPowerRankItem(PlayerId(i), "name_"+strconv.Itoa(i), power)
		items = append(items, item)

	}
	return items
}

func TestGameRankArr(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	sz := 1000
	rank := NewGameRankArr(sz)
	items := buildRankItem(r, sz)
	// add
	for i := range items {
		//fmt.Println("add", items[i])
		rank.SetPlayerScore(items[i], items[i].GetScore())

		//printRank(rank)
	}
	// change
	for i := 0; i < sz; i++ {
		np := r.Intn(sz)
		item := items[r.Intn(len(items))]
		rank.SetPlayerScore(item, np)

		//fmt.Printf("change %v power %v\n", item, np)
		//printRank(rank)
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].GetScore() > items[j].GetScore()
	})
	list := rank.GetRankList()

	//fmt.Println()
	//printRank(rank)
	//fmt.Println(items)

	if !equalSlicePowerRankItem(items[0:sz], list) {
		t.Error("不一致")
	}

}
func equalSlicePowerRankItem(a []IRankItem, b []IRankItem) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i].GetScore() != b[i].GetScore() {
			return false
		}
	}
	return true
}

func printRank(rank *GameRankArr) {
	for _, v := range rank.GetRankList() {
		fmt.Println(v)
	}

	fmt.Println("mapSize", len(rank.rankMap))
	fmt.Println("--------------------------")

}
