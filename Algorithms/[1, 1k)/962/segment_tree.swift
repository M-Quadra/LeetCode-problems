class Solution {
    
    var tree = Array(repeating: -1, count: 1 << 17)
    
    func update(index: Int, offset: Int) {
        if index <= 0 || offset <= tree[index] { return }
        
        tree[index] = offset
        update(index: index >> 1, offset: offset)
    }
    
    func find(index: Int=1, range: Range<Int>=0..<(1<<16), target: Int) -> Int {
        if target >= range.upperBound { return -1 }
        if target <= range.lowerBound || tree[index] == -1 { return tree[index] }
        
        let mid = (range.lowerBound + range.upperBound) >> 1
        return max(
            find(index: index<<1, range: range[..<mid], target: target),
            find(index: index<<1|1, range: range[mid...], target: target)
        )
    }
    
    func maxWidthRamp(_ nums: [Int]) -> Int {
        for i in 0..<tree.count {
            tree[i] = -1
        }
        
        var opt = 0
        for (i, v) in nums.enumerated().reversed() {
            let j = max(i, find(target: v))
            opt = max(opt, j - i)
            update(index: (1 << 16) + v, offset: i)
        }
        return opt
    }
}