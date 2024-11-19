from typing import List, Set

class ThroneInheritance:
    def __init__(self, kingName: str):
        self.names = [kingName]
        self.index = {kingName: 0}
        self.children: List[List[int]] = [[]]
        self.died: Set[int] = set()

    def birth(self, parentName: str, childName: str) -> None:
        iC = len(self.names)
        self.names.append(childName)
        self.index[childName] = iC
        self.children.append([])

        iP = self.index[parentName]
        self.children[iP].append(iC)

    def death(self, name: str) -> None:
        i = self.index[name]
        self.died.add(i)

    def getInheritanceOrder(self) -> List[str]:
        opt: List[str] = []
        stack = [0]
        while len(stack) > 0:
            i = stack.pop()
            stack += self.children[i][::-1]
            if i not in self.died:
                opt.append(self.names[i])
        return opt

# Your ThroneInheritance object will be instantiated and called as such:
# obj = ThroneInheritance(kingName)
# obj.birth(parentName,childName)
# obj.death(name)
# param_3 = obj.getInheritanceOrder()
