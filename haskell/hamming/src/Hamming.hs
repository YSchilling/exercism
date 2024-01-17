module Hamming (distance) where

distance :: String -> String -> Maybe Int
distance xs ys =
    if check_error then Just 1 else Nothing
    where
        check_error = True