## 278.First Bad Version

[https://leetcode.com/problems/first-bad-version/](https://leetcode.com/problems/first-bad-version/)

*Easy*

You are a product manager and currently leading a team to develop a new

product. Unfortunately, the latest version of your product fails the
quality check. Since each version is developed based on the previous
version, all the versions after a bad version are also bad.

Suppose you have `n` versions `[1, 2, ..., n]` and you want to find out the first bad one, which causes all the following ones to be bad.

You are given an API `bool isBadVersion(version)` which returns whether `version` is bad. Implement a function to find the first bad version. You should minimize the number of calls to the API.

**Example 1:**

<pre><strong>Input:</strong> n = 5, bad = 4
<strong>Output:</strong> 4
<strong>Explanation:</strong>
call isBadVersion(3) -> false
call isBadVersion(5) -> true
call isBadVersion(4) -> true
Then 4 is the first bad version.
</pre>

**Example 2:**

<pre><strong>Input:</strong> n = 1, bad = 1
<strong>Output:</strong> 1
</pre>

**Constraints:**

* `1 <= bad <= n <= 2<sup>31</sup> - 1`
