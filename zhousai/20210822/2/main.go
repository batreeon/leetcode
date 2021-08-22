func findDifferentBinaryString(nums []string) string {
    n := len(nums)
    m := map[int]bool{}
    for _,num := range nums {
        m[str2int(num)] = true
    }
    // 注意这个i的范围，他只有n个数，因此你取n+1个数肯定能找到不存在的那个
    for i := 0 ; i < n+1 ; i++ {
        if !m[i] {
            return int2str(i,n)
        }
    }
    return string("")
}
func str2int(s string) int {
    result := 0
    for _,b := range s {
        result <<= 1
        result += int(b-'0')
    }
    return result
}
func int2str(num int,n int) string {
    result := string("")
    for i := 0 ; i < n ; i++ {
        if num != 0 && num&1 == 1 {
            result = "1" + result
        }else{
            result = "0" + result
        }
        if num != 0 {
            num >>= 1
        }
    }
    return result
}