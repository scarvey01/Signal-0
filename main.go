package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var player Character

type Choice struct {
	Text   string
	Action func()
}

func typewriter(text string, delay time.Duration) {
	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(delay)
	}
	fmt.Println()
}

type Character struct {
	Name   string
	Skills Skills
}

type Skills struct {
	Scavenging int
	Medicine   int
	Combat     int
	Barter     int
	Crafting   int
}

func printAsciiArt(lines []string, delay time.Duration) {
	for _, line := range lines {
		fmt.Println(line)
		time.Sleep(delay)
	}
}
func presentChoices(prompt string, choices []Choice) {
	fmt.Println("\n" + prompt)
	for i, choice := range choices {
		fmt.Printf("%d. %s\n", i+1, choice.Text)
	}

	var input int
	for {
		fmt.Print("Choose an option: ")
		_, err := fmt.Scanln(&input)
		if err == nil && input >= 1 && input <= len(choices) {
			clearScreen()
			choices[input-1].Action()
			break
		}
		fmt.Println("Invalid input. Try again.")
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
	fmt.Print("\033[H\033[2J")
	fmt.Println("Welcome To")
	time.Sleep(1 * time.Second)
	typewriter("........", 200*time.Millisecond)
	time.Sleep(2 * time.Second)
	printAsciiArt(strings.Split(title, "\n"), 150*time.Millisecond)
	time.Sleep(2 * time.Second)
	var answer string

	for {
		fmt.Println("\nBEGIN? [Y/N]")
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
	var answer string
	fmt.Print("\033[H\033[2J")
	typewriter("The date is June 17th, 2054.", 50*time.Millisecond)
	time.Sleep(750 * time.Millisecond)
	typewriter("At least that's what we think.", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)
	typewriter("If that's true, it's the 15-year anniversary of all electronics going out.", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)
	typewriter("Feels like it's been a lot longer...", 50*time.Millisecond)
	time.Sleep(900 * time.Millisecond)
	typewriter("The last time I ate must have been 2 days ago when Ron ordered all the pets be shot.", 50*time.Millisecond)
	time.Sleep(1000 * time.Millisecond)
	typewriter("....", 700*time.Millisecond)
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
	typewriter("My brother says he knows a guy who will help us steal a working car with some gas.", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)
	typewriter("If that's true, we can both leave.", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)
	typewriter("Hopefully, we can find some kind of military camp.", 50*time.Millisecond)
	time.Sleep(750 * time.Millisecond)
	typewriter("If thats even a thing anymore...", 50*time.Millisecond)
	time.Sleep(900 * time.Millisecond)
	typewriter("But I'll do anything to leave this hellhole.", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)
	time.Sleep(1000 * time.Millisecond)
	for {
		fmt.Println("\nCONTINUE? [Y/N]")
		fmt.Scanln(&answer)
		answer = strings.TrimSpace(answer)

		if strings.EqualFold(answer, "y") {
			time.Sleep(500 * time.Microsecond)
			char_creator()
			break
		} else if strings.EqualFold(answer, "n") {
			fmt.Println("Exiting...")
			break
		} else {
			fmt.Println("Invalid input. Please type Y or N.")
		}
	}

}

func clearScreen() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func char_creator() {
	clearScreen()
	skills := Skills{
		Scavenging: 1,
		Medicine:   1,
		Combat:     1,
		Barter:     1,
		Crafting:   1,
	}

	skillMap := map[int]string{
		1: "Scavenging",
		2: "Medicine",
		3: "Combat",
		4: "Barter",
		5: "Crafting",
	}

	descriptions := map[int]string{
		1: "Scavenging helps you find better loot when exploring ruins.",
		2: "Medicine increases your ability to heal and treat injuries.",
		3: "Combat improves your fighting ability and accuracy.",
		4: "Barter gives you better trade deals and influences people.",
		5: "Crafting allows you to craft more items and makes crafted items more durable",
	}

	remainingPoints := 5
	reader := bufio.NewReader(os.Stdin)

	for remainingPoints > 0 {
		fmt.Println("\033[H\033[2J")
		fmt.Println("\n=== SKILL ALLOCATION ===")
		fmt.Printf("POINTS REMAINING: %d \n", remainingPoints)
		for i := 1; i <= 5; i++ {
			fmt.Printf("%d: %s (Level: %d)\n", i, skillMap[i], getSkillLevel(skills, i))
		}

		fmt.Print("Choose a skill to view and add points (1–5): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > 5 {
			fmt.Println("Invalid choice. Try again.")
			continue
		}

		fmt.Printf("\n%s: %s\n", skillMap[choice], descriptions[choice])
		fmt.Print("How many points would you like to add? (0 to cancel): ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		pointsToAdd, err := strconv.Atoi(input)
		if err != nil || pointsToAdd < 0 {
			fmt.Println("Invalid number. Please enter a valid amount.")
			continue
		}

		currentLevel := getSkillLevel(skills, choice)
		if currentLevel+pointsToAdd > 5 {
			fmt.Println("Skill level cannot go above 5.")
			continue
		}
		if pointsToAdd > remainingPoints {
			fmt.Println("You don't have enough points.")
			continue
		}

		skills = addSkillPoints(skills, choice, pointsToAdd)
		remainingPoints -= pointsToAdd
	}

	// Confirm final allocation
	clearScreen()
	fmt.Println("\nFinal Skill Levels:")
	fmt.Printf("Scavenging: %d\n", skills.Scavenging)
	fmt.Printf("Medicine:   %d\n", skills.Medicine)
	fmt.Printf("Combat:     %d\n", skills.Combat)
	fmt.Printf("Barter:     %d\n", skills.Barter)
	fmt.Printf("Crafting:   %d\n", skills.Crafting)

	for {
		fmt.Print("Are you sure about your choices? THIS CAN NOT BE CHANGED [Y/N]: ")
		confirm, _ := reader.ReadString('\n')
		confirm = strings.ToLower(strings.TrimSpace(confirm))
		if confirm == "y" {
			fmt.Println("Skills locked in.")
			time.Sleep(500 * time.Millisecond)

			// Get the name
			player.Name = name_char()

			// Assign skills to player
			player.Skills = skills

			fmt.Println("Hello,", player.Name)
			time.Sleep(2000 * time.Millisecond)
			chap1_1()

			return
		} else if confirm == "n" {
			fmt.Println("Restarting allocation...")
			time.Sleep(500 * time.Millisecond)
			char_creator() // Restart allocation
			return
		} else {
			fmt.Println("Please type 'y' or 'n'.")
		}
	}
}

func getSkillLevel(s Skills, skillID int) int {
	switch skillID {
	case 1:
		return s.Scavenging
	case 2:
		return s.Medicine
	case 3:
		return s.Combat
	case 4:
		return s.Barter
	case 5:
		return s.Crafting
	}
	return 0
}

func addSkillPoints(s Skills, skillID int, points int) Skills {
	switch skillID {
	case 1:
		s.Scavenging += points
	case 2:
		s.Medicine += points
	case 3:
		s.Combat += points
	case 4:
		s.Barter += points
	case 5:
		s.Crafting += points
	}
	return s
}
func GainSkillPoint(skills *Skills, skillName string) {
	switch skillName {
	case "Scavenging":
		if skills.Scavenging < 5 {
			skills.Scavenging++
			fmt.Println("Scavenging +1 (New level:", skills.Scavenging, ")")
		}
	case "Medicine":
		if skills.Medicine < 5 {
			skills.Medicine++
			fmt.Println("Medicine +1 (New level:", skills.Medicine, ")")
		}
	case "Combat":
		if skills.Combat < 5 {
			skills.Combat++
			fmt.Println("Combat +1 (New level:", skills.Combat, ")")
		}
	case "Barter":
		if skills.Barter < 5 {
			skills.Barter++
			fmt.Println("Barter +1 (New level:", skills.Barter, ")")
		}
	case "Crafting":
		if skills.Crafting < 5 {
			skills.Crafting++
			fmt.Println("Crafting +1 (New level:", skills.Crafting, ")")
		}
	}
}

func name_char() string {
	reader := bufio.NewReader(os.Stdin)
	var name string
	for {
		clearScreen()
		fmt.Println("====== NAME YOUR CHARACTER ======")
		fmt.Print("Enter your character's name: ")
		name, _ = reader.ReadString('\n')
		name = strings.TrimSpace(name)
		if name != "" {
			break
		}
		fmt.Println("Name cannot be blank.")
	}
	return name
}

//CHAPTER 1

func chap1_1() {
	chap1 := `
╭━━━┳╮╱╭┳━━━┳━━━┳━━━━┳━━━┳━━━╮╱╭╮
┃╭━╮┃┃╱┃┃╭━╮┃╭━╮┃╭╮╭╮┃╭━━┫╭━╮┃╭╯┃
┃┃╱╰┫╰━╯┃┃╱┃┃╰━╯┣╯┃┃╰┫╰━━┫╰━╯┃╰╮┃
┃┃╱╭┫╭━╮┃╰━╯┃╭━━╯╱┃┃╱┃╭━━┫╭╮╭╯╱┃┃
┃╰━╯┃┃╱┃┃╭━╮┃┃╱╱╱╱┃┃╱┃╰━━┫┃┃╰╮╭╯╰╮
╰━━━┻╯╱╰┻╯╱╰┻╯╱╱╱╱╰╯╱╰━━━┻╯╰━╯╰━━╯
`

	printAsciiArt(strings.Split(chap1, "\n"), 150*time.Millisecond)
	time.Sleep(5 * time.Second)
	clearScreen()
	print("Silas: ")
	typewriter("He said he would meet us here.", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)

	print("Silas: ")
	typewriter("It's been like an hour.", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)

	typewriter("...", 2*time.Second)

	print("Silas: ")
	typewriter("Let's just leave.", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)

	println("A short, stocky man covered in dirt and grime walks in front of you and your brother.")

	print("?: ")
	typewriter("Hey Silas, what's goin' on, man?", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)

	print("Silas: ")
	typewriter("I thought we were going to meet an hou—", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)

	print("?: ")
	typewriter("Yo, who's this guy?", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)

	print("Silas: ")
	typewriter("He's my brother. I told you he was coming.", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)

	print("Mason: ")
	typewriter("Yo, what's up, lil' man? Name's Mason.", 50*time.Millisecond)
	time.Sleep(1500 * time.Millisecond)

	clearScreen()
	presentChoices("The man extends his hand, covered in... well, you're not quite sure what, for a handshake.", []Choice{
		{
			Text: "Shake his hand",
			Action: func() {
				chap1_2()
			},
		},
		{
			Text: "Just tell him your name",
			Action: func() {
				chap1_3()
			},
		},
	})
}

func chap1_2() {
	clearScreen()
	println("You shake his hand, and whatever was on it has now rubbed onto yours.")
	print(player.Name + ": ")
	typewriter("I'm "+player.Name+".", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)

	print("Mason: ")
	typewriter("Nice meetin' you, man.", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)

	chap1_4()
}
func chap1_3() {
	print(player.Name + ": ")
	typewriter("I'm "+player.Name+".", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)

	println("Mason looks a little disappointed and embarrassed.")
	time.Sleep(500 * time.Millisecond)

	chap1_4()
}

func chap1_4() {
	print("Silas: ")
	typewriter("So, what's the plan?", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)

	print("Mason: ")
	typewriter("I stole a key to the car garage. If we can get past the guard, we're in.", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)

	print("Silas: ")
	typewriter("So, how do you recommend we do that?", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)

	println("Mason pulls a small revolver from his waistband.")
	time.Sleep(500 * time.Millisecond)

	print("Mason: ")
	typewriter("Got this about two years ago.", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)

	print("Silas: ")
	typewriter("We need to do this quietly. If you shoot that gun, we're not getting out.", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)

	print("Mason: ")
	typewriter("Well, if you can find another way, let me know.", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)

	print("Silas: ")
	typewriter("Fine. Let's go.", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)

	print("Mason: ")
	typewriter("Hold on, man. You remember the deal, right? I'm coming with you guys.", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)

	print("Silas: ")
	typewriter("Yeah, that's fine. Let's go.", 50*time.Millisecond)
	time.Sleep(500 * time.Millisecond)

	var answer string
	for {
		fmt.Println("\nCONTINUE? [Y/N]")
		fmt.Scanln(&answer)
		answer = strings.TrimSpace(answer)

		if strings.EqualFold(answer, "y") {
			time.Sleep(3 * time.Second)
			chap1_5()
			break
		} else if strings.EqualFold(answer, "n") {
			fmt.Println("Exiting...")
			break
		} else {
			fmt.Println("Invalid input. Please type Y or N.")
		}
	}
}
func chap1_5() {
	clearScreen()

}
