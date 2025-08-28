#
# @lc app=leetcode.cn id=684 lang=python3
#
# [684] 冗余连接
#

# @lc code=start
class UnionFind:
    def __init__(self, n: int):
        self.parent = [i for i in range(n+1)]
        self.rank = [1] * (n+1)
    
    def find(self, i: int) -> int:
        if self.parent[i] != i:
            self.parent[i] = self.find(self.parent[i])
        return self.parent[i]

    def union(self, i: int, j: int) -> bool:
        pi, pj = self.find(i), self.find(j)
        if pi == pj:
            return True
        
        if pi != pj:
            if self.rank[pi] < self.rank[pj]:
                pi, pj = pj, pi
            self.rank[pi] += self.rank[pj]
            self.parent[pj] = pi
            return False

class Solution:
    def findRedundantConnection(self, edges: List[List[int]]) -> List[int]:
        uf = UnionFind(len(edges))
        result = []
        for edge in edges:
            if uf.union(edge[0], edge[1]):
                result = edge
        return result
# @lc code=end

