package strand

func ToRNA(dna string) string {
	transcriptionMap := map[string]string{
		"G": "C",
		"C": "G",
		"T": "A",
		"A": "U",
	}

	str := ""
	for _, c := range dna {
		str = str + transcriptionMap[string(c)]
	}

	return str
}
