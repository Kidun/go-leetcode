## 219.Contains Duplicate II

[https://leetcode.com/problems/contains-duplicate-ii/](https://leetcode.com/problems/contains-duplicate-ii/)

*Easy*

Given an integer array `nums` and an integer `k`, return `true` *if there are two **distinct indices** * `i`* and * `j`* in the array such that * `nums[i] == nums[j]`* and * `abs(i - j) <= k`.

**Example 1:**

<pre><strong>Input:</strong> nums = [1,2,3,1], k = 3
<strong>Output:</strong> true
</pre>

**Example 2:**

<pre><strong>Input:</strong> nums = [1,0,1,1], k = 1
<strong>Output:</strong> true
</pre>

**Example 3:**

<pre><strong>Input:</strong> nums = [1,2,3,1,2,3], k = 2
<strong>Output:</strong> false
</pre>

**Constraints:**

* `1 <= nums.length <= 10<sup>5</sup>`
* `-10<sup>9</sup> <= nums[i] `
