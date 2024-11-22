package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type MyStr string

func (s MyStr) IsUpperCase() bool {
	ett := 1
	fmt.Println(strconv.Itoa(ett))
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
	names := strings.Split(strings.ToUpper(s), ";")
	result := []string{}
	for _, name := range names {
		parts := strings.Split(name, ":")
		result = append(result, "("+parts[1]+", "+parts[0]+")")
	}
	sort.Strings(result)
	return strings.Join(result, "")
}
func GetCount(str string) (count int) {
	s := "aeiouAEIOU"
	for _, char := range str {
		for _, vowel := range s {
			if char == rune(vowel) {
				count++
			}
		}
	}
	return count
}

// replace DNA T to A, A to T, C to G, G to C
func DNAStrand(dna string) string {
	dnaMap := map[string]string{
		"A": "T",
		"T": "A",
		"C": "G",
		"G": "C",
	}
	var result string
	for _, char := range dna {
		result += dnaMap[string(char)]
	}
	return result
}
func PartList(arr []string) (res string) {
	for i := 1; i < len(arr); i++ {
		// join all elements before i to index i [:i], join all elements after i to the end [i:]
		res += fmt.Sprintf("(%s, %s)", strings.Join(arr[:i], " "), strings.Join(arr[i:], " "))
	}
	return res
}
func CountPositivesSumNegatives(numbers []int) []int {
	res := []int{0, 0}

	for _, num := range numbers {
		switch {
		case num > 0:
			res[0]++
		case num < 0:
			res[1] += num
		}
	}
	return res
}
func Arithmetic(a int, b int, operator string) int {
	switch {
	case operator == "add":
		return a + b
	case operator == "subtract":
		return a - b
	case operator == "multiply":
		return a * b
	case operator == "divide":
		return a / b
	}
	return 0
}

func WordsToMarks(s string) int {
	var sum int
	for _, char := range s {
		sum += int(char) - 96
	}
	return sum
}

func FindNb(m int) int {
	for n := 1; m > 0; n++ {
		m -= n * n * n
		if m == 0 {
			return n
		}
	}
	return -1
}
func TowerBuilder(nFloors int) []string {
	tower := make([]string, nFloors)
	if nFloors == 1 {
		return []string{"*"}
	}
	for i := 0; i < nFloors; i++ {
		tower[i] = strings.Repeat(" ", nFloors-i-1) + strings.Repeat("*", 2*i+1) + strings.Repeat(" ", nFloors-i-1) + "\n"
	}
	return tower
}
func sum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}
func Max(arr []int) int {
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}
func minIndex(arr []int) int {
	min := arr[0]
	minIndex := 0
	for i, num := range arr {
		if num < min {
			min = num
			minIndex = i
		}
		return minIndex
	}
	return minIndex
}

func QueueTime1(customers []int, n int) int {
	if n == 1 {
		return sum(customers)
	}
	if n >= len(customers) {
		return Max(customers)
	}
	tills := make([]int, n)
	for _, customer := range customers {
		minTill := minIndex(tills)
		if tills[len(tills)-1] < tills[minTill] {
			tills[len(tills)-1] += customer
		} else {
			tills[minTill] += customer
		}
	}
	return Max(tills)
}
func QueueTime2(customers []int, n int) int {
	result := 0
	if n == 1 {
		for _, v := range customers {
			result += v
		}
		return result
	}
	queues := make([]int, n)
	for _, v := range customers {
		fmt.Println(queues[0])
		queues[0] += v
		sort.Ints(queues)
	}
	result = queues[n-1]
	return result
}

func QueueTime(customers []int, n int) int {
	queueTime := make([]int, n)

	for _, customerTime := range customers {
		queueTime[0] += customerTime
		sort.Ints(queueTime)
	}
	return queueTime[n-1]
}

func main() {
	l := []int{9, 7, 5, 3, 1}
	fmt.Println(l)
	fmt.Println(QueueTime(l, 2))                       //.To(Equal(16))
	fmt.Println(QueueTime([]int{}, 1))                 //.To(Equal(0))
	fmt.Println(QueueTime([]int{1, 2, 3, 4}, 1))       //.To(Equal(10))
	fmt.Println(QueueTime([]int{2, 2, 3, 3, 4, 4}, 2)) //.To(Equal(9))
	fmt.Println(QueueTime([]int{1, 2, 3, 4, 5}, 100))  //.To(Equal(5))

	//
	//	fmt.Println(TowerBuilder(0))
	//	fmt.Println(TowerBuilder(1))
	//	fmt.Println(TowerBuilder(2))
	//	sList := []string{"I", "wish", "I", "hadn't", "come"}
	//	fmt.Println("(I, wish I hadn't come)(I wish, I hadn't come)(I wish I, hadn't come)(I wish I hadn't, come)")
	//	fmt.Println(PartList(sList))
	//
	//	sList2 := []string{"cdIw", "tzIy", "xDu", "rThG"}
	//	fmt.Println("(cdIw, tzIy xDu rThG)(cdIw tzIy, xDu rThG)(cdIw tzIy xDu, rThG)")
	//	fmt.Println(PartList(sList2))

	//	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	//	res := []int{10, -65}
	//	result := CountPositivesSumNegatives(arr)
	//	fmt.Println(result, res)

	//fmt.Println(GetCount("abracadabra"))

	//	fmt.Println(Meeting("Alexis:Wahl;John:Bell;Victoria:Schwarz;Abba:Dorny;Grace:Meta;Ann:Arno;Madison:STAN;Alex:Cornwell;Lewis:Kern;Megan:Stan;Alex:Korn"))
	//	//"(ARNO, ANN)(BELL, JOHN)(CORNWELL, ALEX)(DORNY, ABBA)(KERN, LEWIS)(KORN, ALEX)(META, GRACE)(SCHWARZ, VICTORIA)(STAN, MADISON)(STAN, MEGAN)(WAHL, ALEXIS)")
	//	fmt.Println(Meeting("John:Gates;Michael:Wahl;Megan:Bell;Paul:Dorries;James:Dorny;Lewis:Steve;Alex:Meta;Elizabeth:Russel;Anna:Korn;Ann:Kern;Amber:Cornwell"))
	//	//"(BELL, MEGAN)(CORNWELL, AMBER)(DORNY, JAMES)(DORRIES, PAUL)(GATES, JOHN)(KERN, ANN)(KORN, ANNA)(META, ALEX)(RUSSEL, ELIZABETH)(STEVE, LEWIS)(WAHL, MICHAEL)")

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
