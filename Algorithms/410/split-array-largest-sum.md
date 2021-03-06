# LeetCode 410 分割数组的最大值

## SPFA

> 以我低劣算法水平当然不首先想出最优解

在理想情况下数组被分成m份, 最大值即平均数

$$
avg = \frac{\sum_{i=1}^n a_i}{m}
$$

因此每段子数组的分割位置

$$
avg \approx \sum_{j =st_i}^{ed_i} a_j
$$

考虑到实际情况中, 大概率是要出现高于`avg`的情况, 此时引入`max[ed][m]`描述从`[1, ed]`分割`m`时的最大子数组和

如果`max[ed][m] > avg`, 则优先满足`max[ed][m]`, 使后续分割时的拆解更加零散

当然, 向后满足`max[ed][m]`这个极大值还不够完善, 之前的数据是优先满足`avg`而不是极大值。当极大值发生更新的时, 需要追溯旧数据更新

加入优先队列可以缩短搜索次数, 第一次搜索终点即最优解

## 优先队列

> 别问, 问就是头铁

在上一步中, 每次完整的遍历都需要执行多次二分搜索与回溯

如果合并遍历步骤达到`O(n)`, 将回溯与极大值压缩为独立的遍历, 不再拆解为`SPFA`的单步搜索

那么在只保留优先队列的前提下, 可以精简单步搜索的步骤

为了再缩小一点点遍历次数, 之前的起点设置在`avg`附近, 如果再考虑上`n == m`的情况, 则起点如下

$$
max (\frac{\sum_{i=1}^n a_i}{m}, max(a_i))
$$

通过失败的遍历结果可以不断缩小最小值范围

## 二分搜索

> 稳定打倒一切...

从上我们知晓了优先队列的搜索范围为

$$
[max (\frac{\sum_{i=1}^n a_i}{m}, max(a_i)), \sum_{i=1}^n a_i]
$$

等等, 在某个连续区间中寻找某个符合条件的最小值, 自然就是二分了

优先队列根据每次发现的最小值遍历, 虽然部分情况下遍历次数低于二分搜索, 但终究没有辣么稳定\_(:з」∠)_
