class Solution {
    
    func countHighestScoreNodes(_ parents: [Int]) -> Int {
        var childs = Array(repeating: [-1, -1], count: parents.count)
        var indegrees = [Int](repeating: 0, count: parents.count)
        for (i, v) in parents.enumerated() where i > 0 {
            indegrees[v] += 1
            let j = childs[v][0] == -1 ? 0 : 1
            childs[v][j] = i
        }
        
        var scores = [Int](repeating: 1, count: parents.count)
        var arr = indegrees.indices.filter { indegrees[$0] == 0 }
        while let v = arr.popLast() {
            let p = parents[v]
            if p < 0 { continue }
            
            scores[p] += scores[v]
            indegrees[p] -= 1
            
            if indegrees[p] > 0 { continue }
            arr.append(p)
        }
        
        var maxScore = 0, opt = 0
        for i in parents.indices {
            let p = max(1, parents.count - scores[i])
            let l = childs[i][0] >= 0 ? scores[childs[i][0]] : 1
            let r = childs[i][1] >= 0 ? scores[childs[i][1]] : 1
            let s = p * l * r
            if s < maxScore {
                continue
            } else if s == maxScore {
                opt += 1
            } else {
                maxScore = s
                opt = 1
            }
        }
        return opt
    }
}