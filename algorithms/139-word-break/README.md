# 139. 单词拆分

给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

说明：

- 拆分时可以重复使用字典中的单词。
- 你可以假设字典中没有重复的单词。

示例 1：

```
输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
```

示例 2：

```
输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
     注意你可以重复使用字典中的单词。
```

示例 3：

```
输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false
```

来源：力扣（LeetCode）

链接：https://leetcode-cn.com/problems/word-break

著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

## 解题思路（动态规划）

这道题使用动态规划进行解答，思路不是很直观。

对于长度为`n`的字符串`s`，将问题分成`n`个子问题，我们使用一个长度为`n+1`的布尔数组`res`，`res[0]`设置为`true`，从`1`开始，分别求`s[1]`, `s[1..2]`, `s[1..3]`, `s[1..n]`的解。

这样分解后子问题符合动态规划的多阶段决策特征，而且包含重叠子问题，下面是具体思路：

假设当前到达长度为`i`的子问题，也就是求`s[1..i]`能否符合要求时，我们可以这样做：

```
遍历所给的字典 wordDict，假设每次循环获得的单词为 word，单词长度为lenOfWord，

我们只需要判断两个条件：

1. res[i-lenOfWord]是否符合条件，也就是切除尾部的word后，前面的字符串是否满足要求
2. word是否等于s[i-lenOfWord:i]

那么有下面的公式成立：

res[i] = res[i-lenOfWord] && word == s[i-lenOfWord:i]

```

采用自底向上的方式，一步一步求解，最终res[n]就是我们想要的答案。

具体解答参考代码实现。
