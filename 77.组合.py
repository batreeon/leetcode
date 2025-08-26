#
# @lc app=leetcode.cn id=77 lang=python3
#
# [77] 组合
#

# @lc code=start
class Solution:
    def combine(self, n: int, k: int) -> List[List[int]]:
        res = []
        l = []
        nums = [num for num in range(1,n+1)]
        self.backtrack(nums, l, k, res)
        return res


    def backtrack(self, nums: [int], l: [int], k: int, res: [int]):
        if len(l) == k:
            res.append(copy.deepcopy(l))
            return
        
        for i, num in enumerate(nums):
            l.append(num)
            self.backtrack(nums[i+1:], l, k, res)
            l.pop()
# @lc code=end

