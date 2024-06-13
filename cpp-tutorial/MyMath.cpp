#include <cstdlib>

int Multiply(int a, int b)
{
	return a * b;
}

void Increment(int& value)
{
	value++;
}

float MySqrt(float x)
{
	float y = x;
	float z = (y + (x / y)) / 2;

	while (std::abs(y - z) >= 0.000001)
	{
		y = z;
		z = (y + (x / y)) / 2;
	}

	return z;
}

float Pythagoras(float a, float b)
{
	return MySqrt(a * a + b * b);
}

float RandomNumberBetween(float a, float b)
{
	float random_number = (float) rand()/RAND_MAX;
	random_number = random_number * (b - a) + a;

	return random_number;
}

