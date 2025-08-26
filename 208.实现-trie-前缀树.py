#
# @lc app=leetcode.cn id=208 lang=python3
#
# [208] 实现 Trie (前缀树)
#

# @lc code=start
class Node:
    def __init__(self):
        self.child = [None] * 26
        self.is_word = False

class Trie:

    def __init__(self):
        self.root = Node()

    def insert(self, word: str) -> None:
        cur = self.root
        for b in word:
            i = ord(b) - ord('a')
            if not cur.child[i]:
                cur.child[i] = Node()
            cur = cur.child[i]
        cur.is_word = True

    def search(self, word: str) -> bool:
        cur = self.root
        for b in word:
            i = ord(b) - ord('a')
            if not cur.child[i]:
                return False
            cur = cur.child[i]
        return cur.is_word

    def startsWith(self, prefix: str) -> bool:
        cur = self.root
        for b in prefix:
            i = ord(b) - ord('a')
            if not cur.child[i]:
                return False
            cur = cur.child[i]
        return True


# Your Trie object will be instantiated and called as such:
# obj = Trie()
# obj.insert(word)
# param_2 = obj.search(word)
# param_3 = obj.startsWith(prefix)
# @lc code=end

