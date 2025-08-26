#
# @lc app=leetcode.cn id=211 lang=python3
#
# [211] 添加与搜索单词 - 数据结构设计
#

# @lc code=start
class WordDictionary:

    def __init__(self):
        self.child = [None for _ in range(26)]
        self.is_word = False

    def addWord(self, word: str) -> None:
        cur = self
        for b in word:
            i = ord(b)-ord('a')
            if not cur.child[i]:
                cur.child[i] = WordDictionary()
            cur = cur.child[i]
        cur.is_word = True

    def search(self, word: str) -> bool:
        if not word:
            if self.is_word:
                return True
            return False
        
        if word[0] == '.':
            for c in self.child:
                if c and c.search(word[1:]):
                    return True
        else:
            i = ord(word[0])-ord('a')
            if self.child[i] and self.child[i].search(word[1:]):
                return True
        
        return False
                


# Your WordDictionary object will be instantiated and called as such:
# obj = WordDictionary()
# obj.addWord(word)
# param_2 = obj.search(word)
# @lc code=end

