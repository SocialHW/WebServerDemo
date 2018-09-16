/* Barebones web server which serves up a static HTML file
 * Based on the example given here:
 * https://www.alexedwards.net/blog/serving-static-sites-with-go
 *
 * This can be the base for a file upload service. 
 */

package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("."))
	http.Handle("/", fileServer)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)

}


