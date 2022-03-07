class Solution {
    
    var cnt = Array(repeating: Int.max, count: 0x9999 + 1)
    var que = Array(repeating: 0, count: 1_0000)
    
    func strToInt(_ str: String) -> Int {
        str.map { c in
            Int(c.asciiValue ?? 48) - 48
        }.reduce(0) { partialResult, v in
            partialResult << 4 | v
        }
    }
    
    func openLock(_ deadends: [String], _ target: String) -> Int {
        let targetNum = strToInt(target)
        
        for i in 0..<cnt.count {
            cnt[i] = Int.max
        }
        cnt[0] = 0
        for v in deadends {
            cnt[strToInt(v)] = -1
        }
        
        var l = 0, r = 1
        que[l] = 0
        while l != r {
            let now = que[l]
            l += 1
            l %= que.count
            if cnt[now] == -1 { continue }
            if cnt[now] > cnt[targetNum] { continue }
            let nCnt = cnt[now] + 1
            
            for i in 0..<4 {
                let a = now & (0xF << (i*4))
                let n = now ^ a
                let v = a >> (i*4)
                
                for j in 0..<2 {
                    let next = n | (((v+1 + 8*j) % 10) << (i*4))
                    if nCnt < cnt[next] {
                        cnt[next] = nCnt
                        que[r] = next
                        r += 1
                        r %= que.count
                    }
                }
            }
        }
        
        return cnt[targetNum] == Int.max ? -1 : cnt[targetNum];
    }
}