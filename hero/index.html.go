// Code generated by hero.
// source: /Users/Lime/Documents/workspace/GoProject/src/goTemplateBenchmark/hero/index.html
// DO NOT EDIT!
package hero

import (
	"bytes"
	"html"
	"strconv"

	"github.com/SlinSo/goTemplateBenchmark/model"
)

import "github.com/shiyanhui/hero"

func Index(u *model.User, nav []*model.Navigation, title string) *bytes.Buffer {
	_buffer := hero.GetBuffer()
	_buffer.WriteString(`

<!DOCTYPE html>
<html>
<body>

<header>
    `)
	_buffer.WriteString(`<title>`)
	_buffer.WriteString(html.EscapeString(title))
	_buffer.WriteString(`'s Home Page</title>
<div class="header">Page Header</div>
`)

	_buffer.WriteString(`
</header>

<nav>
    `)
	_buffer.WriteString(`<ul class="navigation">
`)
	for _, item := range nav {
		_buffer.WriteString(`
	<li><a href="`)
		_buffer.WriteString(html.EscapeString(item.Link))
		_buffer.WriteString(`">`)
		_buffer.WriteString(html.EscapeString(item.Item))
		_buffer.WriteString(`</a></li>
`)
	}
	_buffer.WriteString(`
</ul>
`)

	_buffer.WriteString(`
</nav>

<section>
<div class="content">
	<div class="welcome">
		<h4>Hello `)
	_buffer.WriteString(html.EscapeString(u.FirstName))
	_buffer.WriteString(`</h4>

		<div class="raw">`)
	_buffer.WriteString(u.RawContent)
	_buffer.WriteString(`</div>
		<div class="enc">`)
	_buffer.WriteString(html.EscapeString(u.EscapedContent))
	_buffer.WriteString(`</div>
	</div>

`)
	for i := 1; i <= 5; i++ {
		if i == 1 {
			_buffer.WriteString(`
			<p>`)
			_buffer.WriteString(html.EscapeString(u.FirstName))
			_buffer.WriteString(` has `)
			_buffer.WriteString(html.EscapeString(strconv.FormatInt(int64(i), 10)))
			_buffer.WriteString(` message</p>
		`)
		} else {
			_buffer.WriteString(`
			<p>`)
			_buffer.WriteString(html.EscapeString(u.FirstName))
			_buffer.WriteString(` has `)
			_buffer.WriteString(html.EscapeString(strconv.FormatInt(int64(i), 10)))
			_buffer.WriteString(` messages</p>
		`)
		}
	}
	_buffer.WriteString(`
</div>
</section>

<footer>
    `)
	_buffer.WriteString(`<div class="footer">copyright 2016</div>
`)

	_buffer.WriteString(`
</footer>

</body>
</html>
`)

	return _buffer
}