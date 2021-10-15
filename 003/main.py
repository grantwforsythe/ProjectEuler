with open('primes-to-100k.txt', 'r') as f:
    prime_nums = list(map(int, f.readlines()))

def main():
    num = int(input("Find the largest prime factors of a nonzero number: "))
    print("Answer: ", prime_factorization(num))

def prime_factorization(num : int) -> int:
    """
        Description:    Find the largest prime factors of a nonzero number.
        Parameters:     num             - int
        Return:         the largest prime factor
        Example of Usage:
        >> prime_facorization(13195)
        ans =
            29
    """ 

    # check the number is nonzero
    if num == 0:
        return None

    for prime_num in prime_nums:
        if num % prime_num == 0:
            num /= prime_num
            if num == 1:
                return prime_num
            else:
                prime_factorization(num) 

if __name__ == '__main__':
    main()