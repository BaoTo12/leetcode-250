class Solution:
  @staticmethod
  def twoSum(numbers: list[int], target: int) -> list[int]:
    left = 0
    right = len(numbers) -1
    while left < right:
      sum = numbers[left] + numbers[right]
      if sum > target:
        right -= 1
      elif sum < target:
        left += 1
      else:
        return [left + 1, right + 1]