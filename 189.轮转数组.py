#
# @lc app=leetcode.cn id=189 lang=python3
#
# [189] 轮转数组
#

# @lc code=start
class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        k %= len(nums)
        nums.reverse()
        sub1 = nums[:k]
        sub1.reverse()
        nums[:k] = sub1
        sub2 = nums[k:]
        sub2.reverse()
        nums[k:] = sub2
# @lc code=end

