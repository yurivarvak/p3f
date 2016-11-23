
#include <stdio.h>
#include <time.h>
#include <stdint.h>
#include <string.h>

int is_prime(int n, int *primes) 
{
	int i = 0, p;
	do {
		p = primes[i++];
		if (n%p == 0)
			return 0;
	}
	while ((int64_t)p*p < n);
	return 1;
}

#define MAX_PRIMES 10000000
static int p[MAX_PRIMES];
int primes(int n)
{
	
	int m = n/6 + 1;
	int i, np = 3;  /* start with 3 primes */
	memset(p, 0, MAX_PRIMES*sizeof(int));
	p[0] = 2; p[1] = 3; p[2] = 5;
	for (i = 7; i < m; i += 2)
	{
		if (is_prime(i, p))
		{
			p[np++] = i;
			if (np >= MAX_PRIMES)
				return 0;
		}
	}
	printf("nprimes = %d\n", np);
	return np;
}

typedef struct { int64_t count, sum; } result;
 
result prime3factor(int n, int pn) 
{
	result r = {0, 0};
	int64_t num = n;
	int ai, bi, ci;
	for (ai = 0; ai < pn-2; ai++) 
	{
		int64_t a, ax;
		ax = a = p[ai];
		do {
			for (bi = ai+1; bi < pn-1; bi++)
			{
				int64_t b = p[bi];
				int64_t axby = ax * b;
				if (axby >= num)
					break;
				do {
					for (ci = bi + 1; ci < pn; ci++)
					{
						int64_t c = p[ci];
						int64_t nn = axby * c;
						if (nn >= num)
							break;
						do {
							r.count++;
							r.sum += nn;
							nn *= c;
						} while (nn < num);
					}
					axby *= b;
				} while (axby < num && axby*p[bi+1] < num);
			}
			ax *= a;
		} while (ax < num && ax*p[ai+1] < num && ax*p[ai+1]*p[ai+2] < num);
	}
	return r;
}

int main()
{
	int n = 1000000000;
	clock_t st = clock();
	int pn = primes(n);
	clock_t pt = clock();
	result r;
	printf("setup: %5.3lf\n", ((double)pt-st) / CLOCKS_PER_SEC);
	r = prime3factor(n, pn);
	printf("calc / counter / sum : %5.3lf / %lld / %lld\n", ((double)clock()-pt) / CLOCKS_PER_SEC, r.count, r.sum);
	return 0;
}
