from typing import List

class Solution:
  @staticmethod
  def longestCommonPrefix(strs: List[str]) -> str:
    prefix: str = strs[0]
    prefix_length: int = len(prefix)
    
    for s in strs[1:]:
      while prefix != s[0:prefix_length]:
        prefix_length -= 1
        if prefix_length == 0:
          return ""
        prefix = prefix[0:prefix_length]
    
    return prefix
  
  
result =  Solution.longestCommonPrefix(["flower","flow","flight"])

print(result)

s = "flow"
a = s[0: 6]
print(a)
print(len(a))