#include<cstdio>
using namespace std;
typedef long long LL;
int main(){
    int t;
    LL n,a,a1,a2,ans,sum;
    scanf("%d",&t);
    while(t--){
        scanf("%lld",&n);
        ans = 0;
        for(a=2; a*a<=n; a++){
            a1 = (a+1); //batas bawah
            a2 = n/a;   //batas atas
            if(a1<=a2){
                sum = (a2*(a2+1))/2 - (a1*(a1-1))/2; //untuk menambahkan deret aritmatika pasangan dari a
                ans += sum;
                ans += a*(a2-a1+1); //untuk mencari banyaknya deret jumlah a
            }
            ans += a; // untuk menambahkan nilai a waktu ketemu akar kuadrat n
        }
        ans += n-1; //untuk menambahkan total jumlah nilai 1
        printf("%lld\n",ans);
    }
    return 0;
}