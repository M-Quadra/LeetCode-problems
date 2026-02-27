fileprivate class Node {
    
    var indices: (Int, Int)
    var left: Node?, right: Node?
    init(indices: (Int, Int)) {
        self.indices = indices
    }
    
    func find(range: Range<Int>, st: Int) -> (Int, Int) { // [l, r)
        if st <= range.lowerBound { return self.indices }
        let mid = (range.lowerBound+range.upperBound) >> 1
        
        let li = if st < mid, let left = self.left {
            left.find(range: range.lowerBound..<mid, st: st)
        } else { (Int.max, Int.max) }
        let ri = if let right = self.right {
            right.find(range: mid..<range.upperBound, st: st)
        } else { (Int.max, Int.max) }
        
        let arr = [li.0, li.1, ri.0, ri.1].sorted()
        return (arr[0], arr[1])
    }
    
    func update(i: Int) {
        switch true {
        case i == self.indices.0: break
        case i < self.indices.0: self.indices = (i, self.indices.0)
        case i < self.indices.1: self.indices.1 = i
        default: break
        }
    }
    
    func update(range: Range<Int>, st: Int, i: Int) {
        self.update(i: i)
        if st <= range.lowerBound { return }
        let mid = (range.lowerBound+range.upperBound) >> 1
        
        if st < mid {
            let left = self.left ?? Node(indices: (i, Int.max))
            left.update(range: range.lowerBound..<mid, st: st, i: i)
            self.left = consume left
        } else {
            let right = self.right ?? Node(indices: (i, Int.max))
            right.update(range: mid..<range.upperBound, st: st, i: i)
            self.right = consume right
        }
    }
}

class Solution {
    func secondGreaterElement(_ nums: [Int]) -> [Int] { // 线段树
        var arr = [Int](repeating: -1, count: nums.count)
        
        let head = Node(indices: (Int.max, Int.max))
        let dic = [Int: Int](uniqueKeysWithValues: Set(nums).sorted().enumerated().map { ($0.element, $0.offset) })
        for i in nums.indices.reversed() {
            let v = dic[nums[i]]!
            let j = head.find(range: 0..<dic.count, st: v+1).1
            if j != .max { arr[i] = nums[j] }
            head.update(range: 0..<dic.count, st: v, i: i)
        }
        return arr
    }
}
