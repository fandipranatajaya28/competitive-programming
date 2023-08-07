#include<cstdio>
#include<algorithm>
using namespace std;
 
int main(){
    int t,a,b,x,y;
    scanf("%d",&t);
    while(t--){
        scanf("%d %d",&a,&b);
        x = b / __gcd(a,b);
        y = a / __gcd(a,b);
        printf("%d %d\n",x,y);
    }
    return 0;
}