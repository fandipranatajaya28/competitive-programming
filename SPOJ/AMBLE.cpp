#include<bits/stdc++.h>
using namespace std;
 
int main(){
    int n;
    const double inf = numeric_limits<double>::max(); // an infinite number
    pair<int,int>temp;
    vector<pair<int,int> >titik;
    scanf("%d",&n);
    double dp[n][n], dist[n][n];
    for(int i=0; i<n; i++){
        scanf("%d %d", &temp.first, &temp.second);
        titik.push_back(temp);
    }
    sort(titik.begin(),titik.end());
    for(int i=0; i<n; i++){
        for(int j=0; j<n; j++){
            dist[i][j]=sqrt((titik[i].first-titik[j].first)*(titik[i].first-titik[j].first)+(titik[i].second-titik[j].second)*(titik[i].second-titik[j].second));
            dp[i][j]=inf;
        }
    }
    dp[0][0]=0;
    for (int i = 1; i < n; ++i){
		dp[0][i] = dp[0][i-1] + dist[i-1][i];
	}
	for (int i = 1; i < n; i++){
		for (int j = i; j < n; j++){
		  	if (i == j || i == j - 1){
		        for (int k = 0; k < i; k++){
		          dp[i][j] = min(dp[i][j], dp[k][i] + dist[k][j]);
				}
		    }
			else{
		    	dp[i][j] = dp[i][j-1] + dist[j-1][j];
		  	}
		}
	}
 
    printf("%.2lf\n", dp[n-1][n-1]);
    return 0;
}