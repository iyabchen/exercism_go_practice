package strand

const testVersion = 3

// Given a DNA strand, return its RNA complement (per RNA transcription).
// Given a DNA strand, its transcribed RNA strand is formed by replacing
// each nucleotide with its complement:
//   DNA -> RNA
// * `G` -> `C`
// * `C` -> `G`
// * `T` -> `A`
// * `A` -> `U`

func ToRNA(dna string) string {
	strArr := []byte{}
	for _, v := range dna {
		switch v {
		case 'G':
			strArr = append(strArr, 'C')
		case 'C':
			strArr = append(strArr, 'G')
		case 'T':
			strArr = append(strArr, 'A')
		case 'A':
			strArr = append(strArr, 'U')
		default: // illegal
			return ""

		}
	}
	return string(strArr)

}

// This method is slower
// import "bytes"
// func ToRNA(dna string) string {
// 	var buf bytes.Buffer
// 	for _, v := range dna {
// 		switch v {
// 		case 'G':
// 			buf.WriteRune('C')
// 		case 'C':
// 			buf.WriteRune('G')
// 		case 'T':
// 			buf.WriteRune('A')
// 		case 'A':
// 			buf.WriteRune('U')
// 		default: // illegal
// 			return ""

// 		}
// 	}
// 	return buf.String()

// }
