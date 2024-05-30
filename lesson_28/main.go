package main

import (
	"log"
	"github.com/Go11Group/at_lesson/lesson28/model"
	"github.com/Go11Group/at_lesson/lesson28/storage/postgres"
)

func main() {

	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	st := postgres.NewStudentRepo(db)
	cr := postgres.NewCourseRepo(db)


	users, err := st.GetAllStudents()
	if err != nil {
		log.Fatal(err)
	}

	GetByID, err := st.GetByID(users[2].ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("GetByID:", GetByID)

	newUser := model.User{
		Name:     "New User",
		Age:      20,
		Gender:   "Non-binary",
		Course_id: 1,
	}

	createdUser, err := st.Create(newUser)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created User:", createdUser)

	createdUser.Name = "Updated User"
	err = st.Update(createdUser)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Updated User:", createdUser)

	err = st.Delete(createdUser.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Deleted User with ID:", createdUser.ID)

	newCourse := model.Course{
		Name:  "Computer Science",
		Field: "Engineering",
	}

	err = cr.Create(&newCourse)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created Course:", newCourse)

	// Get course by ID
	courseByID, err := cr.GetByID(newCourse.Id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("GetByID:", courseByID)

	// Update course
	newCourse.Name = "Updated Course Name"
	err = cr.Update(&newCourse)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Updated Course:", newCourse)

	err = cr.Delete(newCourse.Id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Deleted Course with ID:", newCourse.Id)

	allCourses, err := cr.GetAllCourses()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("All Courses:", allCourses)
}
