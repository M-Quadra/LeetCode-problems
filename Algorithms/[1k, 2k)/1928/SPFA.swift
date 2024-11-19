class Queue<T> {
    private var arr: [T]
    private var (i, len) = (0, 0)
    
    let defaultValue: T
    init(_ val: T, _ size: Int = 1) {
        defaultValue = val
        arr = Array(repeating: val, count: max(1, size))
    }
    
    func push(_ v: T) {
        if len+1 >= arr.count {
            var arrB = Array(repeating: defaultValue, count: arr.count*2)
            for j in 0..<len {
                arrB[j] = arr[(i+j)%arr.count]
            }
            arr = arrB
            i = 0
        }
        arr[(i+len)%arr.count] = v
        len += 1
    }
    
    func pop() -> T {
        let opt = arr[i]
        i += 1
        len -= 1
        i %= arr.count
        return opt
    }
    
    var isEmpty: Bool {
        return len <= 0
    }
}

class Solution {
    struct Edge {
        let to: Int
        let time: Int
        let next: Int
    }
    
    func createTable(_ edges: [[Int]], _ n: Int) -> ([Int], [Edge]) {
        var head = Array(repeating: -1, count: n)
        var table = [Edge]()
        for edge in edges {
            table.append(.init(to: edge[1], time: edge[2], next: head[edge[0]]))
            head[edge[0]] = table.count-1
            table.append(.init(to: edge[0], time: edge[2], next: head[edge[1]]))
            head[edge[1]] = table.count-1
        }
        return (head, table)
    }
    
    func createTimeLimits(_ maxTime: Int, _ head: [Int], _ table: [Edge]) -> [Int] {
        let n = head.count
        
        var timeLimits = Array(repeating: -1, count: n)
        timeLimits[n-1] = maxTime
        
        let que = Queue(-1, n/2)
        que.push(n-1)
        
        var used = Array(repeating: false, count: n)
        
        while !que.isEmpty {
            let u = que.pop()
            used[u] = false
            
            var it = head[u]
            while it != -1 {
                let e = table[it]
                let t = timeLimits[u] - e.time
                if t > timeLimits[e.to] {
                    timeLimits[e.to] = t
                    if !used[e.to] {
                        used[e.to] = true
                        que.push(e.to)
                    }
                }
                
                it = e.next
            }
        }
        
        if timeLimits[0] >= 0 {
            timeLimits[0] = 0
        }
        return timeLimits
    }
    
    func minCost(_ maxTime: Int, _ edges: [[Int]], _ passingFees: [Int]) -> Int {
        let n = passingFees.count
        
        let (head, table) = createTable(edges, n)
        let timeLimits = createTimeLimits(maxTime, head, table)
        if timeLimits[0] < 0 {
            return -1
        }
        
        var frees = Array(repeating: [Int](), count: n)
        for i in 0..<n {
            frees[i] = Array(repeating: Int.max, count: timeLimits[i]+1)
        }
        frees[0][0] = passingFees[0]
        
        struct Status {
            let u: Int
            let t: Int
        }
        let que = Queue(Status(u: 0, t: 0), n)
        que.push(Status(u: 0, t: 0))
        
        var used = Array(repeating: [Bool](), count: n-1)
        for i in 0..<used.count {
            used[i] = Array(repeating: false, count: frees[i].count)
        }
        
        var opt = Int.max
        while !que.isEmpty {
            let now = que.pop()
            used[now.u][now.t] = false
            if frees[now.u][now.t] >= opt {
                continue
            }
            
            var it = head[now.u]
            while it != -1 {
                let e = table[it]
                
                let nxt = Status(u: e.to, t: now.t+e.time)
                if nxt.t <= timeLimits[nxt.u] {
                    let f = frees[now.u][now.t] + passingFees[nxt.u]
                    if f < frees[nxt.u][nxt.t] && f < opt {
                        if nxt.u == n-1 {
                            opt = f
                        } else {
                            frees[nxt.u][nxt.t] = f
                            if !used[nxt.u][nxt.t] {
                                used[nxt.u][nxt.t] = true
                                que.push(nxt)
                            }
                        }
                    }
                }
                
                it = e.next
            }
        }
        
        return opt == Int.max ? -1 : opt
    }
}