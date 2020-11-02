# 算法

[互联网公司最常见的面试算法题有哪些？ - 知乎](https://www.zhihu.com/question/24964987)
 
[最长上升子序列（动态规划 + 二分查找，清晰图解）](https://leetcode-cn.com/problems/longest-increasing-subsequence/solution/zui-chang-shang-sheng-zi-xu-lie-dong-tai-gui-hua-2/)
 
[腾讯面试题：编辑距离详解](https://leetcode-cn.com/problems/edit-distance/solution/bian-ji-ju-chi-mian-shi-ti-xiang-jie-by-labuladong/)
 
[BFS使用场景-层序遍历、最短路径问题](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/solution/bfs-de-shi-yong-chang-jing-zong-jie-ceng-xu-bian-l/)

[监控二叉树](https://leetcode-cn.com/problems/binary-tree-cameras)  2020-09-22 每日1题

## 思考路径
**对一组数据进行排序**
* 有没有可能包含大量重复元素？
* 是否大部分数据距离它正确的位置很近？是否近乎有序？
* 是否数据的取值范围非常有限？比如对学生成绩排序
* 是否需要稳定排序？
* 是否是使用链表存储的？
* 数据的大小是否可以装载在内存中？

## 数据规模的概念
如果想要在1s之内解决问题：
* O(n^2)的算法可以处理大约10^4级别的数据；
* O(n)的算法可以处理大约10^8级别的数据；
* O(nlogn)的算法可以处理大约10^7级别的数据；

## 范围
* 排序算法
* 基础数据结构和算法的实现：堆，二叉树，图等；
* 基础数据结构的使用：链表，栈，队列，哈希表，图，Trie，并查集等；
* 基础算法：深度优先，广度优先，二分查找，递归；
* 基本算法思想：递归，分治，回溯搜索，贪心，动态规划；

## Big O


## 基础算法思路的应用
**真正理解循环不变量的意义**

### 数组相关问题
- [x] 283 [移动零](https://leetcode-cn.com/problems/move-zeroes) 
- [x] 27  [移除元素](https://leetcode-cn.com/problems/remove-element) 
- [x] 26  [删除排序数组中的重复项](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array)
- [x] 80  [删除排序数组中的重复项 II](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii) 
- [x] 75  [颜色分类](https://leetcode-cn.com/problems/sort-colors)  三路快排/计数排序
- [x] 88  [合并两个有序数组](https://leetcode-cn.com/problems/merge-sorted-array)  归并
- [x] 215 [数组中的第K个最大元素](https://leetcode-cn.com/problems/kth-largest-element-in-an-array)  快排/::堆排序::
- [x] offer40  [最小的k个数](https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof)   [最小K个数](https://leetcode-cn.com/problems/smallest-k-lcci) 
- [x] offer51  [数组中的逆序对](https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof) 
#### 对撞指针
- [x] 167 [两数之和 II - 输入有序数组](https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted)  ::二分搜索::/对撞指针
- [x] 125 [验证回文串](https://leetcode-cn.com/problems/valid-palindrome)  
- [x] 344 [反转字符串](https://leetcode-cn.com/problems/reverse-string) 
- [x] 345 [反转字符串中的元音字母](https://leetcode-cn.com/problems/reverse-vowels-of-a-string) 
- [x] 11  [盛最多水的容器](https://leetcode-cn.com/problems/container-with-most-water) 
#### 滑动窗口
- [ ] 209 [长度最小的子数组](https://leetcode-cn.com/problems/minimum-size-subarray-sum) 
- [ ] 3  [无重复字符的最长子串](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters) 
- [ ] 438  [找到字符串中所有字母异位词](https://leetcode-cn.com/problems/find-all-anagrams-in-a-string) 
- [ ] 76  [最小覆盖子串](https://leetcode-cn.com/problems/minimum-window-substring) 

### map和set不同的底层实现的区别
**hash表缺点：失去了数据的顺序性**

**C++中的Set和Map：**
* map和set的底层实现是平衡二叉树
* unordered_set和unordered_map的底层实现是哈希表

### set和map
- [x] 349  [两个数组的交集](https://leetcode-cn.com/problems/intersection-of-two-arrays) 
- [x] 350  [两个数组的交集 II](https://leetcode-cn.com/problems/intersection-of-two-arrays-ii) 
- [ ] 242  [有效的字母异位词](https://leetcode-cn.com/problems/valid-anagram) 
- [ ] 202  [快乐数](https://leetcode-cn.com/problems/happy-number) 
- [ ] 290  [单词规律](https://leetcode-cn.com/problems/word-pattern) 
- [ ] 205  [同构字符串](https://leetcode-cn.com/problems/isomorphic-strings) 
- [ ] 451  [根据字符出现频率排序](https://leetcode-cn.com/problems/sort-characters-by-frequency) 
- [x] 804  [唯一摩尔斯密码词](https://leetcode-cn.com/problems/unique-morse-code-words) 

### 二分查找相关问题
- [x] 875  [爱吃香蕉的珂珂](https://leetcode-cn.com/problems/koko-eating-bananas) 
- [x] 1011  [在 D 天内送达包裹的能力](https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days) 
- [x] 69  [x 的平方根](https://leetcode-cn.com/problems/sqrtx) 
- [x] 374  [猜数字大小](https://leetcode-cn.com/problems/guess-number-higher-or-lower) 

## 不同树的定义
* **完全二叉树**除了最后一层，每一层的节点数达到最大，同时最后一层的所有节点都在最左侧(堆)；
* **满二叉树**所有层的节点数达到最大；
* **平衡二叉树**每个节点的左右子树的高度差不超过1；

### 二叉树
- [x] 104 [二叉树的最大深度](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree) 
- [x] 111 [二叉树的最小深度](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree) 
- [x] 226 [翻转二叉树](https://leetcode-cn.com/problems/invert-binary-tree) 
- [x] 100 [相同的树](https://leetcode-cn.com/problems/same-tree) 
- [x] 101 [对称二叉树](https://leetcode-cn.com/problems/symmetric-tree) 
- [x] 222 [完全二叉树的节点个数](https://leetcode-cn.com/problems/count-complete-tree-nodes) 
- [x] 257 [二叉树的所有路径](https://leetcode-cn.com/problems/binary-tree-paths) 
- [x] 129 [求根到叶子节点数字之和](https://leetcode-cn.com/problems/sum-root-to-leaf-numbers) 
- [x] 112 [路径总和](https://leetcode-cn.com/problems/path-sum) 
- [x] 113 [路径总和 II](https://leetcode-cn.com/problems/path-sum-ii) 
- [x] 437 [路径总和 III](https://leetcode-cn.com/problems/path-sum-iii) 
- [x] 404 [左叶子之和](https://leetcode-cn.com/problems/sum-of-left-leaves) 
- [x] 107 [二叉树的层次遍历 II](https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii) 
- [ ] 106 [从中序与后序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal)  2020-09-25

### 二分搜索树
- [x] 中序遍历，后续遍历和层序遍历的非递归实现
- [x] 235  [二叉搜索树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree) 
- [x] 108  [将有序数组转换为二叉搜索树](https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree) 
- [ ] 98  [验证二叉搜索树](https://leetcode-cn.com/problems/validate-binary-search-tree) 
- [ ] 450 [删除二叉搜索树中的节点](https://leetcode-cn.com/problems/delete-node-in-a-bst) 
- [ ] 230 [二叉搜索树中第K小的元素](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst) 
- [ ] 236 [二叉树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree)     LCA

### 查找表
- [x] 1 [两数之和](https://leetcode-cn.com/problems/two-sum) 
- [x] 15 [三数之和](https://leetcode-cn.com/problems/3sum) 
- [x] 18 [四数之和](https://leetcode-cn.com/problems/4sum) 
- [ ] 16 [最接近的三数之和](https://leetcode-cn.com/problems/3sum-closest) 
- [x] 454 [四数相加 II](https://leetcode-cn.com/problems/4sum-ii) 
- [ ] 49  [字母异位词分组](https://leetcode-cn.com/problems/group-anagrams) 
- [x] 447 [回旋镖的数量](https://leetcode-cn.com/problems/number-of-boomerangs) 
- [x] 149 [直线上最多的点数](https://leetcode-cn.com/problems/max-points-on-a-line) 
- [x] 219 [存在重复元素 II](https://leetcode-cn.com/problems/contains-duplicate-ii)   查找表+滑动窗口
- [x] 217 [存在重复元素](https://leetcode-cn.com/problems/contains-duplicate) 
- [ ] 220 [存在重复元素 III](https://leetcode-cn.com/problems/contains-duplicate-iii) 

### 数学
- [ ] 7 [整数反转](https://leetcode-cn.com/problems/reverse-integer) 
- [x] 9 [回文数](https://leetcode-cn.com/problems/palindrome-number) 

### 位运算
- [ ] 136 [只出现一次的数字](https://leetcode-cn.com/problems/single-number) 

### 链表
- [x] 206 [反转链表](https://leetcode-cn.com/problems/reverse-linked-list) 
- [x] 92  [反转链表 II](https://leetcode-cn.com/problems/reverse-linked-list-ii) 
- [x] 83  [删除排序链表中的重复元素](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list) 
- [ ] 86  [分隔链表](https://leetcode-cn.com/problems/partition-list) 
- [x] 328  [奇偶链表](https://leetcode-cn.com/problems/odd-even-linked-list) 
- [ ] 2  [两数相加](https://leetcode-cn.com/problems/add-two-numbers) 
- [ ] 445 [两数相加 II](https://leetcode-cn.com/problems/add-two-numbers-ii) 
- [x] 203 [移除链表元素](https://leetcode-cn.com/problems/remove-linked-list-elements) 
- [ ] 82 [删除排序链表中的重复元素 II](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii) 
- [x] 21 [合并两个有序链表](https://leetcode-cn.com/problems/merge-two-sorted-lists) 
- [x] 24 [两两交换链表中的节点](https://leetcode-cn.com/problems/swap-nodes-in-pairs) 
- [ ] 25 [K 个一组翻转链表](https://leetcode-cn.com/problems/reverse-nodes-in-k-group) 
- [x] 147 [对链表进行插入排序](https://leetcode-cn.com/problems/insertion-sort-list) 
- [ ] 148 [排序链表](https://leetcode-cn.com/problems/sort-list) 
- [x] 237 [删除链表中的节点](https://leetcode-cn.com/problems/delete-node-in-a-linked-list)  不仅仅是穿针引线
- [x] 19 [删除链表的倒数第N个节点](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list)  链表与双指针
- [x] 61 [旋转链表](https://leetcode-cn.com/problems/rotate-list) 
- [x] 143 [重排链表](https://leetcode-cn.com/problems/reorder-list)  转换为数组/找链表中点+反转链表+归并链表
- [x] 234 [回文链表](https://leetcode-cn.com/problems/palindrome-linked-list) 
- [ ] 109 [有序链表转换二叉搜索树](https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree)   分治

### 栈
- [x] 150 [逆波兰表达式求值](https://leetcode-cn.com/problems/evaluate-reverse-polish-notation) 
- [x] 71 [简化路径](https://leetcode-cn.com/problems/simplify-path) 
- [ ] 341 [扁平化嵌套列表迭代器](https://leetcode-cn.com/problems/flatten-nested-list-iterator) 
- [x] 844 [比较含退格的字符串](https://leetcode-cn.com/problems/backspace-string-compare) 

### 队列
- [x] 107  [二叉树的层次遍历 II](https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii) 
- [ ] 103 [二叉树的锯齿形层次遍历](https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal) 
- [ ] 199 [二叉树的右视图](https://leetcode-cn.com/problems/binary-tree-right-side-view) 

### BFS和图的最短路径
- [ ] 279  [完全平方数](https://leetcode-cn.com/problems/perfect-squares) 
	
### 优先队列
- [x] offer40 [最小的k个数](https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof) 
- [x] 347 [前 K 个高频元素](https://leetcode-cn.com/problems/top-k-frequent-elements) 
- [ ] 23 [合并K个升序链表](https://leetcode-cn.com/problems/merge-k-sorted-lists)  延伸：K分排序

### 回溯
- [x] 17 [电话号码的字母组合](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number)   回溯法：暴力解法的一个主要实现手段
- [ ] 93 [复原IP地址](https://leetcode-cn.com/problems/restore-ip-addresses) 
- [ ] 131 [分割回文串](https://leetcode-cn.com/problems/palindrome-partitioning) 
- [x] 46 [全排列](https://leetcode-cn.com/problems/permutations) 
- [x] 47 [全排列 II](https://leetcode-cn.com/problems/permutations-ii) 
- [x] 77 [组合](https://leetcode-cn.com/problems/combinations)  通过**剪枝**来优化回溯
- [x] 39 [组合总和](https://leetcode-cn.com/problems/combination-sum) 
- [x] 40 [组合总和 II](https://leetcode-cn.com/problems/combination-sum-ii) 
- [x] 216 [组合总和 III](https://leetcode-cn.com/problems/combination-sum-iii) 
- [x] 78 [子集](https://leetcode-cn.com/problems/subsets) 
- [x] 90 [子集 II](https://leetcode-cn.com/problems/subsets-ii) 
- [ ] 401 [二进制手表](https://leetcode-cn.com/problems/binary-watch) 
- [x] 79 [单词搜索](https://leetcode-cn.com/problems/word-search) 
- [x] 200 [岛屿数量](https://leetcode-cn.com/problems/number-of-islands)     floodfill
- [ ] 130 [被围绕的区域](https://leetcode-cn.com/problems/surrounded-regions) 
- [ ] 417 [太平洋大西洋水流问题](https://leetcode-cn.com/problems/pacific-atlantic-water-flow) 
- [x] 51 [N 皇后](https://leetcode-cn.com/problems/n-queens) 
- [x] 52 [N皇后 II](https://leetcode-cn.com/problems/n-queens-ii) 
- [ ] 37 [解数独](https://leetcode-cn.com/problems/sudoku-solver) 

### 动态规划
[[动态规划（DP）]]
[[0-1背包问题]]

### 贪心算法
- [x] 455 [分发饼干](https://leetcode-cn.com/problems/assign-cookies) 
- [x] 392 [判断子序列](https://leetcode-cn.com/problems/is-subsequence) 
- [x] 435 [无重叠区间](https://leetcode-cn.com/problems/non-overlapping-intervals) 
- [x] 763 [划分字母区间](https://leetcode-cn.com/problems/partition-labels) 
**贪心选择性质**
### 设计
- [x] 155 [最小栈](https://leetcode-cn.com/problems/min-stack) 



