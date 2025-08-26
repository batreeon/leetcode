#
# @lc app=leetcode.cn id=17 lang=python3
#
# [17] 电话号码的字母组合
#

# @lc code=start
class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        if not digits:
            return []
        res = []
        l = ''
        self.backtrack(digits, l, res)
        return res

    
    def backtrack(self, digits: str, l: str, res: [str]):
        m = {
            2:'abc',
            3:'def',
            4:'ghi',
            5:'jkl',
            6:'mno',
            7:'qprs',
            8:'tuv',
            9:'wxyz'
        }
        if not digits:
            res.append(l)
            return
        for b in m[int(digits[0])]:
            l = l + b
            self.backtrack(digits[1:], l, res)
            l = l[:-1]
            
# @lc code=end

