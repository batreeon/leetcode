/*
 * @lc app=leetcode.cn id=165 lang=c
 *
 * [165] 比较版本号
 */

// @lc code=start
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int compareVersion(char * version1, char * version2){
    char v1s[500][500], v2s[500][500];

    int i = 0;
    for (char *p = version1; *p; p++) {
        // 注意消除的是前导0（因此要判断是否p==version1），不是后面的0
        if (*p == '0' && p == version1) {
            version1++;
        }
        if (*p == '.') {
            *p = '\0';
            strcpy(v1s[i++],version1);
            version1 = p+1;
        }
    }
    strcpy(v1s[i],version1);
    int len1 = i+1;
    int maxIndex = i+1;

    i = 0;
    for (char *p = version2; *p; p++) {
        if (*p == '0' && p == version2) {
            version2++;
        }
        if (*p == '.') {
            *p = '\0';
            strcpy(v2s[i++],version2);
            version2 = p+1;
        }
    }
    strcpy(v2s[i],version2);

    int len2 = i+1;
    if (i+1 < maxIndex) {
        maxIndex = i+1;
    }
    
    for (i = 0; i < maxIndex; i++) {
        int num1 = atoi(v1s[i]);
        int num2 = atoi(v2s[i]);
        if (num1 < num2) {
            return -1;
        }else if(num1 > num2) {
            return 1;
        }
    }
    if (len1 == len2) {
        return 0;
    }
    if (len1 < len2) {
        for (i = maxIndex; i < len2; i++) {
            if (v2s[i][0] != '\0') {
                return -1;
            }
        }
    }else{
        for (i = maxIndex; i < len1; i++) {
            if (v1s[i][0] != '\0') {
                return 1;
            }
        }
    }
    return 0;
}

// int main() {
//     /*case1
//     char v1[] = "1.0";
//     char v2[] = "1.0.0";
//     */
//     char v1[] = "1.0";
//     char v2[] = "1.10";
//     int res = compareVersion(v1,v2);
//     printf("%d",res);
// }
// @lc code=end

