package main

// Chars used to display LCD digits and the separator string
const (
	CellH     = '_' // Horizontal segment
	CellV     = '|' // Vertical segment
	CellO     = '.' // Off segment
	Separator = " "
)

// Digit templates for each number with placeholders for separators
var (
	Digits0 = string([]byte{CellO, CellH, CellO, '%', 's',
		CellV, CellO, CellV, '%', 's',
		CellV, CellH, CellV, '%', 's'})

	Digits1 = string([]byte{CellO, CellO, CellO, '%', 's',
		CellO, CellO, CellV, '%', 's',
		CellO, CellO, CellV, '%', 's'})

	Digits2 = string([]byte{CellO, CellH, CellO, '%', 's',
		CellO, CellH, CellV, '%', 's',
		CellV, CellH, CellO, '%', 's'})

	Digits3 = string([]byte{CellO, CellH, CellO, '%', 's',
		CellO, CellH, CellV, '%', 's',
		CellO, CellH, CellV, '%', 's'})

	Digits4 = string([]byte{CellO, CellO, CellO, '%', 's',
		CellV, CellH, CellV, '%', 's',
		CellO, CellO, CellV, '%', 's'})

	Digits5 = string([]byte{CellO, CellH, CellO, '%', 's',
		CellV, CellH, CellO, '%', 's',
		CellO, CellH, CellV, '%', 's'})

	Digits6 = string([]byte{CellO, CellH, CellO, '%', 's',
		CellV, CellH, CellO, '%', 's',
		CellV, CellH, CellV, '%', 's'})

	Digits7 = string([]byte{CellO, CellH, CellO, '%', 's',
		CellO, CellO, CellV, '%', 's',
		CellO, CellO, CellV, '%', 's'})

	Digits8 = string([]byte{CellO, CellH, CellO, '%', 's',
		CellV, CellH, CellV, '%', 's',
		CellV, CellH, CellV, '%', 's'})

	Digits9 = string([]byte{CellO, CellH, CellO, '%', 's',
		CellV, CellH, CellV, '%', 's',
		CellO, CellO, CellV, '%', 's'})

	// NumToLcdDigitMapping maps integer digits to their LCD representation
	NumToLcdDigitMapping = map[int]string{
		0: Digits0,
		1: Digits1,
		2: Digits2,
		3: Digits3,
		4: Digits4,
		5: Digits5,
		6: Digits6,
		7: Digits7,
		8: Digits8,
		9: Digits9,
	}
)
