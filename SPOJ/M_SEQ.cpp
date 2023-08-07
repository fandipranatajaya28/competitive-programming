#include<cstdio>
using namespace std;
int main(){
    int t;
    double n,ans;
    scanf("%d",&t);
    while(t--){
        scanf("%lf",&n);
        ans=(2.0*n+1.0)/n;
        printf("%.8lf\n",ans);
    }
    return 0;
}