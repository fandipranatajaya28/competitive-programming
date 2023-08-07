#include<cstdio>
#include<iostream>
#include<algorithm>
#include<map>
#include<vector>
#include<cstring>
#include<string>
#include<cstdlib>
 
using namespace std;
 
enum DFSCOLOR {WHITE, GRAY, BLACK};
map< string, int> dfs_num;
map< string, vector<string> > graph;
 
bool TopoVisit(string u)
{
    int length;
    string next;
 
    dfs_num[u] = GRAY;
 
    length = graph[u].size();
    for(int i = 0; i < length; i++)
    {
        next=graph[u][i];
        if(dfs_num.find(next)==dfs_num.end())  return true;
        if(dfs_num[next] == WHITE)  TopoVisit(next);
        if(dfs_num[next] == GRAY) return false;
    }
    dfs_num[u] = BLACK;
    return true;
}
 
int main(){
    string p,q;
    vector<string> index;
    int counter=0;
    while(cin>>p){
        dfs_num[p]=WHITE;
        index.push_back(p);
        graph[p]=vector<string>();
        while(cin>>q && q!="0"){
            graph[p].push_back(q);
        }
    }
    for(int i = 0; i <index.size(); i++){
        if(dfs_num[index[i]] == WHITE) TopoVisit(index[i]);
    }
    for(int i = 0; i <index.size(); i++){
        if(dfs_num[index[i]] == BLACK) counter++;
    }
    printf("%d\n",counter);
    return 0;
}