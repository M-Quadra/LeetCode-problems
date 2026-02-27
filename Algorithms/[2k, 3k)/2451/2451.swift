class Solution {
    func oddString(_ words: borrowing [String]) -> String {
        let arr = words.map { str in
            var arr = str.utf8.map { Int($0) }
            for i in 1..<arr.count {
                arr[i-1] -= arr[i]
            }
            return arr.dropLast().hashValue
        }.enumerated().sorted { $0.element < $1.element }
        
        if arr[0].element == arr[1].element {
            return words[arr.last!.offset]
        }
        return words[arr[0].offset]
    }
}
