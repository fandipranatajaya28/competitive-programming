#include<bits/stdc++.h>
using namespace std;
 
stack<int> mystack;
int main(){
	int t,i;
	char x[401];
	scanf("%d",&t);
	for(i=1;i<=t;i++){
		scanf("%s",x);
		int a=strlen(x);
		int j;
		for(j=0;j<a;j++){
			if(x[j]=='+' || x[j]=='-' || x[j]=='*' || x[j]=='/' || x[j]=='^'){
				mystack.push(x[j]);
			}
			else if(x[j]==')'){
				printf("%c", mystack.top());
				mystack.pop();
			}
			else if(x[j]!='('){
				printf("%c",x[j]);
			}
		}
		printf("\n");
	}
}