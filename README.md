# golang学习

#### 目录说明

- algorithm 主要是一些自己学习算法目录


### 算法

#### 基础算法
- [数组](https://github.com/qiuhoude/my_go_lib/blob/main/algorithm/datastructures/array/array.go)
- [lru缓存算法](https://github.com/qiuhoude/my_go_lib/blob/main/algorithm/datastructures/lru_cache/lru_cache.go)
- [链表](https://github.com/qiuhoude/my_go_lib/blob/main/algorithm/datastructures/linkedlistt/signle_linkedlist.go)
- 栈
    + [*链表结构*栈](https://github.com/qiuhoude/my_go_lib/blob/main/algorithm/datastructures/stack/stack.go)
    + [*数组结构*栈](https://github.com/qiuhoude/my_go_lib/blob/main/algorithm/datastructures/stack/array_stack.go)

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
- 平均情况时间复杂度(average case time complexity),
- 均摊时间复杂度(amortized time complexity) (切片扩容就需要用到,在扩容的那次的操作会比非扩容操作耗时)