//source: https://gist.github.com/ihkN/3c796e613f179ae223bd283775d904f8

//Counting sort används för att sortera listan, inspiration från övning 3

package main

//Lånat från yourbasic.org, jämför slices T(n) är O(n)
// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {

			return false
		}
	}
	return true
}

//Hittar maxelementet i en vektor, tar in en slice och returnerar en int, T(n) är O(n)
func findmax(vek []int) int {
	temp := vek[0]
	for _, v := range vek {
		if temp < v {
			temp = v
		}
	}
	return temp
}

//Hittar maxelementet i en vektor, tar in en slice och returnerar en int, T(n) är O(n)
func findmin(vek []int) int {
	temp := vek[0]
	for _, v := range vek {
		if temp > v {
			temp = v
		}
	}
	if temp >= 0 {
		temp = 0
	}
	return temp
}

//Hittar indexet för maxelementet i en vektor, tar in en slice och returnerar en int, T(n) är O(n)
func findindex(vek []int) int {
	temp := vek[0]
	index := 0
	for i, v := range vek {
		if temp < v {
			temp = v
			index = i
		}
	}
	return index
}

//gör en tom vektor med längd [max], T(n) är O(n)
func makerange(max, min int) []int {
	vek := make([]int, max+1-min)
	for i := range vek {
		vek[i] = 0 //fyller vektorn med nollor
	}
	return vek
}

//Sorterar listan enligt counting sort-algoritmen, tar in en slice och returnerar en slice, T(n) är O(n+k)
func Countingsort(vek []int) []int {
	if len(vek) == 0 {
		panic("Listan är tom!")
	}
	max := findmax(vek)
	min := findmin(vek)
	counter := makerange(max, min)

	for _, v := range vek {
		counter[v-min] += 1
	}

	for i := 1; i < len(counter); i++ {
		counter[i] += counter[i-1]
	}

	result := make([]int, len(vek))

	for i := 0; i < len(vek); i++ {
		j := vek[i]
		k := counter[j-min]
		result[k-1] = j
		counter[j-min] -= 1
	}
	return result
}

//Testfall
func tcount() {
	vek := []int{2, 1}
	testvek := []int{1, 2}
	if !equal(Countingsort(vek), testvek) {
		panic("Fel!")
	}

	vek = []int{200, 0, 0, -200, 2, 1, 3}
	testvek = []int{-200, 0, 0, 1, 2, 3, 200}
	if !equal(Countingsort(vek), testvek) {
		panic("Fel!")
	}

	vek = []int{1, 1, 1, 1, 1, 1, 1}
	testvek = []int{1, 1, 1, 1, 1, 1, 1}
	if !equal(Countingsort(vek), testvek) {
		panic("Fel!")
	}

	vek = []int{1}
	testvek = []int{1}
	if !equal(Countingsort(vek), testvek) {
		panic("Fel!")
	}
}

func main() {
	tcount()
}
