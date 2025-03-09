#
# @lc app=leetcode.cn id=2269 lang=python3
#
# [2269] 找到一个数字的 K 美丽值
#

# @lc code=start
class Solution:
    def divisorSubstrings(self, num: int, k: int) -> int:
        result: int = 0
        numstr: str = str(num)
        for i in range(len(numstr)-(k-1)):
            subnum: int = int(numstr[i:i+k])
            if subnum == 0:
                continue
            if num % subnum == 0:
                result += 1
        
        return result
# @lc code=end

