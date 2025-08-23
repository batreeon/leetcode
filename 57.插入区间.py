#
# @lc app=leetcode.cn id=57 lang=python3
#
# [57] 插入区间
#

# @lc code=start
class Solution:
    def insert(self, intervals: List[List[int]], newInterval: List[int]) -> List[List[int]]:
        result: list[list[int]] = []
        l, r = newInterval

        for i, interval in enumerate(intervals):
            if interval[1] < l:
                result.append(interval)
            elif interval[0] > r:
                result.append([l,r])
                result.extend(intervals[i:])
                return result
            else:
                l, r = min(l, interval[0]), max(r, interval[1])
        
        result.append([l, r])
        return result
# @lc code=end

