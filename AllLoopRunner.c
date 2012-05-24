#include<stdlib.h>
#include<stdio.h>
#include<string.h>
int main(){
	int i = 64;
	char temp[100];
	for(;i<=2048;i*=2){
		printf("\nExecuting loop with i = %d\n",i);
		sprintf(temp,"python3 dataMaker.py --size %d > data1.csv",i);
		printf("Generating first matrix of size %d\n",i);
		system(temp);
		sprintf(temp,"python3 dataMaker.py --size %d > data2.csv",i);
		printf("Generating second matrix of size %d\n",i);
		system(temp);
		sprintf(temp,"./AllInOne >> OutputLog.txt",i);
		printf("Executing all implementations on these matrices\n");
	}
	return 0;
}
