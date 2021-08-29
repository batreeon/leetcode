/*
 * @lc app=leetcode.cn id=1588 lang=c
 *
 * [1588] 所有奇数长度子数组的和
 */

// @lc code=start


int sumOddLengthSubarrays(int* arr, int arrSize){
    int result = 0;
    if (arrSize == 0) {
        return result;
    }
    if (arrSize <= 2) {
        for (int i = 0 ; i < arrSize ; i++) {
            result += arr[i];
        }
        return result;
    }
    
    int sum[arrSize+1];
    for (int i=1 ; i <= arrSize ; i++) {
        sum[i] = sum[i-1]+arr[i-1];
    }

    for (int l = 1 ; l <= arrSize ; l += 2) {
        for (int i = 0 ; i+l-1 < arrSize ; i++) {
            result += sum[i+l]-sum[i];
        }
    }
    return result;
}
// @lc code=end

