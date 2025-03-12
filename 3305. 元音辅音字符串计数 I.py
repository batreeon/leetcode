#
# @lc app=leetcode.cn id=3305 lang=python3
#
# [3305] 元音辅音字符串计数 I
#

# @lc code=start
from collections import defaultdict
class Solution:
    def countOfSubstrings(self, word: str, k: int) -> int:
        return self.count(word, k+1) - self.count(word, k)
    
    # 每个元音字母至少出现一次、至少有k个辅音字母的子串个数
    def count(self,  word: str, k: int) -> int:
        result: int = 0
        vowels: set[str] = {'a','e','i','o','u'}

        vowel = defaultdict(int)
        consonant: int = 0
        j: int = 0
        for i in range(len(word)):
            while j < len(word) and (consonant < k or len(vowel) < 5):
                b: str = word[j]
                if b in vowels:
                    vowel[b] += 1
                else:
                    consonant += 1
                j += 1

            if consonant >= k and len(vowel) == 5:
                result += len(word) - j + 1

            if word[i] in vowels:
                vowel[word[i]] -= 1
                if vowel[word[i]] == 0:
                    del vowel[word[i]]
            else:
                consonant -= 1
        return result
                           
# @lc code=end

