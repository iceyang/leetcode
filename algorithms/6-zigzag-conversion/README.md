# 6. Z 字形变换
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为`"LEETCODEISHIRING"`行数为 3 时，排列如下：

```
L   C   I   R
E T O E S I I G
E   D   H   N
```

之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如：`"LCIRETOESIIGEDHN"`。

请你实现这个将字符串进行指定行数变换的函数：

```
string convert(string s, int numRows);
```

示例 1:

```
输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"
```

示例 2:

```
输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G
```

来源：力扣（LeetCode）

链接：https://leetcode-cn.com/problems/zigzag-conversion

著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

## 解题思路

### 二维数组

将过程存放在二维数组中，一维代表行，二维代表列，想象一下跟Z字形对应起来。

那么关键点就在于如何计算行和列。我采取的办法是：

让`numRows`代表Z字符的行数，`row`代表行，`column`代表列，行列从`0`开始计算，行跟列会慢慢增加。行特殊点在于，会跟随字符串的推进，作Z字形变动（其实也就是后退）。

判断条件是：

- 当`column % (numRows-1)`等于0时，该列位于直线上，行加1
- 当`column % (numRows-1)`不等于0时，该列位于斜线上，行减1
- 当`row`等于`numRows`时，说明Z字形的直线填充满了，列加1，行退后一格

当整个计算完毕后，将计算结果按行拼起来就好了。

### 字符串数组

计算思路与前一个方法一样，经过查看分析，其实二维数组可以看成是一维的string列表，那么每次得到的结果，都保存在对应行的string中即可，最后将string拼接起来得到结果。
