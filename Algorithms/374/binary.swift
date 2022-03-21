/** 
 * Forward declaration of guess API.
 * @param  num -> your guess number
 * @return 	     -1 if the picked number is lower than your guess number
 *			      1 if the picked number is higher than your guess number
 *               otherwise return 0
 * func guess(_ num: Int) -> Int 
 */

class Solution : GuessGame {
    func guessNumber(_ n: Int) -> Int {
        var l = 1, r = n+1
        while l < r {
            let mid = (l+r)/2
            switch guess(mid) {
            case -1: r = mid
            case 1:  l = mid+1
            default: return mid
            }
        }
        
        return l
    }
}
