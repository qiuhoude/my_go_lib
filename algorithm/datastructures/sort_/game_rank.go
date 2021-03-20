package sort_

/*
游戏排行榜 如果只显示前 N 的玩家其实就是top N的问题
需求: 排行榜需要实时变动,每次获取需要自己的名次
思路:
1. 使用 数组 + map<playerId,数组index>,这样可以通过玩家id快速定位index己的index位置
2. 链表 + map , 这种时候排行榜人员变动较大的时候使用

玩家数据改变时:
在榜中
	新值 = 旧值 不变
	新值 > 旧值 往前遍历, 维护rankMap的位置并交换位置
	新值 < 旧值 时王后遍历 维护rankMap的位置并交换位置
不在榜中
	榜满员
		比较最后一名,大于最后一名才能入榜
	没有满
		添加到最后进行上浮
*/

// 玩家id
type PlayerId int64

// 排行的item
type IRankItem interface {
	GetPlayerId() PlayerId
	GetScore() int
	SetScore(newScore int)
}

// 游戏排行榜
type IGameRank interface {
	SetPlayerScore(player IRankItem, newScore int) bool // 返回false 说明没有在排行榜中做修改
	GetRankList() []IRankItem
	GetRankingById(id PlayerId) int
}

// 游戏排行榜 数组结构
type GameRankArr struct {
	size     int // 排行榜的总大小
	curSize  int // 当前大小
	rankList []IRankItem
	rankMap  map[PlayerId]int // <playerId,index>
}

func NewGameRankArr(sz int) *GameRankArr {
	return &GameRankArr{
		size:     sz,
		curSize:  0,
		rankList: make([]IRankItem, sz),
		rankMap:  make(map[PlayerId]int, sz),
	}
}

/*
数组玩家积分
*/
func (g *GameRankArr) SetPlayerScore(player IRankItem, newScore int) bool {
	if index, ok := g.rankMap[player.GetPlayerId()]; ok { // 已经在排行榜中存在
		oldVal := g.rankList[index].GetScore()
		player.SetScore(newScore)
		if newScore > oldVal {
			g.siftUp(index)
			return true
		} else if newScore < oldVal {
			g.siftDown(index)

			return true
		} else {
			// 不变无操作
			return false
		}
	} else { // 不存在榜中
		player.SetScore(newScore)
		if g.curSize >= g.size { // 榜满员
			// 是否能入榜
			if g.rankList[g.curSize-1].GetScore() < newScore {
				// 放到尾部进行上浮
				g.rankList[g.curSize-1] = player
				g.siftUp(g.curSize - 1)
				return true
			} else { //不能入榜
				return false
			}
		} else { //没有满
			g.curSize++
			g.rankList[g.curSize-1] = player
			g.rankMap[player.GetPlayerId()] = g.curSize - 1
			g.siftUp(g.curSize - 1)
			return true
		}
	}
}

// 向下移动, 下面的分数低的
func (g *GameRankArr) siftDown(index int) {
	if index >= g.curSize-1 {
		return
	}
	cur := index
	next := cur + 1
	for next < g.size {
		if g.rankList[cur].GetScore() < g.rankList[next].GetScore() { // 如果积分相同 后掉榜的人的在先掉榜的人的前面
			// 先维护map 后 交换位置
			g.rankMap[g.rankList[cur].GetPlayerId()], g.rankMap[g.rankList[next].GetPlayerId()] = next, cur
			g.rankList[cur], g.rankList[next] = g.rankList[next], g.rankList[cur]
			cur = next
			next++
		} else {
			break
		}
	}
}

// 向上移动 , 上面时分数高的
func (g *GameRankArr) siftUp(index int) {
	if index <= 0 {
		return
	}
	cur := index
	pre := cur - 1
	for pre >= 0 {
		if g.rankList[cur].GetScore() > g.rankList[pre].GetScore() { // 积分相同 后上榜的人在先上榜人的后面
			g.rankMap[g.rankList[cur].GetPlayerId()], g.rankMap[g.rankList[pre].GetPlayerId()] = pre, cur
			g.rankList[cur], g.rankList[pre] = g.rankList[pre], g.rankList[cur]
			cur = pre
			pre--
		} else {
			break
		}
	}
}

func (g *GameRankArr) GetRankList() []IRankItem {
	return g.rankList[0:g.curSize]
}

func (g *GameRankArr) GetRankingById(id PlayerId) int {
	if index, ok := g.rankMap[id]; ok {
		return index + 1
	}
	return 0
}
