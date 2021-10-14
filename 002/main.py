def fib():
    """
    Returns the sum of the even-valued terms in the Fibonacci Sequence
    """
    result = 0
    a = 0
    b = 1
    n = 4*10**6

    while (temp := a + b) < n:
        if (temp % 2 == 0):
            result += temp
        a = b
        b = temp
    return result

if __name__ == '__main__':
    print(f'Answer: {fib()}')
