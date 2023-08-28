package main

import (
	"fmt"
	"net/http"

	"github.com/saktialfansyahp/go-rest-api/routes"
)

func main() {
	routes.DefineRoutes()

	serverAddr := ":8080"
	fmt.Printf("Server is listening on %s\n", serverAddr)
	go http.ListenAndServe(serverAddr, nil)

	// Keep the main goroutine running
	select {}
}

// func generateRandomKey(length int) string {
// 	key := make([]byte, length)
// 	_, err := rand.Read(key)
// 	if err != nil {
// 		panic("Failed to generate key")
// 	}
// 	return base64.URLEncoding.EncodeToString(key)
// }

// func main(){
// 	userID := 123 // Replace with an actual user ID
//     token, err := utils.GenerateToken(userID)
//     if err != nil {
//         fmt.Println("Error generating token:", err)
//         return
//     }

//     fmt.Println("Generated Token:", token)
// }