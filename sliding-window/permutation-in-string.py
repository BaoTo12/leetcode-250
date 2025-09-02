

class Solution:
    @staticmethod
    def checkInclusion(s1: str, s2: str) -> bool:
        counter: list[int] = [0] * 26
        for c in s1:
          idx = ord(c) - ord('a')
          counter[idx] += 1
        
        left = 0
        right = len(s1)
        while right <= len(s2):
          str = s2[left:right]
          checkCounterInclusion: list[int] = [0] * 26
          for c in str:
            idx = ord(c) - ord('a')
            checkCounterInclusion[idx] += 1
          
          if counter == checkCounterInclusion:
            return True
          left += 1
          right += 1
        return False
    @staticmethod
    def checkInclusionV2(s1: str, s2: str) -> bool:
        if len(s1) > len(s2):
            return False

        s1_count = {}
        s2_count = {}

        for i in range(len(s1)):
            s1_count[s1[i]] = 1 + s1_count.get(s1[i], 0)
            s2_count[s2[i]] = 1 + s2_count.get(s2[i], 0)

        if s1_count == s2_count:
            return True

        left = 0
        for right in range(len(s1), len(s2)):
            s2_count[s2[right]] = 1 + s2_count.get(s2[right], 0)
            s2_count[s2[left]] -= 1

            if s2_count[s2[left]] == 0:
                del s2_count[s2[left]]

            left += 1

            if s1_count == s2_count:
                return True

        return False

s1 = "ab"
s2 = "eidbaooo"
# Solution.checkInclusionV2(s1, s2)


s2_count = {
	'a': 1,
	'b': 0
}

print(s2_count)