#include<bits/stdc++.h>
using namespace std;
 
long long GetNum(){
    long long result=0;
    char character;
    while(1){
        character=getchar_unlocked();
        if(character==' '||character=='\n') continue;
        else break;
    }
    result=character-'0';
    bool isDecimal=false;
    int fractionalCount=0;
    while(1){
        character=getchar_unlocked();
        fractionalCount++;
        if(character>='0' && character<='9') result=10*result + character-'0';
        else if(character=='.'){
        	fractionalCount=0;
        	isDecimal=true;
		}
        else break;
    }
    while(isDecimal && fractionalCount<=4){
    	result=10*result;
    	fractionalCount++;
	}
	if(!isDecimal) result=10000*result;
    return result;
}
 
int main(){
	int t;
	scanf("%d", &t);
	while(t--){
		int n;
		long long sumX, sumY, divisor, a, b;
		scanf("%d", &n);
		sumX = 0; sumY = 0;
		while(n--){
			long long x,y;
			x = GetNum();
			y = GetNum();
			sumX += x;
			sumY += y;
		}
		divisor = __gcd(sumX, sumY);
		a = sumY / divisor;
		b = sumX / divisor;
		printf("f(x) = %lldx, g(x) = %lldx\n", a, b);
	}
} 