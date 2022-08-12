class Solution {
    
    func countGoodRectangles(_ rectangles: [[Int]]) -> Int {
        return rectangles.reduce(into: (0, 0)) { counter, arr in
            let edge = min(arr[0], arr[1])
            if edge < counter.0 { return }
            else if edge == counter.0 { counter.1 += 1 }
            else { counter = (edge, 1) }
        }.1
    }
}
