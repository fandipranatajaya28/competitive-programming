#include<cstdio>
using namespace std;
int main(){
	int t;
	int n,m,x,y,vert,horz;
	scanf("%d", &t);
	while(t--){
		scanf("%d%d",&n,&m);
		scanf("%d%d",&x,&y);
		vert=y-1;
		if((m-y)>vert){
			vert=m-y;
		}
		horz=x-1;
		if((n-x)>horz){
			horz=n-x;
		}
		printf("%d\n",vert+horz);
	}
	return 0;
}