package fantacitorio

import (
	"testing"
)

func TestSkip(t *testing.T) {
	token, err := skip(nil, nil)
	if token != nil {
		t.Fatalf("Unexpected token")
	}
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestNewFantacitorioLexer(t *testing.T) {
	lexer := newFantacitorioLexer()
	if lexer == nil {
		t.Fatalf("Nil lexer")
	}
	compilerErr := lexer.Compile()
	if compilerErr != nil {
		t.Fatalf(compilerErr.Error())
	}
}

func TestExtractPoliticiansPoints(t *testing.T) {
	politicians, err := extractPoliticianPoints(`
    400 PUNTI - BARBARA FLORIDIA 
    800 PUNTI A:
    LUIGI NAVE
    ADA LOPREIATO 
    CONCETTA DAMANTE
    GABRIELLA DI GIROLAMO
    SABRINA LICHERI
    ALESSANDRA MAIORINO
    `)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(politicians) != 7 {
		t.Fatalf("Detected politicians are %d, but should be 7", len(politicians))
	}
}
