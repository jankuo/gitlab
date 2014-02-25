package main

import (
	"text/template"
)

func main() {
	println(template.HTMLEscapeString("<script>me he's & are #meand#</script>"))
}
