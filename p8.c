#include <stdio.h>

void popseq(char *fname, char *arr)
{
    FILE *fp;
    char buf[100];
    int i, j = 0;

    fp = fopen(fname, "r");

    // WHY do i have to do 52 and not 51??
    while(fgets(buf, 52, fp) != NULL) {
        for(i = 0; i < 50; i++)
            arr[j+i] = buf[i] - '0';
        j += 50;
    }

}

int maxseqinc(char *arr, int len, int amt, int inc)
{
    int cnt = 0, off = 0, max = 0, tmp;
    int j;

    while(cnt <= len - inc*amt) {
        tmp = 1;
        for(off = 0; off < amt; off += inc)
            if(arr[cnt+off] != 0)
                tmp *= arr[cnt+off];
            else
                break;
        // terminated early due to zero, so just advanced cnt past it
        if(off < amt)
            cnt += off+inc;
        else if(tmp > max) {
            max = tmp;
            printf(" > max is now %d, @ cnt = %d\n", max, cnt);
            printf("cnt vals: ");
            for(j = 0; j < amt; j++)
                printf("%d,", arr[cnt+j*inc]);
            printf("\n");
        }

        cnt += inc;

    }

}

int main(void)
{
    char num[1000];
    int cnt = 0, off=0, max=0, tmp;
    popseq("p8.txt", &num[0]);

    maxseqinc(num, 1000, 5, 1);
}
