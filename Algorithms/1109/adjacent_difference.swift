class Solution {
    func corpFlightBookings(_ bookings: [[Int]], _ n: Int) -> [Int] {
        var opt = Array(repeating: 0, count: n)
        for it in bookings {
            opt[it[0] - 1] += it[2]
            if it[1] < n { opt[it[1]] -= it[2] }
        }
        for i in 1..<n {
            opt[i] += opt[i - 1]
        }
        return opt
    }
}