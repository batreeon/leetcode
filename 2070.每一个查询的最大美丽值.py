#
# @lc app=leetcode.cn id=2070 lang=python3
#
# [2070] 每一个查询的最大美丽值
#
from bisect import bisect_right
# @lc code=start
class Solution:
    def maximumBeauty(self, items: List[List[int]], queries: List[int]) -> List[int]:
        items.sort(key = lambda x : (-x[1], x[0]))
        result: List[int] = []
        for queriey in queries:
            result.append(0)
            for item in items:
                if item[0] <= queriey:
                    result[-1] = item[1]
                    break
        return result
# @lc code=end

