#include<bits/stdc++.h>
#define MOD 1000000007
using namespace std;
 
int t,tc,m,n,flag,acuan,awal,dp[3576225],batas[126],ub[126];
 
int GetNum(){
	int res = 0;
	char c;
	c = getchar_unlocked();
	while(c == '\n' || c == ' ') {
		c = getchar_unlocked();
	}
	res += (c - '0');
	c = getchar_unlocked();
	while(c != '\n' && c != ' ') {
		res = res * 10 + (c - '0');
		c = getchar_unlocked();
	}
	return res;
}
 
int main(){
	t=GetNum();
	awal=1;
    batas[3]=0;
    ub[3]=333333;
    for(int j=4; j<=125; j++){
    	ub[j]=1000000/(j-1);
        batas[j]=batas[j-1]+1+(1000000/(j-1))-(j*(j-1)/2);
    }
	acuan=3;
    for(tc=1; tc<=t; tc++){
    	flag=1;
    	m=GetNum();
        n=GetNum();
		if(n>125 || m<n*(n+1)/2){
        	printf("Case %d: 0\n",tc);
        }
        else if(n==2){
            printf("Case %d: %d\n",tc,(m-1)/2);
        }
        else{
            if(n<acuan && awal==0){
           		flag=0;
            }
            if(flag==1){
	            for(int i=acuan; i<=n; i++){
	                for(int j=i*(i+1)/2; j<=ub[i]; j++){
						if(i==3){
							dp[j-i*(i+1)/2+batas[i]]=(j-4)/2;
	                    }
	                    else{
	                    	dp[j-i*(i+1)/2+batas[i]]=dp[j-i*(i-1)/2+batas[i-1]-i];
	                    }
	                    if(j>=i*(i+1)/2+i){
	                    	dp[j-i*(i+1)/2+batas[i]]=(dp[j-i*(i+1)/2+batas[i]]+dp[j-i*(i+1)/2+batas[i]-i])%MOD;
						}
	                }
	            }
	    		acuan=n+1;
	        }
	       	printf("Case %d: %d\n",tc,dp[m-n*(n+1)/2+batas[n]]);
	    }
	    awal=0;
    }
    return 0;
}  