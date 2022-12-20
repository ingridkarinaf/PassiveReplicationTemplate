package main

import (
	interface "github.com/ingridkarinaf/PassiveReplicationTemplate/interface"
	grpc "google.golang.org/grpc"
	"log"
)

/*
 - Updates and retrieved values from the server
*/

func main() {

	//Creating log file
	f := setLogClient()
	defer f.Close()

	//Establishes connection with primary server 
	FEport := ":" + os.Args[1] 
	conn, err := grpc.Dial(FEport, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}
	connection = conn

	server = hashtable.NewHashTableClient(connection) //creates a connection with an FE server
	defer connection.Close()

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		
		for {
			println("Enter command")
			scanner.Scan()
			textChoice := scanner.Text()
			if (textChoice == "update") {
				println("Enter key and value, separated by a space (integers only!)")
				scanner.Scan()
				text := scanner.Text()
				inputArray := strings.Fields(text)
				
				key, err := strconv.Atoi(inputArray[0])
				if err != nil {
					log.Printf("Couldn't convert key to int: ", err)
					continue
				}
				value, err := strconv.Atoi(inputArray[1])
				if err != nil {
					log.Printf("Couldn't convert key to int: ", err)
					continue
				}
	
				hashtableUpdate := &hashtable.PutRequest{
					Key:   int32(key),
					Value: int32(value),
				}

				result := Put(hashtableUpdate)
				if result.Success == true {
					log.Printf("Hashtable successfully updated to %v for key %v.\n", key, value)
				} else {
					log.Println("Update unsuccessful, please try again.")
				}
				
			} else if (textChoice == "get") {
				println("Enter the key of the value you would like to retireve (integers only!): ")
				scanner.Scan()
				text := scanner.Text()
				key, err := strconv.Atoi(text)
				if err != nil {
					log.Println("Client: Could not convert key to integer: ", err)
					continue
				}
				getReq := &hashtable.GetRequest{
					Key:   int32(key),
				}

				result := Get(getReq) 
				log.Printf("Client: Value of key %s: %v \n",text, int(result))
				fmt.Printf("Value of key %s: %v \n",text, int(result))
			} else {
				log.Println("Sorry, didn't catch that. ")
				fmt.Println("Sorry, didn't catch that. ")
			}
		}
	}()

	for {}
}


func setLogClient() *os.File {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)
	return f
}