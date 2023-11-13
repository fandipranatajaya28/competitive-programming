#include<bits/stdc++.h>
using namespace std;

class Solution {
public:
    int numIslands(vector<vector<char>>& grid) {
        // initiate visited coordinates
        vector<vector<bool>> visited(grid.size(), vector<bool>(grid[0].size(), false));

        int numOfIslands = 0;
        if (grid.size()>0){
            for (int row = 0; row < grid.size(); row++){
                for(int col = 0; col < grid[row].size(); col++){
                    if (grid[row][col] == '1' && !visited[row][col]){
                        bfs(row, col, grid, visited);
                        numOfIslands++;
                    }
                }
            }
        }
        return numOfIslands;
    }

    void bfs(int row, int col, vector<vector<char>>& grid, vector<vector<bool>>& visited){
        vector<vector<int>> directions = {
            // row, col
			{1, 0},  // down
			{-1, 0}, // up
			{0, 1},  // right
			{0, -1}, // left
        };
        queue<pair<int,int>> q;
        q.push({row, col});
        visited[row][col] = true;

        while (!q.empty()){
            int currentRow = q.front().first;
            int currentCol = q.front().second;
            q.pop();

            // traverse the neighbours
            for (int i = 0; i < directions.size(); i++){
                int r = currentRow + directions[i][0];
                int c = currentCol + directions[i][1];

                if (isValid(r, c, grid, visited)){
                    q.push({r, c});
                    visited[r][c] = true;
                }
            }
        }
    }

    bool isValid(int row, int col, vector<vector<char>>& grid, vector<vector<bool>>& visited){
        if (row >= grid.size() || row < 0) return false;
        else if (col >= grid[0].size() || col < 0) return false;
        else if (grid[row][col] == '0') return false;
        else if (visited[row][col]) return false;
        return true;
    }
};

int main(){
    Solution solution_;
    vector<vector<char>> grid = {
        {
			'1', '1', '0', '0', '0',
		},
		{
			'1', '1', '0', '0', '0',
		},
		{
			'0', '0', '1', '0', '0',
		},
		{
			'0', '0', '0', '1', '1',
		},
    };

    cout<<solution_.numIslands(grid)<<endl;
    // output: 3
    return 0;
}