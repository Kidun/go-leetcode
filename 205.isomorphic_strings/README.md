## 205.Isomorphic Strings

[https://leetcode.com/problems/isomorphic-strings/](https://leetcode.com/problems/isomorphic-strings/)

*Easy*

Given two strings `s` and `t`, *determine if they are isomorphic* .

Two strings `s` and `t` are isomorphic if the characters in `s` can be replaced to get `t`.

All occurrences of a character must be replaced with another
character while preserving the order of characters. No two characters
may map to the same character, but a character may map to itself.

**Example 1:**

<pre><strong>Input:</strong> s = "egg", t = "add"
<strong>Output:</strong> true
</pre>

**Example 2:**

<pre><strong>Input:</strong> s = "foo", t = "bar"
<strong>Output:</strong> false
</pre>

**Example 3:**

<pre><strong>Input:</strong> s = "paper", t = "title"
<strong>Output:</strong> true
</pre>

**Constraints:**

* `1 <= s.length <= 5 * 10<sup>4</sup>`
* `t.length == s.length`
* `s` and `t` consist of any valid ascii character.
