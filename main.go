package main

import (
	"os"
	"text/template"
)

type Item struct {
	Id   int
	Name string
}

type Inventory struct {
	Items []Item
}

const inventoryTmpl = `
{
  "items":[
  {{- range $index, $item := .Items -}}
    {{- if $index -}},{{ end }}
    {
      "id": {{ $item.Id }}, "name": {{ $item.Name }}
    }
    {{- end }}
  ]
}
`

func main() {
	inventory := Inventory{
		Items: []Item{
			{
				Id:   1,
				Name: "Self drilling screws",
			},
			{
				Id:   2,
				Name: "Metric machine screws",
			},
			{
				Id:   3,
				Name: "Self tapping construction screws",
			},
		},
	}

	t, err := template.New("inventory").Parse(inventoryTmpl)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, inventory)
	if err != nil {
		panic(err)
	}
}
