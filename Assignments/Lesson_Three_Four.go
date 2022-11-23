package main
 
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)
 
// Check if string extists in a given array
func exists(s []string, str string) bool {
    for _, v := range s {
        if strings.ToLower(v) == strings.ToLower(str) {
            return true
        }
    }
    return false
}
 
// Remove a string from a given array
func remove(s []string, r string) []string {
    for i, v := range s {
        if strings.ToLower(v) == strings.ToLower(r) {
            return append(s[:i], s[i+1:]...)
        }
    }
    return s
}
 
func main() {
    // Skapa en variabel travelbag som är en array
    // Gör en program med följande fungerande meny:
    // Lägg till saker i en travelbag
    // Ta bort saker från travelbag
    // Kolla i travelbag (skriv ut innehållet)
    // Avsluta programmet
 
    programRunning := true
    travelBag := []string{"Water", "Pen", "Apple", "Book"}
    options := [4]string{"What's in my travelbag", "Add Object", "Remove Object", "Exit"}
 
    for programRunning{
        scanner := bufio.NewScanner(os.Stdin)
        fmt.Println("\n\n\n\nEPIC TRAVELBAG OPTIONS")
        // Show the options with their index
        for index, options := range options{
            fmt.Println(index, options)
        }
 
        fmt.Print("Pick an option: ")
        scanner.Scan()
        optionNum, _ := strconv.ParseInt(scanner.Text(), 10, 64)
 
        switch optionNum{
        // Check what's in the bag.
        // Index is not nessecerry unless the user removes later on the object based on the index.
        case 0:
            fmt.Println("\n \n In your bag you have:")
            for index, travelBag := range travelBag{
                fmt.Println(index, travelBag)
            }
            fmt.Print("Press any key to go back")
            scanner.Scan()
        // Append something in the bag.
        case 1:
            fmt.Print("\n \nWhat would you like to add in your bag?: ")
            scanner.Scan()
            add := scanner.Text()
            travelBag = append(travelBag, add)
        // Remove something from the bag.
        case 2:
            fmt.Print("Type the name of the object you want removed: ")
            scanner.Scan()
            removedItem := scanner.Text()
            // Check if obects exists in the bag.
            if exists(travelBag, removedItem){
                travelBag = remove(travelBag, removedItem)
                fmt.Println("You just removed ", removedItem, "from your travelbag")
            }else{
                fmt.Println("It does not exist")
            }
        // Close the loop and exit the program.
        case 3:
            fmt.Println("\nTime to go!\nhGOODBYE!")
            programRunning = false
        // For default just ask again for input
        default:
            fmt.Println("\n\nINVALID OPTION PLEASE TRY AGAIN")
        }
 
    }
 
}
