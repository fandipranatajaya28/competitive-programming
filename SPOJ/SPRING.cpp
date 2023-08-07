#include<bits/stdc++.h>
using namespace std;
#define oo INT_MAX
#define pb push_back
#define mp make_pair
int N, M, D;
vector<pair<double,int> > Node[20002];
 
double Dijkstra(int from, int to){
    vector<double> dist(N+2,oo);
    priority_queue< pair<double, int>, vector<pair<double, int> >, greater<pair<double,int> > > Q;
    //aturan untuk mensort priority queue secara ascending dengan greater
    dist[from] = 0.0;
    Q.push(make_pair(0.0,from));
    while (!Q.empty()){
        pair<double,int> curr = Q.top();
        Q.pop();
        double ftemp = curr.first;
        int stemp = curr.second;
        if(stemp == to) return dist[to];
        int node_size = Node[stemp].size();
        for(int i = 0; i<node_size; i++){
            double cost = Node[stemp][i].first;
            int next = Node[stemp][i].second;
            if(cost + ftemp < dist[next]){
                dist[next] = cost + ftemp;
                Q.push(make_pair(dist[next],next));
            }
        }
    }
    return -1;
}
 
int main(){
    while(scanf("%d%d%d", &N, &M, &D) == 3){
        memset(Node,0.0,sizeof(Node));
        while(M--){
            int a, b, w;
            scanf("%d%d%d", &a, &b, &w);
            double w_ratio = (double)1/w;
            Node[a].pb(mp((double)w_ratio,b));
            Node[b].pb(mp((double)w_ratio,a));
        }
 
        double result = Dijkstra(0,N-1);
        double Fmax = D / result;
        if(Fmax<0.0)    Fmax = 0.0;
        printf("%.2f\n",Fmax);
    }
    return 0;
}