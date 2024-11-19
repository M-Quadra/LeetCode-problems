class Solution {
    func findOcurrences(_ text: String, _ first: String, _ second: String) -> [String] {
        let arr = text.split(separator: " ")
        guard arr.count >= 3 else { return [] }
        
        var opt = [String]()
        for i in 2..<arr.count {
            if arr[i-2] != first || arr[i-1] != second { continue }
            opt.append(String(arr[i]))
        }
        return opt
    }
}