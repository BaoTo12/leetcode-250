class Solution:
  @staticmethod
  def calPoints(operations: list[str]) -> int:
    stack = []
    for str in operations:
      if str == "C":
        stack.pop()
      elif str == "D":
        stack.append(stack[-1] * 2)
      elif str == "+":
        stack.append(stack[-1] + stack[-2])
      else:
        stack.append(int(str))
        
    return sum(stack)