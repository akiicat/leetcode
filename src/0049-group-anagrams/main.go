package main

// T: O(n * k)
// M: O(n * k)
// -- start --

type bytemap [26]byte

func strToBytemap(s string) bytemap {
	r := [26]byte{}
	for _, v := range s {
		r[v-'a']++
	}
	return bytemap(r)
}

func groupAnagrams(strs []string) [][]string {

	c := 0
	m := make(map[bytemap]int)
	rtn := [][]string{}

	for _, v := range strs {
		bm := strToBytemap(v)

		if _, ok := m[bm]; !ok {
			m[bm] = c
			c++
			rtn = append(rtn, []string{})
		}

		rtn[m[bm]] = append(rtn[m[bm]], v)
	}

	return rtn
}

// -- end --
