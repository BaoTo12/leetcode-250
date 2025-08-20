from typing import List
import heapq
from collections import Counter


class Solution:
    @staticmethod
    def topKFrequentV1(nums: List[int], k: int) -> List[int]:
        heap = []
        counter = {}
        for n in nums:
            counter[n] = counter.get(n, 0) + 1

        for key, val in counter.items():
            heapq.heappush(heap, (-val, key))
        print(heap)

        res = []
        while len(res) < k:
            res.append(heapq.heappop(heap)[1])
        return res

    @staticmethod
    def topKFrequentV2(nums: List[int], k: int) -> List[int]:
        counter = {}
        for num in nums:
            counter[num] = counter.get(num, 0) + 1
        freq = [[] for _ in range(len(nums) + 1)]

        for val, key in counter.items():
            freq[key].append(val)

        print((freq))
        res = []

        for i in range(len(freq) - 1, -1, -1):
            for val in freq[i]:
                res.append(val)
                if len(res) == k:
                    return res
        return res

    @staticmethod
    def topKFrequentV3(nums: List[int], k: int) -> List[int]:
        counts = Counter(nums)
        return heapq.nlargest(k, counts, key=lambda num: counts[num])


result = Solution.topKFrequentV3([10, 10, 20, 20, 20, 30], 2)
print(result)
