#include "Playground.h"

int arr[3]{ 3,3,3 };

bool isFillable(int n) {
	if (n == 0) {
		return true;
	}
	bool first = false;
	bool second = false;
	bool third = false;
	if (n >= arr[0]) {
		first = isFillable(n - arr[0]);
	}
	if (n >= arr[1]) {
		second = isFillable(n - arr[1]);
	}
	if (n >= arr[2]) {
		third = isFillable(n - arr[2]);
	}
	return first || second || third;
}
