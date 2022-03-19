class MyHashSet {

    var arr = Array(repeating: false, count: Int(1e6)+1)
    init() {}
    
    func add(_ key: Int) {
        arr[key] = true
    }
    
    func remove(_ key: Int) {
        arr[key] = false
    }
    
    func contains(_ key: Int) -> Bool {
        return arr[key]
    }
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * let obj = MyHashSet()
 * obj.add(key)
 * obj.remove(key)
 * let ret_3: Bool = obj.contains(key)
 */
