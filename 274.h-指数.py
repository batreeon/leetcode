#
# @lc app=leetcode.cn id=274 lang=python3
#
# [274] H 指数
#

# @lc code=start
class Solution:
    def hIndex(self, citations: List[int]) -> int:
        l, r = 0, len(citations)
        while l < r:
            mid = (l+r+1)//2 # mid等于0的时候，会出现死循环
            if self.h(mid, citations):
                l = mid
            else:
                r = mid - 1
        return l
    
    def h(self, mid, citations)->bool:
        cnt  = 0
        for citation in citations:
            if citation >= mid:
                cnt += 1
        if cnt >= mid:
            return True
        return False
# @lc code=end

