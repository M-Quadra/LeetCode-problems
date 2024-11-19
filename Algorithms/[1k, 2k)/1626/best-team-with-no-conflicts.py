from typing import List
from operator import attrgetter

class Solution:
    class P:
        def __init__(self, score: int, age: int):
            self.score = score
            self.age = age
            self.allScore = score

    def bestTeamScore(self, scores: List[int], ages: List[int]) -> int:
        pAry: List[self.P] = []
        for i, nowS in enumerate(scores):
            nowA = ages[i]
            pAry.append(self.P(nowS, nowA))
        pAry = sorted(pAry, key=attrgetter('age', 'score'))
        
        for i in range(len(pAry)-1, -1, -1):
            maxS = pAry[i].allScore
            for p in pAry[i+1:]:
                if p.score >= pAry[i].score:
                    maxS = max(maxS, p.allScore + pAry[i].allScore)
            pAry[i].allScore = maxS
        return max(pAry, key=lambda x: x.allScore).allScore