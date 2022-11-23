package main
 
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)
 
func main() {
 
    // LEKTION 1
 
    fmt.Println("Hello world")
 
    // Låt användaren skriva in sitt namn och födelseår.
    // Skriv ut ett välkomnande meddelande och berätta hur gammal användaren är
    scanner := bufio.NewScanner(os.Stdin)
 
    fmt.Print("What's your name?: ")
    scanner.Scan()
    name := scanner.Text()
    fmt.Print("How old are you?: ")
    scanner.Scan()
    age := scanner.Text()
    fmt.Println("Hello,", name,". You are ",age," years old. Have a very nice day!")
 
    // Låt användaren skriva in ett valfritt tal. Skriv sedan ut vad det talet blir i kvadrat:
    fmt.Print("Type a number and I will tell you it's square root: ")
    scanner.Scan()
    inputNum, _ := strconv.ParseInt(scanner.Text(), 10, 64)
    fmt.Println("The square root of", inputNum, "is: ", inputNum*inputNum)
 
    // Fråga användaren hur många personer de är som ska dela och hur många objekt (äpplen. pengar t.ex.) som ska delas på.
    // Programmet ska sedan skriva ut hur många t.ex. äpplen varje person får(utan att dela några äpplen) och säga hur många äpplen som är kvar.
    // Tips: Använd floor division (//) och modulus (%)
 
    fmt.Println("Let's give apples to people!")
    fmt.Print("How many people are there?: ")
    scanner.Scan()
    numPeople, _:= strconv.ParseInt(scanner.Text(), 10, 64)
    fmt.Print("How many apples are available?: ")
    scanner.Scan()
    numApples, _ := strconv.ParseInt(scanner.Text(), 10, 64)
    fmt.Println("There are ", numPeople, " people and ", numApples, " apples")
 
    fmt.Println("Each person can take ", numApples / numPeople, " apples")
   
   
    // LEKTION 2
    // Gör ett frågesport
 
    fmt.Println("Hello Lets play a game!")
    fmt.Print("How many days are there in a week?: ")
    scanner.Scan()
    qOne, _ := strconv.ParseInt(scanner.Text(), 10, 64)
 
    if qOne == 7{
        fmt.Println(qOne, " is the correct answer!")
    }else{
        fmt.Println("Wrong answer!")
    }
   
    fmt.Print("What is the powerhouse of the cell?: ")
    scanner.Scan()
    cellAnswer := scanner.Text()
 
    switch cellAnswer{
    case "mitochondria":
        println("You are correct")
    case "mayonnaise":
        println("No mayonnaise is not the powerhouse of the cell....")
    default:
        fmt.Println("Wrong answer! ",cellAnswer," is not the powerhouse of the cell!")
    }
 
    fmt.Print("How old did you say you are?: ")
    scanner.Scan()
    verifyAge := scanner.Text()
 
    switch verifyAge{
    case age:
        fmt.Println("Correct answer, you do remember your age!")
    default:
        fmt.Println("No that was incorrect. You are ", age, " years old!")
 
    }
}
