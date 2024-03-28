package fizzbuzz

import "strconv"

// 時間複雜 O(), 空間複雜 O()
func fizzBuzz(n int) []string {
	var result []string
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}

func fizzBuzz2(n int) []string {
	var result []string
	for i := 1; i <= n; i++ {
		var str string
		if i%3 == 0 {
			str += "Fizz"
		}
		if i%5 == 0 {
			str += "Buzz"
		}
		if len(str) <= 0 {
			str = strconv.Itoa(i)
		}
		result = append(result, str)
	}
	return result
}

// 最佳
func fizzBuzz3(n int) []string {
	var result []string
	fizz := 0
	buzz := 0
	for i := 1; i <= n; i++ {
		fizz++
		buzz++
		if fizz == 3 && buzz == 5 {
			result = append(result, "FizzBuzz")
			fizz = 0
			buzz = 0
		} else if fizz == 3 {
			result = append(result, "Fizz")
			fizz = 0
		} else if buzz == 5 {
			result = append(result, "Buzz")
			buzz = 0
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}
