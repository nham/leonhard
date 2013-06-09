#include <stdio.h>

int ispal(int i)
{
    int dig[50], j = 0, k;

    for(; i > 0; i = i/10) {
        dig[j] = i%10;
        j += 1;
    }

    for(k = 0; k <= (j-1)/2; k++) {
        if(dig[k] != dig[j-1-k])
            return 0;
    }

    return 1;
}

int has3digfactors(int c)
{
    int i;
    for(i = 999; i > 1; i -= 1)
        if(c%i == 0 && c/i < 1000)
            return i;

    return 0;
}

int main(void)
{
    int c, f;

    for(c = 999*999; c > 0; c-=1)
        if(ispal(c) && (f = has3digfactors(c))) {
            printf("winner, winner: %d, factor = %d\n", c, f);
            break;
        }
}
