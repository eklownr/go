package main

import (
	"fmt"
	"strconv"
	"strings"
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
}
