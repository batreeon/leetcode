/*
 * @lc app=leetcode.cn id=528 lang=c
 *
 * [528] 按权重随机选择
 */

// @lc code=start
// #include <stdlib.h>

typedef struct {
    int *sum;
    int len;
} Solution;


Solution* solutionCreate(int* w, int wSize) {
    Solution *obj = malloc(sizeof(Solution));
    obj->sum = malloc(sizeof(int)*wSize);
    obj->len = wSize;
    
    obj->sum[0] = w[0];
    for (int i = 1; i < wSize; i++) {
        obj->sum[i] = obj->sum[i-1] + w[i];
    }

    return obj;
}

// 返回target是在obj->sum的第几个区间
// 寻找第一个sum[i] >= target的下标i
int binSearch(Solution *obj, int target) {
    int l = 0,r = obj->len-1;
    for (; l < r; ) {
        int mid = (r - l) / 2 + l;
        if (target <= obj->sum[mid]) {
            r = mid;
        }else {
            l = mid+1;
        }
    }
    return l;
}
int solutionPickIndex(Solution* obj) {
    int max = obj->sum[obj->len-1];
    // 得到的数据范围是[1,max]
    int num = rand() % max + 1;
    return binSearch(obj,num);
}

void solutionFree(Solution* obj) {
    free(obj->sum);
    free(obj);
}

// int main() {
//     int w[1] = {1};
//     Solution *obj = solutionCreate(w,1);
//     printf("%d\n", solutionPickIndex(obj));
// }
/**
 * Your Solution struct will be instantiated and called as such:
 * Solution* obj = solutionCreate(w, wSize);
 * int param_1 = solutionPickIndex(obj);
 
 * solutionFree(obj);
*/
// @lc code=end