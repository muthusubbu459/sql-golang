package operations
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
const (
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
}

func(m  Employee) List()( data []Employee,i int) {
db := Opendb()
//rows,err:=db.Query("select Emp_id, LastName,FirstName,Department from employee where Emp_id = ?",3)
	data =make([]Employee,20)
	 i =0
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
//log.Println(user.EmpID, user.LastName, user.FirstName, user.Department)
data[i]=user
i++
}
err = rows.Err()
if err != nil {
log.Fatal(err)
}
return data,i
}
func(m  Employee) Insert(user * Employee) {
db := Opendb()
//stmt,err:=db.Exec("Insert into employee VALUES {3,'Narayanan','Eshwaran',Chemical}")

/*fmt.Println("Enter the id,last name, first name, department of the person ")
user := m
fmt.Scanf("%d\t%s\t%s\t%s\t", &user.EmpID, &user.LastName, &user.FirstName, &user.Department)*/
_, err := db.Exec("Insert into employee VALUES (?,?,?,?)", user.EmpID, user.LastName, user.FirstName, user.Department)
if err != nil {
log.Fatal(err)
}
/*	id,err:=stmt.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(id)*/
//list()
}

func Delete(id int) {
db := Opendb()
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
fmt.Scanf("%d\t",&id)
Delete(id)
} else {
fmt.Println(" number of rows deleted ", affect)
//list()
}
}

func Update(id int,lastname string,firstname string,dept string){
db := Opendb()
stmt, err := db.Prepare("update employee set lastname =?,firstname=?,department =? where Emp_ID =?  ")
if err != nil {
log.Fatal(err)
}
res, err := stmt.Exec(lastname,firstname,dept,id)
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
Update(id,lastname,firstname,dept)
} else {
fmt.Println(" number of rows updated ", affect)
/*fmt.Println(" the entire list of records are")
list()
*/
}
}
