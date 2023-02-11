package main
// Created By bytehx aka Chan Nyein Wai
import (
	"fmt"
	"net/url"
	"strings"
	"bufio"
	"os"
)


func main() {

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() { 
		
		u, err := url.Parse(sc.Text())
	if err != nil {
		fmt.Println(err)
		return
	}

	queryParams := u.Query()
	for key, values := range queryParams {
		for _, value := range values {
			fmt.Println(strings.Replace(u.String(), u.RawQuery, key+"="+value, 1))
		}
	}

	} 

	
}
