#
# @lc app=leetcode.cn id=46 lang=python3
#
# [46] 全排列
#

# @lc code=start
class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        res = []
        l = []
        visited = set()
        self.backtrack(nums, visited, l, res)
        return res
        
    def backtrack(self, nums: list[int], visited: set(), l: list[int], res: list[int]):
        if len(l) == len(nums):
            res.append(copy.deepcopy(l))
            return
        for i, num in enumerate(nums):
            if i in visited:
                continue
            l.append(num)
            visited.add(i)
            self.backtrack(nums, visited, l, res)
            visited.remove(i)
            l.pop()
# @lc code=end

