class FrontMiddleBackQueue {

    private class Node {
        let val: Int
        weak var prev: Node?
        var next: Node?
        
        init(_ val: Int) {
            self.val = val
        }
    }
    
    private class Deque {
        
        var count = 0
        private var front: Node?
        private weak var back: Node?
        
        func pushFront(_ n: Node) {
            n.prev = nil
            n.next = nil
            
            n.next = front
            front?.prev = n
            front = n
            
            count += 1
            if count == 1 { back = front }
        }
        
        func pushBack(_ n: Node) {
            n.prev = nil
            n.next = nil
            
            back?.next = n
            n.prev = back
            back = n
            
            count += 1
            if count == 1 { front = back }
        }
        
        func popFront() -> Node? {
            guard let n = front else { return nil }
            front = n.next
            
            n.prev = nil
            n.next = nil
            front?.prev = nil
            
            count -= 1
            if count <= 0 { back = nil }
            return n
        }
        
        func popBack() -> Node? {
            guard let n = back else { return nil }
            back = n.prev
            
            n.prev = nil
            n.next = nil
            back?.next = nil
            
            count -= 1
            if count <= 0 { front = nil }
            return n
        }
    }
    
    private let front = Deque()
    private let back = Deque()
    
    init() { }
    
    func pushFront(_ val: Int) {
        front.pushFront(Node(val))
    }
    
    func pushMiddle(_ val: Int) {
        while front.count < back.count {
            guard let n = back.popFront() else { break }
            front.pushBack(n)
        }
        while front.count > back.count {
            guard let v = front.popBack() else { break }
            back.pushFront(v)
        }
        front.pushBack(Node(val))
    }
    
    func pushBack(_ val: Int) {
        back.pushBack(Node(val))
    }
    
    func popFront() -> Int {
        return front.popFront()?.val ?? back.popFront()?.val ?? -1
    }
    
    func popMiddle() -> Int {
        while front.count > back.count {
            guard let n = front.popBack() else { break }
            back.pushFront(n)
        }
        while front.count < back.count {
            guard let n = back.popFront() else { break }
            front.pushBack(n)
        }
        return front.popBack()?.val ?? -1
    }
    
    func popBack() -> Int {
        return back.popBack()?.val ?? front.popBack()?.val ?? -1
    }
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * let obj = FrontMiddleBackQueue()
 * obj.pushFront(val)
 * obj.pushMiddle(val)
 * obj.pushBack(val)
 * let ret_4: Int = obj.popFront()
 * let ret_5: Int = obj.popMiddle()
 * let ret_6: Int = obj.popBack()
 */