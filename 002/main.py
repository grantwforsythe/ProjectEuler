def fib(a, b, n = 4*10**6):
	"""
	Returns the sum of the even-valued terms in the Fibonacci Sequence
	"""
	result = 0

	while (temp := a + b) < n:
		# TODO: missing the first two terms in the sequence
		if (temp % 2 == 0):
			result += temp
		
		a = b
		b = temp
	# Answer off by two, missing the second term
	return result + 2



if __name__ == '__main__':
	print(f'Answer: {fib(1,2)}')