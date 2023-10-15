#include<bits/stdc++.h>
using namespace std;

class Solution {
public:
    struct TwoSum
    {
        int value;
        int index;

        bool operator() (TwoSum a, TwoSum b){
            return a.value < b.value;
        }
    }T;
    vector<int> twoSum(vector<int>& nums, int target) {
        vector<TwoSum> twoSum_;
        vector<int> result;
                
        for(int i=0; i<nums.size(); i++){
            TwoSum temp;
            temp.index = i;
            temp.value = nums[i];
            twoSum_.push_back(temp);
        }

        sort(twoSum_.begin(), twoSum_.end(), T);

        int left = 0;
        int right = twoSum_.size() - 1;

        while (left < right){
            if (twoSum_[left].value + twoSum_[right].value == target){
                result.push_back(twoSum_[left].index);
                result.push_back(twoSum_[right].index);
                break;
            }
            else if (twoSum_[left].value + twoSum_[right].value < target){
                left++;
            }
            else{
                right--;
            }
        }
        return result;
    }
};