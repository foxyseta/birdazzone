package fantacitorio

import (
	"sort"
	"strconv"
	"strings"

	"git.hjkl.gq/team13/birdazzone-api/model"
	"git.hjkl.gq/team13/birdazzone-api/twitter"
	"github.com/timtadh/lexmachine"
	"github.com/timtadh/lexmachine/machines"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const MALUS_TOKEN = 0
const POINTS_TOKEN = 1
const NAME_TOKEN = 2

var lexer *lexmachine.Lexer = newFantacitorioLexer()
var properNames cases.Caser = cases.Title(language.Italian)

// a lexmachine.Action function which skips the match.
func skip(*lexmachine.Scanner, *machines.Match) (interface{}, error) {
	return nil, nil
}

// a lexmachine.Action function with constructs a Token of the given token type by
// the token type's name.
func token(tokenType int) lexmachine.Action {
	return func(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
		return s.Token(tokenType, string(m.Bytes), m), nil
	}
}

func newFantacitorioLexer() *lexmachine.Lexer {
	res := lexmachine.NewLexer()
	res.Add([]byte("MALUS"), token(MALUS_TOKEN))
	res.Add([]byte("[\\+\\-]?\\d+"), token(POINTS_TOKEN))
	res.Add([]byte("(\\W*(RT|PUNTI|A\\W|PRIMI|DI|PER|TOT|[@#]\\w+)\\W*)+"), skip)
	res.Add([]byte("\\W+"), token(99))
	res.Add([]byte("[A-Z\\'\\ ]+[A-ZÀÈÌÒÙ]"), token(NAME_TOKEN))
	err := res.Compile()
	if err != nil {
		panic(err.Error())
	}
	return res
}

func updatePoliticiansList(points *int, names *[]string, malus *bool, res *map[string]int) {
	if *points != -1 && len(*names) > 0 {
		if *malus && *points > 0 {
			*points = -*points
		}
		for _, name := range *names {
			_, ok := (*res)[name]
			if !ok {
				(*res)[name] = *points
			}
		}
		*malus, *points, *names = false, -1, []string{}
	}
}

func extractPoliticianPoints(text string) (map[string]int, error) {
	text = strings.ToUpper(text)
	res := make(map[string]int, 0)
	scanner, err := lexer.Scanner([]byte(text))
	if err != nil {
		return nil, err
	}
	malus, points, names := false, -1, []string{}
	for tok, err, eos := scanner.Next(); !eos; tok, err, eos = scanner.Next() {
		if ui, is := err.(*machines.UnconsumedInput); is {
			// skip the error via:
			scanner.TC = ui.FailTC
		} else if err != nil {
			return nil, err
		}
		token := tok.(*lexmachine.Token)
		switch token.Type {
		case MALUS_TOKEN:
			malus = true
		case POINTS_TOKEN:
			updatePoliticiansList(&points, &names, &malus, &res)
			points, _ = strconv.Atoi(string(token.Lexeme))
			updatePoliticiansList(&points, &names, &malus, &res)
		case NAME_TOKEN:
			names = append(names, properNames.String(string(token.Lexeme)))
		}
	}
	updatePoliticiansList(&points, &names, &malus, &res)
	return res, nil
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
	// From Twitter
	result, err := twitter.GetManyRecentTweetsFromQuery("(from:fanta_citorio OR from:birdazzone) punti -squadre -squadra", "", "")
	if err != nil {
		return nil, err
	}
	tweets := result.Data
	// Map every politician name to its score
	dict := make(map[string]int, 0)
	for i := 0; i < result.Meta.ResultCount; i++ {
		politicians, err := extractPoliticianPoints(tweets[i].Text)
		if err != nil {
			return nil, err
		}
		for k, v := range politicians {
			_, ok := dict[k]
			if !ok {
				dict[k] = v
			}
		}
	}
	// Map to sorted array of politicians
	arr := make(Politicians, len(dict))
	i := 0
	for k, v := range dict {
		arr[i] = model.Politician{Name: k, Score: v}
		i++
	}
	sortPoliticiansByPoints(&arr)
	return arr, err
}
