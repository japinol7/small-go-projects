package main

var (
	openMarks  = []rune{'(', '{', '['}
	closeMarks = []rune{')', '}', ']'}
)

// AreParenthesesBalanced checks if the input string has balanced parentheses
func AreParenthesesBalanced(str string) bool {
	// Create a map of opening to closing marks
	marksMap := make(map[rune]rune)
	for i := 0; i < len(openMarks); i++ {
		marksMap[openMarks[i]] = closeMarks[i]
	}

	// Stack to keep track of opening marks
	stack := make([]rune, 0)

	for _, ch := range str {
		// Check if it's an opening mark
		isOpen := false
		for _, open := range openMarks {
			if ch == open {
				stack = append(stack, marksMap[ch])
				isOpen = true
				break
			}
		}
		if isOpen {
			continue
		}

		// Check if it's a closing mark
		for _, closeMark := range closeMarks {
			if ch == closeMark {
				if len(stack) == 0 || ch != stack[len(stack)-1] {
					return false
				}
				// Pop from stack
				stack = stack[:len(stack)-1]
				break
			}
		}
	}

	// Stack should be empty for balanced parentheses
	return len(stack) == 0
}
