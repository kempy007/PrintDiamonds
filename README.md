# PrintDiamonds

Given a letter print a diamond starting with 'A' with the supplied letter at the widest point.

This is the print diamond for 'E'.

....A....
...B.B...
..C...C..
.D.....D.
E.......E
.D.....D.
..C...C..
...B.B...
....A....


This is the print diamond for 'C'.

..A..
.B.B.
C...C
.B.B.
..A..

This is the print diamond for 'A'.

A

Note: These examples use dots in place of spaces only for readability.
Your diamond must contain spaces where there are dots.

# My Custom Notes
Each Letter can represent a value, We ask the func for a Diamond of Letter $letter, Given that the letter represents the highest value we can deduce the spacing requirements.
A = 0
B = 1
C = 2
D = 3
E = 4
F = 5
...
Z = 25


RowWidth := (LetterValue * 2) + 1
TotalLength := RowWidth * RowWidth

(C * 2) + 1 = 5

A 1x1
B 3x3
C 5x5
D 7x7
E 9x9

F 11x11      Loop A
.....A.....  0 increment counter
....B.B....  1
...C...C...  2
..D.....D..  3
.E.......E.  4
F.........F  5 
             Loop B
.E.......E.  (LetterValue - 1) = 4 decrement counter
..D.....D..  3
...C...C...  2
....B.B....  1
.....A.....  0

if LoopCount == 0 { 
    // write spacing (LetterValue - LoopCount)
    // Write letter valueAlphabet(LoopCount)
    // write spacing
}
if LoopCount != 0 {
    if LoopCount !- LetterValue {
        // write spacing
        // Write letter
        // write middle space   (((LoopCount -1) *2) +1)
        // Write letter
        // write spacing
    }
}
if LoopCount == LetterValue {
    // Write letter
    // write middle space
    // Write letter
 }

No matter the grid size the center space are consistent. In a Loop take ((LoopCount - 1) * 2) + 1
