class Solution {
    func halfQuestions(_ questions: [Int]) -> Int {
        var cnt = Array(repeating: 0, count: 1000+1)
        for v in questions {
            cnt[v] += 1
        }
        cnt = cnt.sorted().reversed()
        
        var n = questions.count/2
        var opt = 0
        for c in cnt {
            opt += 1
            n -= c
            if n <= 0 {
                break
            }
        }
        return opt
    }
}