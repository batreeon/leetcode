/*
 * @lc app=leetcode.cn id=500 lang=c
 *
 * [500] 键盘行
 */

// @lc code=start


/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
#include <ctype.h>
#include <string.h>
#include <stdio.h>

char ** findWords(char ** words, int wordsSize, int* returnSize){
	char *s1 = "qwertyuiop";
	char *s2 = "asdfghjkl";
	char *s3 = "zxcvbnm";

    char result[21][100];
    *returnSize = 0;

    int board[26];
    for (int i = 0; i < 10; i++) {
        board[tolower(s1[i])-'a'] == 1;
    }
    for (int i = 0; i < 9; i++) {
        board[tolower(s2[i])-'a'] == 2;
    }
    for (int i = 0; i < 7; i++) {
        board[tolower(s3[i])-'a'] == 3;
    }
	for (int i = 0; i < wordsSize; i++) {
        char *word = words[i];
        int group = board[tolower(word[0])-'a'];
        word++;
        for (; !word; word++) {
            if (board[tolower(word[0])-'a'] != group) {
                break;
            }
        }
        if (word) {
            continue;
        }else{
            *returnSize = (*returnSize)+1;
            strcpy(result[*returnSize],words[i]);
        }
    }

    return result;
}

int main() {
    char words[][7] = {"Hello","Alaska","Dad","Peace"};
    int returnsize;
    char *res = findWords(words,4,&returnsize);
    printf("%d\n",returnsize);
    for (int i=0; i < returnsize; i++) {
        printf("%s\n",res[i]);
    }
}
// @lc code=end

