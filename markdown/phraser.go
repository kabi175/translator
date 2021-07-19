package markdown

import (
	"fmt"
	"regexp"
	"sort"
)

type engine struct {
	header    *regexp.Regexp
	bold      *regexp.Regexp
	italic    *regexp.Regexp
	underline *regexp.Regexp
	img       *regexp.Regexp
	link      *regexp.Regexp
	br        *regexp.Regexp
}

func NewEngine() *engine {
	return &engine{
		header:    regexp.MustCompile(`(?m)^#+.*\n`),
		bold:      regexp.MustCompile(`\*\*.*\*\*`),
		italic:    regexp.MustCompile(`_[^_].*[^_]_`),
		underline: regexp.MustCompile(`__[^_].*[^_]__`),
		img:       regexp.MustCompile(`!\[.*\]\(.*\)`),
		link:      regexp.MustCompile(`\[.*\]:\(.*\)`),
		br:        regexp.MustCompile(`(?m)^---+\n`),
	}
}

func (e *engine) process(markdown string, tag int8) contaniers {
	var (
		index  [][]int
		htmlCn contaniers
		exp    *regexp.Regexp
	)
	switch tag {
	case Header:
		exp = e.header
		break
	case Bold:
		exp = e.bold
		break
	case Italic:
		exp = e.italic
		break
	case UnderLine:
		exp = e.underline
		break
	case Img:
		exp = e.img
		break
	case Link:
		exp = e.link
		break
	case Break:
		exp = e.br
		break
	default:
		return htmlCn
	}

	index = exp.FindAllStringIndex(markdown, -1)
	for _, pos := range index {
		htmlCn = append(htmlCn, contanier{pos[0], pos[1], tag})
	}
	return htmlCn
}

func (e engine) phraseTag(c contanier, markdown string) string {
	switch c.tag {
	case Header:
		{
			headerCount := 0
			for _, char := range markdown[c.start:c.end] {
				if char == '#' {
					headerCount++
				}
			}

			header := func() string {
				switch headerCount {
				case 1:
					return h1
				case 2:
					return h2
				case 3:
					return h3
				case 4:
					return h4
				case 5:
					return h5
				case 6:
					return h6
				}
				return h1
			}()

			return fmt.Sprintf(header, markdown[c.start+headerCount:c.end-1])
		}
	case Bold:
		return fmt.Sprintf(strong, markdown[c.start+2:c.end-2])
	case Italic:
		return fmt.Sprintf(i, markdown[c.start+1:c.end-1])
	case UnderLine:
		return fmt.Sprintf(u, markdown[c.start+2:c.end-2])
	case Img:
		r := regexp.MustCompile(`\[.*\]`)
		src := r.FindString(markdown[c.start:c.end])
		r = regexp.MustCompile(`\(.*\)`)
		alt := r.FindString(markdown[c.start:c.end])
		return fmt.Sprintf(img, alt[1:len(alt)-1], src[1:len(src)-1])
	case Link:
		r := regexp.MustCompile(`\[.*\]`)
		src := r.FindString(markdown[c.start:c.end])
		r = regexp.MustCompile(`\(.*\)`)
		text := r.FindString(markdown[c.start:c.end])
		return fmt.Sprintf(ahref, text[1:len(text)-1], src[1:len(src)-1])
	case Break:
		return br
	default:
		return ""
	}
}

func (e engine) construct(markdown string, htmlCn contaniers) string {
	var (
		html string
		last int = -1
	)

	for _, c := range htmlCn {
		if last < c.start && c.start != 0 && len(markdown[last+1:c.start-1]) > 0 {
			html += fmt.Sprintf(p, markdown[last+1:c.start-1])
		}
		html += e.phraseTag(c, markdown)
		last = c.end
	}

	if last == -1 {
		return fmt.Sprintf(p, markdown)
	}

	if last+1 < len(markdown) && len(markdown[last:]) > 0 {
		html += fmt.Sprintf(p, markdown[last:])
	}

	return html
}

func (e *engine) Phrase(markdown string) string {
	var htmlCn contaniers
	htmlCn = append(htmlCn, e.process(markdown, Header)...)
	htmlCn = append(htmlCn, e.process(markdown, Bold)...)
	htmlCn = append(htmlCn, e.process(markdown, Italic)...)
	htmlCn = append(htmlCn, e.process(markdown, UnderLine)...)
	htmlCn = append(htmlCn, e.process(markdown, Img)...)
	htmlCn = append(htmlCn, e.process(markdown, Link)...)
	htmlCn = append(htmlCn, e.process(markdown, Break)...)
	sort.Sort(htmlCn)
	htmlCn = htmlCn.Validate()
	return e.construct(markdown, htmlCn)
}
