#
# @lc app=leetcode.cn id=56 lang=python3
#
# [56] 合并区间
#

# @lc code=start
class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        intervals.sort(key=lambda interval: interval[0])
        result: list[list[int]] = []
        for interval in intervals:
            if len(result) == 0 or interval[0] > result[-1][1]:
                result.append(interval)
            else:
                result[-1][1] = max(interval[1], result[-1][1])

        return result
# @lc code=end

