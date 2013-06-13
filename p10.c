#include <stdio.h>

int isprime(int c)
{
    int i;
    if(c%2 == 0)
        return 0;

    for(i = 3; i*i <= c; i += 2)
        if(c%i == 0)
            return 0;

    return 1;
}


int main(void)
{
    int i;
    unsigned long long sum = 2;

    for(i = 3; i < 2000000; i += 2)
        if(isprime(i))
            sum += i;

    printf("sum: %lu\n", sum);
}
