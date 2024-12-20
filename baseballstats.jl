#=
File: baseballstats.jl
Will Hooker
Fall 2024
Ran on Windows 11

Purpose:
This program reads baseball player statistics from an input file and stores them in a matrix.
Using the stats of each player, their batting average, slugging percentage, and on base percentage is calculated.
Once all stats have been calculated, the program sorts the players ordered by their slugging percentage in descending order and
in a separate list sorted by their last name in ascneding order. The results are then printed out to the console.


Input file format:
firstname lastname plateappearances atbats singles doubles triples homeruns walks hitbypitch

=#
using Printf
using Statistics

############################
# Function to return the conversion of text to Int64
############################
function getInt(a)
    return parse(Int64,a)
end

############################
# Function to calculate the batting average, slugging percentage, and on base percentage
############################
function calculateStats(plateAppearances, atBats, singles, doubles, triples, homeruns, walks, hitByPitch)
    battingAvg = (singles + doubles + triples + homeruns) / atBats
    slugging = (singles + 2*doubles + 3*triples + 4*homeruns) / atBats
    onBase = (singles + doubles + triples + homeruns + walks + hitByPitch) / plateAppearances
    return battingAvg, slugging, onBase
end

############################
# Function to print the player data of a given list
############################
function printPlayers(players)
    println("	PLAYER NAME	:	AVERAGE	SLUGGING ONBASE%")
    println("--------------------------------------------------------------")
    for player in players
        @printf("\t%s, %s \t:\t %.3f \t %.3f \t %.3f\n", player[2], player[1], player[3], player[4], player[5])
    end
end


############################
# Main function:
# open a file and read lines containing baseball player stats
# store the player data in a matrix
# calls function to calculate the batting avg, slugging%, and OBP
# completes the necessary sorts
# calls a function to print the results
############################

print("\nEnter the name of the input file: ")
myfile = nothing
filename = readline()

try
    global myfile = open(filename)

catch err
    println("\nUnable to open the file: $filename")
    println("Exiting the program\n")
    exit(0)
end

# If the file was opened, process each player and store

# Initializing an empty array to contain all the player data, and a counter for each player read
players = []
playerCount = 0

for line in eachline(myfile)
    # split each line into separate components
    data = split(line)        

    # Separate the first and last name, then convert the rest of line to int
    firstname, lastname = data[1], data[2]
    stats = map(getInt, data[3:end])
    
    # Array for each player
    player = [firstname, lastname, stats...]

    # Push the player data to the list containing every player
    push!(players, player)
    global playerCount += 1
end

# Array to store the calculated player stats
playerStats = []

for player in players
    firstname, lastname = player[1], player[2]
    plateAppearances, atBats, singles, doubles, triples, homeruns, walks, hitByPitch = player[3:end]

    # Calling function to calculate batting average, slugging percentage, and on base percentage
    battingAvg, slugging, onBase = calculateStats(plateAppearances, atBats, singles, doubles, triples, homeruns, walks, hitByPitch)

    # Storing the names and resulting stats in a list, and pushing it into the list containing every player
    results = [firstname, lastname, battingAvg, slugging, onBase]
    push!(playerStats, results)
end

# Sorting the list by slugging percentage, and then reversing the list to get it in decreasing order
sortedSlugging = sort(playerStats, by=x -> x[4])
sortedSlugging = reverse(sortedSlugging)

# Sorting the list by last name, and then first name in case of a tie
sortedName = sort(playerStats, by=x -> (x[2], x[1]))

@printf("\nBASEBALL STATS REPORT --- %d PLAYERS FOUND IN FILE\n", playerCount)
println("\nREPORT ORDERED BY SLUGGING %\n\n")
printPlayers(sortedSlugging)
println("\n\nREPORT ORDERED BY LAST NAME\n\n")
printPlayers(sortedName)
