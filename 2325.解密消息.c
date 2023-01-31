/*
 * @lc app=leetcode.cn id=2325 lang=c
 *
 * [2325] 解密消息
 */

// @lc code=start


char * decodeMessage(char * key, char * message){
    int exist[26], map[26];
    for (int i = 0; i < 26; ++i) {
        exist[i] = 0;
        map[i] = 0;
    }

    // space用来记录空格和重复字母数，这俩都会导致字母下标不等于其次序
    int space = 0;
    for (int i = 0; key[i]; ++i) {
        char c = key[i];
        if (c == ' ') {
            ++space;
            continue;
        }
        if (!exist[c-'a']) {
            exist[c-'a'] = 1;
            map[c-'a'] = i-space;
        }else{
            ++space;
        }
    }

    for (int i = 0; message[i]; ++i) {
        char c = message[i];
        if (c == ' ') {
            continue;
        }
        message[i] = map[c-'a'] + 'a';
    }

    return message;
}
// @lc code=end

