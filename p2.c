#include <stdio.h>
#define N 4000000

int main(void)
{
    long bprev = 0, bcurr, tmp,
        sum = 0;

    for(bcurr = 2; bcurr <= N; 
        tmp = bcurr, bcurr = bprev + bcurr*4,  bprev = tmp)
    {
        sum += bcurr;
//        printf("%d\n", bcurr);
    }

    printf("%d\n", sum);
}
