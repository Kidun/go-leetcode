## 326.Power of Three

[https://leetcode.com/problems/power-of-three/](https://leetcode.com/problems/power-of-three/)

*Easy*

Given an integer `n`, return *`true` if it is a power of three. Otherwise, return `false`* .

An integer `n` is a power of three, if there exists an integer `x` such that `n == 3<sup>x</sup>`.

**Example 1:**

<pre><strong>Input:</strong> n = 27
<strong>Output:</strong> true
<strong>Explanation:</strong> 27 = 3<sup>3</sup>
</pre>

**Example 2:**

<pre><strong>Input:</strong> n = 0
<strong>Output:</strong> false
<strong>Explanation:</strong> There is no x where 3<sup>x</sup> = 0.
</pre>

**Example 3:**

<pre><strong>Input:</strong> n = -1
<strong>Output:</strong> false
<strong>Explanation:</strong> There is no x where 3<sup>x</sup> = (-1).
</pre>

**Constraints:**

* `-2<sup>31</sup> <= n <= 2<sup>31</sup> - 1`

**Follow up:** Could you solve it without loops/recursion?
