#include<bits/stdc++.h>
using namespace std;

class Solution {
public:
    bool checkValidString(string s) {
        vector<int> stack, wildcard;

        for (int idx = 0; idx<s.length(); idx++){
            if (s[idx] == '('){
                stack.push_back(idx);
                continue;
            } else if (s[idx] == '*'){
                wildcard.push_back(idx);
                continue;
            }
            if (stack.size() == 0 && wildcard.size() == 0){
                return false;
            }
            if (stack.size() > 0){
                stack.pop_back();
                continue;
            }
            wildcard.pop_back();
        }

        int index = 0;
        for (int i = 0; i<stack.size(); i++){
            bool found = false;
            for (int j = index; j<wildcard.size(); j++){
                if (wildcard[j] > stack[i]){
                    index = j+1;
                    found = true;
                    break;
                }
            }
            if (!found) return false;
        }
        
        return true;
    }
};

int main(){
    Solution solution_;
    
    string s1 = "((((()(()()()*()(((((*)()*(**(())))))(())()())(((())())())))))))(((((())*)))()))(()((*()*(*)))(*)()";
	cout<<solution_.checkValidString(s1)<<endl;
	// true

	string s2 = "(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())";
	cout<<solution_.checkValidString(s2)<<endl;
    // false

    return 0;
}