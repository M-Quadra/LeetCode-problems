class Solution {
    
    struct point {
        var cnt = 0
        var grid = 0
    }
    
    func possibleToStamp(_ grid: [[Int]], _ stampHeight: Int, _ stampWidth: Int) -> Bool {
        var arr = Array(repeating: Array(repeating: point(), count: grid[0].count), count: grid.count)
        for i in grid.indices {
            for j in grid[i].indices {
                arr[i][j].grid = grid[i][j]
            }
        }
        
        var st = point()
        for y in 0..<min(grid.count, stampHeight-1) {
            for x in 0...min(grid[y].count-1, stampWidth-1) {
                st.cnt += grid[y][x]
            }
        }
        if st.cnt == 0 && grid.count>=stampHeight && grid[0].count>=stampWidth {
            for y in 0..<(stampHeight-1) {
                for x in 0...(stampWidth-1) {
                    arr[y][x].grid = -1
                }
            }
        }
        
        for y in grid.indices {
            if grid[y].count < stampWidth { break }
            
            let last = y > 0 ? arr[y-1][0] : st
            arr[y][0].cnt = last.cnt;
            if y-1 >= 0 {
                for x in 0..<stampWidth {
                    arr[y][0].cnt -= grid[y-1][x]
                }
            }
            let i = y + (stampHeight-1)
            guard grid.indices ~= i else { break }
            for x in 0..<stampWidth {
                arr[y][0].cnt += grid[i][x]
            }
            if arr[y][0].cnt == 0 {
                if last.cnt == 0 {
                    for x in 0..<stampWidth {
                        arr[i][x].grid = -1;
                    }
                } else {
                    for y in y...i {
                        for x in 0..<stampWidth {
                            arr[y][x].grid = -1
                        }
                    }
                }
            }
            
            for x in grid[y][1...].indices {
                let last = arr[y][x-1]
                arr[y][x].cnt = last.cnt
                let j = x + stampWidth-1
                guard grid[y].indices ~= j else { break }
                
                for i in y..<(y + stampHeight) {
                    arr[y][x].cnt -= grid[i][x-1]
                    arr[y][x].cnt += grid[i][j]
                }
                if arr[y][x].cnt == 0 {
                    if last.cnt == 0 {
                        for i in y..<(y + stampHeight) {
                            arr[i][j].grid = -1
                        }
                    } else {
                        for i in y..<(y + stampHeight) {
                            for j in x...j {
                                arr[i][j].grid = -1
                            }
                        }
                    }
                }
            }
        }
        
        for y in grid.indices {
            for x in grid[y].indices {
                if arr[y][x].grid == 0 {
                    return false
                }
            }
        }
        return true
    }
}