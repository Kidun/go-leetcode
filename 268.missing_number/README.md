## 268.Missing Number

[https://leetcode.com/problems/missing-number/](https://leetcode.com/problems/missing-number/)

*Easy*

Given an array `nums` containing `n` distinct numbers in the range `[0, n]`, return *the only number in the range that is missing from the array.*

**Example 1:**

<pre><strong>Input:</strong> nums = [3,0,1]
<strong>Output:</strong> 2
<strong>Explanation:</strong> n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the missing number in the range since it does not appear in nums.
</pre>

**Example 2:**

<pre><strong>Input:</strong> nums = [0,1]
<strong>Output:</strong> 2
<strong>Explanation:</strong> n = 2 since there are 2 numbers, so all numbers are in the range [0,2]. 2 is the missing number in the range since it does not appear in nums.
</pre>

**Example 3:**

<pre><strong>Input:</strong> nums = [9,6,4,2,3,5,7,0,1]
<strong>Output:</strong> 8
<strong>Explanation:</strong> n = 9 since there are 9 numbers, so all numbers are in the range [0,9]. 8 is the missing number in the range since it does not appear in nums.
</pre>

**Constraints:**

* `n == nums.length`
* `1 <= n <= 10<sup>4</sup>`
* `0 <= nums[i] <= n`
* All the numbers of`nums` are**unique** .

**Follow up:** Could you implement a solution using only `O(1)` extra space complexity and `O(n)` runtime complexity?
