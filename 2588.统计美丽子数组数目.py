#
# @lc app=leetcode.cn id=2588 lang=python3
#
# [2588] 统计美丽子数组数目
#

# @lc code=start
class Solution:
    def beautifulSubarrays(self, nums: List[int]) -> int:
        xors: list[int] = [0] * (len(nums)+1)
        for i in range(1, len(nums)+1):
            xors[i] = xors[i-1] ^ nums[i-1]
        
        count: dict[int, int] = {}
        result: int = 0
        for i in range(len(xors)):
            c: int = count.get(xors[i], 0)
            result += c
            count[xors[i]] = c+1
        
        return result
# @lc code=end

