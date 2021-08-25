package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/devries/advent_of_code_2020/utils"
	"github.com/spf13/pflag"
)

var verbose = pflag.BoolP("verbose", "v", false, "verbose output")

type Ingredients map[string]bool

type Label struct {
	Contents  Ingredients
	Allergens []string
}

func main() {
	pflag.Parse()
	f, err := os.Open("input.txt")
	utils.Check(err, "error opening file")
	defer f.Close()

	labels := parseInput(f)
	solution := solve(labels)
	fmt.Println(solution)
}

func solve(labels []Label) string {
	allIngredients := make(Ingredients)
	possibilities := make(map[string]Ingredients)

	for _, label := range labels {
		allIngredients = union(allIngredients, label.Contents)
		for _, allergen := range label.Allergens {
			if possibilities[allergen] == nil {
				possibleAllergens := make(Ingredients)
				possibleAllergens.extend(label.Contents.get())
				possibilities[allergen] = possibleAllergens
			} else {
				possibilities[allergen] = intersection(possibilities[allergen], label.Contents)
			}
		}
	}

	for {
		if r := reduceAllergens(possibilities); r == false {
			break
		}
	}

	allergenList := []string{}
	for k := range possibilities {
		allergenList = append(allergenList, k)
	}

	sort.Strings(allergenList)

	allergicIngredients := []string{}

	for _, allergen := range allergenList {
		if *verbose {
			fmt.Printf("%s contains %s\n", strings.Join(possibilities[allergen].get(), ", "), allergen)
		}
		allergicIngredients = append(allergicIngredients, possibilities[allergen].get()[0])
	}

	return strings.Join(allergicIngredients, ",")
}

func reduceAllergens(possibilities map[string]Ingredients) bool {
	result := false

	for allergen, ing := range possibilities {
		if len(ing) == 1 {
			// Get the only possibility
			var cause string
			for k := range ing {
				cause = k
			}
			for otherAllergen, otherIng := range possibilities {
				if otherAllergen == allergen {
					continue
				}
				if otherIng.contains(cause) {
					result = true
					otherIng.remove(cause)
				}
			}
		}
	}

	return result
}

func parseInput(r io.Reader) []Label {
	lines := utils.ReadLines(r)
	result := []Label{}

	for _, line := range lines {
		sections := strings.Split(line, " (contains ")

		ingredients := strings.Split(sections[0], " ")
		allergens := strings.Split(sections[1][:len(sections[1])-1], ", ")

		ing := make(Ingredients)
		ing.extend(ingredients)

		result = append(result, Label{ing, allergens})
	}

	return result
}

func (ing Ingredients) add(v string) {
	ing[v] = true
}

func (ing Ingredients) remove(v string) {
	delete(ing, v)
}

func (ing Ingredients) extend(values []string) {
	for _, v := range values {
		ing[v] = true
	}
}

func (ing Ingredients) contains(v string) bool {
	return ing[v]
}

func (ing Ingredients) get() []string {
	result := []string{}

	for k, v := range ing {
		if v {
			result = append(result, k)
		}
	}

	return result
}

func union(ing1 Ingredients, ing2 Ingredients) Ingredients {
	result := make(Ingredients)

	for k, v := range ing1 {
		result[k] = v
	}

	for k, v := range ing2 {
		result[k] = v
	}

	return result
}

func intersection(ing1 Ingredients, ing2 Ingredients) Ingredients {
	result := make(Ingredients)

	for k, v := range ing1 {
		if v && ing2[k] {
			result[k] = true
		}
	}

	return result
}
