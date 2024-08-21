class Solution {
    func simplifyPath(_ path: String) -> String {
        var stack = [Substring]()
        for v in path.split(separator: "/") {
            switch v {
            case ".": continue
            case "..": _ = stack.popLast()
            default: stack.append(v)
            }
        }
        return "/" + stack.joined(separator: "/")
    }
}