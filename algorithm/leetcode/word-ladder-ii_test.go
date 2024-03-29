package leetcode

import (
	"testing"
)

// 126. 单词接龙 II https://leetcode-cn.com/problems/word-ladder-ii/

/*
按字典 wordList 完成从单词 beginWord 到单词 endWord 转化，一个表示此过程的 转换序列 是形式上像 beginWord -> s1 -> s2 -> ... -> sk 这样的单词序列，并满足：

每对相邻的单词之间仅有单个字母不同。
转换过程中的每个单词 si（1 <= i <= k）必须是字典 wordList 中的单词。注意，beginWord 不必是字典 wordList 中的单词。
sk == endWord
给你两个单词 beginWord 和 endWord ，以及一个字典 wordList 。请你找出并返回所有从 beginWord 到 endWord 的 最短转换序列 ，如果不存在这样的转换序列，返回一个空列表。每个序列都应该以单词列表 [beginWord, s1, s2, ..., sk] 的形式返回。



示例 1：

输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
输出：[["hit","hot","dot","dog","cog"],["hit","hot","lot","log","cog"]]
解释：存在 2 种最短的转换序列：
"hit" -> "hot" -> "dot" -> "dog" -> "cog"
"hit" -> "hot" -> "lot" -> "log" -> "cog"
示例 2：

输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
输出：[]
解释：endWord "cog" 不在字典 wordList 中，所以不存在符合要求的转换序列。


提示：

1 <= beginWord.length <= 7
endWord.length == beginWord.length
1 <= wordList.length <= 5000
wordList[i].length == beginWord.length
beginWord、endWord 和 wordList[i] 由小写英文字母组成
beginWord != endWord
wordList 中的所有单词 互不相同

思路:
与 127题一致, 构建图,使用广度优先找到多路径path,然后使用dfs+stack解析path得到结果
使用数组记录路径, 下标表示节点id,值存的上一个节点的id



*/

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordId := make(map[string]int, len(wordList)+1) // <word,id> 单词到id的映射
	idWord := make(map[int]string, len(wordList)+1) // <id,word> id到单词的映射
	var graph [][]int

	// 分配id
	addVertexFn := func(word string) int {
		id, has := wordId[word]
		if !has {
			id = len(wordId)
			wordId[word] = id
			idWord[id] = word
			graph = append(graph, []int{})
		}
		return id
	}

	beginId := addVertexFn(beginWord)
	// 构建邻接表图,添加顶点
	for _, word := range wordList {
		addVertexFn(word)
	}

	var res [][]string
	if _, has := wordId[endWord]; !has {
		return res
	}
	// 添加边
	for id1 := 0; id1 < len(idWord)-1; id1++ {
		for id2 := id1 + 1; id2 < len(idWord); id2++ {
			if isDiffOneCharacter(idWord[id1], idWord[id2]) {
				graph[id1] = append(graph[id1], id2)
				graph[id2] = append(graph[id2], id1)
			}
		}
	}
	// 广度优先查找
	endId := wordId[endWord]
	paths := map[int]map[int]bool{} // <id,set<preId>> 记录多条路径
	visited := map[int]bool{}
	queue := []int{beginId}

	for len(queue) > 0 {
		idArr := queue
		queue = nil
		for _, id := range idArr {
			visited[id] = true
		}
		for _, id := range idArr {
			if id == endId { // 广度优先本层级找到最小值
				break
			}
			for _, nextId := range graph[id] {
				if !visited[nextId] {
					if _, has := paths[nextId]; has {
						paths[nextId][id] = true
					} else {
						paths[nextId] = map[int]bool{id: true}
					}
					queue = append(queue, nextId)
					//fmt.Printf("%v -> %v\n", idWord[id], idWord[nextId])
				}
			}
		}

	}
	// dfs 解析path 得到结果
	//stack := []int{endId}
	//dfsMultiPathFn(paths, idWord, beginId, endId, 0, &stack, &res)
	return dfsMultiPathFn2(paths, idWord, beginId, endId)
}

func dfsMultiPathFn2(paths map[int]map[int]bool, idWord map[int]string, beginId, curId int) [][]string {
	if curId == beginId {
		return [][]string{{idWord[curId]}}
	}
	var res [][]string
	if li, has := paths[curId]; has {
		for precursorId := range li {
			pathRes := dfsMultiPathFn2(paths, idWord, beginId, precursorId)
			for _, v := range pathRes {
				res = append(res, append(v, idWord[curId]))
			}
		}
	}
	return res
}

func isDiffOneCharacter(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	cnt := 0
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			cnt++
		}
		if cnt > 1 {
			return false
		}
	}
	return cnt == 1
}

func dfsMultiPathFn(paths map[int]map[int]bool, idWord map[int]string, beginId, curId, depth int, stack *[]int, res *[][]string) {
	if curId == beginId {
		size := len(*stack)
		path := make([]string, 0, size)
		for i := len(*stack) - 1; i >= 0; i-- {
			path = append(path, idWord[(*stack)[i]])
		}
		*res = append(*res, path)
		return
	}

	if li, has := paths[curId]; has {
		for precursorId := range li {
			*stack = append(*stack, precursorId)
			dfsMultiPathFn(paths, idWord, beginId, precursorId, depth+1, stack, res)
			*stack = (*stack)[:len(*stack)-1]
		}
	}
}

func Test_findLadders(t *testing.T) {
	res := findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	//res := findLadders("a", "c", []string{"a", "b", "c"})
	//res := findLadders("red", "tax", []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"})
	t.Logf("%v\n", res)
}
