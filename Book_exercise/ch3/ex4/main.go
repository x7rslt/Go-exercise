package main

import (
	"fmt"
	"log"
	"net/http"
)

//前面为演示svg图已经代入web功能，此处只体现构建web服务的功能：
func main() {
	fmt.Println("http://localhost:8000/?color=gray&f=f")
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			switch k {
			case "color":
				color = v[0]
			case "f":
				switch v[0] {
				case "eggbox":
					f1 = eggbox
				case "saddle":
					f1 = saddle
				default:
					f1 = f
				}
			}
		}
		if color == "" {
			color = defaultColor
		}
		if f1 == nil {
			f1 = f
		}
		w.Header().Set("Content-Type", "image/svg+xml")
		svg(w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return

}
