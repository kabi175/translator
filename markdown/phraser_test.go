package markdown

import "testing"

func TestEngine(t *testing.T) {
	e := NewEngine()
	if e.header == nil {
		t.Error("engine header regex not found")
	}
	if e.bold == nil {
		t.Error("engine bold regex not found")
	}
	if e.italic == nil {
		t.Error("engine italic regex not found")
	}
	if e.underline == nil {
		t.Error("engine underline regex not found")
	}
	if e.img == nil {
		t.Error("engine img regex not found")
	}
	if e.link == nil {
		t.Error("engine link regex not found")
	}
	if e.br == nil {
		t.Error("engine break regex not found")
	}
}

func TestPhraser(t *testing.T) {
	phraser := NewEngine()
	type testDataType struct {
		input string
		want  string
	}
	var testCases []testDataType = []testDataType{
		{"hello", "<p>hello</p>"},
		{"#hello\n", "<h1>hello</h1>"},
		{"hello\n#hello\n", "<p>hello</p><h1>hello</h1>"},
		{"#hello\nhello", "<h1>hello</h1><p>hello</p>"},
	}
	for _, test := range testCases {
		got := phraser.Phrase(test.input)
		if got != test.want {
			t.Errorf("Want '%s' : Got '%s' \n", test.want, got)
		}
	}
}
