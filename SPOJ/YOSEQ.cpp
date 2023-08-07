#include<bits/stdc++.h>
using namespace std;
const int N = 100001;
string nums;
int lnums, arr[N];
bool stats[N];
 
int main(){
    cin.sync_with_stdio(0);
    cin.tie(0); cout.tie(0);
    cin>>nums;
    lnums = nums.size();
    for(int ii = 0; ii<lnums; ii++)
        arr[ii] = nums[ii] - '0';
    for(int ii = 0; ii<lnums; ii++){
        int curr = 0;
        for(int jj = 0; jj<5;   jj++){
            curr = curr * 10 + arr[ii+jj];
            stats[curr] = true;
        }
    }
    for(int ii = 0; ii<=1000000; ii++){
        if(stats[ii] == false){
            printf("%d\n",ii);
            ii = 1000001;
        }
    }
    return 0;
}