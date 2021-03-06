# 1040 移动石子直到连续 II

模...本来想归类到模拟的, 看了看下面的标签还是放到滑动窗口好了

- 最大移动次数

选择一边的端点石子, 每次都移动到与原位置尽可能近的位置

除了首次选择的起点, 剩余空位置的总会被填充一次石子

由此可以得到最大值的公式, 用间距之和减去首尾较短的那段间距

```
maxNum = sumGap - min(gap.front(), gap.back())
```

- 最小移动次数

从结果反推, 最终所有石子都会落在一个连续区间中

为了移动次数最小, 区间外每个石子只移动一次就到达最终位置

随着区间不断缩小, 最后区间的端点石子至少有一颗没有移动过

此时稍微转换一下问题, 就能依靠原始石子的位置选择区间范围, 区间外的石子可视为自由石子, 最终目的为尽可能填充满区间（忽略细节）

为了便于判断, 划分出2个我们需要的变量

1. `sum`: 区间外的自由石子数, 可以用来自由移动
2. `gap`: 区间内需要填充的空白位置数

在不断变化区间（移动窗口）的过程中, 再划分区间改变的几个条件

1. `sum < gap`: 自由石子数量小于空白位置, 需要额外移动一次端点石子, 完成移动需要的次数为`sum + 1`, 期望减小区间
2. `sum == gap`: 自由石子数量等于空白位置, 完成移动需要的次数恰好为为`sum`, 期望移动区间
3. `sum > gap`: 自由石子数量大于空白位置, 期望增大区间
4. `sum == 0`: 自由石子数量不足, 期望减小区间

要是不嫌烦的话, 还能对具体实现做一下优化...(懒癌发作)