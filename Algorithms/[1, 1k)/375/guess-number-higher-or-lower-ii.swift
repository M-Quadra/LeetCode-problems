class Solution {
    func getMoneyAmount(_ n: Int) -> Int {
        if n < 2 {
            return 0
        }
        var payAry = Array.init(repeating: Array.init(repeating: 0, count: n + 5), count: n + 5)
        
        for len in 2...n {
            for st in 1...n {
                let ed = st + len-1
                if ed > n {
                    continue
                }
                
                var pay = Int.max
                for i in st...ed {
                    let tmp = i + max(payAry[st][i - st], payAry[i+1][ed - i])
                    pay = min(pay, tmp)
                }
                payAry[st][len] = pay
            }
        }
        
        return payAry[1][n]
    }
}