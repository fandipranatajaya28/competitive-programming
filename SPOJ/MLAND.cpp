#include<cstdio>
#include<algorithm>
using namespace std;
 
#define EPS 1e-9
const int M = 825;
const int N = 125;
 
struct edge
{
	int u;
	int v;
	double a;
	double b;
	double val;
 
	bool operator < (edge other) const
	{
		return val < other.val;
	}
};
 
edge arr[M];
int pset[N];
int n, m, tc;
double t1, t2;
 
void _reset()
{
	for(int i = 0; i < n; i++) pset[i] = i;
}
 
int finds(int x)
{
	return x == pset[x] ? x : pset[x] = finds(pset[x]);
}
 
double work(double t)
{
	_reset();
	for(int i = 0; i < m; i++) arr[i].val = t * arr[i].a + arr[i].b;
	sort(arr, arr + m);
	double res = 0;
	for(int i = 0; i < m; i++)
	{
		int uu = arr[i].u;
		int vv = arr[i].v;
		if(finds(uu) != finds(vv))
		{
			pset[finds(uu)] = vv;
			res += arr[i].val;
		}
	}
	return res;
}
 
int main()
{
	scanf("%d", &tc);
	while(tc--)
	{
		scanf("%d %d", &n, &m);
		scanf("%lf %lf", &t1, &t2);
		for(int i = 0; i < m; i++)
		{
			scanf("%d %d %lf %lf", &arr[i].u, &arr[i].v, &arr[i].a, &arr[i].b);
		}
		while(t2 - t1 > EPS)
		{
			double m1 = t1 + (t2 - t1) / 3.0;
			double m2 = t1 + (t2 - t1) * 2.0 / 3.0;
			double r1 = work(m1);
			double r2 = work(m2);
			if(r1 + EPS > r2) t2 = m2;
			else t1 = m1;
		}
		printf("%.3lf %.3lf\n", t1, work(t1));
	}
 
	return 0;
}