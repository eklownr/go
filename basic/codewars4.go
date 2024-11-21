package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type MyStr string

func (s MyStr) IsUpperCase() bool {
	return s == MyStr(strings.ToUpper(string(s)))
}

func Derive(coefficient, exponent int) string {
	return fmt.Sprintf("%dx^%d", coefficient*exponent, exponent-1)
}

func NearestSq(n int) int {
	root := math.Sqrt(float64(n))
	//	roundRoot := math.Round(math.Sqrt(float64(n)))
	//	fmt.Printf("Root: %v, rouned: %v, squere: %v\n", root, roundRoot, roundRoot*roundRoot)

	if int((root)+1)*(int(root)+1)-n < n-int(root)*int(root) {
		return int((root)+1) * (int(root) + 1)
	}
	return int(root) * int(root)
}
func FindUniq(arr []float32) float32 {
	// take care of the first elements
	if arr[0] != arr[1] && arr[1] == arr[2] {
		return arr[0]
	}
	// arr[0] is not unique so we can compare the rest of the array
	for _, num := range arr[1:] {
		if num != arr[0] {
			return num
		}
	}
	return arr[0]
}

func FindUniq2(arr []float32) float32 {
	sort.SliceStable(arr, func(i, j int) bool { return arr[i] < arr[j] })
	if arr[0] == arr[1] {
		return arr[len(arr)-1]
	}
	return arr[0]
}

func UniqueFloat32(arr []float32) []float32 {
	uniqMap := make(map[float32]bool)
	uniqSlice := make([]float32, 0, len(arr))
	for _, num := range arr {
		if _, ok := uniqMap[num]; !ok {
			uniqMap[num] = true
			uniqSlice = append(uniqSlice, num)
		}
	}
	return uniqSlice
}

// Upper sorted lastname, firstname
func Meeting(s string) string {
	s = strings.ToUpper(s)
	names := strings.Split(s, ";")
	for i := 0; i < len(names); i++ {
		firstname, lastname := strings.Split(names[i], ":")[0], strings.Split(names[i], ":")[1]
		//names[i] = strings.Split(names[i], ":")[1] + strings.Split(names[i], ":")[0]
		fmt.Println("(", lastname, ",", firstname, ")")
	}
	sort.SliceStable(names, func(i, j int) bool { return names[i] < names[j] })
	//for i := 0; i < len(names)-1; i++ {
	//	names[i] = strings.Split(names[i], " ")[1] + ", " + strings.Split(names[i], " ")[0]
	//}
	return strings.Join(names, ", ")
}

func main() {
	fmt.Println(Meeting("Alexis:Wahl;John:Bell;Victoria:Schwarz;Abba:Dorny;Grace:Meta;Ann:Arno;Madison:STAN;Alex:Cornwell;Lewis:Kern;Megan:Stan;Alex:Korn"))
	//"(ARNO, ANN)(BELL, JOHN)(CORNWELL, ALEX)(DORNY, ABBA)(KERN, LEWIS)(KORN, ALEX)(META, GRACE)(SCHWARZ, VICTORIA)(STAN, MADISON)(STAN, MEGAN)(WAHL, ALEXIS)")
	fmt.Println(Meeting("John:Gates;Michael:Wahl;Megan:Bell;Paul:Dorries;James:Dorny;Lewis:Steve;Alex:Meta;Elizabeth:Russel;Anna:Korn;Ann:Kern;Amber:Cornwell"))
	//"(BELL, MEGAN)(CORNWELL, AMBER)(DORNY, JAMES)(DORRIES, PAUL)(GATES, JOHN)(KERN, ANN)(KORN, ANNA)(META, ALEX)(RUSSEL, ELIZABETH)(STEVE, LEWIS)(WAHL, MICHAEL)")

	//	fmt.Println(FindUniq([]float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0}))
	//	fmt.Println(FindUniq([]float32{0, 0, 0.55, 0, 0}))
	//	fmt.Println(FindUniq([]float32{8, 7, 7, 7, 7, 7, 7, 7, 7}))
	//	fmt.Println(FindUniq([]float32{7, 7, 7, 7, 7, 7, 7, 8}))
	//	fmt.Println(FindUniq([]float32{1, 0, 0, 0, 0, 0, 0, 0, 0}))
	//	fmt.Println(FindUniq([]float32{0, 1, 1, 1, 1, 1, 1, 1, 1}))
	//
	//	fmt.Println(UniqueFloat32([]float32{7, 7, 7, 7, 7, 7, 7, 8}))

	// fmt.Println("Hello play ground")
	// fmt.Println(MyStr.IsUpperCase("c"))
	// fmt.Println(MyStr.IsUpperCase("C"))
	// fmt.Println(MyStr.IsUpperCase("ABC"))
	//
	// fmt.Println(NearestSq(9))
	// fmt.Println(NearestSq(10))
	// fmt.Println(NearestSq(8))
	// fmt.Println(NearestSq(111))
}
