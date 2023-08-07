#include<cstdio>
using namespace std;
int main(){
	char wd;
	double dist,cp,tp,cbs,ws;
	scanf("%lf %lf %lf %lf %c",&cp,&tp,&cbs,&ws,&wd);
	if(wd=='L'||wd=='R'){
        if(cp>tp){
            cbs=cbs*-1;
            if(wd=='L'){
                ws=ws*-1;
                dist=(tp-cp)/(cbs+ws);
                if(dist>0){
                    printf("%.6lf\n",dist);
                }
                else{
                    printf("Impossible\n");
                }
            }
            else{
                dist=(tp-cp)/(cbs+ws);
                if(dist>0){
                    printf("%.6lf\n",dist);
                }
                else{
                    printf("Impossible\n");
                }
            }
        }
        else if(cp<tp){
            if(wd=='L'){
                ws=ws*-1;
                dist=(tp-cp)/(cbs+ws);
                if(dist>0){
                    printf("%.6lf\n",dist);
                }
                else{
                    printf("Impossible\n");
                }
            }
            else{
                dist=(tp-cp)/(cbs+ws);
                if(dist>0){
                    printf("%.6lf\n",dist);
                }
                else{
                    printf("Impossible\n");
                }
            }
        }
        else{
            printf("0.000000\n");
        }
 
	}
	return 0;
}