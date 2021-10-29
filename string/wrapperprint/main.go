package main

import "fmt"
//每10个字符换行打印
func main() {
	const text = `永和有永和路，中和有中和路，中和的中和路有接永和的中和路，永和的永和路没接中和的永和路；永和的中和路有接永和的永和路，中和的永和路没接中和的中和路。永和有中正路，中和有中正路，永和的中正路用景平路接中和的中正路；永和有中山路，中和有中山路，永和的中山路直接接上了中和的中山路。永和的中正路接上了永和的中山路，中和的中正路却不接中和的中山路。中正桥下来不是中正路，但永和有中正路；秀朗桥下来也不是秀朗路，但永和也有秀朗路。永福桥下来不是永福路，永和没有永福路；福和桥下来不是福和路，但福和路接的是永福桥。`

	const maxWidth = 10

	var lw int // line width

	ps := ""
	for _, r := range text {
		//fmt.Printf("%q\n", r)
		ps = ps + string(r)
		lw += 1
		if lw == 10{
			fmt.Println(ps)
			//fmt.Printf("\n")
			ps = ""
			lw =0
		}
	}

}