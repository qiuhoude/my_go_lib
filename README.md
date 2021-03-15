# golang学习库
学习算法 golang版本

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
- [链表][signle_linkedlist]
    + 问题
        + 判断单链表是否有环
        + 反转链表
        + 找到链表的中间节点
        + 两个有序单链表合并
        + 删除倒数第N个节点
        + 答案在 [链表结构][signle_linkedlist]文件中
- 栈
    + [**链表结构**栈][stack]
    + [**数组结构**栈](https://github.com/qiuhoude/my_go_lib/blob/main/algorithm/datastructures/stack/array_stack.go)
- 队列
    + [数组队列](https://github.com/qiuhoude/my_go_lib/blob/main/algorithm/datastructures/queue/array_queue.go)
    + [链表队列](https://github.com/qiuhoude/my_go_lib/blob/main/algorithm/datastructures/queue/linked_queue.go)
    + [循环队列](https://github.com/qiuhoude/my_go_lib/blob/main/algorithm/datastructures/queue/circular_queue.go)
    + [优先级队列][priority_queue] 使用最小堆实现
- 排序
    + [冒泡排序][sort_test] (两两比较数组中的值,不满足就换位置)
    + [插入排序][sort_test] (分区思想,将未排序的区域第一个值插入到已排序区域中)
    + [选择排序][sort_test] (分区思想,每次在未排序区域中找到最小值,放到排序区域的最后)
    + [归并排序][sort_test] (分治思想,把数组从中间分成前后两部分，然后对前后两部分分别排序，再将排好序的两部分合并在一起，这样整个数组就都有序了)
    + [快速排序][sort_test]
    + [如何在 O(n) 的时间复杂度内查找一个无序数组中的第 K 大元素?][sort_test]
- [二分查找][binary_search] 似于数学中的夹逼定理,两边不断逼近某个值
    + 二分法求平方根
    + 查找第一个值等于给定值的元素(有重复元素的切片)
    + 查找最后一个值等于给定值的元素
    + 查找第一个大于等于给定值的元素
    + 查找第一个大于给定值的元素
    + 查找最后一个小于等于给定值的元素
- [权重随机算法(游戏中很常用)][weighted]
    + 权重随机普通算法 只要随机的数量不多可以进行 O(n)级别
    + 优化版: 使用计数排序,后通过二分查找 O(logn)级别
    + 优化版2: 使用 Alias method 将时间复杂度降到 O(1)级别, 应用于有项比较多的情况
    
    


[signle_linkedlist]: https://github.com/qiuhoude/my_go_lib/blob/main/algorithm/datastructures/linkedlist/signle_linkedlist.go
[stack]: https://github.com/qiuhoude/my_go_lib/blob/main/algorithm/datastructures/stack/stack.go
[priority_queue]: https://github.com/qiuhoude/my_go_lib/blob/main/algorithm/datastructures/queue/priority_queue.go
[sort_test]: https://github.com/qiuhoude/my_go_lib/blob/main/algorithm/datastructures/sort_/sort_test.go
[binary_search]: https://github.com/qiuhoude/my_go_lib/blob/main/algorithm/datastructures/binary_search/binary_search.go
[weighted]: https://github.com/qiuhoude/my_go_lib/blob/main/algorithm/datastructures/randweighted/randweighted.go