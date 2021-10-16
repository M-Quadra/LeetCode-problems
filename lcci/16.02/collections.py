from typing import List
import collections

class WordsFrequency:
    def __init__(self, book: List[str]):
        self.dic = collections.Counter(book)
    def get(self, word: str) -> int:
        return self.dic[word]

# Your WordsFrequency object will be instantiated and called as such:
# obj = WordsFrequency(book)
# param_1 = obj.get(word)