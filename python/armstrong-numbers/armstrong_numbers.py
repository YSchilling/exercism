def is_armstrong_number(number):
    num_of_digits = len(str(number))
    number_to_str_list = list(str(number))
    pow_list = [int(x)**num_of_digits for x in number_to_str_list]
    endsum = sum(pow_list)
    return endsum == number

is_armstrong_number(9474)