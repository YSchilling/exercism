def equilateral(sides):
    if not isTriangle(sides): return False
    a, b, c = sides
    return a == b == c


def isosceles(sides):
    if not isTriangle(sides): return False
    a, b, c = sides
    return a == b or a == c or b == c


def scalene(sides):
    if not isTriangle(sides): return False
    a, b, c = sides
    return a != b and a != c and b != c

def isTriangle(sides):
    a, b, c = sides
    if a <= 0 or b <= 0 or c <= 0: return False
    if a + b < c or a + c < b or b + c < a: return False
    return True