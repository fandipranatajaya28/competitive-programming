#include<iostream>
#include<algorithm>
using namespace std;
 
int main(){
    int t,k,n,h,arr[20002],i,suma,sumb,sum;
    cin>>t;
    while(t--){
        cin>>n>>k;
        for(i=0; i<n; i++){
            cin>>h;
            arr[i]=h;
        }
        sort(arr,arr+i);
        for(i=0; i<(n-k+1); i++){
            suma = arr[i+(k-1)] - arr[i];
            if(i==0) sum = suma;
            if(suma<sum) sum = suma;
        }
        cout<<sum<<endl;
    }
    return 0;
}