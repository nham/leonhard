#include <stdio.h>
#define M 1000

int main(void)
{
    int c, a, inc;
    /* a^2 + b^2 = c^2 & a + b + c = M
     * => M(M - 2c) = 2ab
     * Since a, b > 0, M > 2c, or c < M/2
     *
     * Also, a^2 + (M - a - c)^2 = c^2
     * => 1/2 M^2 = (M-a)*(a+c)
     *
     * Note that this implies that M can never be odd. (M odd iff M^2 odd)
     * So M is even and 1/2 M^2 is even, which implies that if c is even,
     * a cannot be odd, since (a+c) would be odd, as would (M-a), and
     * odd*odd = even
     *
     * Perhaps the latter analysis is not really worthwhile, since it only
     * cuts down operations by 1/4 in the worst case (for even c we only need
     * examine the even a)
     */

    if(M % 2 != 0)
        return 0;

    for(c = 1; c < M/2; c++) {
        if(c % 2 == 0)
            inc = 2;
        else
            inc = 1;

        for(a = inc; a <= (M - c)/2; a += inc)
            if(M*M/2 == (M-a)*(a+c)) {
                printf("(%d, %d, %d) => %d\n", a, (M-a-c), c, a*c*(M-a-c));
            }

    }

}
