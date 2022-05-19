package main
import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    var choice string
    fmt.Println("Hey, welcome to Rock Paper Scissor game!")
    fmt.Println("Press p to play or q to quit")
    fmt.Scanln(&choice)
    i := 69
    if choice == "p" {
        var sel string
        fmt.Println("Great! Choose p for paper, r for rock and s for scissor or other keys to quit")
        for i > 14 {
            fmt.Scanln(&sel)
            rand.Seed(time.Now().UnixNano())
            randint := rand.Intn(4 - 1) + 1 // 1 = Paper, 2 = Scissor, 3 = Rock
            if sel == "p" {
                if randint == 1 {
                    fmt.Println("It's a tie, computer also chose paper!")
                } else if randint == 2 {
                    fmt.Println("Oh! you lost, computer chose scissor")
                } else {
                    fmt.Println("Congrats! You won, computer chose rock")
                }
            } else if sel == "r" {
                if randint == 3 {
                    fmt.Println("It's a tie, computer also chose rock!")
                } else if randint == 1 {
                    fmt.Println("Oh! you lost, computer chose paper")
                } else {
                    fmt.Println("Congrats! You won, computer chose scissor")
                }
            } else if sel == "s" {
                if randint == 2 {
                    fmt.Println("It's a tie, computer also chose scissor!")
                } else if randint == 3 {
                    fmt.Println("Oh! you lost, computer chose rock")
                } else {
                    fmt.Println("Congrats! You won, computer chose paper")
                }
            } else {
                fmt.Println("Thanks for using our service! will see you soon :)")
                break
            }
        }
    } else if choice == "q" {
        fmt.Println("Thanks for using our service, hope to see you soon :)")
    } else {
        fmt.Println("Wrong choice :) try again!")
    }
}
