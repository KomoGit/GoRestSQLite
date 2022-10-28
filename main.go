package main

func main() {
	//insertIntoDB("Billy Bob", "bibp@yeehaw.com", "+1 043 2231-1211")
	startServer()
	//insertIntoDB(fName,email,pNum) // Insert into the database. Requires full-name & email & pNum.
	//updateName(fName,uuid) // Updates the name based on id. Obvious what it requires.
	//updateEmail(email,uuid)
	//updateNumber(number,uuid)
	//deleteAllFromDB()
	//deleteWithId(uuid) //Requires UUID
	//getAllFromDB()
	//getWithId(uuid)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
