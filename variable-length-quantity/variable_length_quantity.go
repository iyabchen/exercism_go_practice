// Implement [VLQ] encoding/decoding
// https://en.wikipedia.org/wiki/Variable-length_quantity

package variablelengthquantity

const testVersion = 1

func EncodeVarint(input uint32) (byteArr []byte) {
	r := input % 128
	byteArr = []byte{byte(r)}
	input = input / 128
	for input > 0 {
		r := input % 128
		input = input / 128
		byteArr = append([]byte{byte(r + 128)}, byteArr...)
	}
	return

}

// no idea what the size means
func DecodeVarint(byteArr []byte) (output uint32, size int) {
	lenArr := len(byteArr)
	for ix, b := range byteArr {
		output = output << 7
		if ix != lenArr-1 {
			b = b - 128
			output += uint32(b)
		} else { // last byte
			output += uint32(b)
		}
	}
	if lenArr <= 2 {
		size = 1
	} else {
		size = lenArr - 1
	}
	return

}
