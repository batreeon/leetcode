/*
 * @lc app=leetcode.cn id=63 lang=typescript
 *
 * [63] 不同路径 II
 */

// @lc code=start
// 简单动态规划
function uniquePathsWithObstacles(obstacleGrid: number[][]): number {
    let m: number = obstacleGrid.length;
    let n: number = obstacleGrid[0].length;
    
    let dp: number[] = new Array<number>(n).fill(0);
    dp[0] = obstacleGrid[0][0]^1;
    for (let j = 1; j < n; j++) {
        if (obstacleGrid[0][j] === 1) {
            dp[j] = 0;
        }else{
            dp[j] = dp[j-1];
        }
    }

    for (let i = 1; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (obstacleGrid[i][j] === 1) {
                dp[j] = 0;
            }else if (j === 0) {
                dp[j] = dp[j];
            }else{
                dp[j] = dp[j-1] + dp[j]
            }
        }
    }
    return dp[n-1];
};
// @lc code=end

