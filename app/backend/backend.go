package backend

import (
	"strings"
)

func DigitToString(num int) string{
	switch num {	
		case 1 : 
			return "satu"
		case 2 : 
			return "dua"
		case 3 :
			return "tiga"
		case 4 : 
			return "empat"
		case 5 : 
			return "lima"
		case 6 :
			return "enam"
		case 7 :
			return "tujuh"
		case 8 : 
			return "delapan"
		case 9 : 
			return "sembilan"
	}
	return ""
}

func ParseNumeral(num int) []int {
	var numbers []int
	for num != 0 {
		numbers = append(numbers, num % 1000)
		num = num/1000
	}
	return numbers
}

func ParseRatusan(ratusan int) string {
	result := ""
	ratus := ratusan / 100
	puluh := ratusan % 100 / 10
	satuan := ratusan % 10
	belas := false

	/** ratus **/
	if (ratus == 1) {
		result += "seratus "
	} else if (ratus > 1) {
		result += DigitToString(ratus) + " ratus "
	}

	/** puluhan **/
	if (puluh == 1) {
		if (satuan == 0) {
			result += "sepuluh "
		} else if (satuan == 1) {
			result += "sebelas "
			belas = true;
		} else {
			result += DigitToString(satuan) + " belas "
			belas = true;
		}
	} else if (puluh > 1) {
		result += DigitToString(puluh) + " puluh "
	} 

	/** satuan **/
	if (!belas  && satuan != 0) {
		result += DigitToString(satuan) + " "
	}

	return result
}

func NumeralSpeller(number int) string {
	num := ParseNumeral(number)
	delimiter := [3]string{"ribu ","juta ","milyar "}
	result := ""
	i := len(num) - 1
	for i >= 0 {
		if (num[i] != 0) {
			if (i == 1) {
				if (ParseRatusan(num[i]) == "satu ") {
					result += "seribu "
				} else {
					result += ParseRatusan(num[i]) + delimiter[i-1]
				}
			} else {
				if (i-1>=0) {
					result += ParseRatusan(num[i]) + delimiter[i-1]
				} else {
					result +=  ParseRatusan(num[i])
				}
			}
		}
		i--
	}
	if len(result) > 0 && result[len(result) - 1] == ' ' {
		result = result[:len(result) - 1]
	}
	return result
}

func SpellToNumeral(s string) int {
	spl := strings.Fields(s)
	result := 0
	temp := 0

	for _, spel := range spl {
		if (spel == "satu") {
			temp+=1
		} else if (spel == "dua") {
			temp+=2
		} else if (spel == "tiga") {
			temp+=3
		} else if (spel == "empat") {
			temp+=4
		} else if (spel == "lima") {
			temp+=5
		} else if (spel == "enam") {
			temp+=6
		} else if (spel == "tujuh") {
			temp+=7
		} else if (spel == "delapan") {
			temp+=8
		} else if (spel == "sembilan") {
			temp+=9
		} else if (spel == "sepuluh") {
			temp+=10
		} else if (spel == "sebelas") {
			temp+=11
		} else if (spel == "seratus") {
			temp+=100
		} else if (spel == "seribu") {
			temp+=1000
		} else if (spel == "puluh") {
			temp2 := temp % 10
			temp2 *= 10
			temp += temp2 - temp % 10
		} else if (spel == "ratus") {
			temp2 := temp % 10
			temp2 *= 100
			temp += temp2 - temp % 10
		} else if (spel == "belas") {
			temp2 := temp % 10
			temp2 += 10
			temp += temp2 - temp % 10
		}		
		if (spel == "milyar") {
			temp *= 1000000000
			result += temp
			temp = 0
		} else if (spel == "juta") {
			temp *= 1000000
			result += temp
			temp = 0
		} else if (spel == "ribu") {
			temp *= 1000
			result += temp
			temp = 0
		}
	}
	result += temp

	return result
}
 
// func main() {
// 	consoleReader := bufio.NewReader(os.Stdin)
// 	n, _ := consoleReader.ReadString('\n') 
// 	fmt.Println(SpellToNumeral(n))
	// var n int
	// fmt.Scanf("%d", &n)

	// fmt.Println(NumeralSpeller(n))

	// for _, num := range numbers {
	// 	fmt.Println(ParseRatusan(num))
	// }
// }