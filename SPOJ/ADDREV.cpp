#include<iostream>
using namespace std;
int main(){
    int t,a,b,c,suma,sumb,sum;
    cin>>t;
    while(t--){
        cin>>a>>b;
        suma = 0;
        while (a){
            suma = suma*10;
            suma += a%10;
            a=a/10;
        }
        sumb = 0;
        while (b){
            sumb = sumb*10;
            sumb += b%10;
            b=b/10;
        }
        c = suma+sumb;
        sum = 0;
        while (c){
            sum = sum*10;
            sum += c%10;
            c=c/10;
        }
        cout<<sum<<endl;
    }
    return 0;
}