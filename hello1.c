#include <stdio.h>

int main() {
    printf("hello world!\n");
    char s[] = "A";
    printf("%d\n",tolower(s[0])-'a');
    putchar(tolower(s[0]));
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