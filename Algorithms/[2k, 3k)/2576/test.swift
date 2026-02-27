import Testing

@Test func case0() {
    let arr = [3, 5, 2, 4]
    let opt = Solution().maxNumOfMarkedIndices(arr)
    #expect(opt == 2)
}

@Test func case1() {
    let arr = [9, 2, 5, 4]
    let opt = Solution().maxNumOfMarkedIndices(arr)
    #expect(opt == 4)
}

@Test func case2() {
    let arr = [7, 6, 8]
    let opt = Solution().maxNumOfMarkedIndices(arr)
    #expect(opt == 0)
}

@Test func case3() {
    let arr = [2, 2, 4, 4]
    let opt = Solution().maxNumOfMarkedIndices(arr)
    #expect(opt == 4)
}

@Test func case4() {
    let arr = [1]
    let opt = Solution().maxNumOfMarkedIndices(arr)
    #expect(opt == 0)
}

@Test func case5() {
    let arr = [42,83,48,10,24,55,9,100,10,17,17,99,51,32,16,98,99,31,28,68,71,14,64,29,15,40]
    let opt = Solution().maxNumOfMarkedIndices(arr)
    #expect(opt == 26)
}

@Test func case6() {
    let arr = [1,78,27,48,14,86,79,68,77,20,57,21,18,67,5,51,70,85,47,56,22,79,41,8,39,81,59,74,14,45,49,15,10,28,16,77,22,65,8,36,79,94,44,80,72,8,96,78,39,92,69,55,9,44,26,76,40,77,16,69,40,64,12,48,66,7,59,10]
    let opt = Solution().maxNumOfMarkedIndices(arr)
    #expect(opt == 64)
}

@Test func case7() {
    let arr = [57,40,57,51,90,51,68,100,24,39,11,85,2,22,67,29,74,82,10,96,14,35,25,76,26,54,29,44,63,49,73,50,95,89,43,62,24,88,88,36,6,16,14,2,42,42,60,25,4,58,23,22,27,26,3,79,64,20,92]
    let opt = Solution().maxNumOfMarkedIndices(arr)
    #expect(opt == 58)
}

@Test func case8() {
    let arr = [13,41,84,21,64,75,3,44,73,41,83,1,4,52,41,72,4,60,15,94,34,9,10,4,81,94,50,45,38,46,14,11,79,70,31,70,53,9,18,32,32,19]
    let opt = Solution().maxNumOfMarkedIndices(arr)
    #expect(opt == 42)
}
