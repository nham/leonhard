#include <stdio.h>

int main(void)
{
    int num[1000], cnt = 0, off=0, max=0, tmp;

    while(cnt <= 995) {
        tmp = 1;
        for(off = 0; off < 5; off++)
            if(num[cnt+off] != 0)
                tmp *= num[cnt+off]
            else
                break;
        // terminated early due to zero, so just advanced cnt past it
        if(off < 5)
            cnt += off+1;
        else if(tmp > max)
            max = tmp;


    }

    

}
