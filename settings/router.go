package settings

import (
	"../Http/Routes"
	"fmt"
)

func Router() {
	routes := Routes.Api()
	Process(routes)
}

func Process(routes []Routes.Route) {
	for _, route := range routes {
		fmt.Println(route)
	}
}
