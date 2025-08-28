class Solution:
    @staticmethod
    def fourSum( nums: list[int], target: int) -> list[list[int]]:
        res = []
        nums.sort()
        n = len(nums)
        for i in range(n - 3):
            if i > 0 and nums[i] == nums[i - 1]:
                continue
            for j in range(n - 1, i + 2, -1):   
                if j < n - 1 and nums[j] == nums[j + 1]:
                    continue
                l2, r2 = i + 1, j - 1
                while l2 < r2:
                    total = nums[i] + nums[l2] + nums[j] + nums[r2]
                    if total == target:
                        res.append([nums[i], nums[l2], nums[r2], nums[j]])
                        while l2 < r2 and nums[l2] == nums[l2 + 1]:
                            l2 += 1
                        while l2 < r2 and nums[r2] == nums[r2 - 1]:
                            r2 -= 1
                        l2 += 1
                        r2 -= 1
                    elif total > target:
                        r2 -= 1
                    else:
                        l2 += 1
        return res

Solution.fourSum([1,0,-1,0,-2,2], 0)