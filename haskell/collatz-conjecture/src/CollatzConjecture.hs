module CollatzConjecture (collatz) where

collatz :: Integer -> Maybe Integer
collatz n = if n <= 0
    then Nothing
    else Just (innerCollatz n 0)

innerCollatz :: Integer -> Integer -> Integer
innerCollatz n i = if n == 1
    then i
    else if even n
        then innerCollatz (n `div` 2) (i+1)
        else innerCollatz (3*n+1) (i+1)