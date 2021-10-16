from typing import Dict, List

class WordsFrequency:
    dic: Dict[str, int] = {}
    def __init__(self, book: List[str]):
        self.dic.clear()
        for v in book:
            self.dic[v] = self.dic.get(v, 0) + 1
    def get(self, word: str) -> int:
        return self.dic.get(word, 0)

# Your WordsFrequency object will be instantiated and called as such:
# obj = WordsFrequency(book)
# param_1 = obj.get(word)