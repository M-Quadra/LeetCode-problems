# LeetCode 375 猜数字大小 II

如果使用一维dp, 只记录不同n的花费在n>7还是8来着就会因为被中途选择次数不同导致花费不同, 进而得到错误的结果

此时就需要使用二维dp, 加入位置信息后得到数组`payAry[st][len]`

st 为当前区间的起点, len 为区间长度

len=1时, 由于最后一次选择是白给的, 所以归零, 归零, 归归归零

得到最初的dp组

 len\st | 1 | 2 | 3 | 4 | 5 | ... | n
:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:
  1  | 0   | 0 | 0 | 0 | 0 | ... | 0
  2  | ? | ? | ? |? | ?
  3  | ? | ? | ? | ?
  4  | ? | ? | ?
  5  | ? | ?
 ... | ...
  n  | x

对于剩下的每个`payAry[st][len]`, 状态转移方程如下

因为太长, 下面用`dp`指代`payAry`\_(:з」∠)_

对于每一个可选区间`[st, st + len)`即`dp[st][len]`, 满足:

$$
\left\{\begin{matrix}
pay(i) = i + max(dp[st][i - st], dp[i+1][ed - i]) \\
st \leq i \leq n \\
st + len \leq n \\
dp[st][len] \leq pay(i)
\end{matrix}\right.
$$

将状态转移方程带入`payAry`, 得到最终结果`payAry[1][n]`
