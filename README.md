//# sql-golang
//insert update operations


package main

import (
	//"database/sql"
	//"fmt"
	_ "github.com/go-sql-driver/mysql"

	"log"
	"data/operations"
	//"container/list"
	"fmt"
	"github.com/gorilla/mux"

	"net/http"
	"encoding/json"
	"strconv"
)

var emps = make(map[string]operations.Employee)

/*const (
	DB_HOST = "tcp(127.0.0.1:3306)"
	DB_NAME = "emp"
	DB_USER = "root"
	DB_PASS = "wazzarooney"
)

type Employee struct {
	EmpID     int
	LastName   string
	FirstName  string
	Department string
}

type Empty interface {
	Insert()
	List()
	Delete()
	Update()

}

func Opendb() *sql.DB {
	dsn := DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}*/
/*func(m  Employee) List() {
	db := Opendb()
	//rows,err:=db.Query("select Emp_id, LastName,FirstName,Department from employee where Emp_id = ?",3)
	rows, err := db.Query("select Emp_id, LastName,FirstName,Department from employee")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		user := m
		err := rows.Scan(&user.EmpID, &user.LastName, &user.FirstName, &user.Department)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(user.EmpID, user.LastName, user.FirstName, user.Department)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}*/
/*func(m  Employee) Insert() {
	db := Opendb()
	//stmt,err:=db.Exec("Insert into employee VALUES {3,'Narayanan','Eshwaran',Chemical}")

	fmt.Println("Enter the id,last name, first name, department of the person ")
	user := m
	fmt.Scanf("%d\t%s\t%s\t%s\t", &user.EmpID, &user.LastName, &user.FirstName, &user.Department)
	_, err := db.Exec("Insert into employee VALUES (?,?,?,?)", user.EmpID, user.LastName, user.FirstName, user.Department)
	if err != nil {
		log.Fatal(err)
	}
	/*	id,err:=stmt.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(id)*/
	//list()*/
/*func Delete() {
	db := Opendb()
	var id int
	fmt.Println("enter the user id you want to delete")
	fmt.Scanf("%d\t", &id)
	stmt, err := db.Prepare("delete from employee where Emp_ID=?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if affect == 0 {
		fmt.Println("there is no such id in record")
		fmt.Println("Please enter a valid id to delete")
		Delete()
	} else {
		fmt.Println(" number of rows deleted ", affect)
		//list()
	}
}*/
/*func(m Employee) Update(){
	db := Opendb()
	var id int
	var dept string
	fmt.Println("enter the user id you and department so that departments can updated ")
	fmt.Scanf("%d\t%s\t", &id, &dept)
	stmt, err := db.Prepare("update employee set department =? where Emp_ID =?  ")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(dept,id)
	if err != nil {
		log.Fatal(err)
	}
	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if affect == 0 {
		fmt.Println("there is no such id in record")
		fmt.Println("Please enter a valid id to update")
		m.Update()
	} else {
		fmt.Println(" number of rows updated ", affect)
		fmt.Println(" the entire list of records are")
		//list()
	}
}*/
func main() {
	db := operations.Opendb()
	/*	dsn:=DB_USER+":"+DB_PASS+"@"+DB_HOST+"/"+DB_NAME+"?charset=utf8"

		db,err:=sql.Open("mysql",dsn)
		if err != nil {
		log.Fatal(err)
		}*/
	defer db.Close()
	user :=new(operations.Employee)
	value,i :=user.List()
	fmt.Println("the list of records are below")
	for j:=0;j<i;j++{
		log.Println(value[j].EmpID, value[j].LastName, value[j].FirstName, value[j].Department)
	}
	/*user =new(operations.Employee)
	fmt.Println("enter the user id you want to delete")
	fmt.Scanf("%d\t",&user.EmpID)
	//var i empty
	//i = new
	//i.list()
	//i.insert()
	//list()
	//insert()
	operations.Delete(user.EmpID)
	user=new(operations.Employee)
	fmt.Println("enter the user id you and department so that departments can updated ")
	fmt.Scanf("%d\t%s\t", &user.EmpID, &user.Department)
	operations.Update(user.EmpID,user.Department)
	fmt.Println("Enter the id,last name, first name, department of the person to be inserted")
	user =new(operations.Employee)
	fmt.Scanf("%d\t%s\t%s\t%s\t", &user.EmpID, &user.LastName, &user.FirstName, &user.Department)
	//data:=operations.Employee{9,"bob","simpson","chemistry"}
	//data.insert(data)
	user.Insert(user)
	user= new(operations.Employee)
	value,i =user.List()
	fmt.Println("the list of records are below")
	for j:=0;j<i;j++{
		log.Println(value[j].EmpID, value[j].LastName, value[j].FirstName, value[j].Department)
	}*/

	//router
	router:=mux.NewRouter()
	router.HandleFunc("/data",data).Methods("GET")
	router.HandleFunc("/data/{id}",dataid).Methods("GET","POST","DELETE","PUT")
	http.ListenAndServe(":1880",router)

}
func dataid(res http.ResponseWriter,req*http.Request){
	res.Header().Set("content type","application/json")
	vars:=mux.Vars(req)
	id :=vars["id"]
	switch req.Method {
	case "GET": emp,ok := emps[id]
				if !ok {
				res.WriteHeader(http.StatusNotFound)
				fmt.Fprint(res, string("data not found"))
				}
				outgoingJSON,error:= json.Marshal(emp)
				if error != nil {
					log.Println(error.Error())
					http.Error(res, error.Error(), http.StatusInternalServerError)
				return
				}
				fmt.Fprint(res, string(outgoingJSON))
	case "POST":emp := new(operations.Employee)
				decoder := json.NewDecoder(req.Body)
				error := decoder.Decode(&emp)
				if error != nil {
					log.Println(error.Error())
					http.Error(res, error.Error(), http.StatusInternalServerError)
					return
				}
				emp.Insert(emp)
				emps[id] = * emp

				outgoingJSON, err := json.Marshal(emp)
				if err != nil {
					log.Println(error.Error())
					http.Error(res, err.Error(), http.StatusInternalServerError)
					return
				}
				res.WriteHeader(http.StatusCreated)
				fmt.Fprint(res, string(outgoingJSON))
	case "DELETE":delete(emps,id)
			//	i,err:=strconv.Atoi(id)
			//	if err!=nil{
			//		log.Fatal(err)
			//	}
				//operations.Delete(i)
				res.WriteHeader(http.StatusNoContent)
	case "PUT": emp,ok := emps[id]
				if !ok {
					res.WriteHeader(http.StatusNotFound)
					fmt.Fprint(res, string("data not found to modify"))
				}
				empnew := new(operations.Employee)
				decoder := json.NewDecoder(req.Body)
				error := decoder.Decode(&empnew)
				if error != nil {
					log.Println(error.Error())
					http.Error(res, error.Error(), http.StatusInternalServerError)
					return
				}
				operations.Update(emp.EmpID,empnew.LastName,empnew.FirstName,empnew.Department)
				//emps[id] = emp
				user:=new(operations.Employee)
				value,i:=user.List()
				k,err:=strconv.Atoi(id)
				if err!=nil{
					log.Println(error.Error())
					http.Error(res, err.Error(), http.StatusInternalServerError)
					return
				}
				fmt.Println(k)
				for j:=0;j<=i;j++{
				if(value[j].EmpID ==k) {
						fmt.Println(j,k)
						emp= value[j]
					}
				}
				fmt.Println(emp.FirstName,emp.LastName,emp.EmpID,emp.Department)
				if !ok {
					res.WriteHeader(http.StatusNotFound)
					fmt.Fprint(res, string("data not found to modify"))
				}
				outgoingJSON,error:= json.Marshal(emp)
				if error != nil {
					log.Println(error.Error())
					http.Error(res, error.Error(), http.StatusInternalServerError)
					return
				}
				fmt.Fprint(res, string(outgoingJSON))
	}

}
func data( res http.ResponseWriter,req* http.Request){
res.Header().Set("content type","application/json")
	user:=new(operations.Employee)
	value,i:=user.List()
	//emp:= make(map[string]operations.Employee)
	for j := 0; j < i; j++ {
		K:=strconv.Itoa(value[j].EmpID)
		emps[K] = value[j]
	}

	outgoingJSON,error:=json.Marshal(emps)
	if error != nil{
		log.Println(error.Error())
		http.Error(res,error.Error(),http.StatusInternalServerError)
		return
	}
	//json.MarshalIndent(emp," ","    ")
	fmt.Fprintf(res,string(outgoingJSON))
	/*slc := make([]operations.Employee)
	for _, val := range value {
		slc = append(out, val)
	}
	json, _ := json.Marshal(slc)
	fmt.Fprintf(res,string(json))*/


}

