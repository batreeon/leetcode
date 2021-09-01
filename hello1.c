#include <stdio.h>

int main() {
    printf("hello world!\n");
    char *s = "Ab";
    printf("%d\n",tolower((s+1)[0])-'a');
    putchar(tolower(s[1]));
    printf("\n");

    int i = 0;
    char c;
    char str[] = "RUNOOB";

    while( str[i] ) 
    {
        putchar(tolower(str[i]));
        i++;
    }

    return(0);
}