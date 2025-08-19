#
# @lc app=leetcode.cn id=26 lang=python3
#
# [26] 删除有序数组中的重复项
#

# @lc code=start
class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        tail: int = 0
        for i in range(1, len(nums)):
            if nums[i] != nums[tail]:
                tail += 1
                nums[tail] = nums[i]
        return tail+1
# @lc code=end

