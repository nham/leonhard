#include <stdio.h>
#include <stdlib.h>

void popseq(char *fname, char *arr)
{
    FILE *fp;
    char buf[100];
    int i, j = 0;
    char tmp[3];
    tmp[2] = '\0';

    fp = fopen(fname, "r");

    // WHY do i have to do 52 and not 51??
    while(fgets(buf, 61, fp) != NULL) {
        for(i = 0; i < 20; i++) {
            tmp[0] = buf[3*i];
            tmp[1] = buf[3*i + 1];
            arr[j+i] = (char) atoi(tmp);
        }
        j += 20;
    }

}

int maxseqinc(char *arr, int len, int amt, int inc)
{
    if(len < amt) return 0;

    int cnt = 0, off = 0, max = 0, tmp;
    int j;

    while(cnt <= inc*len - inc*amt) {
        tmp = 1;
        for(off = 0; off < inc*amt; off += inc)
            if(arr[cnt+off] != 0)
                tmp *= arr[cnt+off];
            else
                break;
        // terminated early due to zero, so just advanced cnt past it
        if(off < inc*amt)
            cnt += off+inc;
        else if(tmp > max) {
            max = tmp;
        }

        cnt += inc;
    }

    return max;

}

int main(void)
{
    char num[400];
    popseq("p11.txt", &num[0]);
    int N = 20, i;
    int max = 0,tmp;

    for(i = 0; i < N; i++) {
        tmp = maxseqinc(&num[i], N, 4, 20);
        if(tmp > max) {
            max = tmp;
            printf("new max is %d, vert %d\n", max, i);
        }
        tmp = maxseqinc(&num[i*N], N, 4, 1);
        if(tmp > max) {
            max = tmp;
            printf("new max is %d, horiz %d\n", max, i);
        }
        tmp = maxseqinc(&num[i], N-i, 4, 21);
        if(tmp > max) {
            max = tmp;
            printf("new max is %d, right diag %d\n", max, i);
        }
        tmp = maxseqinc(&num[i], N-19+i, 4, 19);
        if(tmp > max) {
            max = tmp;
            printf("new max is %d, left diag %d\n", max, i);
        }
    }

    printf("max %d\n", max);
}
