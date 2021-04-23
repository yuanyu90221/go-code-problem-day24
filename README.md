# go-code-problem-day24

將給定的 encoded string 進行解碼 (decode)，return 解碼後的 string。

Encoded string 可表示為 k[encodedString]，解碼後所得到的 string 為括號內的 string 重複 k 次
Constraints:

k 為用字串表示的正整數
encodedString 中，可能存在的字母範圍為 a-z 及 A-Z
Example 1:
```script==
Input: "a2[B2[c]]"
Output: "aBccBcc"
Explanation: "a2[B2[c]]" -> "a2[Bcc]" -> "aBccBcc"
```
Example 2:
```script==
Input: "a2[bc]12[d]"
Output: "abcbcdddddddddddd"
Explanation: "a2[bc]12[d]" -> "abcbc12[d]" -> "abcbcdddddddddddd"
```
## 題目分析

給定一個 string s

要把 s 根據Pattern 轉換成另一個 string r

首先根據 example 可以知道 需要把string s 根據 pattern把 轉換成 token

"a2[bc]12[d]" => {"a", "bc", "d"}, {2, 12}

然後再根據 stack count

pop up 最後的字串

可以建立兩個 stacks: 

tokenStack: 用來收集分曾的token

numberStack: 把遇到的number收集


先建立一個 inputBuffer 收集 

要放入 tokenStack的token還有要放入numberStack的number

收集rune的規則如下

遇到 number:

則停止收集, 

先確認 inputBuffer的值是否為number

若否 放入 tokenstack

若是 放入 numberstack

清空inputBuffer

遇到 [:

則停止收集, 

先確認 inputBuffer的值是否為number

若否 放入 tokenstack

若是 放入 numberstack

清空inputBuffer

遇到 ] :

則先確認 目前 inputBuffer是否有值

如果有值 

取出 numberStack 最後一筆

取出 tokenStack 最後一筆

inputBuffer之值為需要重複append的token

根據 numberStack最後一筆的數值

以及 inputBuffer的值 組出所需要append的token值 appendToken

然後把 最後的token 就是 tokenStack的最後一筆的值 concat appendToken

如果 目前以loop到最後一個字元 則 返回這個值為答案

否則存入 tokenstack

清空inputBuffer

如果沒值

取出 numberStack 最後一筆

取出 tokenStack 最後2筆

tokenStack的最後一筆之值為需要重複append的token

根據 numberStack最後一筆的數值

以及 tokenStack的最後一筆之值 組出所需要append的token值 appendToken

然後把 最後的token 就是 tokenStack的最後第2筆的值 concat appendToken

如果 目前以loop到最後一個字元 則 返回這個值為答案

否則存入 tokenstack