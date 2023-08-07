#include<iostream>
#include<cmath>
using namespace std;
 
int main(){
    int t,n,temp,ans;
    cin>>t;
    while(t--){
        ans=0;
        cin>>n;
        for(int i=1; i<n; i++){
            temp=n/pow(5,i);
            ans+=temp;
            if(temp==0) break;
        }
        cout<<ans<<endl;
    }
    return 0;
}