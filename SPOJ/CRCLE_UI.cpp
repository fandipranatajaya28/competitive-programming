#include<bits/stdc++.h>
using namespace std;
typedef long long ll;
ll MOD = 1000000007;
 
ll bigmod(char num[1007], ll mod)
{
    ll res = 0;
 
    for (ll i = 0; i < strlen(num); i++)
         res = (res*10 + (ll)num[i] - '0') %mod;
 
    return res;
}
 
ll modexp(ll b, ll e, ll m){
    ll r = 1;
    while(e>0){
        if((e&1)==1){
            r =(r*b)%m;
        }
        e>>=1;
        b = (b*b)%m;
    }
    return (ll)r;
}
 
int main(){
    int t;
    ll a, b, ans;
    char n[1007],k[1007];
    scanf("%d", &t);
    while(t--){
        scanf("%s %s", &n, &k);
        b = bigmod(n,MOD-1);
        a = bigmod(k,MOD)-1;
        ans = modexp(a,b,MOD);
        if((b&1) == 1){
            ans -= a;
        }
        else{
            ans += a;
        }
        ans = ans % MOD;
        if(ans<0) ans = ans + MOD;
        printf("%lld\n", ans);
    }
    return 0;
}