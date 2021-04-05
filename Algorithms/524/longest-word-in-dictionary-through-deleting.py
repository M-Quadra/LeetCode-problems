from typing import List

class Solution:
    def findLongestWord(self, s: str, dictionary: List[str]) -> str:
        dictionary.sort()
        edAry = [0] * len(dictionary)
        for c in s:
            for i, word in enumerate(dictionary):
                if edAry[i]<len(word) and c==word[edAry[i]]:
                    edAry[i] += 1
        for i, word in enumerate(dictionary):
            if edAry[i] != len(word):
                edAry[i] = -1
        maxEd = max(edAry)
        for i, word in enumerate(dictionary):
            if edAry[i]==maxEd and len(word)==maxEd:
                return word
        return ""