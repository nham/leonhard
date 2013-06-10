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
    int count = 1, i;

    printf("test: %d\n", isprime(104744));

    for(i = 3; count < 10001; i +=1)
        if(isprime(i))
            count += 1;

    printf("10001st prime is %d\n", i-1);

}
