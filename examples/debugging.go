package main

import (
	"fmt"

	"simplereach/go-handlebars"
)

func main() {
    fmt.Println("Rendering test.html.hbs")
    debugFileRender()

    fmt.Println("Rendering test string")
    debugStringRender()
}

func debugFileRender() {
	renderContext := map[string]interface{}{
		"list:test": []interface{}{
			map[string]string{"firstName": "Eric", "lastName": "Lubow"},
			map[string]string{"firstName": "Matt", "lastName": "Reiferson"},
		},
		"title": "Person List",
	}
	fmt.Printf("DEBUG input: %v\n\n", renderContext)
	fmt.Printf("DEBUG result:\n%v\n",
		handlebars.RenderFile("test.html.hbs", renderContext))
}

func debugStringRender() {
	renderContext := map[string]interface{}{
		"list:test": []interface{}{
			map[string]string{"firstName": "Eric", "lastName": "Lubow"},
			map[string]string{"firstName": "Matt", "lastName": "Reiferson"},
		},
		"title": "Person List",
	}
	fmt.Printf("DEBUG input: %v\n\n", renderContext)
	fmt.Printf("DEBUG result:\n%v\n",
		handlebars.Render("<ul>{{#list:test}}<li>{{firstName}} {{lastName}}</li>{{/list:test}}</ul>", renderContext))
}
