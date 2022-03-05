class Solution {
    
    var arr = [Int]()
    
    func find(_ x: Int) -> Int {
        if arr[x] == x {
            return x
        }
        
        arr[x] = find(arr[x])
        return arr[x]
    }
    
    func union(_ x: Int, _ y: Int) {
        let xHat = find(x), yHat = find(y)
        let v = min(xHat, yHat)
        arr[xHat] = v
        arr[yHat] = v
    }
    
    func findRedundantConnection(_ edges: [[Int]]) -> [Int] {
        arr = Array(0..<edges.count+1)
        for v in edges {
            if find(v[0]) == find(v[1]) {
                return v
            }
            
            union(v[0], v[1])
        }
        
        return []
    }
}