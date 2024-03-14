## 383.Ransom Note

[https://leetcode.com/problems/ransom-note/](https://leetcode.com/problems/ransom-note/)

*Easy*

Given two strings `ransomNote` and `magazine`, return `true`* if * `ransomNote`* can be constructed by using the letters from * `magazine`* and * `false`* otherwise* .

Each letter in `magazine` can only be used once in `ransomNote`.

**Example 1:**

<pre><strong>Input:</strong> ransomNote = "a", magazine = "b"
<strong>Output:</strong> false
</pre>

**Example 2:**

<pre><strong>Input:</strong> ransomNote = "aa", magazine = "ab"
<strong>Output:</strong> false
</pre>

**Example 3:**

<pre><strong>Input:</strong> ransomNote = "aa", magazine = "aab"
<strong>Output:</strong> true
</pre>

**Constraints:**

* `1 <= ransomNote.length, magazine.length <= 10<sup>5</sup>`
* `ransomNote` and`magazine` consist of lowercase English letters.
