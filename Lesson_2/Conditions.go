package main

func grade(a int) string {
	if a >= 95 && a <= 100 {
		return "A+"
	} else if a < 95 && a > 90 {
		return "A"
	} else if a >= 85 && a <= 90 {
		return "B+"
	} else if a < 85 && a > 80 {
		return "B"
	} else if a <= 80 && a >= 75 {
		return "C+"
	} else if a < 75 && a > 70 {
		return "C"
	} else if a >= 65 && a <= 70 {
		return "D"
	} else if a <= 64 {
		return "FAIL"
	} else {
		return "Failed"
	}
}
