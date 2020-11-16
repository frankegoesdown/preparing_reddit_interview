package main

func main() {

}

func calculate(s string) int {
	ans, cur := 0, 0
	stack := []int{}
	op := '+'
	for i, ch := range s {
		if ch-'0' >= 0 && ch-'0' <= 9 {
			cur = cur*10 + int(ch-'0')
		}
		if ch == '+' || ch == '-' || ch == '*' || ch == '/' || i == len(s)-1 {
			switch op {
			case '+':
				stack = append(stack, cur)
			case '-':
				stack = append(stack, -cur)
			case '*':
				t := cur * stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, t)
			case '/':
				t := stack[len(stack)-1] / cur
				stack = stack[:len(stack)-1]
				stack = append(stack, t)

			}
			op = rune(ch)
			cur = 0
		}
	}

	for _, n := range stack {
		ans += n
	}
	return ans
}
