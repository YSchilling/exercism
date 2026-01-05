def steps(number):
    if number <= 0: raise ValueError("Only positive integers are allowed")
    counter = 0
    while True:
        if number == 1: return counter

        if number % 2 == 0: number /= 2
        else: number = 3 * number + 1
        counter += 1
