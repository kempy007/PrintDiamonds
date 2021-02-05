package main

import "fmt"

const (
	A = iota
	B
	C
	D
	E
	F
	G
	H
	I
	J
	K
	L
	M
	N
	O
	P
	Q
	R
	S
	T
	U
	V
	W
	X
	Y
	Z
)

type rows []struct {
	row string
}

func alphabetValue(s string) int {
	switch s {
	case "A":
		return A
	case "B":
		return B
	case "C":
		return C
	case "D":
		return D
	case "E":
		return E
	case "F":
		return F
	case "G":
		return G
	case "H":
		return H
	case "I":
		return I
	case "J":
		return J
	case "K":
		return K
	case "L":
		return L
	case "M":
		return M
	case "N":
		return N
	case "O":
		return O
	case "P":
		return P
	case "Q":
		return Q
	case "R":
		return R
	case "S":
		return S
	case "T":
		return T
	case "U":
		return U
	case "V":
		return V
	case "W":
		return W
	case "X":
		return X
	case "Y":
		return Y
	case "Z":
		return Z
	default:
		return 0
	}
}

func valueAlphabet(s int) string {
	switch s {
	case A:
		return "A"
	case B:
		return "B"
	case C:
		return "C"
	case D:
		return "D"
	case E:
		return "E"
	case F:
		return "F"
	case G:
		return "G"
	case H:
		return "H"
	case I:
		return "I"
	case J:
		return "J"
	case K:
		return "K"
	case L:
		return "L"
	case M:
		return "M"
	case N:
		return "N"
	case O:
		return "O"
	case P:
		return "P"
	case Q:
		return "Q"
	case R:
		return "R"
	case S:
		return "S"
	case T:
		return "T"
	case U:
		return "U"
	case V:
		return "V"
	case W:
		return "W"
	case X:
		return "X"
	case Y:
		return "Y"
	case Z:
		return "Z"
	default:
		return ""
	}
}

func calcSpacing(letterValue int, loop int) int {
	return letterValue - loop
}

func calcMidSpacing(loop int) int {
	return (((loop - 1) * 2) + 1)
}

func calcGridSize(letterValue int) int {
	return (letterValue * 2) + 1
}

func stringSpacing(count int) string {
	tmpString := ``
	loopCount := 1
	for loopCount <= count {
		tmpString = tmpString + ` `
		loopCount++
	}
	return tmpString
}

func drawDiamond(s string) []string {
	LetterValue := alphabetValue(s)
	MaxWidthHeight := calcGridSize(LetterValue)
	TotalLength := (MaxWidthHeight * MaxWidthHeight)
	fmt.Println("GridSize: ", MaxWidthHeight, " StrLength: ", TotalLength)
	retData := []string{}

	// Loop A
	loopAcount := 0
	for loopAcount <= LetterValue {
		//fmt.Println("Loop A ", loopAcount)
		retData = append(retData, innerDrawLoop(loopAcount, LetterValue))
		loopAcount++
	}
	// Loop B
	loopBcount := LetterValue - 1
	for loopBcount >= 0 {
		//fmt.Println("Loop B ", loopBcount)
		retData = append(retData, innerDrawLoop(loopBcount, LetterValue))
		loopBcount--
	}

	return retData
}

func innerDrawLoop(loop int, LetterValue int) string {
	if loop == 0 {
		res := ""
		res = stringSpacing(calcSpacing(LetterValue, loop))       // write spacing
		res = res + valueAlphabet(loop)                           // Write letter
		res = res + stringSpacing(calcSpacing(LetterValue, loop)) // write spacing
		//fmt.Println(res)
		return res
	}
	if loop != 0 {
		if loop != LetterValue {
			res1 := ""
			res1 = stringSpacing(calcSpacing(LetterValue, loop))        // write spacing
			res1 = res1 + valueAlphabet(loop)                           // Write letter
			res1 = res1 + stringSpacing(calcMidSpacing(loop))           // write middle space
			res1 = res1 + valueAlphabet(loop)                           // Write letter
			res1 = res1 + stringSpacing(calcSpacing(LetterValue, loop)) // write spacing
			//fmt.Println(res1)
			return res1
		}
	}
	if loop == LetterValue {
		res2 := ""
		res2 = res2 + valueAlphabet(loop)                 // Write letter
		res2 = res2 + stringSpacing(calcMidSpacing(loop)) // write middle space
		res2 = res2 + valueAlphabet(loop)                 // Write letter
		//fmt.Println(res2)
		return res2
	}
	return ""
}

// func main() {
// 	data := drawDiamond("E")
// 	if data != nil {
// 		//fmt.Println(data) // just one line
// 	}
// }
