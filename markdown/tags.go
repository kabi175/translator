package markdown

const (
	Header = iota
	Bold
	Italic
	UnderLine
	Img
	Link
	Para
	Break
)

const (
	h1     = `<h1>%s</h1>`
	h2     = `<h2>%s</h2>`
	h3     = `<h3>%s</h3>`
	h4     = `<h4>%s</h4>`
	h5     = `<h5>%s</h5>`
	h6     = `<h6>%s</h6>`
	strong = `<strong>%s</strong>`
	i      = `<i>%s</i>`
	u      = `<u>%s</u>`
	p      = `<p>%s</p>`
	img    = `<img src="%s" alt="%s"></img>`
	ahref  = `<a href="%s">%s</a>`
	br     = `<br>`
)
