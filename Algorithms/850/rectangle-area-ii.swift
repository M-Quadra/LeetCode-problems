class Solution {
    
    func rectangleArea(_ rectangles: [[Int]]) -> Int {
        let mod: Int64 = 1_0000_0000_7
        var recAry = rectangles
        var opt: Int64 = 0
        
        while recAry.count > 0 {
            let area = Int64(recAry[0][2] - recAry[0][0]) * Int64(recAry[0][3] - recAry[0][1])
            opt += area % mod
            opt %= mod
            
            var i = 1;
            let ed = recAry.count
            while i < ed {
                let subRecAry = self.subRec(rec1: recAry[0], rec2: recAry[i])
                if subRecAry.count >= 1 {
                    recAry[i] = subRecAry.first!
                    recAry += subRecAry[1..<subRecAry.count]
                }
                i += 1
            }
            
            recAry.removeFirst()
        }
        return Int(opt);
    }
    
    func subLine(line1: [Int], line2: [Int]) -> [Int] {
        let opt = [max(line1[0], line2[0]), min(line1[1], line2[1])]
        return opt[0] < opt[1] ? opt : []
    }
    
    func subRec(rec1: [Int], rec2: [Int]) -> [[Int]] {
        let x = self.subLine(line1: [rec1[0], rec1[2]], line2: [rec2[0], rec2[2]])
        if x.count <= 0 {
            return []
        }
        let y = self.subLine(line1: [rec1[1], rec1[3]], line2: [rec2[1], rec2[3]])
        if y.count <= 0 {
            return []
        }
        
        var optAry = [[Int]]()
        
        var subRec = [rec2[0], rec2[1], x[0], rec2[3]]// left
        if subRec[0] < subRec[2] && subRec[1] < subRec[3] {
            optAry.append(subRec)
        }
        
        subRec = [x[0], y[1], x[1], rec2[3]]// up
        if subRec[0] < subRec[2] && subRec[1] < subRec[3] {
            optAry.append(subRec)
        }
        
        subRec = [x[0], rec2[1], x[1], y[0]]// down
        if subRec[0] < subRec[2] && subRec[1] < subRec[3] {
            optAry.append(subRec)
        }
        
        subRec = [x[1], rec2[1], rec2[2], rec2[3]]// right
        if subRec[0] < subRec[2] && subRec[1] < subRec[3] {
            optAry.append(subRec)
        }
        
        if optAry.count <= 0 {
            optAry.append([0, 0, 0, 0])
        }
        return optAry
    }
}