package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

	var version = flag.Int("version", 1, "first or second part of the assignment")

	flag.Parse()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []int{}

	for scanner.Scan() {
		l := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}
		iL, _ := strconv.Atoi(l)

		lines = append(lines, iL)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	res := getChargers(lines)

	log.Printf("Version: %d ,Result: %d , Result v2: %d 🎄", *version, res.getDiffNumber(), len(getCombinitionsPossible(lines)))

}

func getChargers(input []int) chargersummary {

	curJolt := 0

	sort.Ints(input)

	hInput := input[len(input)-1]

	input = append(input, hInput+3)

	sort.Ints(input)
	cs := chargersummary{
		chargers: []charger{},
	}

	for _, i := range input {
		okInput := false
		c := charger{
			value:      i,
			differance: i - curJolt,
		}

		if i-curJolt == 1 {
			okInput = true
		}
		if i-curJolt == 2 {
			okInput = true
		}
		if i-curJolt == 3 {
			okInput = true
		}

		if okInput {
			curJolt = i
			cs.chargers = append(cs.chargers, c)
		}
	}

	return cs
}

func getCombinitionsPossible(input []int) map[string]bool {

	sort.Ints(input)

	hInput := input[len(input)-1]

	input = append(input, hInput+3)

	sort.Ints(input)

	curJolt := 0
	initalTry := map[int]int{}
	validIndex := 0

	for _, i := range input {
		if validateInput(curJolt, i) {
			initalTry[validIndex] = i
			validIndex++
			curJolt = i
		}

	}
	log.Printf("what the current try %+v", initalTry)

	deviceJolt := 0
	intialTryArr := []int{}
	for _, v := range initalTry {
		intialTryArr = append(intialTryArr, v)
		if v > deviceJolt {
			deviceJolt = v
		}
	}

	sort.Ints(intialTryArr)

	validCombination := map[string]bool{}

	// for startIndex := 0; startIndex < len(initalTry); startIndex++ {
	// 	for skipIndex := 0; skipIndex < len(initalTry); skipIndex++ {
	// 		for startSkipAt := 0; startSkipAt < len(initalTry); startSkipAt++ {
	// 			for endSkipAt := 0; endSkipAt < len(initalTry); endSkipAt++ {

	findpool := 0
	for start := 0; start < hInput; start += 3 {
		findpool++
	}

	for pool := 1; pool < len(intialTryArr); pool++ {
		log.Printf("current pool %d", pool)
		// p := Combinations(intialTryArr[:len(intialTryArr)-1], pool)
		p := Pool(pool, intialTryArr[:len(intialTryArr)-1], findpool, deviceJolt)

		// skippedStart := 0
		// skippedSkipEnd := 0
		for _, sp := range p {
			curJolt := 0
			currentTry := map[int]int{}
			validIndex := 0
			sp = append(sp, deviceJolt)
			for _, v := range sp {
				// if startSkipAt == i && skipIndex > skippedStart {
				// 	skippedStart++
				// 	continue
				// }

				// if endSkipAt == i && skipIndex > skippedSkipEnd {
				// 	skippedSkipEnd++
				// 	continue
				// }
				if validateInput(curJolt, v) {
					currentTry[validIndex] = v
					validIndex++
					curJolt = v
				}
			}
			tryDeviceJolt := 0
			sortedTry := []int{}
			for _, v := range currentTry {
				sortedTry = append(sortedTry, v)
				if v > tryDeviceJolt {
					tryDeviceJolt = v
				}
			}

			if tryDeviceJolt == deviceJolt {
				combination := ""
				sort.Ints(sortedTry)
				for _, v := range sortedTry {
					combination += strconv.Itoa(v)
				}
				validCombination[combination] = true
				// log.Printf("valid: %d", len(validCombination))
			}
		}
	}

	return validCombination
}

func rPool(p int, n []int, c []int, cc [][]int, minLength int, high int) [][]int {
	if len(n) == 0 || p <= 0 {
		return cc
	}
	p--

	for i := range n {
		r := make([]int, len(c)+1)
		copy(r, c)

		r[len(r)-1] = n[i]
		if p == 0 {
			if len(r) > minLength && high-r[len(r)-1] < 4 && r[0]-0 < 4 {
				//log.Printf("adding %+v", r)
				cc = append(cc, r)
			}
		}

		if len(r) > 1 {
			if r[len(r)-1]-r[len(r)-2] > 3 {
				break
			}
			if r[0]-0 > 3 {
				break
			}
			diff := 0
			for i, v := range r[1:] {
				//log.Printf("prev %d , next %d", r[i], v)
				diff += v - r[i]
			}

			if diff+((p+1)*3) < high-3 {
				//log.Printf("breaking %d , %d , sum: %d , high %d , r: %+v", diff, p, diff+((p)*3), high, r)
				continue
			}

			//log.Printf("diff: %d whats p: %d, r:%+v , high: %d", diff, p, r, high)
		}

		cc = rPool(p, n[i+1:], r, cc, minLength, high)

	}
	return cc
}

func Pool(p int, n []int, minLength, high int) [][]int {
	return rPool(p, n, nil, nil, minLength, high)
}

func validateInput(cur, v int) bool {
	okInput := false

	if v-cur == 1 {
		okInput = true
	}
	if v-cur == 2 {
		okInput = true
	}
	if v-cur == 3 {
		okInput = true
	}

	return okInput
}

func (cs *chargersummary) getDiffNumber() int {
	diff1 := 0
	diff3 := 0

	for _, c := range cs.chargers {
		if c.differance == 1 {
			diff1++
		}

		if c.differance == 3 {
			diff3++
		}
	}
	return diff1 * diff3
}

type (
	chargersummary struct {
		chargers []charger
	}

	charger struct {
		value      int
		differance int
	}
)
