## Container with most water
<p class="has-line-data" data-line-start="0" data-line-end="1">You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).</p>
<p class="has-line-data" data-line-start="2" data-line-end="3">Find two lines that together with the x-axis form a container, such that the container contains the most water.</p>
<p class="has-line-data" data-line-start="4" data-line-end="5">Return the maximum amount of water a container can store.</p>
<p class="has-line-data" data-line-start="6" data-line-end="7">Notice that you may not slant the container.</p>
<p class="has-line-data" data-line-start="10" data-line-end="11">Example 1:</p>
<p class="has-line-data" data-line-start="13" data-line-end="17">Input: height = [1,8,6,2,5,4,8,3,7]<br>
Output: 49<br>
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.<br>
Example 2:</p>
<p class="has-line-data" data-line-start="18" data-line-end="20">Input: height = [1,1]<br>
Output: 1</p>
<p class="has-line-data" data-line-start="22" data-line-end="23">Constraints:</p>
<p class="has-line-data" data-line-start="24" data-line-end="27">n == height.length<br>
2 &lt;= n &lt;= 105<br>
0 &lt;= height[i] &lt;= 104</p>