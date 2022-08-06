class Solution {
    
    var nameDic = [Substring: Substring]()
    
    func find(_ name: Substring) -> Substring {
        if (nameDic[name] ?? name) != name {
            nameDic[name] = find(nameDic[name] ?? name)
        }
        return nameDic[name] ?? name
    }
    
    func trulyMostPopular(_ names: [String], _ synonyms: [String]) -> [String] {
        defer { nameDic.removeAll() }
        
        for str in synonyms {
            let arr = str.dropFirst().dropLast()
                .split(separator: ",")
                .map { find($0) }
                .sorted()
            guard let first = arr.first else { continue }
            guard let last = arr.last else { continue }
            nameDic[last] = first
        }
        
        return names.reduce(into: [Substring: Int]()) { dic, str in
            let arr = str.split(separator: "(")
            guard var name = arr.first else { return }
            name = find(name)
            guard let cnt = Int(arr.last?.dropLast() ?? "0") else { return }
            dic[name] = (dic[name] ?? 0) + cnt
        }.map { (k, v) in
            return k + "(\(v))"
        }
    }
}