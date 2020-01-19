class Solution {
    func diStringMatch(_ S: String) -> [Int] {
        var opt = [Int]()
        var min = 0, max = S.count
        
        for char in S {
            if char == "I" {
                opt.append(min)
                min += 1
            } else {
                opt.append(max)
                max -= 1
            }
        }
        opt.append(min)
        
        return opt
    }
}