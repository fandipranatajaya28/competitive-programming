#include<bits/stdc++.h>
using namespace std;

class Solution {
public:
    int search(vector<int>& nums, int target) {
        int left = 0;
        int right = nums.size() - 1;
        int mid = (left + right) / 2;

        while (left <= right){
            if (nums[mid] == target){
                return mid;
            }
            else if (nums[mid] < target){
                left = mid + 1;
                mid = (left + right) / 2;
            }
            else{
                right = mid -1;
                mid = (left + right) / 2;
            }
        }

        return -1;
    }
};

int main(){
    vector<int>nums = {1, 2, 3, 4, 5};
	int firstTarget = 1;
	int secondTarget = 5;
    Solution solution_;
    cout<<solution_.search(nums, firstTarget)<<endl;
    cout<<solution_.search(nums, secondTarget)<<endl;
    return 0;
}