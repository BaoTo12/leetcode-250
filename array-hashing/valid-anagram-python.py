

#! Version 1
from typing import Dict


class Solution:
    @staticmethod
    def isAnagram(s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        counter: Dict[str, int] = {}

        for char in s:
            counter[char] = counter.get(char, 0) + 1

        for char in t:
            if char not in counter or counter[char] == 0:
                return False
            counter[char] -= 1

        return True


# print(Solution.isAnagram("anagram", "nagaram"))


#! Version 2
class Solution1:
    @staticmethod
    def isAnagram(s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        counter = [0] * 26

        for char in s:
            counter[ord(char) - ord('a')] += 1

        for char in t:
            if counter[ord(char) - ord('a')] == 0:
                return False
            counter[ord(char) - ord('a')] -= 1

        return True


#! Version 3
class Solution3:
    @staticmethod
    def isAnagram(s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        
        s_count : Dict[str, int] = {}
        t_count : Dict[str, int]= {}
        
        for i in range(len(s)):
            print(s[i])
            s_count[s[i]] = 1 + s_count.get(s[i], 0)
            t_count[t[i]] = 1 + t_count.get(t[i], 0)
            
        print(s_count)
        print(t_count)
        return s_count == t_count


Solution3.isAnagram("anagram", "nagaram")

count : Dict[str, int] = {}

print(count)

count['a'] = count.get('a', 0) + 1

print(count)