class Solution:
    def complexNumberMultiply(self, num1: str, num2: str) -> str:
        v1, i1 = [int(x) for x in num1[:-1].split('+')]
        v2, i2 = [int(x) for x in num2[:-1].split('+')]
        return f'{v1*v2 - i1*i2}+{v1*i2 + i1*v2}i'