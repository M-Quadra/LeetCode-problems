# 面试题 16.03 交点

原本我想使用无限叠`if-else`的方式估计能解决, 但这样总归是不够优雅。为了跑路自然得强迫自己动脑

使用`点斜式`在斜率上可能存在问题, 为了少写判断, 使用`一般式`描述直线。并且将坐标带入`一般式`就能直观判断点与直线的位置关系, 以此判断线段的位置关系

完成关系判断后, 剩下的事就相当于求直线交点, 直接解方程又得判断条件, 不行, 这很麻烦...

搜索? 二分x轴或y轴？同时考虑xy又得加条件, 太累了

既然给定目标与样本...等等, 机器学习? 线性回归?

反向传播...不对, 题目的条件明确(样本完全), 可以直接求最优解(xy)

接下来就是转化成矩阵形式

$$
\begin{cases}
a_1x + b_1y+c_1 = 0 \\
a_2x + b_2y+c_2 = 0
\end{cases}
$$

$$
\begin{cases}
a_1x + b_1y = -c_1 \\
a_2x + b_2y = -c_2
\end{cases}
$$

$$
\begin{pmatrix}
x&y
\end{pmatrix}
\begin{pmatrix}
a_1&a_2\\
b_1&b_2
\end{pmatrix}
=
\begin{pmatrix}
-c_1&-c_2
\end{pmatrix} \\
\begin{pmatrix}
x&y
\end{pmatrix}
\begin{pmatrix}
a_1&a_2\\
b_1&b_2
\end{pmatrix}
\begin{pmatrix}
a_1&a_2\\
b_1&b_2
\end{pmatrix}^{-1}
=
\begin{pmatrix}
-c_1&-c_2
\end{pmatrix}
\begin{pmatrix}
a_1&a_2\\
b_1&b_2
\end{pmatrix}^{-1} \\
\begin{pmatrix}
x&y
\end{pmatrix}
=
\begin{pmatrix}
-c_1&-c_2
\end{pmatrix}
\begin{pmatrix}
a_1&a_2\\
b_1&b_2
\end{pmatrix}^{-1}
$$

接下来就是二阶矩阵求逆矩阵

$$
\begin{pmatrix}
a&b\\
c&d
\end{pmatrix}^{-1}
=
\frac{1}{det}
\begin{pmatrix}
d&-b\\
-c&a
\end{pmatrix} \\
\begin{pmatrix}
a&b\\
c&d
\end{pmatrix}^{-1}
=
\frac{1}{ad - bc}
\begin{pmatrix}
d&-b\\
-c&a
\end{pmatrix}
$$

由于在`一般式`阶段已经可以完成一些判断, 过滤特殊情况, 所以行列式`det`已经不为`0`, 直接计算逆矩阵带入后就能得到交点(x, y)

\_(:з」∠)_

如果从行列式角度直接讨论无解/多解/唯一解, 似乎需要线性变换, 太麻烦...不管了