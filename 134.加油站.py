#
# @lc app=leetcode.cn id=134 lang=python3
#
# [134] 加油站
#

# @lc code=start
class Solution:
    def canCompleteCircuit(self, gas: List[int], cost: List[int]) -> int:
        minGas = 0
        gs = 0
        result = 0
        for i, (g, c) in enumerate(zip(gas, cost)):
            gs += g - c
            if gs < minGas:
                result = i + 1
                minGas = gs
        
        if gs < 0:
            return -1
        return result
# @lc code=end

