# LeetCode 962. 恰好移动 k 步到达某一位置的方法数目

假设 `startPos` <= `endPos`, 若满足恰好`k`步到达, 则最终执行`a`次+1, `b`次-1。

此时

$$
\left\{
\begin{array}{ll}
k &= a + b \\
k &= (pos_{start} - pos_{end})*2 + b
\end{array}
\right.
$$

忽略可行判断, 则等价组合数问题:

给定a个0与b个1, a+b=k, 能够组合出多少种不同数?

为个人看组合数方便, 定义:

$$
\begin{aligned}
n' &= k \\
k' &= pos_{start} - pos_{end}
\end{aligned}
$$

则答案为:

$$
\begin{aligned}
ans &= \binom{n'}{k'} \\
    &= \frac{n'!}{k'!(n'-k')!} \\
    &= n'!/k'!/(n'-k')!
\end{aligned}
$$

有除法, 上逆元。

- 时间复杂度: $O(\log{N})$
- 空间复杂度: $O(N)$