module ProteinTranslation(proteins) where

proteins :: String -> Maybe [String]
proteins x = if length x `mod` 3 == 0
    then Just (codonsToProteins (rnaToCodons x))
    else Nothing

-- RNA into codons list (stop if there is a stop codon)
rnaToCodons :: String -> [String]
rnaToCodons rna = if codonToProtein first == "STOP"
    then []
    else if length rest == 0
    then first : []
    else first : (rnaToCodons rest)
    where
        (first, rest) = splitAt 3 rna

-- codons list into protein list
codonsToProteins :: [String] -> [String]
codonsToProteins codons = [codonToProtein codon | codon <- codons]

-- codon into protein
codonToProtein :: String -> String
codonToProtein "AUG" = "Methionine"
codonToProtein "UUU" = "Phenylalanine"
codonToProtein "UUC" = "Phenylalanine"
codonToProtein "UUA" = "Leucine"
codonToProtein "UUG" = "Leucine"
codonToProtein "UCU" = "Serine"
codonToProtein "UCC" = "Serine"
codonToProtein "UCA" = "Serine"
codonToProtein "UCG" = "Serine"
codonToProtein "UAU" = "Tyrosine"
codonToProtein "UAC" = "Tyrosine"
codonToProtein "UGU" = "Cysteine"
codonToProtein "UGC" = "Cysteine"
codonToProtein "UGG" = "Tryptophan"
codonToProtein "UAA" = "STOP"
codonToProtein "UAG" = "STOP"
codonToProtein "UGA" = "STOP"