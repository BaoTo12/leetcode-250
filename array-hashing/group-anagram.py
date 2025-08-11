
from collections import defaultdict
class Solution:
  @staticmethod
  def groupAnagrams(strs: list[str]) ->  list[list[str]]: 
    ans: dict[str ,list[str]] = {}
    
    for s in strs:
      sortedKey = "".join(sorted(s))
      if sortedKey not in ans:
        ans[sortedKey] = []
      ans[sortedKey].append(s)
    print(ans.values())
    return list(ans.values())
  @staticmethod
  def groupAnagrams2(strs: list[str]) -> list[list[str]]: 
    ans: dict[tuple[int,...], list[str]] = defaultdict(list)
    for s in strs:
      count: list[int] = [0] * 26
      for c in s:
        count[ord(c) - ord('a')] += 1
      ans[tuple(count)].append(s)
    return list(ans.values())
  
  
# Solution.groupAnagrams(["eat","tea","tan","ate","nat","bat"])
Solution.groupAnagrams2(["eat","tea","tan","ate","nat","bat"])