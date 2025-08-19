#
# @lc app=leetcode.cn id=27 lang=python3
#
# [27] 移除元素
#

# @lc code=start
class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        tail: int = 0
        for num in nums:
            if num != val:
                nums[tail] = num
                tail += 1
        return tail
# @lc code=end

