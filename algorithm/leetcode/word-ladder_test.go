package leetcode

import (
	"testing"
)

// 127. 单词接龙 https://leetcode-cn.com/problems/word-ladder/

/*
字典 wordList 中从单词 beginWord 和 endWord 的 转换序列 是一个按下述规格形成的序列：

序列中第一个单词是 beginWord 。
序列中最后一个单词是 endWord 。
每次转换只能改变一个字母。
转换过程中的中间单词必须是字典 wordList 中的单词。
给你两个单词 beginWord 和 endWord 和一个字典 wordList ，找到从 beginWord 到 endWord 的 最短转换序列 中的 单词数目 。如果不存在这样的转换序列，返回 0。

输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
输出：5
解释：一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog", 返回它的长度 5。

输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
输出：0
解释：endWord "cog" 不在字典中，所以无法进行转换。

1 <= beginWord.length <= 10
endWord.length == beginWord.length
1 <= wordList.length <= 5000
wordList[i].length == beginWord.length
beginWord、endWord 和 wordList[i] 由小写英文字母组成
beginWord != endWord
wordList 中的所有字符串 互不相同

思路:
将问题转换成无向图, 每个str当成一个顶点, 将str的变化组成新的字符串也当成虚顶点加入到图中,这样顶点与顶点的连接都是通过虚节点进行连接
使用广度优先求解起点到终点的最短路径

改进
使用双广度优先, 同时从启动和终点开始找
*/

// 双向广度优先
func ladderLength3(beginWord string, endWord string, wordList []string) int {
	// 将字符串列表放到hash表中
	wordListMap := make(map[string]bool, len(wordList)+1)
	for _, s := range wordList {
		wordListMap[s] = true
	}
	if _, ok := wordListMap[endWord]; !ok { // 没有结束字符串
		return 0
	}
	visitedBegin := map[string]bool{beginWord: true}
	visitedEnd := map[string]bool{endWord: true}

	queueBegin := []string{beginWord}
	queueEnd := []string{endWord}
	step := 1

	// 双向广度优找路径
	for len(queueBegin) > 0 && len(queueEnd) > 0 {
		beginArr := queueBegin[:]
		queueBegin = nil
		for _, word := range beginArr {
			wordByte := []byte(word)
			// 因为全是小写字母,可以直接将每个位置的字符替换成其他的小写字母在hash中看是否有存在
			for i := 0; i < len(wordByte); i++ {
				originC := wordByte[i] // 原来的字符
				for j := 'a'; j <= 'z'; j++ {
					c := byte(j)
					if originC == c {
						continue
					}
					wordByte[i] = c
					nextWord := string(wordByte)
					if _, has := wordListMap[nextWord]; has {
						if visitedEnd[nextWord] {
							return step + 1
						}
						if !visitedBegin[nextWord] {
							visitedBegin[nextWord] = true
							queueBegin = append(queueBegin, nextWord)
							//fmt.Printf("begin@ %v: %v -> %v\n", step, word, nextWord)
						}
					}
				}
				wordByte[i] = originC // 还原
			}
		}
		step++
		endArr := queueEnd[:]
		queueEnd = nil
		for _, word := range endArr {
			wordByte := []byte(word)
			for i := 0; i < len(wordByte); i++ {
				originC := wordByte[i]
				for j := 'a'; j <= 'z'; j++ {
					c := byte(j)
					if originC == c {
						continue
					}
					wordByte[i] = c
					nextWord := string(wordByte)
					if _, has := wordListMap[nextWord]; has {
						if visitedBegin[nextWord] {
							return step + 1
						}
						if !visitedEnd[nextWord] {
							visitedEnd[nextWord] = true
							queueEnd = append(queueEnd, nextWord)
							//fmt.Printf("end@ %v: %v -> %v\n", step, word, nextWord)
						}
					}
				}
				wordByte[i] = originC // 还原
			}
		}
		step++
	}
	return 0
}

func ladderLength2(beginWord string, endWord string, wordList []string) int {
	// 将字符串列表放到hash表中
	wordListMap := make(map[string]int, len(wordList)+1) // <word,距离>
	const initDist int = -1
	for _, s := range wordList {
		wordListMap[s] = initDist
	}
	if _, ok := wordListMap[endWord]; !ok { // 没有结束字符串
		return 0
	}
	wordListMap[beginWord] = 1

	queue := []string{beginWord}
	// 广度优找路径
	for len(queue) > 0 {
		word := queue[0]
		queue = queue[1:]
		wordByte := []byte(word)
		// 因为全是小写字母,可以直接将每个位置的字符替换成其他的小写字母在hash中看是否有存在
		for i := 0; i < len(wordByte); i++ {
			originC := wordByte[i] // 原来的字符
			for j := 'a'; j <= 'z'; j++ {
				c := byte(j)
				if originC == c {
					continue
				}
				wordByte[i] = c
				nextWord := string(wordByte)
				if nextWord == endWord {
					return wordListMap[word] + 1
				}
				if dist, has := wordListMap[nextWord]; has {
					if dist == initDist {
						wordListMap[nextWord] = wordListMap[word] + 1
						queue = append(queue, nextWord)
					}
				}
			}
			wordByte[i] = originC // 还原
		}
	}
	return 0
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordId := map[string]int{}
	var graph [][]int // 邻阶表的方式
	// 分配id
	addWord := func(word string) int {
		id, has := wordId[word]
		if !has {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id1 := addWord(word)
		s := []byte(word)
		for i, b := range s {
			s[i] = '*'
			// 创建虚节点
			id2 := addWord(string(s))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			s[i] = b
		}
		return id1
	}
	// 构建图
	for _, word := range wordList {
		addEdge(word)
	}
	beginId := addEdge(beginWord)
	endId, has := wordId[endWord]
	if !has {
		return 0
	}

	const initVal int = -1

	/*
		// 单向广度优先
		dist := make([]int, len(wordId)) // 距离或步数,下标表示起点到该点的距离
		for i := range dist {
			dist[i] = initVal
		}
		dist[beginId] = 0
		queue := []int{beginId}
		for len(queue) > 0 {
			v := queue[0]
			queue = queue[1:]
			if v == endId { // 找到了目标
				return dist[endId]/2 + 1 // 除2 因为图中有虚顶点, 顶点相互连接是通过虚节点进行
			}
			for _, w := range graph[v] {
				if dist[w] == initVal {
					dist[w] = dist[v] + 1 // 到这一步的总步数
					queue = append(queue, w)
				}
			}
		}
	*/

	// 双向广度优先
	distBegin := make([]int, len(wordId))
	for i := range distBegin {
		distBegin[i] = initVal
	}
	distBegin[beginId] = 0

	distEnd := make([]int, len(wordId))
	for i := range distEnd {
		distEnd[i] = initVal
	}
	distEnd[endId] = 0

	queueBegin := []int{beginId}
	queueEnd := []int{endId}

	for len(queueBegin) > 0 && len(queueEnd) > 0 {
		//  begin->end
		beginArr := queueBegin // 取出队列中所有的数据
		queueBegin = nil
		for _, id := range beginArr {
			if distEnd[id] != initVal { // 已经找到结果
				return (distBegin[id]+distEnd[id])/2 + 1
			}
			for _, w := range graph[id] {
				if distBegin[w] == initVal { //没有被访问
					distBegin[w] = distBegin[id] + 1 // 到这一步的总步数
					queueBegin = append(queueBegin, w)
				}
			}
		}
		// end->begin
		endArr := queueEnd // 取出队列中所有的数据
		queueEnd = nil
		for _, id := range endArr {
			if distBegin[id] != initVal { // 已经找到结果
				return (distBegin[id]+distEnd[id])/2 + 1
			}
			for _, w := range graph[id] {
				if distEnd[w] == initVal {
					distEnd[w] = distEnd[id] + 1 // 到这一步的总步数
					queueEnd = append(queueEnd, w)
				}
			}
		}
	}
	return 0
}

func Test_ladderLength(t *testing.T) {

	tests := []struct {
		expected int
		arg1     string
		arg2     string
		arg3     []string
	}{
		{5, "hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}},
		{2, "hit", "hot", []string{"hot", "dot", "dog", "lot", "log", "cog"}},
		{5, "hit", "cog", []string{"hot", "cog", "dot", "dog", "hit", "lot", "log"}},
	}
	for _, tt := range tests {
		res := ladderLength(tt.arg1, tt.arg2, tt.arg3)
		if res != tt.expected {
			t.Logf("ladderLength(%v,%v,%v) expected:%v but got:%v", tt.arg1, tt.arg2, tt.arg3, tt.expected, res)
		}
	}
}
