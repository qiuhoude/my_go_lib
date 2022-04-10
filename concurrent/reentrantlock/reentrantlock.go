package reentrantlock

import (
	"fmt"
	"github.com/petermattis/goid"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

/*
golang 重入锁的实现
方式1：使用记录 goroutine id
	获取 goroutine id :
		1 使用runtime.Stack()方法获取栈帧信息，栈帧信息里包含 goroutine id
		2 我们获取运行时的 g 指针，反解出对应的 g 的结构。每个运行的 goroutine 结构的
		g 指针保存在当前 goroutine 的一个叫做 TLS 对象中
		第一步：我们先获取到 TLS 对象；
		第二步：再从 TLS 中获取 goroutine 结构的 g 指针；
		第三步：再从 g 指针中取出 goroutine id。
		可以使用开源库 github.com/petermattis/goid

方式2：自己维护token，每次加锁解锁都需要传入token才能进行重复加锁和解锁
*/

func GoId() int64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	// 得到id字符串
	//fmt.Println(string(buf[:n]))
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))
	id, err := strconv.ParseInt(idField[0], 10, 64)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

type ReentrantLock struct {
	sync.Mutex
	owner     int64 // 当前持有锁的goroutine id
	recursion int32 // 这个goroutine 重入的次数
}

func (m *ReentrantLock) Lock() {
	gid := goid.Get()
	// 如果当前持有锁的goroutine就是这次调用的goroutine,说明是重入
	if atomic.LoadInt64(&m.owner) == gid {
		m.recursion++
		return
	}
	m.Mutex.Lock()
	// 获得锁的goroutine第一次调用，记录下它的goroutine id,调用次数加1
	atomic.StoreInt64(&m.owner, gid)
	m.recursion = 1
}
func (m *ReentrantLock) Unlock() {
	gid := goid.Get()
	// 非持有锁的goroutine尝试释放锁，错误的使用
	if atomic.LoadInt64(&m.owner) != gid {
		panic(fmt.Sprintf("wrong the owner(%d): %d!", m.owner, gid))
	}
	// 调用次数减1
	m.recursion--
	if m.recursion != 0 { // 如果这个goroutine还没有完全释放，则直接返回
		return
	}
	// 此goroutine最后一次调用，需要释放锁
	atomic.StoreInt64(&m.owner, -1)
	m.Mutex.Unlock()
}
