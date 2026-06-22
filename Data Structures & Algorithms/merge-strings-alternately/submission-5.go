func mergeAlternately(word1 string, word2 string) string {
	idx1, idx2 := 0, 0
	var word3 string
	for idx1 < len(word1) || idx2 < len(word2)  {
		if idx1 < len(word1) {
			word3 += string(word1[idx1])
			idx1++
		} 
		if idx2 < len(word2) {
			word3 += string(word2[idx2])
			idx2++
		}
	}
 	return word3
}
