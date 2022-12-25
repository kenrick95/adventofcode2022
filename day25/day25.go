package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

const factor = 5

func toSnafu(num int) string {
	chars := []string{}

	fiveAry := []int{}
	for cur, i := num, 0; cur > 0; i++ {
		rem := cur % factor
		fiveAry = append(fiveAry, rem)
		cur = cur / factor
	}
	// fmt.Printf("fiveAry: %v -> %v\n", num, fiveAry)

	tmpAry := []int{}
	for i := 0; i < len(fiveAry); i++ {
		tmpAry = append(tmpAry, fiveAry[i])
	}

	for i := 0; i < len(tmpAry); i++ {
		v := tmpAry[i]

		if v == 0 || v == 1 || v == 2 {
			chars = append(chars, strconv.Itoa(v))
		} else if v == 4 {
			chars = append(chars, "-")
			found := false
			for j := i + 1; j < len(tmpAry); j++ {
				if tmpAry[j] < 4 {
					tmpAry[j] += 1
					found = true
					break
				} else {
					tmpAry[j] = 0
				}
			}
			if !found {
				tmpAry = append(tmpAry, 1)
			}
		} else if v == 3 {
			chars = append(chars, "=")
			found := false
			for j := i + 1; j < len(tmpAry); j++ {
				if tmpAry[j] < 4 {
					tmpAry[j] += 1
					found = true
					break
				} else {
					tmpAry[j] = 0
				}
			}
			if !found {
				tmpAry = append(tmpAry, 1)
			}
		}

	}

	// 2022      ------- 3 1 0 4 2
	//                       1 -
	//                 1 =

	 //  1 = 1 1 - 2


	//        4 4 4 2
	//      1 0 0 -  

	/*
	     Decimal          SNAFU
	           1              1
	           2              2
	           3             1=
	           4             1-
	           5             10
	           6             11
	           7             12
	           8             2=
	           9             2-
	          10             20
	          15            1=0
	          20            1-0
	   	   25			 100
	        2022         1=11-2
	       12345        1-0---0
	   314159265  1121-1110-1=0

	   4890      2=-1=0

	    SNAFU  Decimal
	   1=-0-2     1747
	    12111      906
	     2=0=      198
	       21       11
	     2=01      201
	      111       31
	    20012     1257
	      112       32
	    1=-1=      353
	     1-12      107
	       12        7
	       1=        3
	      122       37
	*/

	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	

	return strings.Join(chars, "")
}
func fromSnafu(str string) int {
	chars := strings.Split(str, "")
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	res := 0
	for i := 0; i < len(chars); i++ {
		chVal := 0
		switch chars[i] {
		case "1":
			{
				chVal = 1
			}
		case "2":
			{
				chVal = 2
			}
		case "-":
			{
				chVal = -1
			}
		case "=":
			{
				chVal = -2
			}
		}

		val := fivePower(i) * chVal
		res += val
	}
	return res
}
func fivePower(exp int) int {
	res := 1
	for i := 0; i < exp; i++ {
		res *= factor
	}
	return res
}

func main() {
	sum := 0
	lines := utils.ReadFileToLines("day25.real.in")

	// for i := 1; i <= 25; i++ {
	// 	fmt.Println("toSnafu", i, toSnafu(i))
	// }
	// fmt.Println("toSnafu", 2022, toSnafu(2022))
	// fmt.Println("toSnafu", 12345, toSnafu(12345))
	// fmt.Println("toSnafu", 314159265, toSnafu(314159265))

	for _, line := range lines {
		res := fromSnafu(line)
		sum += res
	}
	fmt.Printf("sum: %v\n", sum)
	ansPart1 := toSnafu(sum)
	fmt.Printf("ansPart1: %v\n", ansPart1)
}
