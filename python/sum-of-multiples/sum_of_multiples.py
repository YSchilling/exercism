def sum_of_multiples(limit, multiples):
    result_set = set()
    for n in multiples:
        for i in [i for i in range(1, limit) if n != 0 and i != 0 and i % n == 0]:
            result_set.add(i)
    return sum(result_set)
