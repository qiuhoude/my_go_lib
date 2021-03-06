package randweighted

import (
	"math/rand"
	"testing"
)

func buildData() [][2]int {
	var ret [][2]int

	for i := 0; i < 10000; i++ {
		m := rand.Intn(10)
		ret = append(ret, [2]int{i, m})
	}
	ret = append(ret, [2]int{100, 100})
	return ret

}
func TestWeightedRandom1(t *testing.T) {
	weight, err := WeightedRandom3(buildData())
	if err != nil {
		t.Log("出错..")
	}
	for i := 0; i < 10; i++ {
		t.Log(weight())
	}
}

func BenchmarkWeightedRandom1(b *testing.B) {
	weight, err := WeightedRandom1(buildData())
	if err != nil {
		b.Fail()
	}
	for i := 0; i < b.N; i++ {
		weight()
	}
}

func BenchmarkWeightedRandom2(b *testing.B) {
	weight, err := WeightedRandom2(buildData())
	if err != nil {
		b.Fail()
	}
	for i := 0; i < b.N; i++ {
		weight()
	}
}

func BenchmarkWeightedRandom3(b *testing.B) {
	weight, err := WeightedRandom3(buildData())
	if err != nil {
		b.Fail()
	}
	for i := 0; i < b.N; i++ {
		weight()
	}
}
