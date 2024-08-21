from queue import LifoQueue

class Solution:
    def scoreOfParentheses(self, s: str) -> int:
        stackS, stackN = LifoQueue[str](), LifoQueue[int]()
        for v in s:
            if v == '(':
                stackS.put(v)
                continue
            else:
                n = 0
                while not stackS.empty():
                    top = stackS.get_nowait()
                    if top == '(':
                        break
                    
                    n += stackN.get_nowait()
                stackN.put(max(1, n*2))
                stackS.put('-')
        
        opt = 0
        while not stackN.empty():
            opt += stackN.get_nowait()
        return opt