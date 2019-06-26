import "strings"

//This is a leetcode practice problem
func myAtoi(str string) int {
   	INT_MAX := 2147483647
	INT_MIN := -2147483648

	if str == "" {
		return 0
	}
	//a := regexp.MustCompile(`0123456789.`)
	var word string
	str1 := strings.Split(str, " ")
	for i := range str1 {
		if str1[i] != "" {
			word = str1[i]
			break
		}
	}
	//fmt.Println(word)
	//j := 0
	var str2 []byte
	for i := range word {
		if (word[i] == '+' || word[i] == '-') && i == 0 {
			str2 = append(str2, word[i])

		} else {
	//		fmt.Println(rune(word[i]))
			if (word[i] >= '0' && word[i] <= '9') || word[i] == '.' {
				str2 = append(str2, word[i])

			} else {
				break
			}

		}
	}
	
	//Uncomment the below to debug
	//fmt.Println(string(str2))
	str3 := string(str2[:])
	//fmt.Println(str3)
	str4 := strings.Split(str3, ".")
	//fmt.Println(str4)
	number, _ := strconv.Atoi(str4[0])
	//fmt.Println(number)
	
		if number > INT_MAX {
			return INT_MAX
		} else if number < INT_MIN {
			return INT_MIN
		} else {
			return number
		}
	
}
