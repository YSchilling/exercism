let square_of_sum n =
    let rec sum_n n = if n <= 0
    then 0 else n + sum_n(n - 1) in
    let sum = sum_n(n) in
    sum * sum

let rec sum_of_squares n =
    let square n = n * n in
    if n <= 0
    then 0 else square(n) + sum_of_squares(n - 1)

let difference_of_squares n =
    let sqsum = square_of_sum n in
    let sumqs = sum_of_squares n in
    if sqsum > sumqs
    then sqsum - sumqs else sumqs - sqsum
