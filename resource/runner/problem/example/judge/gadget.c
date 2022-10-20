#include<stdio.h>
void construct() __attribute__((constructor(101)));
void construct() { puts("Greetings from gadgets..."); }
