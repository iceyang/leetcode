# 96. 不同的二叉搜索树

给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？

示例:

```
输入: 3
输出: 5
解释:
给定 n = 3, 一共有 5 种不同结构的二叉搜索树:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
```

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-binary-search-trees
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

## 解题思路

给定`n`，根据序列`1...n`构建二叉搜索树，我们可以将问题分解成：

```
指定i从1开始遍历，以i作为根，1...(i-1)构成左子树，(i+1)...n构成右子树。
那么当前i的组合 = 左子树组合数 * 右子树组合数，即笛卡尔积。
```

那么给定`n`的二叉搜索树的数量为所有i的子树数量之和。

又因为该问题符合动态规划的特性，我们可以避免子问题的重复计算，采用自底向上的方法解决问题。

其他详见解题代码。