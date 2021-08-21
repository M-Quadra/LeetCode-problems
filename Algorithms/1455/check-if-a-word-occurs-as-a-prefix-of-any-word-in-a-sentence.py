class Solution:
    def isPrefixOfWord(self, sentence: str, searchWord: str) -> int:
        for i, v in enumerate(sentence.split(' '), 1):
            if v.startswith(searchWord):
                return i
        return -1
