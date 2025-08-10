from __future__ import annotations
from typing import Optional, List


class Node:
    def __init__(self, key: int = -1, val: int = -1, next: Optional[Node] = None) -> None:
        self.key = key
        self.val = val
        self.next = next


class MyHashMap:
    def __init__(self) -> None:
        self.map: List[Node] = []
        for _ in range(1000):
            self.map.append(Node())

    def put(self, key: int, val: int)-> None:
        cur = self.map[self.key(key)]

        while cur.next:
            if cur.next.key == key:
                cur.next.val = val
                return
            cur = cur.next

        cur.next = Node(key, val)

    def retrieve(self, key: int) -> int:
        cur: Node = self.map[self.key(key)]
        while cur.next:
            if cur.next.key == key:
                return cur.val
            cur = cur.next
        return -1
      
    def remove(self, key: int) -> None:
      cur: Node = self.map[self.key(key)]
      while cur.next:
        if cur.next.key == key:
          cur.next = cur.next.next
          return
        cur = cur.next

    def key(self, key: int) -> int:
        return key % len(self.map)
