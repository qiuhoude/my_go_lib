package reentrantlock

import (
	"github.com/bmizerany/assert"
	"github.com/petermattis/goid"
	"testing"
)

func TestGoId(t *testing.T) {
	gid1 := GoId()
	gid2 := goid.Get()
	assert.Equal(t, gid1, gid2)
}

func TestReentrantLock(t *testing.T) {
	var rtLock ReentrantLock
	rtLock.Lock()
	rtLock.Lock()
	rtLock.Unlock()
	rtLock.Unlock()
}
