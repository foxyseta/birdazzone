package fantacitorio

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"git.hjkl.gq/team13/birdazzone-api/model"
	"git.hjkl.gq/team13/birdazzone-api/twitter"
)

func extractPoliticianPoints(text string) []model.Politician {
	var ret []model.Politician
	lines := strings.Split(text, "\n")
	for i := 0; i < len(lines); i++ {
		re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`) //regex for points
		a := re.FindString(lines[i])
		lines[i] = strings.ToUpper(lines[i])

		if len(a) > 1 { // excludes all points made of 1 char (usually coming from links) or empty strings
			lines[i] = strings.ReplaceAll(lines[i], a, "")
			if strings.Contains(strings.ToUpper(lines[i]), "MALUS") {
				a = "-" + a
			}
			lines[i] = strings.ReplaceAll(lines[i], " PUNTI", "")
			lines[i] = strings.ReplaceAll(lines[i], "MALUS DI ", "")
			lines[i] = strings.ReplaceAll(lines[i], " PER ", "")
			lines[i] = strings.ReplaceAll(lines[i], " - ", "")
			points, err := strconv.Atoi(a)
			if err == nil {
				ret = append(ret, model.Politician{Name: lines[i], Score: points})
			}
		}

	}
	return ret
}

func (this *Politicians) Len() int {
	return len(*this)
}

func (this *Politicians) Less(i int, j int) bool {
	firstPolitician, secondPolitician := (*this)[i], (*this)[j]
	firstScore, secondScore := firstPolitician.Score, secondPolitician.Score
	if firstScore < secondScore {
		return false
	}
	if firstScore > secondScore {
		return true
	}
	firstName := strings.Split(firstPolitician.Name, " ")
	secondName := strings.Split(secondPolitician.Name, " ")
	return firstName[1] < secondName[1] || firstName[1] == secondName[1] &&
		firstName[0] < secondName[0]
}

func (this *Politicians) Swap(i int, j int) {
	(*this)[i], (*this)[j] = (*this)[j], (*this)[i]
}

type Politicians []model.Politician

func sortPoliticiansByPoints(politicians *Politicians) {
	sort.Sort(politicians)
}

func getPoliticiansPoints() ([]model.Politician, error) {
	result, err := twitter.GetManyRecentTweetsFromQuery("from:fanta_citorio punti -squadre -squadra", "", "")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var arr Politicians
	tweets := result.Data
	for i := 0; i < result.Meta.ResultCount; i++ {
		arr = append(arr, extractPoliticianPoints(tweets[i].Text)...)
	}
	sortPoliticiansByPoints(&arr)
	return arr, err
}
