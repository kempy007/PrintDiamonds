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


# timelog
after fews hours had clear idea of howto solve. 6hrs
Solved in vscode. 2hrs
Total 8hrs


https://cyber-dojo.org/kata/edit/TAKQTP

https://cyber-dojo.org/kata/edit/Q2eY4A
