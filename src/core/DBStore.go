package main

import (
	"errors"
	"fmt"
	log "logger"
	"math/rand"
	"net/http"
	"sort"
)

var countticket = 1

type lotteryticket struct {
	ticketID   int
	ticketdata []randomDataType
	status     bool
}
type randomDataType struct {
	ticketline []int
	Outcome    int
}

func getoutcome(data []int) int {

	if len(data) != 3 {
		return -1
	}
	var total int
	for _, val := range data {
		total += val
	}
	switch {
	case total == 2:
		return 10
	case data[0] == data[1] && data[1] == data[2]:
		return 5
	case data[0] != data[1] && data[0] != data[2]:
		return 1
	default:
		return 0
	}
}
func insertDatatoDB(quantity int, lines int) error {

	if database == nil {
		log.Error.Println("error database is nil")
		return errors.New("error database is nil")
	}
	statement, err := database.Prepare("INSERT INTO ticket (status) VALUES (?)")
	if err != nil {
		log.Error.Println("error preparing data base  is:", err)
		return err
	}
	ticketDataTable, err := database.Prepare("INSERT INTO ORDERS (ID2, value,outcome) VALUES (?, ?,?)")
	if err != nil {
		log.Error.Println("error preparing data base is:", err)
		return err
	}
	log.Info.Println("quantity of tickets going to create are :", quantity, "lines for each ticket are :", lines)
	for i := 0; i < quantity; i++ {
		_, err := statement.Exec(false)
		if err != nil {
			log.Error.Println("Exec is failed err :", err)
			return err
		}
	}
	for i := ticketCount; i < ticketCount+quantity; i++ {
		for j := 0; j < lines; j++ {
			for k := 0; k < 3; k++ {
				data := rand.Intn(3)
				ticketDataTable.Exec(i, data, -1)
			}
		}
	}
	ticketCount += quantity
	cmd := "UPDATE ticketcount SET count = ?"
	updateDate, err := database.Prepare(cmd)
	if err != nil {
		log.Error.Println("error preparing the table :", err)
		return err
	}
	_, err = updateDate.Exec(ticketCount)
	if err != nil {
		log.Error.Println("updateDate.Exec is failed")
		return err
	}
	log.Info.Println("total ticket count is:", ticketCount)
	return nil
}

func getListOfTicketsfromDB(w http.ResponseWriter) error {
	if database == nil {
		log.Error.Println("error database is nil")
		return errors.New("error database is nil")
	}

	ticketRows, err := database.Query("select C.id from ticket C ")
	if err != nil {
		log.Error.Println("error querrying", err)
		return err
	}
	var id int
	var listOftickets []int
	for ticketRows.Next() {
		ticketRows.Scan(&id)
		listOftickets = append(listOftickets, id)
	}
	fmt.Fprintf(w, "list of tickets in below line \n%+v\n", listOftickets)
	return nil
}

func getticketFromDB(id int, w http.ResponseWriter) error {
	if database == nil {
		log.Error.Println("error database is nil")
		return errors.New("error database is nil")
	}

	if id >= ticketCount {
		log.Info.Println("invalid id for getting ticket id is:", id)
		fmt.Fprintf(w, "invalid id")
		return errors.New("invalid id")
	}
	ticketRows, err := database.Query("select C.status from ticket C where C.id=?", id)
	if err != nil {
		log.Error.Println("error querrying", err)
		return err
	}
	var status bool
	for ticketRows.Next() {
		ticketRows.Scan(&status)
	}

	var ticketlines int
	var ticketid int
	var lineoutcome int
	var linevalue int
	var ticket lotteryticket
	Query := "select O.ID2 from ORDERS O where  O.ID2=?"
	rows, err := database.Query(Query, id)
	if err != nil {
		log.Error.Println("error querrying", err, "query is :", Query)
		return err
	}
	for rows.Next() {
		ticketlines++
	}
	log.Info.Println("ticket lines are:", ticketlines)
	if (ticketlines % 3) != 0 {
		log.Error.Println("invalid ticketlines from data base ")
		return errors.New("invalid ticketlines from data base ")
	}
	Query = "select O.ID2, O.value ,O.outcome from ORDERS O where O.ID2=?"
	rows1, err := database.Query(Query, id)
	if err != nil {
		log.Error.Println("error querrying", err, "query is:", Query)
		return err
	}
	ticket.ticketID = id
	ticket.ticketdata = make([]randomDataType, ticketlines/3, ticketlines/3)
	for i := 0; i < ticketlines/3; i++ {
		ticket.ticketdata[i].ticketline = make([]int, 3, 3)
	}
	l := 0
	for rows1.Next() {
		err := rows1.Scan(&ticketid, &linevalue, &lineoutcome)
		if err != nil {
			fmt.Println("there is a error in scanning:", err)
			return err
		}
		ticket.ticketdata[l/3].ticketline[l%3] = linevalue
		ticket.ticketdata[l/3].Outcome = lineoutcome
		l++
	}
	ticket.status = status
	fmt.Fprintf(w, "***the data for ticket is %+v***\n", ticket)
	return nil
}

func amendLinetoticketInDB(id int, lines int, w http.ResponseWriter) error {
	if database == nil {
		log.Error.Println("error database is nil")
		return errors.New("error database is nil")
	}
	if id >= ticketCount {
		log.Info.Println("invalid id is:", id)
		fmt.Fprintf(w, "invalid ticket id")
		return errors.New("invalid id")
	}
	Query := "select C.status from ticket C where C.id=?"
	ticketRows, err := database.Query(Query, id)
	if err != nil {
		log.Error.Println("error querrying", err, "query is:", Query)
		return err
	}
	var status bool
	for ticketRows.Next() {
		ticketRows.Scan(&status)
	}
	if status == true {
		fmt.Fprintf(w, "ticket cannot be amended ,status has been checked out")
		return nil
	}
	log.Info.Println(lines, " lines are getting amended to ticket id :", id)
	statement, err := database.Prepare("INSERT INTO ORDERS (ID2, value,outcome) VALUES (?, ?,?)")
	if err != nil {
		log.Error.Println("preparing data base is failed :err:", err)
		return err
	}
	for j := 0; j < lines; j++ {
		for k := 0; k < 3; k++ {
			rand := rand.Intn(3)
			statement.Exec(id, rand, -1)
		}
	}
	fmt.Fprintf(w, "amending %d lines to the ticket id %d is successful", lines, id)
	return nil
}

func statusCheckoutinDB(id int, w http.ResponseWriter) error {
	if database == nil {
		log.Error.Println("error database is nil")
		return errors.New("error database is nil")
	}

	if id >= ticketCount {
		log.Info.Println("invalid id is:", id)
		fmt.Fprintf(w, "invalid ticket id")
		return errors.New("invalid id")
	}
	var ticketlines int
	var ticketid int
	var lineoutcome int
	var linevalue int
	var ticket lotteryticket

	cmd := "UPDATE ticket  SET status = ? where id=?"
	updateTicket, err := database.Prepare(cmd)
	if err != nil {
		log.Error.Println("error preparing the table:", err, "query is:", cmd)
		return err
	}
	_, err = updateTicket.Exec(true, id)
	if err != nil {
		log.Error.Println("updateDate.Exec is failed err:", err)
		return err
	}

	rows, err := database.Query("select O.ID2 from ORDERS O where  O.ID2=?", id)
	if err != nil {
		log.Error.Println("first time error querrying", err)
		return err
	}
	for rows.Next() {
		ticketlines++
	}
	log.Info.Println(" no.of ticket lines are:", ticketlines, "for ticket id:", id)

	if (ticketlines % 3) != 0 {
		log.Error.Println("invalid ticketlines from data base ")
		return errors.New("invalid ticketlines from data base ")
	}
	rows, err = database.Query("select O.ID2, O.value ,O.outcome from ORDERS O where O.ID2=?", id)
	if err != nil {
		log.Error.Println("first time error querrying", err)
		return err
	}
	ticket.ticketID = id
	ticket.ticketdata = make([]randomDataType, ticketlines/3, ticketlines/3)
	for i := 0; i < ticketlines/3; i++ {
		ticket.ticketdata[i].ticketline = make([]int, 3, 3)
	}
	l := 0
	for rows.Next() {
		err := rows.Scan(&ticketid, &linevalue, &lineoutcome)
		if err != nil {
			log.Error.Println("there is a error in scanning:", err)
			return err
		}
		ticket.ticketdata[l/3].ticketline[l%3] = linevalue
		if (l % 3) == 2 {
			ticket.ticketdata[l/3].Outcome = getoutcome(ticket.ticketdata[l/3].ticketline)
		}
		l++
	}
	sort.SliceStable(ticket.ticketdata, func(i, j int) bool {
		return ticket.ticketdata[i].Outcome < ticket.ticketdata[j].Outcome
	})
	ticket.status = true
	fmt.Fprintf(w, "***the data for ticket is %+v***\n", ticket)
	return nil
}
