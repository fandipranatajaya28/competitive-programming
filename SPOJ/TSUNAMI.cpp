#include <bits/stdc++.h>
using namespace std;
#define fi first
#define se second
#define mp make_pair
#define MAX_N 1001
 
pair <int,int> pii[1005];
vector <pair <double,pair <int,int> > > vc;
double dist (pair <double, double> a,pair <double,double> b)
{
	return sqrt ((a.fi - b.fi) * (a.fi-b.fi) + (a.se - b.se) * (a.se - b.se));
}
 
int par[1005];
 
int findSet(int a)
{
	if (par[a] == a) return a;
	return par[a] = findSet(par[a]);
}
 
void initSet()
{
	for(int i=0; i < MAX_N; i++)
	{
		par[i] = i;
	}
}
 
void unionSet(int i, int j)
{
	par[findSet(i)] = findSet(j);
}
 
int sameSet(int i, int j)
{
	return findSet(i) == findSet(j);
}
 
int main()
{
	int n;
	while (scanf("%d", &n) != EOF)
	{
		if (n == 0) break;
		initSet();
		for (int i = 1; i <= n; i++)
		{
			scanf("%d %d", &pii[i].se, &pii[i].fi);
		}
		sort(pii+1,pii+n+1);
		
		double ans = 0.0;
		int prev = 0;
		for (int i = 1; i <= n+1; i++)
		{
			if (pii[i].fi != prev)
			{
				sort (vc.begin(), vc.end());
				for (int j = 0; j < vc.size(); j++)
				{
					int u = vc[j].se.fi, v = vc[j].se.se;
					if (!sameSet(u,v))
					{
						unionSet(u,v);
						ans += vc[j].fi;
					}
				}
				vc.clear();
				prev = pii[i].fi;
			}
			
			for(int j = 1; j < i; j++)
			{
				vc.push_back( mp( dist(pii[i],pii[j]),mp(i,j) ));
			}				
		}
		vc.clear();
		printf("%.2lf\n", ans);
	}
	
	return 0;
}