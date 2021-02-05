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
	tmpString := ""
	loopCount := 1
	for loopCount <= count {
		tmpString = tmpString + " "
	}
	return tmpString
}

func drawDiamond(s string) rows {
	LetterValue := alphabetValue(s)
	MaxWidthHeight := calcGridSize(LetterValue)
	TotalLength := (MaxWidthHeight * MaxWidthHeight)
	fmt.Println("GridSize: ", MaxWidthHeight, " StrLength: ", TotalLength)

	// Loop A
	loopAcount := 0
	for loopAcount <= LetterValue {
		fmt.Println("Loop A ", loopAcount)

		if loopAcount == 0 {
			res := ""
			res = stringSpacing(calcSpacing(LetterValue, loopAcount))
			res = res + valueAlphabet(loopAcount)
			res = res + stringSpacing(calcSpacing(LetterValue, loopAcount))
			fmt.Println(res)
			// write spacing (LetterValue - LoopCount)
			// Write letter valueAlphabet(LoopCount)
			// write spacing
		}
		if loopAcount != 0 {
			if loopAcount != LetterValue {
				// write spacing
				// Write letter
				// write middle space   (((LoopCount -1) *2) +1)
				// Write letter
				// write spacing
			}
		}
		if loopAcount == LetterValue {
			// Write letter
			// write middle space
			// Write letter
		}

		loopAcount++
	}
	// Loop B
	loopBcount := LetterValue - 1
	for loopBcount >= 0 {
		fmt.Println("Loop B ", loopBcount)
		loopBcount--
	}

	data := rows{{row: s}}
	return data
}

func main() {
	data := drawDiamond("A")
	fmt.Print(data)
}
