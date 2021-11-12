# LeetCode 1775. 通过最少操作次数使数组的和相等

由于只出现正整数`[1, 6]`, 辣么明显先用`计数排序`处理, 得到数组和的表达式

$$
sum_1 = \begin{bmatrix}
a_0 \\ a_1 \\ a_2 \\ a_3 \\ a_4 \\ a_5
\end{bmatrix} \cdot \begin{bmatrix}
1 \\ 2 \\ 3 \\ 4 \\ 5 \\ 6
\end{bmatrix}^{-1}

sum_2 = \begin{bmatrix}
b_0 \\ b_1 \\ b_2 \\ b_3 \\ b_4 \\ b_5
\end{bmatrix} \cdot \begin{bmatrix}
1 \\ 2 \\ 3 \\ 4 \\ 5 \\ 6
\end{bmatrix}^{-1}
$$

朴素观察, 则题意为通过改变$a_i$ $b_i$, 使得两数组和相等, 即

$$
\begin{cases}
sum_1 = sum_2 \\
\sum_{i=0}^5 a_i = C_1 \\
\sum_{i=0}^5 b_i = C_2
\end{cases}
$$

$C_1$ $C_2$均为常量, 通过调整$a_i$ $b_i$使得$sum_1$ $sum_2$相等

假设初始状态下 $sum_1 < sum_2$

由于每次值变化最终是单调函数, 则$a_i$只向$a_{[i+1, 5]}$转移, 同理$b_i$只向$b_{[i+1, 5]}$转移

更进一步, 同时观察$a_i$ $b_i$, 变化值`v`如下

$$
v_1 = \begin{bmatrix}
a_i \rightarrow a_{i+5} \\
a_i \rightarrow a_{i+4} \\
a_i \rightarrow a_{i+3} \\
a_i \rightarrow a_{i+2} \\
a_i \rightarrow a_{i+1} \\
\end{bmatrix} \cdot \begin{bmatrix}
5 \\ 4 \\ 3 \\ 2 \\ 1
\end{bmatrix}^{-1}

v_2 = \begin{bmatrix}
b_i \rightarrow b_{i-5} \\
b_i \rightarrow b_{i-4} \\
b_i \rightarrow b_{i-3} \\
b_i \rightarrow b_{i-2} \\
b_i \rightarrow b_{i-1} \\
\end{bmatrix} \cdot \begin{bmatrix}
-5 \\ -4 \\ -3 \\ -2 \\ -1
\end{bmatrix}^{-1}
$$

在$|sum_2 - sum_1|$减少$n$时, $a_i \rightarrow a_{i+n}$等价$b_i \rightarrow b_{i-n}$

即$a_0 \rightarrow a_5$等价$b_5 \rightarrow b_0$

合并$a_i$与$b_i$表达式

$$
\begin{bmatrix}
c_0 \\ c_1 \\ c_2 \\ c_3 \\ c_4 \\ c_5
\end{bmatrix} = \begin{bmatrix}
a_0 \\ a_1 \\ a_2 \\ a_3 \\ a_4 \\ a_5
\end{bmatrix} + \begin{bmatrix}
b_5 \\ b_4 \\ b_3 \\ b_2 \\ b_1 \\ b_0
\end{bmatrix} = \begin{bmatrix}
a_0 + b_5 \\ a_1 + b_4 \\ a_2 + b_3 \\ a_3 + b_2 \\ a_4 + b_1 \\ a_5 + b_0
\end{bmatrix}
$$

问题转换为通过削减 $c_i$ 使 $|sum_2 - sum_1|$ 接近0, 则设$c_i$到最终值为$c'_i$

$$
\begin{bmatrix}
c'_0 \\ c'_1 \\ c'_2 \\ c'_3 \\ c'_4 \\ c'_5
\end{bmatrix} \cdot \begin{bmatrix}
5 \\ 4 \\ 3 \\ 2 \\ 1 \\ 0
\end{bmatrix} = |sum_2 - sum_1|
$$

$c'_i$限制条件如下

$$
c'_i \in \begin{cases}
[0, c_0] &i=0 \\
[0, c_i+\sum_0^{i-1}c'_i] &i\gt 0
\end{cases}
$$

最终结果

$$
ans = min(\sum_{i=0}^4 c'_i)
$$

接着`贪心`, 为使总步骤数最小, 每次尽可能削减较大的值, 从5到1依次遍历, 就能得出结果了
