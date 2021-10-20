def main():
    digits = list(range(1,101))
    result = difference(digits) 
    print(f"Answer: {result}")


def difference(digits: list) -> int:
    return square_of_sums(digits) - sum_of_squares(digits)


def square_of_sums(digits: list) -> int:
    return sum(digits)**2


def sum_of_squares(digits: list) -> int:
    return sum(map(lambda x: x**2, digits))


if __name__ == '__main__':
    main()