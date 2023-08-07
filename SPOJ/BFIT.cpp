#include<cstdio>
#include<cassert>
#define MAX 10005
using namespace std;
 
int getNum(){
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
 
int N;
int s[MAX];
 
int main(){
    for(int cases=getNum(); cases--;){
        N=getNum();
        double xbar=(N+1)/2.0;
        double ybar=0;
        for(int i=1; i<=N; i++){
            s[i]=getNum();
            ybar+=s[i];
        }
        ybar/=N;
        double den=((1.0*N*(N+1)*(2*N+1))/6 - (1.0*N*(N+1)*(N+1))/4);
        double num=0.0;
        for(int i=1; i<=N; i++){
            num+=((i-xbar)*(s[i]-ybar));
        }
        double beta=num/den;
        double alpha=ybar-beta*xbar;
        printf("%.4lf %.4lf\n",beta,alpha);
    }
    return 0;
}