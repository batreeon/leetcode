#
# @lc app=leetcode.cn id=2080 lang=python3
#
# [2080] 区间内查询数字的频率
#

# @lc code=start
from collections import defaultdict
from bisect import bisect_left, bisect_right
class RangeFreqQuery:

    def __init__(self, arr: List[int]):
        # 记录每个元素的下表序列
        self.indexs = defaultdict(list[int])
        for i, num in enumerate(arr):
            self.indexs[num].append(i)

    def query(self, left: int, right: int, value: int) -> int:
        indexs: list[int] = self.indexs.get(value, [])
        if len(indexs) == 0:
            return 0
        lo: int = bisect_left(indexs, left)
        hi: int = bisect_right(indexs, right)
        return hi - lo


# Your RangeFreqQuery object will be instantiated and called as such:
# obj = RangeFreqQuery(arr)
# param_1 = obj.query(left,right,value)
# @lc code=end

