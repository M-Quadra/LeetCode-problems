# LeetCode 1425. 带限制的子序列和

考虑拼凑子序列时的状态转移方向, 定义 $dp_i$ 为符合限制且最后一个元素为 $n_i$ 的子序列和。

如果从 i 向 [i+1, i+k] 转移, $O(n^k)$ 显然不行。

考虑 [i-k, i-1] 向 i 转移, 方程如下:

$$\begin{aligned}
max(x, y) = n_i &, i \in [x, y] \\
n_i \ge n_j &, j \in [x, y]
\end{aligned}$$

$$
dp_i = \begin{cases}
max(i-k, i-1) + n_i &, max(i-k, i-1)+n_i \ge n_i \\
n_i &, n_i \ge max(i-k, i-1)+n_i
\end{cases}
$$

下一步求 dp[i-k, i-1] 区间最大值。线段树, 堆, 单调双端队列, 都是可以的。
