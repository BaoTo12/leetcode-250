from typing import List, Dict


class Solution:
    @staticmethod
    def majorityElement(nums: List[int]) -> int:
        hash: Dict[int, int] = {}
        for n in nums:
            hash[n] = 1 + hash.get(n, 0)
            if hash[n] > len(nums) // 2:
                return n
        return 0

    @staticmethod
    def majorityElement1(nums: List[int]) -> int:
        res = majority = 0
        for n in nums:
                if majority == 0:
                    res = n
                majority += 1 if n == res else -1
        return res

print(Solution.majorityElement1([2, 2, 1, 3, 1, 2, 4]))