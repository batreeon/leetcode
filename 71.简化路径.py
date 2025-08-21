#
# @lc app=leetcode.cn id=71 lang=python3
#
# [71] 简化路径
#

# @lc code=start
class Solution:
    def simplifyPath(self, path: str) -> str:
        paths = path.split('/')
        temp = []
        up = 0
        while len(paths) > 0:
            item = paths.pop()
            if item == '' or item == '.':
                continue
            if item == '..':
                up += 1
                continue
            if up > 0:
                up -= 1
                continue
            temp.append(item)
        temp.reverse()
        return '/' + '/'.join(temp)
# @lc code=end

