/*
 * @lc app=leetcode.cn id=997 lang=c
 *
 * [997] 找到小镇的法官
 */

// @lc code=start


int findJudge(int n, int** trust, int trustSize, int* trustColSize){
    // trustMap[i][j]表示i是否信任j
    int trustMap[n+1][n+1];
    for (int i = 0 ; i < trustSize ; i++) {
        trustMap[trust[i][0]][trust[i][1]] = 1;
    }
    for (int i = 1 ; i <= n ; i++) {
        for (int j = 1 ; j <= n ; j++) {
            // 法官i的trustMap[i][0]应等于0
            trustMap[i][0] = trustMap[i][0] || trustMap[i][j];
        }
    }
    for (int i = 1 ; i <= n ; i++) {
        trustMap[0][i] = 1;
        for (int j = 1 ; j <= n ; j++) {
            if (j == i) {
                continue;
            }
            trustMap[0][i] = trustMap[0][i] && trustMap[j][i];
        }
        if (trustMap[0][i] == 1 && trustMap[i][0] == 0) {
            return i;
        }
    }
    return -1;
}
// @lc code=end

