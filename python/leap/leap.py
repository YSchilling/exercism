def leap_year(year):
    is_div_4 = (year % 4 == 0)
    is_not_div_100 = not (year % 100 == 0)
    is_div_100_and_400 = (year % 100 == 0 and year % 400 == 0)
    return is_div_4 and (is_not_div_100 or is_div_100_and_400)
