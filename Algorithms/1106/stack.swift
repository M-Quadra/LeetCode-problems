class Solution {
    
    func parseBoolExpr(_ expression: String) -> Bool {
        var stackV = [Bool]()
        var stackC = [Character]()
        for c in expression {
            if c == "," { continue }
            if c != ")" {
                stackC.append(c)
                continue
            }
            
            while let last = stackC.popLast(), last != "(" {
                switch last {
                case "t": stackV.append(true)
                case "f": stackV.append(false)
                default: continue
                }
            }
            
            switch stackC.last {
            case "&":
                stackV[0] = stackV[1...].reduce(stackV[0]) { prev, ok in
                    prev && ok
                }
            case "|":
                stackV[0] = stackV[1...].reduce(stackV[0]) { prev, ok in
                    prev || ok
                }
            case "!":
                stackV[0] = !stackV[0]
            default: break
            }
            _ = stackC.popLast()
            stackC.append(stackV[0] ? "t" : "f")
            stackV.removeAll()
        }
        
        return stackC.first == "t" ? true : false
    }
}