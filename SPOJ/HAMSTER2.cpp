#include<bits/stdc++.h>
using namespace std;
 
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
    int t;
    t=GetInt();
    while(t--){
        int n, i, j, k, ans, temp;
        double v, x, y1, y2, tan1, tan2, tan3, tan4, batasmin[40000], batasmax[40000], atas, bawah;
        v=(double)GetInt();
        n=GetInt();
        i=0;
        while(n--){
            x=(double)GetInt();
            y1=(double)GetInt();
            y2=(double)GetInt();
            tan1 = ((v*v)-sqrt(v*v*v*v-20*y1*v*v-100*x*x))/(10*x);
            tan2 = ((v*v)+sqrt(v*v*v*v-20*y1*v*v-100*x*x))/(10*x);
            tan3 = ((v*v)-sqrt(v*v*v*v-20*y2*v*v-100*x*x))/(10*x);
            tan4 = ((v*v)+sqrt(v*v*v*v-20*y2*v*v-100*x*x))/(10*x);
            if(isnan(tan1)==0 && isnan(tan2)==0){
                if(isnan(tan3)==1 || isnan(tan4)==1){
                    batasmin[i]=tan1;
                    batasmax[i]=tan2;
                    i++;
                }
                else{
                    batasmin[i]=tan1;
                    batasmax[i]=tan3;
                    i++;
                    batasmin[i]=tan4;
                    batasmax[i]=tan2;
                    i++;
                }
            }
        }
        sort(batasmin, batasmin+i);
        sort(batasmax, batasmax+i);
        ans=0;
        temp=0;
        j=0;
        k=0;
        while(1){
            atas=batasmax[j];
            bawah=batasmin[k];
            if(bawah<=atas){
                temp++;
                k++;
            }
            else{
                temp--;
                j++;
            }
            if(temp>ans){
                ans=temp;
            }
            if(k==(i-1))break;
        }
        printf("%d\n",ans);
    }
    return 0;
}