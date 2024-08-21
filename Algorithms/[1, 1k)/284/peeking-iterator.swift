// Swift IndexingIterator refernence:
// https://developer.apple.com/documentation/swift/indexingiterator

class PeekingIterator {
    
    var arr: IndexingIterator<Array<Int>>
    var tmp: Int?
    
    init(_ arr: IndexingIterator<Array<Int>>) {
        self.arr = arr
    }
    
    func next() -> Int {
        if let tmp = tmp {
            defer { self.tmp = nil }
            return tmp
        }
        return arr.next()!
    }
    
    func peek() -> Int {
        if let tmp = tmp {
            return tmp
        }
        
        tmp = arr.next()
        return tmp ?? 0
    }
    
    func hasNext() -> Bool {
        if tmp == nil {
            tmp = arr.next()
        }
        
        return tmp != nil
    }
}

/**
 * Your PeekingIterator object will be instantiated and called as such:
 * let obj = PeekingIterator(arr)
 * let ret_1: Int = obj.next()
 * let ret_2: Int = obj.peek()
 * let ret_3: Bool = obj.hasNext()
 */