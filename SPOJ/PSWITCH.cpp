#include<cstdio>
using namespace std;
int main(){
    int lamp[8][6] = {0, 0, 0, 0, 0, 0,
                      0, 0, 1, 1, 1, 0,
                      0, 1, 0, 1, 0, 1,
                      0, 1, 1, 0, 1, 1,
                      1, 0, 0, 1, 0, 0,
                      1, 0, 1, 0, 1, 0,
                      1, 1, 0, 0, 0, 1,
                      1, 1, 1, 1, 1, 1};
    bool checkLamp[8], cek = false;
    int N, C, on[2000], off[2000], a, b;
    for(int i = 0; i<8; i++){
        checkLamp[i] = true;
    }
    scanf("%d", &N);
    scanf("%d", &C);
    if(C == 1){
        checkLamp[1] = false;
        checkLamp[4] = false;
        checkLamp[6] = false;
        checkLamp[7] = false;
    }
    else if(C == 2)
        checkLamp[3] = false;
    a=0;
    while(scanf("%d", &on[a]) && on[a] != -1){
        on[a] = on[a] - 1;
        on[a] = on[a] % 6;
        a++;
    }
    b=0;
    while(scanf("%d", &off[b]) && off[b] != -1){
        off[b] = off[b] - 1;
        off[b] = off[b] % 6;
        b++;
    }
    for(int i = 0; i < 8; i++)
        if(checkLamp[i])
            for(int j = 0; j < a; j++)
                if(lamp[i][on[j]] == 0)
                    checkLamp[i] = false;
    for(int i = 0; i < 8; i++)
        if(checkLamp[i])
            for(int j = 0; j < b; j++)
                if(lamp[i][off[j]] == 1)
                    checkLamp[i] = false;
    for(int i = 0; i < 8; i++)
        if(checkLamp[i]){
            cek = true;
            for(int j = 0; j < N; j++)
                printf("%d", lamp[i][j%6]);
            printf("\n");
        }
    if(!cek) printf("Impossible\n");
    return 0;
}