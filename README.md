# golang学习库

### 数据结构和算法
主要通过 极客时间-数据结构与算法之美 和 慕课网的算法课学习的算法知识, 自己整理出golang版本的基础的算法和数据结构;
基本上每个方法都是单元测试;
里面也有些自己做的leetcode的题目和课程中题目,解答思路都在代码的注释中

#### 时间复杂度

##### 量级
量级: 随数量递增而递增
公式会把常量和低次项都忽略就得到下面

- O(1)      常熟阶
- O(logn)   对数阶
- O(n)      线性阶
- O(nlogn)  对数线性阶
- O(n^2)    指数阶
- O(n!)     阶乘阶

##### 分类
- 最好情况时间复杂度(best case time complexity)
- 最坏情况时间复杂度(worst case time complexity)
- 平均情况时间复杂度(average case time complexity)
- 均摊时间复杂度(amortized time complexity) (golang的slice扩容就需要用到均摊,在扩容的那次的操作会比非扩容操作耗时)

#### 基础数据结构和算法

- [数组](https://github.com/qiuhoude/my_go_lib/blob/main/algorithm/datastructures/array/array.go)
- [lru缓存算法](https://github.com/qiuhoude/my_go_lib/blob/main/algorithm/datastructures/lru_cache/lru_cache.go)
- [链表](https://github.com/qiuhoude/my_go_lib/blob/main/algorithm/datastructures/linkedlistt/signle_linkedlist.go)
- 栈
    + [**链表结构**栈][stack]
    + [**数组结构**栈](https://github.com/qiuhoude/my_go_lib/blob/main/algorithm/datastructures/stack/array_stack.go)
    + 问题 :
        + 反转链表
        + 判断单链表是否有环
        + 找到链表的中间节点
        + 两个有序单链表合并
        + 删除倒数第N个节点
        + 答案在 [链表结构][stack]文件中
- 队列
    + [数组队列](https://github.com/qiuhoude/my_go_lib/blob/main/algorithm/datastructures/queue/array_queue.go)
    + [链表队列](https://github.com/qiuhoude/my_go_lib/blob/main/algorithm/datastructures/queue/linked_queue.go)
    + [循环队列](https://github.com/qiuhoude/my_go_lib/blob/main/algorithm/datastructures/queue/circular_queue.go)
    + [优先级队列][priority_queue]






[stack]: https://github.com/qiuhoude/my_go_lib/blob/main/algorithm/datastructures/stack/stack.go
[priority_queue]: https://github.com/qiuhoude/my_go_lib/blob/main/algorithm/datastructures/queue/priority_queue.go