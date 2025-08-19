#
# @lc app=leetcode.cn id=80 lang=python3
#
# [80] 删除有序数组中的重复项 II
#

# @lc code=start
class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        tail: int = 1
        for i in range(2, len(nums)):
            if nums[i] != nums[tail-1]:
                tail += 1
                nums[tail] = nums[i]
        
        return tail + 1
        
# @lc code=end

