#
# @lc app=leetcode.cn id=146 lang=python3
#
# [146] LRU 缓存
#

# @lc code=start
class Node:
    def __init__(self, key=0, value=0):
        self.key = key
        self.value = value
        self.prev = None
        self.next = None

class LRUCache:

    def __init__(self, capacity: int):
        self.capacity = capacity
        self.dummy = Node()
        self.dummy.prev = self.dummy
        self.dummy.next = self.dummy
        self.m = {}

    def get(self, key: int) -> int:
        if key not in self.m:
            return -1
        node = self.m[key]

        self.remove(node)
        self.insert_to_front(node)
        return node.value

    def put(self, key: int, value: int) -> None:
        if key in self.m:
            node = self.m[key]
            node.value = value

            self.remove(node)
            self.insert_to_front(node)
            return
        
        node = Node(key, value)
        self.m[key] = node
        self.insert_to_front(node)
        if len(self.m) > self.capacity:
            last_node = self.dummy.prev
            del self.m[last_node.key]
            self.remove(last_node)

        return

    def remove(self, node: Node) -> None:
        node.prev.next = node.next
        node.next.prev = node.prev

    def insert_to_front(self, node: Node) -> None:
        node.next = self.dummy.next
        node.prev = self.dummy
        node.next.prev = node
        node.prev.next = node
# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)
# @lc code=end

