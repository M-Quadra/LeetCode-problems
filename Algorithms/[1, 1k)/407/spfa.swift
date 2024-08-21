class Queue<T> {
    
    class Node<T> {
        let now: T
        var next: Node<T>?
        
        init(_ v: T) {
            now = v
        }
    }
    
    var head: Node<T>?
    var last: Node<T>?
    
    init(v: T) {
        head = Node(v)
        last = head
    }
    
    func pop() -> T {
        let opt = head?.now
        head = head?.next
        return opt!
    }
    
    func push(_ v: T) {
        last?.next = Node(v)
        last = last?.next
        if head == nil {
            head = last
        }
    }
    
    func isEmpty() -> Bool {
        return head == nil
    }
    
}

class Solution {
    func trapRainWater(_ heightMap: [[Int]]) -> Int {
        var hArr = heightMap.map { arr in
            arr.map { v in Int(1e5) }
        }
        
        let que = Queue(v: (0, 0))
        while !que.isEmpty() {
            let (y, x) = que.pop()
            
            let top = heightMap.indices ~= (y-1) ? hArr[y-1][x] + heightMap[y-1][x] : 0
            let bottom = heightMap.indices ~= (y+1) ? hArr[y+1][x] + heightMap[y+1][x] : 0
            let left = heightMap[0].indices ~= (x-1) ? hArr[y][x-1] + heightMap[y][x-1] : 0
            let right = heightMap[0].indices ~= (x+1) ? hArr[y][x+1] + heightMap[y][x+1] : 0
            
            let minH = min(top, bottom, left, right)
            let h = max(0, minH - heightMap[y][x])
            if h < hArr[y][x] {
                hArr[y][x] = h

                if heightMap.indices ~= (y-1) { que.push((y: y-1, x: x)) }
                if heightMap.indices ~= (y+1) { que.push((y: y+1, x: x)) }
                if heightMap[0].indices ~= (x-1) { que.push((y: y, x: x-1)) }
                if heightMap[0].indices ~= (x+1) { que.push((y: y, x: x+1)) }
            }
        }
        
        return hArr.flatMap { arr in
            arr
        }.reduce(0) { partialResult, v in
            partialResult + v
        }
    }
}