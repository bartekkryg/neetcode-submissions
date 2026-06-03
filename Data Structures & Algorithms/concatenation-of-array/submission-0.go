func getConcatenation(nums []int) []int {
    var concNums = make([]int, 2*len(nums))
    for i, num := range nums {
        concNums[i] = num
        concNums[i+len(nums)] = num
    }
    return concNums
}
