#include <omp.h>
#include <stdio.h>
#include <stdlib.h>  
#include <time.h>
#include "2DArray.h"

#define GRAIN  1024 /* product size below which matmultleaf is used */  

void seqMatMult(int m, int n, int p, int** A, int** B, int** C)
{
  for (int i = 0; i < m; i++)
  		for (int j = 0; j < n; j++)
		{
			C[i][j] = 0.0;
			for (int k = 0; k < p; k++) {
				C[i][j] += A[i][k]*B[k][j]; 
			          //printf("%lf %lf\t",A[i][k],B[k][j]);//added to test code
	        }
	     }
}
 
  
void matmultleaf(int mf, int ml, int nf, int nl, int pf, int pl, int **A, int **B, int **C) 
/*  
  subroutine that uses the simple triple loop to multiply  
  a submatrix from A with a submatrix from B and store the  
  result in a submatrix of C.
*/
// mf, ml; /* first and last+1 i index */
// nf, nl; /* first and last+1 j index */
// pf, pl; /* first and last+1 k index */
{   
	
	for (int i = mf; i < ml; i++)
		for (int j = nf; j < nl; j++) {
		    C[i][j] = 0;
            for (int k = pf; k < pl; k++)
				C[i][j] += A[i][k]*B[k][j];
        }
} 
  

void copyQtrMatrix(int **X, int m, int **Y, int mf, int nf)
{
	
	for (int i = 0; i < m; i++) 
		X[i] = &Y[mf+i][nf];
}

void AddMatBlocks(int **T, int m, int n, int **X, int **Y)
{
	
	for (int i = 0; i < m; i++)
		for (int j = 0; j < n; j++)
			T[i][j] = X[i][j] + Y[i][j];
}

void SubMatBlocks(int **T, int m, int n, int **X, int **Y)
{
	
	for (int i = 0; i < m; i++)
		for (int j = 0; j < n; j++)
			T[i][j] = X[i][j] - Y[i][j];
}


void strassenMMult(int mf, int ml, int nf, int nl, int pf, int pl, int **A, int **B, int **C)
{
	if ((ml-mf)*(nl-nf)*(pl-pf) < GRAIN)
	      matmultleaf(mf, ml, nf, nl, pf, pl, A, B, C); 

	else {
		int m2 = (ml-mf)/2;
		int n2 = (nl-nf)/2;
		int p2 = (pl-pf)/2;

		int **M1 = Allocate2DArray< int >(m2, n2);
		int **M2 = Allocate2DArray< int >(m2, n2);
		int **M3 = Allocate2DArray< int >(m2, n2);
		int **M4 = Allocate2DArray< int >(m2, n2);
		int **M5 = Allocate2DArray< int >(m2, n2);
		int **M6 = Allocate2DArray< int >(m2, n2);
		int **M7 = Allocate2DArray< int >(m2, n2);

		int **A11 = new int*[m2];
		int **A12 = new int*[m2];
		int **A21 = new int*[m2];
		int **A22 = new int*[m2];

		int **B11 = new int*[p2];
		int **B12 = new int*[p2];
		int **B21 = new int*[p2];
		int **B22 = new int*[p2];

		int **C11 = new int*[m2];
		int **C12 = new int*[m2];
		int **C21 = new int*[m2];
		int **C22 = new int*[m2];

		int **tAM1 = Allocate2DArray< int >(m2, p2);
		int **tBM1 = Allocate2DArray< int >(p2, n2);
		int **tAM2 = Allocate2DArray< int >(m2, p2);
		int **tBM3 = Allocate2DArray< int >(p2, n2);
		int **tBM4 = Allocate2DArray< int >(p2, n2);
		int **tAM5 = Allocate2DArray< int >(m2, p2);
		int **tAM6 = Allocate2DArray< int >(m2, p2);
		int **tBM6 = Allocate2DArray< int >(p2, n2);
		int **tAM7 = Allocate2DArray< int >(m2, p2);
		int **tBM7 = Allocate2DArray< int >(p2, n2);

		copyQtrMatrix(A11, m2, A, mf, pf);
		copyQtrMatrix(A12, m2, A, mf, p2);
		copyQtrMatrix(A21, m2, A, m2, pf);
		copyQtrMatrix(A22, m2, A, m2, p2);

		copyQtrMatrix(B11, p2, B, pf, nf);
		copyQtrMatrix(B12, p2, B, pf, n2);
		copyQtrMatrix(B21, p2, B, p2, nf);
		copyQtrMatrix(B22, p2, B, p2, n2);

		copyQtrMatrix(C11, m2, C, mf, nf);
		copyQtrMatrix(C12, m2, C, mf, n2);
		copyQtrMatrix(C21, m2, C, m2, nf);
		copyQtrMatrix(C22, m2, C, m2, n2);

	// M1 = (A11 + A22)*(B11 + B22)
		AddMatBlocks(tAM1, m2, p2, A11, A22);
		AddMatBlocks(tBM1, p2, n2, B11, B22);
		strassenMMult(0, m2, 0, n2, 0, p2, tAM1, tBM1, M1);

	//M2 = (A21 + A22)*B11
		AddMatBlocks(tAM2, m2, p2, A21, A22);
		strassenMMult(0, m2, 0, n2, 0, p2, tAM2, B11, M2);

	//M3 = A11*(B12 - B22)
		SubMatBlocks(tBM3, p2, n2, B12, B22);
		strassenMMult(0, m2, 0, n2, 0, p2, A11, tBM3, M3);

	//M4 = A22*(B21 - B11)
		SubMatBlocks(tBM4, p2, n2, B21, B11);
		strassenMMult(0, m2, 0, n2, 0, p2, A22, tBM4, M4);

	//M5 = (A11 + A12)*B22
		AddMatBlocks(tAM5, m2, p2, A11, A12);
		strassenMMult(0, m2, 0, n2, 0, p2, tAM5, B22, M5);

	//M6 = (A21 - A11)*(B11 + B12)
		SubMatBlocks(tAM6, m2, p2, A21, A11);
		AddMatBlocks(tBM6, p2, n2, B11, B12);
		strassenMMult(0, m2, 0, n2, 0, p2, tAM6, tBM6, M6);

	//M7 = (A12 - A22)*(B21 + B22)
		SubMatBlocks(tAM7, m2, p2, A12, A22);
		AddMatBlocks(tBM7, p2, n2, B21, B22);
		strassenMMult(0, m2, 0, n2, 0, p2, tAM7, tBM7, M7);

		
		for (int i = 0; i < m2; i++)
			for (int j = 0; j < n2; j++) {
				C11[i][j] = M1[i][j] + M4[i][j] - M5[i][j] + M7[i][j];
				C12[i][j] = M3[i][j] + M5[i][j];
				C21[i][j] = M2[i][j] + M4[i][j];
				C22[i][j] = M1[i][j] - M2[i][j] + M3[i][j] + M6[i][j];
			}

		Free2DArray< int >(M1);
		Free2DArray< int >(M2);
		Free2DArray< int >(M3);
		Free2DArray< int >(M4);
		Free2DArray< int >(M5);
		Free2DArray< int >(M6);
		Free2DArray< int >(M7);

		delete[] A11; delete[] A12; delete[] A21; delete[] A22;
		delete[] B11; delete[] B12; delete[] B21; delete[] B22;
		delete[] C11; delete[] C12; delete[] C21; delete[] C22;

		Free2DArray< int >(tAM1);
		Free2DArray< int >(tBM1);
		Free2DArray< int >(tAM2);
		Free2DArray< int >(tBM3);
		Free2DArray< int >(tBM4);
		Free2DArray< int >(tAM5);
		Free2DArray< int >(tAM6);
		Free2DArray< int >(tBM6);
		Free2DArray< int >(tAM7);
		Free2DArray< int >(tBM7);
	}
}
           
void matmultS(int m, int n, int p, int **A, int **B, int **C)
{	
  int i,j;	
  for (i=0; i < m; i++) 
	 for (j=0; j < n; j++)
	 C[i][j] = 0;   
	strassenMMult(0, m, 0, n, 0, p, A, B, C);
}


int CheckResults(int m, int n, int **C, int **C1)
{
#define THRESHOLD 0.001
//
//  May need to take into consideration the floating point roundoff error
// due to parallel execution
//
  for (int i = 0; i < m; i++) {
	 for (int j = 0; j < n; j++) {
	if (abs(C[i][j] - C1[i][j]) > THRESHOLD ) {
	  printf("C[%d][%d]:%d  %d\n",i,j, C[i][j], C1[i][j]);
	  return 1;
	}
	 }
  }
  return 0;
}


  
int main(int argc, char* argv[])
{	
  double before, after;

  int M = atoi(argv[1]);
  int N = atoi(argv[2]);
  int P = atoi(argv[3]);

  int **A = Allocate2DArray< int >(M, P);
  int **B = Allocate2DArray< int >(P, N);
  int **C = Allocate2DArray< int >(M, N);
  int **C4 = Allocate2DArray< int >(M, N);

  int i, j;

  for (i = 0; i < M; i++) {
 for (j = 0; j < P; j++) {   
		A[i][j] = ((int)(rand()%100) /10);  
	 } 
  } 

  for (i = 0; i < P; i++) {	
	 for (j = 0; j < N; j++) {
	B[i][j] = ((int)(rand()%100) / 10.0);   
	 } 
  } 

  printf("Execute Standard matmult\n\n");
  before = omp_get_wtime();
  seqMatMult(M, N, P, A, B, C);
  after = omp_get_wtime();
  printf("Standard matrix function done in %10f secs\n\n\n",(after - before));

  before = omp_get_wtime();
  matmultS(M, N, P, A, B, C4);
  after = omp_get_wtime();
  printf("Strassen matrix function done in %10f secs\n\n\n",(after - before));

  if (CheckResults(M, N, C, C4)) 
	 printf("Error in matmultS\n\n");
  else
	 printf("OKAY\n\n");

  Free2DArray(A);
  Free2DArray(B);
  Free2DArray(C);
  Free2DArray(C4);

  return 0;
}
