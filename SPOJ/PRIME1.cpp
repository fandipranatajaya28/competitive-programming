#include <bits/stdc++.h> 
using namespace std;
 
int binpower(int base, int e, int mod) {
    int result = 1;
    base %= mod;
    while (e) {
        if (e & 1)
            result = (long long)result * base % mod;
        base = (long long)base * base % mod;
        e >>= 1;
    }
    return result;
}
 
bool check_composite(int n, int a, int d, int s) {
    int x = binpower(a, d, n);
    if (x == 1 || x == n - 1)
        return false;
    for (int r = 1; r < s; r++) {
        x = (long long)x * x % n;
        if (x == 1)	 return true;
		if (x == n - 1)
            return false;
    }
    return true;
}
 
bool MillerRabin(int n) { // returns true if n is prime, else returns false.
 
    // Corner cases 
	if (n <= 1 || n == 4) return false; 
	if (n <= 3) return true;
	if ((n & 1) == 0) return false;
 
    int r = 0;
    int d = n - 1;
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
    int m,n,t;
    scanf("%d",&t);
    while(t--){
        scanf("%d %d",&m,&n);
        if (m & 1){
        	if (m==1){
        		cout<<"2"<<endl;
        		m=3;
			}
        	for (int p = m; p <= n; p+=2){
	        	if(MillerRabin(p)){
					cout<<p<<endl;
				}
			}
		}
		else{
			if (m==2){
				cout<<"2"<<endl;
				m=3;
			}
			else{
				m++;
			}
			for (int p = m; p <= n; p+=2){
	        	if(MillerRabin(p)){
					cout<<p<<endl;
				}
			}
		}
    }
    return 0;
} 