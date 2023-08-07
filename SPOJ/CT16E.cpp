#include <bits/stdc++.h>
using namespace std;
double dp[1 << 18];
double p[18][18];
int n;
 
int main()
{
	while(cin >> n)
	{
		for (int i = 0; i < n; i++)
		{
			for(int j = 0; j<n; j++)
			{
				cin >> p[i][j];
			}
		}
 
		dp[(1<<n)-1] = 1;
		for (int cond = (1<<n)-1; cond > 0; cond--)
		{
			int num = __builtin_popcount(cond);
			for(int i=0; i < n; i++)
			{
				if(!((cond >> i) & 1) ) continue;
				for(int j = i+1; j < n; j++)
				{
					if(((cond >> j) & 1) )
					{
						int temp = cond & ~(1 << i);
						dp[temp] += p[j][i] * dp[cond] / (num * (num-1) / 2);
						temp = cond & ~(1 << j);
						dp[temp] += p[i][j] * dp[cond] / (num * (num-1) / 2);
					}
				}
			}
		}
 
		for(int i = 0; i < n; i++)
		{
			cout << dp[1 << i] << " ";
		}
 
		cout << endl;
	}
 
	return 0;
}