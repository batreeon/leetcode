/*
 * @lc app=leetcode.cn id=1109 lang=c
 *
 * [1109] 航班预订统计
 */

// @lc code=start


/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* corpFlightBookings(int** bookings, int bookingsSize, int* bookingsColSize, int n, int* returnSize){
    int *res = malloc(sizeof(int)*n);
    memset(res, 0, sizeof(int)*n);
    *returnSize = n;
    for (int i = 0; i < bookingsSize; i++) {
        int beg = bookings[i][0], end = bookings[i][1], seats = bookings[i][2];
        res[beg-1] += seats;
        if (end < n) {
            res[end] -= seats;
        }
    }
    for (int i = 1 ; i < n ; i++) {
        res[i] += res[i-1];
    }
    return res;
}
// @lc code=end

