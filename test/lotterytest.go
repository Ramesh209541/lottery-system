package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

func main() {

	fmt.Println("Lotter system")
	i := 5
	var hostname,portNumber string
	fmt.Println("enter server ip ")
	fmt.Scanln(&hostname)
	fmt.Println("enter port number")
	fmt.Scanln(&portNumber)
	ip := "http://"+ hostname +":"+ portNumber
	fmt.Println("ip is:",ip)
	for {
		fmt.Println("1.Create New Ticket")
		fmt.Println("2.Return list of tickets")
		fmt.Println("3.get individual ticket")
		fmt.Println("4.amend ticket lines")
		fmt.Println("5.Retrive status of ticket")
		var j int
		fmt.Println("enter input")
		_, err := fmt.Scanln(&j)

		if err != nil {
			fmt.Println("invalid input err:", err)
			clearScreen()
			continue
		}
		if j > i {
			fmt.Println("invalid input give proper input")
			clearScreen()
			continue
		}
		switch j {
		case 1:
			jsondata := `{"Quantity":1,"LinesPerTicket":3}`
			response := postrequest(ip+"/ticket", jsondata)
			fmt.Println("Response is:", response)
		case 2:
			fmt.Println("sending get request")
			getrequest(ip+"/ticket")
		case 3:
			fmt.Println("get the individual ticket")
			fmt.Println("enter a ticket id")
			_, err := fmt.Scanln(&j)
			if err != nil {
				fmt.Println("invalid input err:", err)
				clearScreen()
				continue
			}
			id := strconv.Itoa(j)

			getrequest(ip+"/ticket/" + id)
		case 4:
			fmt.Println("amend lines to ticket")
			fmt.Println("enter a ticket id")
			_, err := fmt.Scanln(&j)
			if err != nil {
				fmt.Println("invalid input err:", err)
				clearScreen()
				continue
			}
			id := strconv.Itoa(j)
			jsondata := `{"LinesCount":2}`

			response := putrequest(ip+"/ticket/"+id, jsondata)
			fmt.Println("Response is:", response)

		case 5:
			fmt.Println("status of ticket")
			fmt.Println("enter a ticket id")
			_, err := fmt.Scanln(&j)
			if err != nil {
				fmt.Println("invalid input err:", err)
				clearScreen()
				continue
			}
			id := strconv.Itoa(j)
			response := putrequest(ip+"/status/"+id, "")
			fmt.Println("Response is:", response)
		default:

		}

	}
}
func clearScreen() {
	cmd := exec.Command("bash", "-c", "clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func getrequest(link string) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println("creating request is failed :", err)
		return
	}
	defer resp.Body.Close()
	if resp.Status != "200 OK" {
		fmt.Println("response status is not 200 status:", resp.Status)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("reading post response body is failed:err:", err)
		return
	}
	fmt.Println("body is:", string(body))
}

func postrequest(url string, jsonData string) string {
	var jsonStr = []byte(jsonData)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func putrequest(url string, jsonData string) string {
	var jsonStr = []byte(jsonData)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}
