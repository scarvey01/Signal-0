package main

import (
	"fmt"
	"strings"
	"time"
)

func typewriter(text string, delay time.Duration) {
	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(delay)
	}
	fmt.Println()
}

func printAsciiArt(lines []string, delay time.Duration) {
	for _, line := range lines {
		fmt.Println(line)
		time.Sleep(delay)
	}
}
func main() {
	title := `

░██████╗██╗░██████╗░███╗░░██╗░█████╗░██╗░░░░░░░░█████╗░
██╔════╝██║██╔════╝░████╗░██║██╔══██╗██║░░░░░░░██╔══██╗
╚█████╗░██║██║░░██╗░██╔██╗██║███████║██║░░░░░░░██║░░██║
░╚═══██╗██║██║░░╚██╗██║╚████║██╔══██║██║░░░░░░░██║░░██║
██████╔╝██║╚██████╔╝██║░╚███║██║░░██║███████╗░░╚█████╔╝
╚═════╝░╚═╝░╚═════╝░╚═╝░░╚══╝╚═╝░░╚═╝╚══════╝░░░░╚════╝                                                                                            
`
	fmt.Println("Welcome To")
	time.Sleep(1 * time.Second)
	typewriter("........", 200*time.Millisecond)
	time.Sleep(2 * time.Second)
	printAsciiArt(strings.Split(title, "\n"), 150*time.Millisecond)
	time.Sleep(2 * time.Second)
	var answer string

	for {
		fmt.Println("\nBegin? [Y/N]")
		fmt.Scanln(&answer)
		answer = strings.TrimSpace(answer)

		if strings.EqualFold(answer, "y") {
			time.Sleep(500 * time.Microsecond)
			begin()
			break
		} else if strings.EqualFold(answer, "n") {
			fmt.Println("Exiting...")
			break
		} else {
			fmt.Println("Invalid input. Please type Y or N.")
		}
	}
}

func begin() {
	fmt.Print("\033[H\033[2J")
	typewriter("The date is June 17th, 2034.", 50*time.Millisecond)
	time.Sleep(750 * time.Millisecond)
	typewriter("At least that's what we think.", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)
	typewriter("If that's true, it's the 2-year anniversary of all electronics going out.", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)
	typewriter("Feels like it's been a lot longer...", 50*time.Millisecond)
	time.Sleep(900 * time.Millisecond)
	typewriter("Feels like it's been a lot longer.", 50*time.Millisecond)
	time.Sleep(700 * time.Millisecond)
	typewriter("The last time I ate must have been 2 days ago when Ron ordered all the pets be shot.", 50*time.Millisecond)
	time.Sleep(1000 * time.Millisecond)
	println()
	typewriter("But that food went quick...", 50*time.Millisecond)
	time.Sleep(1000 * time.Millisecond)
	typewriter("I'm not sure what disease has hit the crops this year, but it's bad.", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)
	typewriter("I don't think any of the crops made it.", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)
	typewriter("They won't let us leave because of the war with the settlement in Skyesville.", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)
	typewriter("But I'm sick of this.", 50*time.Millisecond)
	time.Sleep(750 * time.Millisecond)
	typewriter("My brother says he knows a guy who will help us steal a horse.", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)
	typewriter("If that's true, we can both leave.", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)
	typewriter("Hopefully, we can find some kind of military camp.", 50*time.Millisecond)
	time.Sleep(750 * time.Millisecond)
	typewriter("If thats even a thing anymore...", 50*time.Millisecond)
	time.Sleep(900 * time.Millisecond)
	typewriter("But I'll do anything to leave this hellhole.", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)

}
