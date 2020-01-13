class Solution {
    func topKFrequent(_ nums: [Int], _ k: Int) -> [Int] {
        var ary = Array(Set(nums))
        var cntDic = [Int: Int]()
        
        for num in nums {
            cntDic[num] = (cntDic[num] ?? 0) + 1
        }
        ary.sort {
            (cntDic[$0] ?? 0) > (cntDic[$1] ?? 0)
        }
        
        return Array(ary[0..<k])
    }
}