#include <stdio.h>
#define N 600851475143


int divides(long a, long b)
{
    return b != 0 ? a%b == 0 : -1;
}

int is_prime(int i)
{
    int j;
    for(j = 2; j*j <= i; j++) {
        if(divides(i, j))
            return 0;
    }
    return 1;
}

int main(void)
{
    long n = N, i;
    long bpf = 0;

    if (divides(n, 2)) {
        n /= 2;
        bpf = 2;
    }

    for(i = 3; i <= N && n > 1 && i <= n; i += 2) {
        //printf("divides %d %d: %d\n", n, i, divides(n, i));
        //printf("isprime %d: %d\n\n", i, is_prime(i));
        if(is_prime(i) && divides(n, i)) {
            n /= i;
            bpf = i;
            printf("%d\n", bpf);
        }
    }

    printf("%d\n", bpf);


}
