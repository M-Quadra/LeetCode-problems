# LeetCode 1191 K 次串联后最大子数组之和

贪心

为了方便说明, 定义几个变量。对每个数组`arr`而言:

`maxSub`: 最大子数组之和  
`maxSuf`: 最大后缀数组之和  
`maxPre`: 最大前缀数组之和

$$
opt =
\begin{cases}
maxSub &k=1 \\
max(maxSuf + s_{arr}*(k-2) + maxPre, maxSub) &k>=2
\end{cases} \\

s_{arr} = max(0, sum(arr))
$$

`maxSuf`与`maxPre`均可以通过遍历累加解决

`maxSub`也是通过遍历选取, 不同在于每次都需要贪心选择局部最大值

$$
tmpSub_i = \begin{cases}
max(0, arr_i) &i=0 \\
max(0, arr_i, tmpSub_{i-1} + arr_i) &i>=1
\end{cases} \\
i \in [0, len(arr)) \\
maxSub = max(tmpSub_i)
$$

伪代码如下：

```
maxSub, tmpSub = 0, 0
for i, v in arr:
    tmpSub = max(0, v, tmpSub + v)
    maxSub = max(maxSub, tmpSub)
```
