#include <stdio.h>
#define N 1000

int divides(int a, int b) {
    return b != 0 ? a%b == 0 : -1;
}

int main(void)
{
    int i, count = 0;
    int d1 = 3, d2 = 5;
    for(i = 0; i < N; i++)
    {
        if(divides(i, d1) || divides(i,d2))
            count += i;
    }

    printf("Sum of multiples of %d or %d less than %d: %d\n", d1, d2, N, count);
}
