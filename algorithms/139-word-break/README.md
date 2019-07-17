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

对于长度为n的字符串s，将问题分成n个子问题，我们使用一个长度为n+1的布尔数组res，res[0]设置为true，从1开始，分别求s[1], s[1..2], s[1..3], s[1..n]的解。

仔细思考一下，这样分解后子问题符合动态规划的多阶段决策特征，而且包含重叠子问题。

采用自底向上的方式，一步一步求解，最终res[n]就是我们想要的答案。

具体解答参考代码实现。
