/*
* Http (curl) request in golang
* @author Shashank Tiwari
*/
 
package main
 
import (
	"fmt"
	"net/http"
	"io/ioutil"
)
 
func main() {
 
	url := "https://google.com"
	
	req, _ := http.NewRequest("GET", url, nil)
	
	res, _ := http.DefaultClient.Do(req)
	
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	
	fmt.Println(string(body))
 
}