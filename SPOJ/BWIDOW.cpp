#include<iostream>
using namespace std;
 
int main(){
    int t,n,ans,r1[1007],r2[1007],index,max_r1,max_r2;
    cin>>t;
    while(t--){
        cin>>n;
        ans=0;
        for(int i=1; i<=n; i++){
            cin>>r1[i]>>r2[i];
        }
        max_r1=0;
        max_r2=0;
        for(int i=1; i<=n; i++){
            if(r1[i]>max_r1){
                max_r1=r1[i];
                index=i;
            }
        }
        for(int i=1; i<=n; i++){
            if(r2[i]>max_r2 && i!=index){
                max_r2=r2[i];
            }
        }
        if(max_r2>max_r1)   cout<<-1<<endl;
        else    cout<<index<<endl;
    }
    return 0;
}