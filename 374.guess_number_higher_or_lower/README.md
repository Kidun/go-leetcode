## 374.Guess Number Higher or Lower

[https://leetcode.com/problems/guess-number-higher-or-lower/](https://leetcode.com/problems/guess-number-higher-or-lower/)

*Easy*

We are playing the Guess Game. The game is as follows:

I pick a number from `1` to `n`. You have to guess which number I picked.

Every time you guess wrong, I will tell you whether the number I picked is higher or lower than your guess.

You call a pre-defined API `int guess(int num)`, which returns three possible results:

* `-1`: Your guess is higher than the number I picked (i.e. `num > pick`).
* `1`: Your guess is lower than the number I picked (i.e. `num < pick`).
* `0`: your guess is equal to the number I picked (i.e. `num == pick`).

Return *the number that I picked* .

**Example 1:**

<pre><strong>Input:</strong> n = 10, pick = 6
<strong>Output:</strong> 6
</pre>

**Example 2:**

<pre><strong>Input:</strong> n = 1, pick = 1
<strong>Output:</strong> 1
</pre>

**Example 3:**

<pre><strong>Input:</strong> n = 2, pick = 1
<strong>Output:</strong> 1
</pre>

**Constraints:**

* `1 <= n <= 2<sup>31</sup> - 1`
* `1 <= pick <= n`
