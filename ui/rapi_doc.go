package ui

import "fmt"

func RapiDocHTML(config Config) string {
	return fmt.Sprintf(`
			<!doctype html>
			<html>
			  <head>
				<meta charset="utf-8">
				<title>%s</title>
				<script type="module" src="https://unpkg.com/rapidoc/dist/rapidoc-min.js"></script>
			  </head>
			  <body>
				<rapi-doc
				  spec-url = "%s"
				> </rapi-doc>
			  </body>
			</html>
		`, config.Title, config.SpecUrl)
}
