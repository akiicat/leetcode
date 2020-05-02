package main
import "testing"
import . "main/pkg/testing_helper"

func TestLongestPalindromeSubseq(_t *testing.T) {
  i, o := "bbbab", 4
  T(_t, S(i), S(longestPalindromeSubseq(i)), S(o))

  i, o = "cbbd", 2
  T(_t, S(i), S(longestPalindromeSubseq(i)), S(o))

  i, o = "xybbccfcbaabx", 9
  T(_t, S(i), S(longestPalindromeSubseq(i)), S(o))

  i, o = "euazbipzncptldueeuechubrcourfpftcebikrxhybkymimgvldiwqvkszfycvqyvtiwfckexmowcxztkfyzqovbtmzpxojfofbvwnncajvrvdbvjhcrameamcfmcoxryjukhpljwszknhiypvyskmsujkuggpztltpgoczafmfelahqwjbhxtjmebnymdyxoeodqmvkxittxjnlltmoobsgzdfhismogqfpfhvqnxeuosjqqalvwhsidgiavcatjjgeztrjuoixxxoznklcxolgpuktirmduxdywwlbikaqkqajzbsjvdgjcnbtfksqhquiwnwflkldgdrqrnwmshdpykicozfowmumzeuznolmgjlltypyufpzjpuvucmesnnrwppheizkapovoloneaxpfinaontwtdqsdvzmqlgkdxlbeguackbdkftzbnynmcejtwudocemcfnuzbttcoew", 159
  T(_t, S(i), S(longestPalindromeSubseq(i)), S(o))
}

func TestLongestPalindromeSubseqCacheRecursive(_t *testing.T) {
  i, o := "bbbab", 4
  T(_t, S(i), S(longestPalindromeSubseqCacheRecursive(i)), S(o))

  i, o = "cbbd", 2
  T(_t, S(i), S(longestPalindromeSubseqCacheRecursive(i)), S(o))

  i, o = "xybbccfcbaabx", 9
  T(_t, S(i), S(longestPalindromeSubseqCacheRecursive(i)), S(o))

  i, o = "euazbipzncptldueeuechubrcourfpftcebikrxhybkymimgvldiwqvkszfycvqyvtiwfckexmowcxztkfyzqovbtmzpxojfofbvwnncajvrvdbvjhcrameamcfmcoxryjukhpljwszknhiypvyskmsujkuggpztltpgoczafmfelahqwjbhxtjmebnymdyxoeodqmvkxittxjnlltmoobsgzdfhismogqfpfhvqnxeuosjqqalvwhsidgiavcatjjgeztrjuoixxxoznklcxolgpuktirmduxdywwlbikaqkqajzbsjvdgjcnbtfksqhquiwnwflkldgdrqrnwmshdpykicozfowmumzeuznolmgjlltypyufpzjpuvucmesnnrwppheizkapovoloneaxpfinaontwtdqsdvzmqlgkdxlbeguackbdkftzbnynmcejtwudocemcfnuzbttcoew", 159
  T(_t, S(i), S(longestPalindromeSubseqCacheRecursive(i)), S(o))
}

func TestLongestPalindromeSubseqRecursive(_t *testing.T) {
  i, o := "bbbab", 4
  T(_t, S(i), S(longestPalindromeSubseqRecursive(i)), S(o))

  i, o = "cbbd", 2
  T(_t, S(i), S(longestPalindromeSubseqRecursive(i)), S(o))

  i, o = "xybbccfcbaabx", 9
  T(_t, S(i), S(longestPalindromeSubseqRecursive(i)), S(o))
}

