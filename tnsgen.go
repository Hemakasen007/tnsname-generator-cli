package main

import "fmt"

func main() {
	var host, connection, port, service string
	// var connection string
	// var port string
	// var service string
	fmt.Print("Enter the Connection name: ")
	fmt.Scanln(&connection)
	fmt.Print("Enter the Hostname: ")
	fmt.Scanln(&host)
	fmt.Print("Enter the Port: ")
	fmt.Scanln(&port)
	fmt.Print("Enter the Service: ")
	fmt.Scanln(&service)
	fmt.Println("	...	")
	fmt.Println("	")
	fmt.Printf(` %s =
	(DESCRIPTION =
	  (ADDRESS_LIST =
		(ADDRESS = (PROTOCOL = TCP)(HOST = %s)(PORT = %s))
	  )
	  (CONNECT_DATA =
		(SERVICE_NAME = %s)
	  )
	)`, connection, host, port, service)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Press Enter to Continue.....")
	fmt.Scanln()

}
