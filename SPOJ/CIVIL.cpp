#include<bits/stdc++.h>
using namespace std;
 
int main(){
    double h,w,f,ans;
    while(scanf("%lf %lf %lf", &h, &w, &f)){
        if(h==0 && w==0 && f==0) break;
        ans = w*w*(f*f-1.0)/(16*h);
        printf("%.3lf\n", ans);
    }
    return 0;
}