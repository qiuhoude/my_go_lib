package leetcode

import "testing"

// 202 https://leetcode-cn.com/problems/happy-number/

/*
编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」定义为：
对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
如果 可以变为  1，那么这个数就是快乐数。
如果 n 是快乐数就返回 true ；不是，则返回 false 。

输入：n = 19
输出：true
解释：
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1

输入：n = 2
输出：false

1 <= n <= 231 - 1

思路:
1. 使用hash表记录经历过的数,如发现hash表中有值说明无线循环,不能进行下去
2. 使用快慢指针,和链表找出链表中成环的思路一致
*/

func isHappy1(n int) bool {
	tab := make(map[int]bool)
	for n != 1 && !tab[n] {
		n, tab[n] = nextNum(n), true
	}
	return n == 1
}

func nextNum(n int) int {
	sum := 0
	for ; n > 0; n = n / 10 {
		sum += (n % 10) * (n % 10)
	}
	return sum
}
func isHappy2(n int) bool {
	slow, fast := n, nextNum(n)
	for slow != fast && slow != 1 {
		slow = nextNum(slow)
		fast = nextNum(nextNum(fast))
	}
	return slow == 1
}

func Test_isHappy(t *testing.T) {
	tests := []struct {
		expected bool
		arg1     int
	}{
		{true, 19},
		{false, 2},
		{true, 7},
	}
	for _, tt := range tests {
		res := isHappy2(tt.arg1)
		if res != tt.expected {
			t.Logf("expected:%v but got:%v", tt.expected, res)
		}
	}
}
