package hofs

import (
	"reflect"
	"strings"
	"testing"

	"github.com/pmuens/learn-go-with-tests/generics"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		// Can also use `slices.Equal` (since Go 1.21).
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}

		generics.AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}

		generics.AssertEqual(t, Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})

		generics.AssertTrue(t, found)
		generics.AssertEqual(t, firstEvenNumber, 2)
	})
}

func TestBank(t *testing.T) {
	var (
		riya  = Account{Name: "Riya", Balance: 100}
		chris = Account{Name: "Chris", Balance: 75}
		adil  = Account{Name: "Adil", Balance: 200}

		transactions = []Transaction{
			NewTransaction(chris, riya, 100),
			NewTransaction(adil, chris, 25),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	generics.AssertEqual(t, newBalanceFor(riya), 200)
	generics.AssertEqual(t, newBalanceFor(chris), 0)
	generics.AssertEqual(t, newBalanceFor(adil), 175)
}

func TestProgrammer(t *testing.T) {
	t.Run("find the best programmer", func(t *testing.T) {
		people := []Person{
			{Name: "John Doe"},
			{Name: "Jane Doe"},
		}

		best, found := Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "Jane")
		})

		generics.AssertTrue(t, found)
		generics.AssertEqual(t, best, Person{Name: "Jane Doe"})
	})
}
