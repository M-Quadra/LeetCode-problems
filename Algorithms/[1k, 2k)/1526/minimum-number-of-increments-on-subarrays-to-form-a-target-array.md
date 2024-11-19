# LeetCode 1526 形成目标数组的子数组最少增加次数

题目描述简单易懂, 一开始以为和之前的打印机类似, 思考2s后发觉不是

一发暴力直接超时, 开始考虑分情况讨论

## len = 0

opt = 0

## len = 1

opt = ary[0]

## len = 2

opt = max(ary[0], ary[1])

## len = 3

为了方便说明, 下面使用min, mid, max 三个数表示

- [min, mid, max]

单调递增

[min, mid, max] -> [min, max], opt += 0

- [max, mid, min]

单调递减

[max, mid, min] -> [max, min], opt += 0

- [min, max, mid]

上凸

[min, max, mid] -> [min, mid], opt += max-mid

- [min, mid, max]

上凹

[min, mid, max] -> [min, max], opt += mid-min

## 合并情况

单调: 保留首尾

非单调: 保留首尾, 增加步骤

此时已经可以用`滑动窗口`在O(n)时间复杂度解决了

## 更进一步

虽然已经可以解决了, 但这并不优雅

每次保留首尾, 在$[a_i, a_{i+1}, a_{i+2}]$情况下被操作元素都是$a_1$与$mid(a_i, a_{i+1}, a_{i+2})$

随着遍历进度不断移动, $a_i$始终为定值, 最终结果可以改写如下

$$
opt = max(a_0, a_{n-1}) + \sum_{i=1}^{n-2}|a_i - mid(a_0, a_i, a_{i+1})|
$$

如果要稍微往下一点点, 还可以对$|a_i - mid(a_0, a_i, a_{i+1})|$讨论

此时有

$$
[a_0, a_i, a_{i+1}]
$$

$$
|a_i - mid(a_0, a_i, a_{i+1})| = \begin{cases}
0 & (a_0-a_i)(a_i-a_{i+1}) >= 0 \\
a_i - max(a_0, a_{i+1}) & (a_0-a_i)(a_i-a_{i+1})<0, a_0-a_i<0 \\
min(a_0, a_{i+1}) - a_i & (a_0-a_i)(a_i-a_{i+1})<0, a_0-a_i>0
\end{cases}
$$

然后...然后以我的羸弱数学水平就推导不出最优解了_(ˊཀˋ」∠)_

# DP

上面走到了死胡同, 此时必须换一种思路

用$dp_i$表示第i个数需要增加的操作数, 则

$$
dp_i = \begin{cases}
a_0 &i=0 \\
max(a_i - a_{i-1}, 0) &i>0 \\
\end{cases} \\
$$

最终结果

$$
opt = \sum_{i=0}^{n-1}dp_i
$$
