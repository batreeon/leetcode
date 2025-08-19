#
# @lc app=leetcode.cn id=88 lang=python3
#
# [88] 合并两个有序数组
#

# @lc code=start
class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        tail: int = m + n - 1
        tail1, tail2 = m-1, n-1
        while tail1 >= 0 or tail2 >= 0:
            if tail1 < 0:
                nums1[tail] = nums2[tail2]
                tail2 -= 1
            elif tail2 < 0:
                nums1[tail] = nums1[tail1]
                tail1 -= 1
            elif nums1[tail1] >= nums2[tail2]:
                nums1[tail] = nums1[tail1]
                tail1 -= 1
            else:
                nums1[tail] = nums2[tail2]
                tail2 -= 1
            tail -= 1
        
# @lc code=end

