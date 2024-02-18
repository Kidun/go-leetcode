## 203.Remove Linked List Elements

[https://leetcode.com/problems/remove-linked-list-elements/](https://leetcode.com/problems/remove-linked-list-elements/)

*Easy*

Given the `head` of a linked list and an integer `val`, remove all the nodes of the linked list that has `Node.val == val`, and return *the new head* .

**Example 1:**

<pre><strong>Input:</strong> head = [1,2,6,3,4,5,6], val = 6

<strong>Output:</strong> [1,2,3,4,5]
</pre>

**Example 2:**

<pre><strong>Input:</strong> head = [], val = 1
<strong>Output:</strong> []
</pre>

**Example 3:**

<pre><strong>Input:</strong> head = [7,7,7,7], val = 7
<strong>Output:</strong> []
</pre>

**Constraints:**

* The number of nodes in the list is in the range`[0, 10<sup>4</sup>]`.
* `1 <= Node.val <= 50`
* `0 <= val <= 50`
