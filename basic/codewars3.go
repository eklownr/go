package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func RoundToNext5(n int) int {
	if n == 0 {
		return 0
	}
	if n%5 == 0 || n < 0 {
		return int(n/5) * 5
	}
	return int(n/5)*5 + 5
}

func StockList(listArt []string, listCat []string) (res string) {
	if len(listArt) == 0 || len(listCat) == 0 {
		return ""
	}

	catMap := make(map[string]int, len(listCat))
	for _, val := range listCat {
		catMap[string(val)] = 0
	}

	for i := 0; i < len(listArt); i++ {
		split := strings.Split(listArt[i], " ")
		cat := strings.Split(split[0], "")
		for key, _ := range catMap {
			if cat[0] == key {
				count, _ := strconv.Atoi(split[1])
				catMap[cat[0]] += count
			}
		}
	}

	for key, val := range listCat {
		res += fmt.Sprintf("(%s : %d) - ", listCat[key], catMap[val])
	}
	return res[:len(res)-3]
}

func dotest(book, cat []string, s string) {
	var ans = StockList(book, cat)
	if ans != s {
		fmt.Println("Error. Expected\n", s, "and you gave\n", ans)
	} else {
		fmt.Println("Good job. \nExp ", s, "\ngot ", ans)
	}
}

type Fighter struct {
	Name            string
	Health          int
	DamagePerAttack int
}

func Attack(fighter Fighter, defender Fighter) {
	defender.Health -= fighter.DamagePerAttack
}
func DeclareWinner2(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
	for fighter1.Health >= 0 && fighter2.Health >= 0 {
		if firstAttacker == fighter1.Name {
			fighter2.Health -= fighter1.DamagePerAttack
			if fighter2.Health <= 0 {
				return fighter1.Name
			}
			fighter1.Health -= fighter2.DamagePerAttack
			if fighter1.Health <= 0 {
				return fighter2.Name
			}
		} else {
			fighter1.Health -= fighter2.DamagePerAttack
			if fighter1.Health <= 0 {
				return fighter2.Name
			}
			fighter2.Health -= fighter1.DamagePerAttack
			if fighter2.Health <= 0 {
				return fighter1.Name
			}
		}
	}
	if fighter1.Health > fighter2.Health {
		return fighter1.Name
	}
	return fighter2.Name
}

func DeclareWinner(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
	for {
		if firstAttacker == fighter1.Name {
			fighter2.Health -= fighter1.DamagePerAttack
			firstAttacker = fighter2.Name
		} else {
			fighter1.Health -= fighter2.DamagePerAttack
			firstAttacker = fighter1.Name
		}

		if fighter1.Health <= 0 {
			return fighter2.Name
		}

		if fighter2.Health <= 0 {
			return fighter1.Name
		}
	}
}

func MultiTable(number int) (table string) {
	for i := 1; i <= 10; i++ {
		table += fmt.Sprintf("%d * %d = %d\n", i, number, i*number)
	}
	return table[:len(table)-1]
}

// 10 min walk
func IsValidWalk(walk []rune) bool {
	if len(walk) != 10 {
		return false
	}
	var n, s, w, e int
	for _, val := range walk {
		switch val {
		case 'n':
			n++
		case 's':
			s++
		case 'w':
			w++
		case 'e':
			e++
		}
	}
	if n == s && w == e {
		return true
	}
	return false
}

type MyString string

func (s MyString) IsUpperCase2() bool {
	str := string(s)
	noBlanks := strings.ReplaceAll(str, " ", "")
	count := 0
	for i := 0; i < len(noBlanks); i++ {
		if unicode.IsUpper(rune(noBlanks[i])) {
			count++
		}
	}
	return count == len(noBlanks)
}
func (s MyString) IsUpperCase() bool {
	return s == MyString(strings.ToUpper(string(s)))
}

func main() {
	fmt.Println("Hello play ground")
	count := 12
	strCount := strconv.Itoa(count)
	str := strings.Join([]string{"a", "b", "c"}, "-")
	fmt.Println(str, strCount)

	fmt.Println(RoundToNext5(12))
	fmt.Println(RoundToNext5(5))
	fmt.Println(RoundToNext5(-7))
	fmt.Println(RoundToNext5(0))

	var b = []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
	var c = []string{"A", "B", "C", "D"}
	dotest(b, c, "(A : 0) - (B : 1290) - (C : 515) - (D : 600)")

	var d = []string{"ABAR 200", "CDXE 500", "BKWR 250", "BTSQ 890", "DRTY 600"}
	var e = []string{"A", "B"}
	dotest(d, e, "(A : 200) - (B : 1140)")

	f1 := Fighter{"Lew", 10, 2}
	f2 := Fighter{"Harry", 5, 4}
	fmt.Println(DeclareWinner(f1, f2, "Lew"))
	//.To(Equal("Lew"))
	fmt.Println(DeclareWinner(f1, f2, "Harry"))
	// has to return "Harry"

	fmt.Println(MultiTable(5))
	fmt.Println(MultiTable(45))

	r1 := []rune{'n', 's', 'n', 's', 'n', 's', 'n', 's', 'n', 's'}
	r2 := []rune{'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e'}
	r3 := []rune{'w'}
	r4 := []rune{'n', 'n', 'n', 's', 'n', 's', 'n', 's', 'n', 's'}
	r5 := []rune{'n', 's', 'e', 'w', 'n', 's', 'e', 'w', 'n', 's'}
	fmt.Println(IsValidWalk(r1))
	fmt.Println(IsValidWalk(r2))
	fmt.Println(IsValidWalk(r3))
	fmt.Println(IsValidWalk(r4))
	fmt.Println(IsValidWalk(r5))

	fmt.Println("IsUpperCase:")
	fmt.Println(MyString.IsUpperCase("c"))
	fmt.Println(MyString.IsUpperCase("C"))
	fmt.Println(MyString.IsUpperCase("ABC"))
	fmt.Println(MyString.IsUpperCase("C B A"))
	fmt.Println(MyString.IsUpperCase("hello I AM DONALD"))
	fmt.Println(MyString.IsUpperCase("HELLO I AM DONALD"))
	fmt.Println(MyString.IsUpperCase("ACSKLDFJSggggSKLDFJSKLDFJ"))
}
