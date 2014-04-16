package main

import (
	"github.com/madari/goskirt"
	"os"
)

func main() {

	data := `

#this is a title

##this is second title

1. me
2. he
3. they

+ me
+ hi
+ me

      var me = your



[me]("http://www.google.com")
`

	skirt := goskirt.Goskirt{
		goskirt.EXT_AUTOLINK | goskirt.EXT_STRIKETHROUGH,
		goskirt.HTML_SMARTYPANTS | goskirt.HTML_USE_XHTML,
	}

	skirt.WriteHTML(os.Stdout, []byte(data))
}
