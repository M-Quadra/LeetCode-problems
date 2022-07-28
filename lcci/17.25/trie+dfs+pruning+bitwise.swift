class Node {
    
    private var v = -1
    var marks: [UInt64] = [0, 0]
    var next = [Int](repeating: 0, count: 26)
    
    func rest(_ v: Int) {
        self.v = v
        marks = [0, 0]
        next = [Int](repeating: 0, count: 26)
    }
    
    var isWord: Bool {
        get {
            marks[1] & (1<<63) != 0
        }
        set {
            marks[1] |= (1<<63)
        }
    }
    
    var char: Character {
        Character(Unicode.Scalar(UInt8(97 + v)))
    }
    
    func mark(_ offset: Int) {
        if offset < 64 {
            marks[0] |= 1<<offset
        } else {
            marks[1] |= 1<<(offset-64)
        }
    }
    func check(_ offset: Int) -> Bool {
        return offset < 64 ?
        marks[0] & (1<<offset) != 0 :
        marks[1] & (1<<(offset-64)) != 0
    }
}

let trie = (0..<10_0005).map { _ in Node() }

class Solution {
    
    class Point {
        var ver = trie[0], hor = trie[0]
    }
    
    func prepare(_ words: [String]) -> [[Int]] {
        var words = words.map({ str in
            str.map { Int($0.asciiValue ?? 97) - 97 }
        })
        
        trie[0].rest(-1)
        var cnt = 1
        
        for word in words {
            var cur = trie[0]
            cur.mark(word.count)
            
            for v in word {
                let i = cur.next[v] == 0 ? {
                    trie[cnt].rest(v)
                    cur.next[v] = cnt
                    cnt += 1
                    return cur.next[v]
                }() : cur.next[v]
                
                cur = trie[i]
                cur.mark(word.count)
            }
            cur.isWord = true
        }
        
        words = words.filter({ word in
            word.first { trie[0].next[$0] == 0 } == nil
        }).sorted(by: { a, b in
            a.count > b.count
        })
        return words
    }
    
    let datas = (0..<105).map { _ in
        (0..<105).map { _ in Point() }
    }
    
    func dfs(_ x: Int, _ y: Int, _ w: Int, _ h: Int) -> Bool {
        let top = datas[y-1][x], left = datas[y][x-1]
        
        for (i, j) in zip(top.ver.next, left.hor.next) where i>0 && j>0 {
            let ver = trie[i], hor = trie[j]
            if !ver.check(h) || !hor.check(w) { continue }
            
            if x+1 >= w || y+1 >= h {
                if !ver.isWord && !hor.isWord { continue }
            }
            (datas[y][x].ver, datas[y][x].hor) = (trie[i], trie[j])
            
            if 0..<w ~= x+1 {
                if dfs(x+1, y, w, h) { return true }
            } else if 0..<h ~= y+1 {
                if dfs(1, y+1, w, h) { return true }
            } else { return true }
        }
        return false
    }
    
    func maxRectangle(_ words: [String]) -> [String] {
        let words = prepare(words)
        
        var opt = [String](), area = 0
        
        for iW in words.indices {
            let w = words[iW]

            var wMarks:[UInt64] = [UInt64.max, UInt64.max>>1]
            for (i, v) in w.enumerated() {
                let n = trie[trie[0].next[v]]
                wMarks[0] &= n.marks[0]
                wMarks[1] &= n.marks[1]
                (datas[0][i].ver, datas[0][i].hor) = (n, n)
            }
            if wMarks[0] == 0 && wMarks[1] == 0 { continue }

        loopH:
            for iH in iW..<words.count {
                let h = words[iH]
                let nArea = w.count * h.count
                if nArea <= area { break }
                
                if h.count<64 {
                    if wMarks[0] & (1<<h.count) == 0 { continue }
                } else {
                    if wMarks[1] & (1<<(h.count-64)) == 0 { continue }
                }

                for (i, v) in h.enumerated() {
                    let n = trie[trie[0].next[v]]
                    if !n.check(w.count) { continue loopH }
                    
                    datas[i][0].hor = n
                }
                
                if w.count <= 1 || h.count <= 1 || !dfs(1, 1, w.count, h.count) { continue }
                opt = datas[0..<h.count].map({ arr in
                    String(arr[0..<w.count].map { $0.hor.char })
                })
                area = nArea
            }
        }
        return opt
    }
}