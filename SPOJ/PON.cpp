#include <bits/stdc++.h> 
using namespace std;
typedef unsigned long long u64;
 
u64 binpower(u64 base, u64 e, u64 mod) {
    u64 result = 1;
    base %= mod;
    while (e) {
        if (e & 1)
            result = (__uint128_t)result * base % mod;
        base = (__uint128_t)base * base % mod;
        e >>= 1;
    }
    return result;
}
 
bool check_composite(u64 n, u64 a, u64 d, int s) {
    u64 x = binpower(a, d, n);
    if (x == 1 || x == n - 1)
        return false;
    for (int r = 1; r < s; r++) {
        x = (__uint128_t)x * x % n;
        if (x == 1)	return true;
        if (x == n - 1) return false;
    }
    return true;
}
 
bool MillerRabin(u64 n) { // returns true if n is prime, else returns false.
    // Corner cases 
	if (n <= 1 || n == 4) return false; 
	if (n <= 3) return true;
	if ((n & 1) == 0) return false;
 
    int r = 0;
    u64 d = n - 1;
    while ((d & 1) == 0) {
        d >>= 1;
        r++;
    }
    
    int base[12] = {2,3,5,7,11,13,17,19,23,29,31,37};
    
    for (int i = 0; i < 12; i++){
		if (n==base[i]) return true;
		if (check_composite(n, base[i], d, r)) 
			return false;
	}
	
    return true;
}
 
int main(){
    int t;
    scanf("%d",&t);
    while(t--){
        u64 n;
        cin>>n;
        if(MillerRabin(n)){
			printf("YES\n");
		}
		else{
			printf("NO\n");
		}
    }
    return 0;
} 