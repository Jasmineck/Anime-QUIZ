package main
import "fmt"
type quiz struct{
	question string
	option1 string
	option2 string
	option3 string
	option4 string
	answer int
}
func main(){
	
	var name string
	fmt.Printf("OwO Entwr yo name Onwwiichwann :")
	fmt.Scan(&name)  
	//fmt.Printf("Konichiwa  %v, denki desu ka? >u<\n",name) 
	fmt.Println()
	fmt.Println()
	fmt.Println("⠀⠀⠀⠀⣠⣶⡾⠏⠉⠙⠳⢦⡀⠀⠀⠀⢠⠞⠉⠙⠉⠙⠲⡀")
	fmt.Println("⠀⠀⠀⣴⠿⠏⠀⠀⠀⠀⠀⠀ ⢳⡀⠀⡏⠀⠀⠀⠀     ⢷")
	fmt.Println("⠀⠀⢠⣟⣋⡀⢀⣀⣀⡀⠀⣀  ⣧⢸    OK!   ⡇")
	fmt.Printf("⠀⠀⢸⣯⡭⠁⠸⣛⣟⠆⡴⣻⡲ ⣿⣸ %v   ⡇\n",name)
	fmt.Println("⠀⠀⣟⣿⡭⠀⠀⠀⠀⠀⢱⠀⠀ ⣿⠀⢹⠀⠀⠀⠀⠀    ⡇")
	fmt.Println("⠀⠀⠙⢿⣯⠄⠀⠀⠀⢀⡀⠀⠀⡿⠀⠀⡇⠀⠀⠀⠀    ⡼")
	fmt.Println("⠀⠀⠀⠀⠹⣶⠆⠀⠀⠀⠀⠀⡴⠃⠀⠀⠘⠤⣄⣠⣄⣠⣄⠞⠀")
	fmt.Println("⠀⠀⠀⠀⠀⢸⣷⡦⢤⡤⢤⣞⣁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
	fmt.Println("⠀⠀⢀⣤⣴⣿⣏⠁⠀⠀⠸⣏⢯⣷⣖⣦⡀⠀⠀⠀⠀⠀⠀")
	fmt.Println("⢀⣾⣽⣿⣿⣿⣿⠛⢲⣶⣾⢉⡷⣿⣿⠵⣿⠀⠀⠀⠀⠀⠀")
	fmt.Println("⣼⣿⠍⠉⣿⡭⠉⠙⢺⣇⣼⡏⠀⠀⠀⣄⢸⠀⠀⠀⠀⠀⠀")
	fmt.Println("⣿⣿⣧⣀⣿.........⣀⣰⣏⣘⣆⣀⠀⠀")
   
	
	fmt.Println()
	fmt.Println("   ____   ____       ____             ____    ____                          ____    ") 
	fmt.Println(" _(  OO) (  OO) )   ( OO ).       |  ( OO ) _(  OO)                        ( OO ) ) ")
    fmt.Println(" ,--.     (,------./     '._ (_)---| )       ;-----.|(,------. ,----.     ,-.-') ,--./ ,--,'  ")
    fmt.Println(" |  |.-')  |  .---'|'--...__)/    _ |        | .-.  | |  .---''  .-./-')  |  |OO)|   | |  ||  ")
    fmt.Println(" |  | OO ) |  |    '--.  .--'|  :` `.        | '-' /_)|  |    |  |_( O- ) |  |  ||    ||  | ) ")
    fmt.Println(" |  |`-' |(|  '--.    |  |    '..`''.)       | .-. `.(|  '--. |  | .--, | |  |(_/|  .     |/  ")
    fmt.Println(" |  '---.' |  .--'    |  |   .-._)   |       | |  | ||  .--' (|  | '. (_/,|  |_.'|  ||    |  ") 
    fmt.Println(" |      |  |  `---.   |  |   |       /       | '--'  /|  `---.|  '--'  |(_|  |   |  | |   |  ") 
    fmt.Println(" `------'  `------'   `--'    `-----'        `------' `------' `------'   `--'   `--'  `--' ")  
    fmt.Println()
	fmt.Println()
	

    var classes = make(map[string][]quiz)
	var s []quiz
	t:=quiz{question:"Ok! Here's an easy one to start off with. What do you have to do with Pokemon?",option1: "Snatch'em all!",option2: "Catch 'em all!",option3: "Fetch 'em all!",option4: "Eat 'em all!",answer:2}
	s= append(s,t)
	t=quiz{question:"What is special about the giant cat in the anime My Neighbour Totoro?",option1: " It is also a plane",option2: "It is also a bus",option3: " It is also a ship",option4: " It is also a train",answer:2}
	s= append(s,t)
	t=quiz{question:"In the anime Fairy Tail, which kind of wizard is Lucy?",option1:"Fire Wizard",option2: "Ice Wizard",option3: "Celestial Wizard",option4: "Horny Wizard",answer:3}
	s= append(s,t)
	t=quiz{question:"The Gum-Gum Pistol is a signature attack of which anime character?",option1: "Monkey D. Luffy",option2: "Black Butler",option3: "Chobits",option4: "Chikka",answer:1}
	s= append(s,t)
	t=quiz{question:"In the anime Fullmetal Alchemist, which Homunculi is created when Alphonse and Edward Elric used alchemy to bring their mom back to life?",option1: "Greed",option2: "Sloth",option3: "Envy",option4: "Jelousy",answer:2}
    s= append(s,t)
	t=quiz{question:"In the anime Death Note, who was the first successor of L in the Kira investigation?",option1: "Light",option2: "MisaMisa",option3: "N",option4: "Mello",answer:1}
	s= append(s,t)
	t=quiz{question:"In the anime Hunter X Hunter, who is the father of Gon Freecss?",option1: "Gin Freecss",option2: "Mito Freecss",option3: "Ging Freecss",option4: "David Freecss",answer:3}
    s= append(s,t)
	t=quiz{question:"Which Japanese anime studio is known for famous works, such as Princess Mononoke, Grave of the Fireflies, Spirited Away, and My Neighbour Totoro?",option1: "Kyoto Animation",option2: "Studio Bones",option3: "Rikka Studio",option4: "Studio Ghibli",answer:4}
	s= append(s,t)
    classes["yeeeeee..........Lets GO!GO!GO!"]=s

    fmt.Println("Befowore getting stwrted with da quiz..Lets do our weeb rituwual for good luck >u< !!!")
	fmt.Println("Say UwU !!!")
	var uwu string
	fmt.Scan(&uwu)
	fmt.Println()
	fmt.Println("Say UwU !!!")
	fmt.Scan(&uwu)
	fmt.Println()
	fmt.Println("Say UwU !!!")
	fmt.Scan(&uwu)
	fmt.Println()
	
	total:=0
    

    for key,value:= range classes {
		fmt.Println(key)
		fmt.Println()
		fmt.Println("Press 69 to Begin!!!")
	    fmt.Scan(&uwu)
        arr := value
		
        for z:=0; z<len(arr); z++ {
           quiz := arr[z]
	       fmt.Println("QUESTION",z+1)
		   fmt.Println()
		   fmt.Println(" ^ ^                 ")
		   fmt.Println("(O,O)                ")
		   fmt.Printf("(   ) %v \n",quiz.question    )
		   fmt.Println("----------------------------------------------------------------------------------------------- ")
						
		   fmt.Printf("1: %v \n" ,quiz.option1)
		   fmt.Println()
		   fmt.Printf("2: %v \n" ,quiz.option2)
		   fmt.Println()
		   fmt.Printf("3: %v \n" ,quiz.option3)
		   fmt.Println()
		   fmt.Printf("4: %v \n" ,quiz.option4)
		   fmt.Println()
          
		   fmt.Printf("Enter your awnswer:")
		   var ans int
		   fmt.Scan(&ans)
		   var org_ans int
		   org_ans=quiz.answer
		   if ans==org_ans{
			
			            fmt.Println("⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠁⠀⠀⠀⠀⠠⠤⠶⠞⢻⣿⡿⣿⣿⣿⣿⣿⣿")
			            fmt.Println("⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠟⠁⠀⢀⣠⣤⣤⣴⣶⣄⠀⢸⣿⠇⠻⣿⣿⣿⣿⣿")
			            fmt.Println("⣿⣿⣿⣿⣿⣿⣿⣿⣿⠋⠀⠀⠰⠛⠛⠛⠻⠿⠿⣿⡇⠈⠉⠀⠀⠈⠻⣿⣿⣿")
			            fmt.Println("⣿⣿⣿⣿⣿⣿⣿⣿⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣤⣄⡀⢹⣿")
			            fmt.Println("⣿⣿⣿⣿⣿⣿⢿⢏⡜⢱⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠾⢿⣿⠻⣿⣿⣿")
			            fmt.Println("⣿⣿⣿⣿⣿⡿⢸⡞⣠⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠛⠀⠀⠹⡍")
			            fmt.Println("⣉⠙⠻⣿⣹⡇⡞⢰⡟⠀⣠⠤⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢹")
			            fmt.Println("⠈⡹⣶⢆⣿⢱⡇⢸⡷⠋⣠⣤⡈⢇⠀⠀⠀⠀⠀⠀⠀⢀⡤⠄⢠⡀⠀⠀⠀⠈")
			            fmt.Println("⣇⣿⢸⣿⠸⣷⠀⢧⣾⠋⠈⠻⣾⣦⠀⠀⠀⠀⠀⣴⠋⢀⣦⠀⢿⠀⠀⠀⢀")
			            fmt.Println("⡀⠈⣿⠘⢿⠄⠈⢀⠸⡏⠀⠀⢰⡇⡜⠀⠀⠀⠀⠀⠁⠀⠈⢸⠈⠀⠀⠀⠀⡼")
			            fmt.Println("⣿⣷⣿⠀⠀⠀⠀⡌⠀⢧⣀⡴⠛⢁⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣾⡇")
			            fmt.Println("⣿⣿⣿⡇⠀⠀⠰⢰⠀⠀⠙⠃⢀⡾⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⣿⣿⡇")
			            fmt.Println("⣿⣿⣿⣷⠀⠀⠀⡸⠀⠀⠀⣠⣿⣿⣶⣤⣤⣀⡀⠀⠀⠀⠀⢀⣴⣿⣿⣿⣿⡇")
			            fmt.Println("⣿⣿⣿⠏⠀⠀⢀⡇⢀⣠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣶⡞⠋⢸⣿⣿⣿⣿⡇")
			            fmt.Println("⣿⡿⠃⠀⠐⠶⣿⡿⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣞⢻⣿⣿⣿⣿⡇")

						fmt.Println()
						fmt.Println()

			fmt.Println(" ______                              __   __ __ ")
			fmt.Println("|      |.-----.----.----.-----.----.|  |_|  |  |")
			fmt.Println("|   ---||  _  |   _|   _|  -__|  __||   _|__|__|")
			fmt.Println("|______||_____|__| |__| |_____|____||____|__|__|")
			fmt.Println()
			total=total+1												
		   }else {
			fmt.Println()
			fmt.Println("  W R O N G   A N S W E R !!!")
			fmt.Println()
			fmt.Println("⣿⣿⡿⠟⠛⠛⣿⠛⠛⠛⠛⢻⠛⠛⠛⢻⡟⠛⣿⣿⠛⢻⣿⣿⣿⣿⣿⣿")
	        fmt.Println("⣿⣿⡇⠐⠿⣿⣿⣿⡇⢸⣿⣿⠄⢸⣿⣿⡇ ⣿⣿ ⢸⣿⣿⣿⣿⣿⣿")
	        fmt.Println("⣿⣿⣿⣶⣤ ⣿⣿⡇⢸⣿⣿⠄⢰⣶⣾⡇ ⣿⣿ ⢸⣿⣿⣿⣿⣿⣿")
	        fmt.Println("⣿⣿⣇⣉⣁⣤⣿⣿⣇⣸⣿⣿⣀⣸⣿⣿⣿⣤⣈⣁⣤⣿⣿⣿⣿⣿⣿⣿")
	        fmt.Println("⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿")
	        fmt.Println("⡿⢿⣿⡿⠿⣿⡿⢿⡿⠿⠿⠿⣿⡿⠿⠿⠿⣿⠿⠿⠿⠿⣿⣿⠿⠟⠿⣿")
	        fmt.Println("⡇⠈⣿⠄⠄⣿⠃⣸⡇⠄⣶⣶⣿⡇⢰⣶⣶⣿⠄⢸⡷⠄⣿⡇⠰⢿⣶⣿")
	        fmt.Println("⣿⠄⠋⢰⡆⠸⠄⣿⡇⠄⣤⣤⣿⡇⢠⣤⣤⣿⠄⢰⣤⠈⢻⣿⣦⣄⠈⢿")
	        fmt.Println("⣿⣇⣀⣾⣷⣀⣸⣿⣇⣀⣉⣉⣹⣇⣈⣉⣉⣿⣀⣈⣉⣠⣾⣋⣉⣉⣠⣿")
				
		   }
		   
		   

           if z==(len(arr)-1){
			fmt.Println("Press 69 to PeekPeek OwO your twotal scwore!!!")
			fmt.Println()
			fmt.Println("⣿⣿⡆⠀⠀⢸⣷⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⡇⠀⠀⣾⣿⡆⠀")
			fmt.Println("⣿⣿⡇⠀⠀⢸⣿⢰⣿⡆⠀⣾⣿⡆⠀⣾⣷ ⣿⣿⡇⠀⠀⣿⣿⡇")
			fmt.Println("⣿⣿⡇⠀⠀⢸⣿⠘⣿⣿⣤⣿⣿⣿⣤⣿⡇⢻⣿⡇⠀⠀⣿⣿⡇⠀")
			fmt.Println("⣿⣿⡇⠀⠀⢸⡿⠀⢹⣿⣿⣿⣿⣿⣿⣿⠁⢸⣿⣇⠀⢀⣿⣿⠇⠀")
			fmt.Println("⠙⢿⣷⣶⣶⡿⠁⠀⠈⣿⣿⠟⠀⣿⣿⠇⠀⠈⠻⣿⣶⣾⡿⠋⠀⠀")
			fmt.Println()
			fmt.Scan(&uwu)

		   } else {fmt.Println("Press 69 to manifest the nwxt quwuestion!!!")
		   fmt.Println()
		   fmt.Scan(&uwu)
		}

	}
    if total>3 {
		fmt.Println()
		fmt.Println("  -----T O T A L   S C O R E !!!----- :",total)
		fmt.Println()
		             fmt.Println("   ______ _______ _______ _______ _______ ______ _______      __ __ ")
		             fmt.Println("  |   __ |       |     __|     __|    ___|   __ |     __|    |  |  |")
		             fmt.Println("  |    __/   -   |    |  |    |  |    ___|      <__     |    |__|__|")
		             fmt.Println("  |___|  |_______|_______|_______|_______|___|__|_______|    |__|__|")
					 fmt.Println()
					 fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡀⠀⢀⡴⢶⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣠⣤⣤⣶⣤⣄⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⠟⢷⠀⠀")
					 fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⡍⠛⠾⠀⢼⡇⠀⠀⠀⠀⠀⠀⣠⣴⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⣤⡀⠀⠀⠀⠀⠀⠀⡎⠀⠚⠚⠉⢩⡇⠀⠀⠀⠀⠀⠀")
					 fmt.Println("⠀⠀⠀⠀⠀⠀⣀⠀⠀⠀⠀⠀⠀⠙⢦⡀⠀⣼⠁⠀⠀⠀⠀⣠⣾⣿⣯⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⡀⠀⠀⠀⠀⢧⠀⠀⣠⠴⠛⠁⠀⠀⠀⠀⠀⠀")
					 fmt.Println("⠀⠀⠀⢀⣀⣺⣟⣧⠀⠀⠀⠀⠀⠀⠀⠉⠳⠏⠀⠀⠀⢀⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡄⠀⠀⠀⠸⠗⠋⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀")
					 fmt.Println("⠀⠀⠀⠘⠿⣦⣒⣿⠀⠀⠀⣀⣠⣄⣀⠀⣀⡀⠀⠀⣠⣿⣿⣿⣿⣿⣿⡟⢸⣇⢸⣿⠸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
					 fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⣿⣷⡇⡿⢹⠉⣷⣀⣰⣿⣿⣿⣿⣿⣿⡏⠀⢸⣿⡼⢿⣦⠄⢹⣿⠈⠋⢿⣿⣿⣿⣿⣿⣿⣷⠀⠀⣀⣤⡤⡴⢲⣄⡀⠀⠀⠀⠀⠀")
					 fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⢀⣼⣿⣿⣿⣧⣇⣼⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⠓⠒⠻⣿⣷⣼⣿⠃⢸⣿⡴⠠⠸⣿⣿⣿⣿⣿⣿⣿⣷⡍⢯⠈⠁⢳⣾⣿⣿⡄⠀⠀⠀⠀")
					 fmt.Println("⠀⠀⠀⠀⠀⠀⠀⣰⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠏⠀⠀⠀⢳⣿⣷⣿⣦⣸⡟⣿⡻⣷⣿⣿⡿⣿⣿⣿⣿⣿⣷⣾⣆⣀⣸⣿⣿⣿⣧⠀⠀⠀")
					 fmt.Println("⠀⠀⠀⠀⠀⠀⣸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠏⠁⠀⠀⠀⠀⠀⠀⠙⠻⢿⣿⣇⢻⣷⣻⣿⣿⣶⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣄⠀⠀⠀⠀")
					 fmt.Println("⠀⠀⠀⠀⠀⢠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡏⠀⣠⣤⣤⣴⣶⠄⠀⠀⠀⠀⠈⠛⠀⢻⣿⣿⣿⣹⣿⣿⣿⣿⣿⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⡄⠀")
					 fmt.Println("⠀⠀⠀⠀⢀⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣯⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠈⠉⠀⠀⠀⠀⠀⢠⣟⣲⠄⠀⠙⣿⣿⣿⣿⣿⣿⣿⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣇⠀")
					 fmt.Println("⠀⠀⠀⠀⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠛⢿⣦⠀⣿⣿⣿⣿⣿⣿⣿⣸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀")
					 fmt.Println("⠀⠀⢀⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⡀⠀⠀⠀⢠⣴⣦⣤⣄⡀⠀⠀⠀⠀⠀⠈⢡⣿⣿⣿⣿⣿⣿⢻⣧⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀")
					 fmt.Println("⠀⠀⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣄⠀⠀⠸⣯⡀⠙⣻⡇⠀⠀⠀⠀⠀⣰⣿⣿⣿⣿⣿⣿⠏⠀⡿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀")
					 fmt.Println("⢀⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣤⡀⠀⠉⠉⠉⠀⠀⠀⢀⣠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣟⣁⢹⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀")
					 fmt.Println("⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠿⢿⠏⠓⠒⠒⠒⣶⣶⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠈⠉⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣇⠀")
					 fmt.Println("⠘⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⢃⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀")
					 fmt.Println("⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡲⣤⡀⠀⠀⢀⣰⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡄")
					 fmt.Println("⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⡄⠉⢉⣠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡷")
					 fmt.Println("⠀⠘⠻⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠁")
					 fmt.Println("⠀⠀⠀⠀⠀⠉⠉⠉⠛⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠇⠀")
					 fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠙⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢿⠟⠀⠀")
					 fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠟⠋⠉⠀⠀⠀⠀⠀⠀")
					 fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠟⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠟⠛⠛⠛⠋⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
					 fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
					 fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
					 fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
																		  

	} else {
		fmt.Println()
		fmt.Println("  -----T O T A L   S C O R E !!!----- : %d",total)
		             fmt.Println("         ______               __         __ __ ")
		             fmt.Println("        |   __ |.-----.-----.|  |--.    |  |  |")
		             fmt.Println("        |   __ <|  _  |     ||    <     |__|__|")
		             fmt.Println("        |______||_____|__|__||__|__|    |__|__|")
					 fmt.Println()

		fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣠⣤⣤⣤⣤⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
		fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣤⣶⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⣤⣤⢤⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
		fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣽⡲⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀")
		fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣾⣆⠀⠀⠀⠀⠀⠀⠀⠀")
		fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡀⠀⠀⠀⠀⠀⠀⠀")
		fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣄⠀⠀⠀⠀⠀⠀")
		fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣴⡟⢛⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣆⠀⠀⠀⠀⠀")
		fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣾⣿⠀⣼⠀⣰⠏⣉⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⢿⡛⢿⡛⣞⢧⠀⠀⠀⠀")
		fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⠎⣼⡇⠘⠁⠔⠃⠘⠁⢠⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣄⡀⠙⠢⠙⠼⣞⢧⠀⠀⠀")
		fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⠏⣰⣿⣧⡀⠀⠀⠀⠀⠀⣯⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⣿⢋⣿⡿⣿⣿⡋⢠⣀⠉⠁⠀⠀⠀⣸⠚⣧⠀⠀")
		fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡏⢠⣿⣿⣿⣇⠀⠀⠀⠀⠀⠘⣿⣿⣿⣿⣛⣉⣘⠛⠉⠒⠉⠉⠞⠉⠙⠟⠉⢻⣿⣷⣶⠏⠀⠀⠀⠀⢰⡇⠀⢻⠀⠀")
		fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡼⠀⣾⣿⣿⣿⣿⡄⠀⠀⠀⠀⠀⣿⣿⣿⠟⠛⠻⠿⣿⡆⠀⠀⠀⠀⠀⣠⢤⣤⣼⣿⣿⣿⠀⠀⠀⠀⢀⡞⠀⠀⢸⠀⠀")
		fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⢇⣾⣿⣿⣿⣿⣿⡿⠀⠀⠀⠀⣼⣿⣿⣿⣿⣿⣿⣷⡞⠁⠀⠀⠀⠀⠀⠶⣛⣉⣙⢿⣿⣧⠀⠀⠀⠀⢸⠃⠀⠀⠈⠀⠀")
		fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣎⣞⣿⣿⣿⣿⣿⣿⠳⢤⡀⠀⢰⣿⣿⣿⣿⣿⣿⣿⡿⠀⠀⠀⠀⠀⠀⠀⣾⣿⣿⣿⡌⣿⣿⡄⠀⠀⠀⢹⣤⡀⠀⠀⠀⠀")
		fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⡟⣼⣿⣿⣿⣿⡿⠁⠙⠦⣌⠓⣿⣿⣿⡿⠻⣿⣿⠟⠁⠀⠀⠀⠀⠀⠀⠀⢿⣿⣿⡿⢁⣿⣿⣷⡄⠀⠀⢈⠟⣳⡆⠀⠀⠀")
		fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣸⣻⣿⣿⣿⣿⡿⠁⠀⠀⠀⠈⠳⣾⣿⣿⡏⠚⠉⠀⠀⠀⠀⠀⠀⠠⠄⠀⠀⠈⠛⠿⢿⣾⣿⣿⢏⣹⠴⠚⣡⡼⠋⠹⣄⠀⠀")
		fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⠀⣰⢻⣿⣿⣿⣿⡿⠁⠀⠀⠀⠀⠀⠀⠘⣿⣿⢧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣿⣿⣿⣍⠀⣠⡾⠋⠀⠀⠀⠹⣄⠀")
		fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⢠⢟⣿⣿⣿⣿⡿⠁⠀⠀⠀⠀⠀⠀⢀⣴⣿⣿⢿⣦⣀⠀⠀⠘⠿⠽⠛⠛⠽⠇⠀⠀⣀⣴⣿⣿⣿⣦⡈⠉⠁⠀⠀⠀⠀⠀⠀⢹⡄")
		fmt.Println("⠀⠀⠀⠀⠀⠀⠀⠀⡿⣾⣿⣿⣿⣿⠁⠀⠀⠀⠀⠀⠀⢠⣾⣿⣿⣿⢸⠉⠻⠷⣶⣤⣀⣀⠀⢀⣀⣤⣶⠿⠿⣿⣿⣿⣿⣿⡿⣦⡀⠀⠀⠀⠀⠀⠀⠀⢳")
		fmt.Println("⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⠃⠀⠀⠀⠀⠀⠀⢠⣿⣿⣿⣿⣿⣸⠀⠙⠀⣠⣴⠋⢿⠻⣏⠁⠀⠀⠀⠐⡟⢿⣿⣿⣿⣷⠈⣷⡀⠀⠀⠀⠀⠀⠀⠀")
		fmt.Println("⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⡿⠁⠀⠀⠀⠀⠀⠀⢀⣿⣽⣿⣿⣿⣿⠿⢄⠀⠀⢹⡜⠀⢸⠀⠈⢳⣄⠀⠀⠀⢳⣸⣿⣿⣿⣿⡆⡃⠻⡆⠀⠀⠀⠀⠀⠀")
		fmt.Println("⠀⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⠃⠀⠀⠀⠀⠀⠀⠀⣾⡏⡿⣿⣿⣿⣿⡄⠈⠳⣄⢸⠃⠀⠈⡇⠀⠀⠙⣷⠀⠀⢸⣿⣿⣿⣿⣿⣧⠀⢠⡇⠀⠀⠀⠀⠀⠀")
		fmt.Println("⠀⠀⠀⠀⠀⢰⣿⣿⣿⣿⠃⠀⠀⠀⠀⠀⠀⠀⢰⣿⠀⡏⣿⣿⣿⣿⣇⣠⣤⣞⠛⡷⣆⠀⣷⣶⣤⣴⠟⠀⠀⠸⡿⣿⣿⣿⣿⣿⣷⣼⡇⠀⠀⠀⠀⠀⠀")
		fmt.Println("⠀⠀⠀⠀⢠⣿⣿⣿⣿⠋⠀⠀⠀⠀⠀⠀⠀⢀⡿⠻⢿⣷⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠛⠶⣶⡀⢀⡇⣿⣿⣿⣿⣿⣿⠈⣿⡄⠀⠀⠀⠀⠀")
		fmt.Println("⠀⠀⠀⢠⣿⣿⣿⡿⡟⠀⠀⠀⠀⠀⠀⠀⠀⣸⠇⣤⠾⢻⢿⣿⣿⣿⣿⡟⣿⣿⣿⣿⣿⣿⣿⣿⡿⠀⠀⢹⠀⣆⣇⢹⣿⣿⣿⣿⣿⠀⣿⣷⡀⠀⠀⠀⠀")
		fmt.Println("⠀⠀⢠⣿⣿⣿⣿⡷⣷⣤⠀⠀⠀⠀⠀⠀⢠⡟⠀⢿⡀⠸⡏⣿⣿⣿⣿⣿⣿⣿⣏⣹⣿⣿⣧⡀⠀⠀⠀⢹⠀⢹⣿⡄⣿⣿⣿⣿⣿⡜⡘⡏⢳⣄⠀⠀⠀")
		fmt.Println("⠀⢀⣾⣿⣿⣿⠁⣿⡟⣿⣧⠀⠀⠀⠀⢀⡾⠁⠀⠈⣧⠀⢳⢿⣿⣿⣿⣿⣿⣿⢉⣋⣿⣿⣿⣷⡄⠀⠀⣾⠀⢸⣿⣧⣿⣿⣿⣿⣿⡇⣷⣷⠈⢿⠂⠀⠀")
		fmt.Println("⠀⡼⣽⣿⡿⣿⣸⣿⢸⣾⢻⣆⠀⠀⠀⣼⠇⠀⠀⠀⢸⣇⠘⣿⣿⣿⣿⣿⢿⠅⠙⠉⠘⠛⣿⠛⠛⠂⠀⡟⠀⠘⡇⠀⠹⣿⣿⣿⣿⣿⣾⣾⡆⠀⠀⠀⠀")
		fmt.Println("⢸⢷⣿⡟⣀⣿⣿⣇⠟⠈⣿⡽⣷⣠⣶⠏⠀⠀⠀⠀⠀⢹⡄⢿⢸⣿⣿⣿⠀⠀⠀⠀⠀⠀⢸⡀⠀⠀⢰⠇⠀⢠⢷⠀⠀⣿⣿⣿⣿⣿⢸⠉⠉⠓⠶⠦⠤")
		fmt.Println("⡿⣸⣿⠇⣿⣿⣿⣿⣷⣬⣿⣿⣿⠏⠀⠀⠀⠀⠀⠀⠀⠈⢻⣜⣿⣿⣿⣿⡆⠀⠀⠀⠀⠀⠈⡇⠀⠀⣿⠀⠀⠀⢸⠀⢀⣿⣿⣿⣿⣿⢸⠀⠀⠀⠀⠀⠀")
		fmt.Println("⡇⣿⡏⡆⢻⣿⣿⣿⣿⣿⠋⢸⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⣿⠈⣿⣿⣿⠁⠀⠀⠀⠀⠀⣴⡇⠀⢸⡇⠀⠀⠀⢸⣦⢛⡏⣿⣿⣿⣿⢸⠀⠀⠀⠀⠀⠀")
		fmt.Println("⡇⣿⡇⡇⠸⣿⣿⣿⣿⡟⢠⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡟⣧⣻⣿⡿⠀⠀⠀⠀⠀⡴⠉⣧⠀⡼⠀⠀⢀⣰⢾⠃⣼⠁⣿⣿⣿⡿⡟⠀⠀⠀⠀⠀⠀")
		fmt.Println("⡇⣿⡀⡇⠀⣿⣿⣿⣿⠁⣾⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⠃⠙⣿⣿⡇⠀⠀⠀⣠⠞⠁⢀⣿⢰⡇⠐⠛⠋⠁⢸⣦⠃⢠⣿⣿⣿⣡⠻⣄⠀⠀⠀⠀⠀")
		fmt.Println("⢻⣿⡇⠹⠀⠀⢻⣿⣿⣾⠃⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⡾⠃⠀⢀⣿⡟⠀⠀⣠⠞⠁⢀⡴⠋⣿⣼⠁⠀⠀⣀⣀⣸⠃⠀⢸⣿⣿⣷⠋⠀⠈⠙⠒⠂⠀⠀")
											   

	} 

	
	
}
}