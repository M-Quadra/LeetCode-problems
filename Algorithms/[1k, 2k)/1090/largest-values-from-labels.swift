class Solution {
    
    func largestValsFromLabels(_ values: [Int], _ labels: [Int], _ numWanted: Int, _ useLimit: Int) -> Int {
        var opt = 0, num = 0, cntDic = [Int: Int]()
        for i in Array(values.indices).sorted(by: { a, b in
            values[a] > values[b]
        }) {
            if num >= numWanted { break }
            
            let cnt = cntDic[labels[i]] ?? 0
            if cnt >= useLimit { continue }
            
            num += 1
            cntDic[labels[i]] = cnt + 1
            opt += values[i]
        }
        return opt
    }
}