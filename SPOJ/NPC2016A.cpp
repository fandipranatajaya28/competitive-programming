#include<cstdio>
#include<cmath>
using namespace std;
int main(){
	int x,y,n;
	double dist;
	scanf("%d %d %d", &x, &y, &n);
	dist=2*(n*sqrt(2));
	printf("%lf\n", dist);
    return 0;
}