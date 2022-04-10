package trymutex

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

/*
给 Mutex 加上tryLock方法,
原理其实就参考 Lock 方法

*/

const (
	// 来自 sync.Mutex源码
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	//mutexWaiterShift = iota
)

type TryMutex struct {
	sync.Mutex
}

// TryLock 尝试获取锁
func (m *TryMutex) TryLock() bool {
	// 成功抢到锁
	// Mutex的首字段是 state (*int32)(unsafe.Pointer(&m.Mutex)) 就可以表示 &m.state
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
		return true
	}
	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	// 如果处于唤醒、加锁或者饥饿状态，这次请求就不参与竞争了，返回false
	if old&(mutexLocked|mutexStarving|mutexWoken) != 0 {
		return false
	}
	// 尝试在竞争的状态下请求锁
	newState := old | mutexLocked
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), old, newState)
}

// IsLocked 锁是否被持有
func (m *TryMutex) IsLocked() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexLocked == mutexLocked
}

// IsWoken 是否有等待者被唤醒
func (m *TryMutex) IsWoken() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexWoken == mutexWoken
}

// IsStarving  锁是否处于饥饿状态
func (m *TryMutex) IsStarving() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexStarving == mutexStarving
}
