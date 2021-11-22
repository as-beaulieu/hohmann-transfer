package graphics

import "fmt"

// StartScreen will display the starting title and menu options
func StartScreen() {
	fmt.Println("|                                                                                |")
	fmt.Println("| ~~~  |   |  |~~~     |   |   /->   |   |  |> <|    /|  |>  |  |>  |            |")
	fmt.Println("|  |   |---|  |~~      |---|  |   |  |---|  | V |   /-|  | > |  | > |            |")
	fmt.Println("|  |   |   |  |~~~     |   |   <->   |   |  |   |  /  |  |  >|  |  >|            |")
	fmt.Println("|                                                                                |")
	fmt.Println("| ~~~  |~>|    /|  |>  |  /~>   |~~~  |~~~  |~>|                                 |")
	fmt.Println("|  |   |__|   /-|  | > |  |__   |~~   |~~   |__|                                 |")
	fmt.Println("|  |   |  >  /  |  |  >|  ___|  |     |~~~  |  >                                 |")
	fmt.Println("|                                                                                |")
	fmt.Println("| (S)tart    (O)ptions    (E)xit                                                 |")
}
