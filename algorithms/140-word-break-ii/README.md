# 140. 单词拆分 II

给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，在字符串中增加空格来构建一个句子，使得句子中所有的单词都在词典中。返回所有这些可能的句子。

说明：

- 分隔时可以重复使用字典中的单词。
- 你可以假设字典中没有重复的单词。

示例 1：

```
输入:
s = "catsanddog"
wordDict = ["cat", "cats", "and", "sand", "dog"]
输出:
[
  "cats and dog",
  "cat sand dog"
]
```

示例 2：

```
输入:
s = "pineapplepenapple"
wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
输出:
[
  "pine apple pen apple",
  "pineapple pen apple",
  "pine applepen apple"
]
解释: 注意你可以重复使用字典中的单词。
```

示例 3：

```
输入:
s = "catsandog"
wordDict = ["cats", "dog", "sand", "and", "cat"]
输出:
[]
```

来源：力扣（LeetCode）

链接：https://leetcode-cn.com/problems/word-break-ii

著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

## 解题思路（动态规划）

这道题是题目[139. 单词拆分](https://github.com/iceyang/leetcode/tree/master/algorithms/139-word-break)的困难版本，建议在完成前者的前提下再查看本题解。

本题目的复杂点在于，我们需要找出所有可能的解。在理解题目「139」的前提下，我们做出进一步改动。

我们依旧使用动态规划解题，将问题拆分成`n`个子问题，从`1`开始，分别求`s[1]`, `s[1..2]`, `s[1..3]`, `s[1..n]`的解。

之前的题目中，我们采用一个长度为`n`的布尔数组来保存中间状态，根据题目的需要，我们换成一个长度为`n`的二维数组`res`。

对于`res[i]`，我们有如下定义：

```
res[i]是一个字符串数组，其中存放的字符串，假设为res[i][j]，满足：
res[i][j] 与 res[i - len(res[i][j])] 数组中的字符串可以构成子问题 s[1..i] 的解
```

打个比方，我们的输入：

```
s = applepenapple
wordDict = ["apple", "pen", "applepen"]
```

那么，`res`最终的结果为：

```
[[] [] [] [] [] ["apple"] [] [] ["pen", "applepen"] [] [] [] [] ["apple"]]
```

当`i`等于8，`j`等于0时，

```
res[i][j] = res[8][0] = "pen"

res[i - len(res[i][j])] = res[i - len(res[8][0])] = ["apple"]
```

`"pen"`与`["apple"]`可以构成子问题`s[1..i]`也就是`applepen`的解。

最终题目的输出，需要我们根据`res`进行回溯计算，具体过程详见代码实现。
