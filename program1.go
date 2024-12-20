/*


Will Hooker
10/7/2024
CS424-01 Programming Languages
Programming Assignment 1
program1.go
Tested on go version 1.23.2 on a Windows PC

Purpose:
This program reads baseball player statistics from an input file and stores them in a list of player objects.
Using the stats of each player, their batting average, slugging percentage, and on base percentage is then calculated.
Once the end of the file has been reached, the program sorts the players ordered by the slugging percentage in descending order.
If there were any invalid lines in the input file, error messages are printed out following all of the valid player data.

FORMAT OF INPUT FILE:
firstname lastname plateappearances atbats singles doubles triples homeruns walks hitbypitch

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Struct for the player
type Player struct {
	firstName, lastName                                            string
	plateAppearances                                               int
	atBats, singles, doubles, triples, homeruns, walks, hitByPitch int
	battingAvg, slugging, onBase                                   float64
}

// Function to calculate the batting average, slugging percentage, and on base percentage for each player. The results are stored in the player object
func (p *Player) calculateStats() {
	p.battingAvg = float64(p.singles+p.doubles+p.triples+p.homeruns) / float64(p.atBats)

	p.slugging = float64(p.singles+2*p.doubles+3*p.triples+4*p.homeruns) / float64(p.atBats)

	p.onBase = float64(p.singles+p.doubles+p.triples+p.homeruns+p.walks+p.hitByPitch) / float64(p.plateAppearances)
}

/*
Function to read each player from the input file declared by the user. If the file doesn't exist or if there is an error opening the file, an error message is printed
and the program quits. Each player is read in from the file, stored in a player object, and then appended to a list of players. Any errors encountered will be stored in a list of errors.
Two counter variables keep track of which line of the file the function is on in case of an error, and also the amount of players successfully read in. Once all data is read in and
the calculations from the calculateStats function are returned, the list of players, errors, and total players counted are returned to main.
*/
func readPlayers(fileName string) ([]Player, []string, int) {

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil, 0
	}
	defer file.Close()

	// Initializing the list for players and error messages
	// Also creating a count and playerCount variable to keep track the current line, and the amount of players successfully read
	var players []Player
	var errors []string
	var count, playerCount int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()
		fields := strings.Fields(line)
		count += 1

		player := Player{
			firstName: fields[0],
			lastName:  fields[1],
		}

		if len(fields) != 10 {
			errors = append(errors, fmt.Sprintf("\tline\t%d:\t%s, %s\t: Line contains not enough data.", count, player.lastName, player.firstName))
			continue
		}

		//Converting the data read in as strings to integers and storing in the player object. If an error exists, it is stored to the list of erros and the next player is read

		if plateAppearances, err := strconv.Atoi(fields[2]); err == nil {
			player.plateAppearances = plateAppearances
		} else {
			errors = append(errors, fmt.Sprintf("\tline\t%d:\t%s, %s\t: Line contains invalid data.", count, player.lastName, player.firstName))
			continue
		}

		if atBats, err := strconv.Atoi(fields[3]); err == nil {
			player.atBats = atBats
		} else {
			errors = append(errors, fmt.Sprintf("\tline\t%d:\t%s, %s\t: Line contains invalid data.", count, player.lastName, player.firstName))
			continue
		}

		if singles, err := strconv.Atoi(fields[4]); err == nil {
			player.singles = singles
		} else {
			errors = append(errors, fmt.Sprintf("\tline\t%d:\t%s, %s\t: Line contains invalid data.", count, player.lastName, player.firstName))
			continue
		}

		if doubles, err := strconv.Atoi(fields[5]); err == nil {
			player.doubles = doubles
		} else {
			errors = append(errors, fmt.Sprintf("\tline\t%d:\t%s, %s\t: Line contains invalid data.", count, player.lastName, player.firstName))
			continue
		}

		if triples, err := strconv.Atoi(fields[6]); err == nil {
			player.triples = triples
		} else {
			errors = append(errors, fmt.Sprintf("\tline\t%d:\t%s, %s\t: Line contains invalid data.", count, player.lastName, player.firstName))
			continue
		}

		if homeruns, err := strconv.Atoi(fields[7]); err == nil {
			player.homeruns = homeruns
		} else {
			errors = append(errors, fmt.Sprintf("\tline\t%d:\t%s, %s\t: Line contains invalid data.", count, player.lastName, player.firstName))
			continue
		}

		if walks, err := strconv.Atoi(fields[8]); err == nil {
			player.walks = walks
		} else {
			errors = append(errors, fmt.Sprintf("\tline\t%d:\t%s, %s\t: Line contains invalid data.", count, player.lastName, player.firstName))
			continue
		}

		if hitByPitch, err := strconv.Atoi(fields[9]); err == nil {
			player.hitByPitch = hitByPitch
		} else {
			errors = append(errors, fmt.Sprintf("\tline\t%d:\t%s, %s\t: Line contains invalid data.", count, player.lastName, player.firstName))
			continue
		}

		player.calculateStats()
		players = append(players, player)
		playerCount += 1
	}
	return players, errors, playerCount
}

/*
   This function drives the program, starting out with a welcome message. It then prompts for and reads the name of the input file.
   The program then calls the readPlayers function to read in each player from the input file and stores the data in a list of Player objects.
   If the filename cannot be found, the program prints an error message and exits. Any errors encountered will be stored in a separate list. The function also keeps
   track of the amount of players successfully processed. Once finished, the lists and total count are returned. Once the players are sorted by slugging percentage,
   the results are printed out to the console.
*/

func main() {
	fmt.Println("Welcome to the player statistics calculator test program. I am going to\nread players from an input data file. You will tell me the name of\nyour input file. I will store all of the players in a list,\ncompute each player's averages and then write the resulting team report to\nthe screen.")

	var fileName string
	fmt.Print("\n\nEnter the name of your input file: ")
	fmt.Scan(&fileName)

	players, errors, playerCount := readPlayers(fileName)
	if players == nil {
		return
	}

	sort.Slice(players, func(i, j int) bool {
		return players[i].slugging > players[j].slugging
	})

	fmt.Printf("\nBASEBALL STATS REPORT --- %d PLAYERS FOUND IN FILE\n\n", playerCount)
	fmt.Println("	PLAYER NAME	:	AVERAGE	SLUGGING ONBASE%")
	fmt.Println("--------------------------------------------------------------")
	for _, player := range players {
		fmt.Printf("\t%s %s \t:\t %.3f \t %.3f \t %.3f\n", player.firstName, player.lastName, player.battingAvg, player.slugging, player.onBase)
	}

	if len(errors) > 0 {
		fmt.Println("\nERROR LINES FOUND IN INPUT DATA")
		fmt.Println("--------------------------------------------------------------")
		for _, error := range errors {
			fmt.Println(error)
		}
	}
}
