package markdown

import "testing"

func TestTags(t *testing.T) {
	phraser := NewEngine()
	type testDataType struct {
		input string
		want  string
	}
	var testCases []testDataType = []testDataType{
		{"hello", "<p>hello</p>"},
		{"#hello\n", "<h1>hello</h1>"},
		{"##hello\n", "<h2>hello</h2>"},
		{"###hello\n", "<h3>hello</h3>"},
		{"####hello\n", "<h4>hello</h4>"},
		{"#####hello\n", "<h5>hello</h5>"},
		{"######hello\n", "<h6>hello</h6>"},
		{"**hello**", "<strong>hello</strong>"},
		{"_hello_", "<i>hello</i>"},
		{"__hello__", "<u>hello</u>"},
		{"![alt](http://google.com)", `<img src="http://google.com" alt="alt"></img>`},
		{"[text]:(http://google.com)", `<a href="http://google.com">text</a>`},
	}
	for _, test := range testCases {
		got := phraser.Phrase(test.input)
		if got != test.want {
			t.Errorf("Want '%s' : Got '%s' \n", test.want, got)
		}
	}
}
