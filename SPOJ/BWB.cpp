#include <bits/stdc++.h>
using namespace std;
const int MOD=1000000007;
int black[100][10001], white[100][10001], in[100];
 
int GetInt(){
    int r=0;
    char c;
    while(1){
        c=getchar_unlocked();
        if(c==' '||c=='\n') continue;
        else break;
    }
    r=c-'0';
    while(1){
        c=getchar_unlocked();
        if(c>='0' && c<='9')    r=10*r + c-'0';
        else break;
    }
    return r;
}
 
int main(){
    int t,n,k,ans=0;
    memset(black, 0, sizeof(black[0][0]) * 100 * 10001);
    memset(white, 0, sizeof(white[0][0]) * 100 * 10001);
    for(int i=0; i<99; i++){
        in[i]=2;
    }
    t=GetInt();
    while(t--){
        n=GetInt();
        k=GetInt();
        if(k==1) ans=1;
        else{
            if(black[k-2][n]==0 || white[k-2][n]==0){
                black[k-2][1]=0;
                white[k-2][1]=1;
                for(int j=in[k-2]; j<=n; j++){
                    black[k-2][j]=(white[k-2][j-1]%MOD+black[k-2][j-1]%MOD)%MOD;
                    white[k-2][j]=black[k-2][j];
                    if(j>k) black[k-2][j]-=white[k-2][j-k];
                }
                in[k-2]=n;
            }
            ans=(black[k-2][n]%MOD+white[k-2][n]%MOD)%MOD;
            if(ans<0)ans+=MOD;
        }
        printf("%d\n", ans);
    }
    return 0;
}