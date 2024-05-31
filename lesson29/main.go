/*un 
package main

import (
	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
type Product struct {
	gorm.Model
	Code  string
	Price uint
}
func main() {
	//dbURL :=
	db, _ := gorm.Open(postgres.Open("postgres://postgres:root@localhost:5432/postgres?sslmode=disable"))

	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "D43", Price: 1654})
	db.Create(&Product{Code: "D45433", Price: 56})

	var product Product
	db.First(&product, 1)
	db.First(&product, "code = ? and price = ?", "D42", 100) // find product with code D42
	db.Model(&product).Where("code = ? and price = ?", "D42", 100).Update("Price", 200)

	fmt.Println(product)

	db.Delete(&product, 1)
}
*/



package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User1 struct {
	gorm.Model
	Id         string `gorm:"primaryKey"`
	FirstName  string
	LastName   string
	Email      string
	Password   string
	Age        int
	Field      string
	Gender     string
	IsEmployee bool
}

func main() {
	dsn := "postgres://postgres:root@localhost:5432/postgres?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Databasaga ulanib bolmadi:", err)
	}

	err = db.AutoMigrate(&User1{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Create and insert
		users := []User1{
		{Id: "1", FirstName: "John", LastName: "Doe", Email: "john.doe1@example.com", Password: "password123", Age: 30, Field: "Engineering", Gender: "Male", IsEmployee: true},
		{Id: "2", FirstName: "Jane", LastName: "Smith", Email: "jane.smith@example.com", Password: "password123", Age: 28, Field: "Marketing", Gender: "Female", IsEmployee: true},
		{Id: "3", FirstName: "Emily", LastName: "Johnson", Email: "emily.johnson@example.com", Password: "password123", Age: 35, Field: "Finance", Gender: "Female", IsEmployee: true},
		{Id: "4", FirstName: "Michael", LastName: "Brown", Email: "michael.brown@example.com", Password: "password123", Age: 40, Field: "HR", Gender: "Male", IsEmployee: true},
		{Id: "5", FirstName: "Jessica", LastName: "Davis", Email: "jessica.davis@example.com", Password: "password123", Age: 29, Field: "IT", Gender: "Female", IsEmployee: true},
		{Id: "6", FirstName: "David", LastName: "Martinez", Email: "david.martinez@example.com", Password: "password123", Age: 33, Field: "Sales", Gender: "Male", IsEmployee: true},
		{Id: "7", FirstName: "Sarah", LastName: "Lopez", Email: "sarah.lopez@example.com", Password: "password123", Age: 27, Field: "Support", Gender: "Female", IsEmployee: true},
		{Id: "8", FirstName: "Chris", LastName: "Wilson", Email: "chris.wilson@example.com", Password: "password123", Age: 31, Field: "Engineering", Gender: "Male", IsEmployee: true},
		{Id: "9", FirstName: "Karen", LastName: "Moore", Email: "karen.moore@example.com", Password: "password123", Age: 38, Field: "Marketing", Gender: "Female", IsEmployee: true},
		{Id: "10", FirstName: "James", LastName: "Taylor", Email: "james.taylor@example.com", Password: "password123", Age: 34, Field: "Finance", Gender: "Male", IsEmployee: true},
		{Id: "11", FirstName: "Amanda", LastName: "Anderson", Email: "amanda.anderson@example.com", Password: "password123", Age: 32, Field: "HR", Gender: "Female", IsEmployee: true},
		{Id: "12", FirstName: "Robert", LastName: "Thomas", Email: "robert.thomas@example.com", Password: "password123", Age: 36, Field: "IT", Gender: "Male", IsEmployee: true},
		{Id: "13", FirstName: "Melissa", LastName: "Jackson", Email: "melissa.jackson@example.com", Password: "password123", Age: 28, Field: "Sales", Gender: "Female", IsEmployee: true},
		{Id: "14", FirstName: "Mark", LastName: "White", Email: "mark.white@example.com", Password: "password123", Age: 37, Field: "Support", Gender: "Male", IsEmployee: true},
		{Id: "15", FirstName: "Nancy", LastName: "Harris", Email: "nancy.harris@example.com", Password: "password123", Age: 30, Field: "Engineering", Gender: "Female", IsEmployee: true},
		{Id: "16", FirstName: "Daniel", LastName: "Clark", Email: "daniel.clark@example.com", Password: "password123", Age: 35, Field: "Marketing", Gender: "Male", IsEmployee: true},
		{Id: "17", FirstName: "Laura", LastName: "Lewis", Email: "laura.lewis@example.com", Password: "password123", Age: 33, Field: "Finance", Gender: "Female", IsEmployee: true},
		{Id: "18", FirstName: "Steven", LastName: "Robinson", Email: "steven.robinson@example.com", Password: "password123", Age: 39, Field: "HR", Gender: "Male", IsEmployee: true},
		{Id: "19", FirstName: "Susan", LastName: "Walker", Email: "susan.walker@example.com", Password: "password123", Age: 29, Field: "IT", Gender: "Female", IsEmployee: true},
		{Id: "20", FirstName: "Paul", LastName: "Hall", Email: "paul.hall@example.com", Password: "password123", Age: 34, Field: "Sales", Gender: "Male", IsEmployee: true},
		{Id: "21", FirstName: "Linda", LastName: "Allen", Email: "linda.allen@example.com", Password: "password123", Age: 37, Field: "Support", Gender: "Female", IsEmployee: true},
		{Id: "22", FirstName: "Joshua", LastName: "Young", Email: "joshua.young@example.com", Password: "password123", Age: 31, Field: "Engineering", Gender: "Male", IsEmployee: true},
		{Id: "23", FirstName: "Betty", LastName: "King", Email: "betty.king@example.com", Password: "password123", Age: 38, Field: "Marketing", Gender: "Female", IsEmployee: true},
		{Id: "24", FirstName: "Edward", LastName: "Wright", Email: "edward.wright@example.com", Password: "password123", Age: 36, Field: "Finance", Gender: "Male", IsEmployee: true},
		{Id: "25", FirstName: "Barbara", LastName: "Scott", Email: "barbara.scott@example.com", Password: "password123", Age: 32, Field: "HR", Gender: "Female", IsEmployee: true},
	}

	//Create
	result := db.Create(&users)
	if result.Error != nil {
		log.Fatal("Create users:", result.Error)
	} else {
		fmt.Println("Users created")
	}

	//Read
	var user User1
	err = db.First(&user, "id = ?", "1").Error
	if err != nil {
		log.Fatal("Failed to read user:", err)
	}
	fmt.Printf("User: %+v\n", user)

	// Update 
	err = db.Model(&user).Update("Email", "john.updated@example.com").Error
	if err != nil {
		log.Fatal("Failed to update user email:", err)
	} else {
		fmt.Printf("Updated User Email: %s\n", user.Email)
	}

	// Delete 
	err = db.Delete(&user).Error
	if err != nil {
		log.Fatal("Failed to delete user:", err)
	} else {
		fmt.Println("User deleted successfully")
	}
}
