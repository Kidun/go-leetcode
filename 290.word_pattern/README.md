## 290.Word Pattern

[https://leetcode.com/problems/word-pattern/](https://leetcode.com/problems/word-pattern/)

*Easy*

Given a `pattern` and a string `s`, find if `s` follows the same pattern.

Here **follow** means a full match, such that there is a bijection between a letter in `pattern` and a **non-empty** word in `s`.

**Example 1:**

<pre><strong>Input:</strong> pattern = "abba", s = "dog cat cat dog"
<strong>Output:</strong> true
</pre>

**Example 2:**

<pre><strong>Input:</strong> pattern = "abba", s = "dog cat cat fish"
<strong>Output:</strong> false
</pre>

**Example 3:**

<pre><strong>Input:</strong> pattern = "aaaa", s = "dog cat cat dog"
<strong>Output:</strong> false
</pre>

**Constraints:**

* `1 <= pattern.length <= 300`
* `pattern` contains only lower-case English letters.
* `1 <= s.length <= 3000`
* `s` contains only lowercase English letters and spaces`' '`.
* `s`**does not contain** any leading or trailing spaces.
* All the words in`s` are separated by a**single space** .
