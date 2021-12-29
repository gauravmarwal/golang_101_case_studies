package main

import "fmt"

// Creating a Simple Interface "Employee"
type Employee interface {
	GetDetails() string
	GetEmployeeSalary() int
}

// Creating a new type "Manager" containing all functions required by "Employee" Interface
type Manager struct {
	Name        string
	Age         int
	Designation string
	Salary      int
}

func (mgr Manager) GetDetails() string {
	return mgr.Name
}

func (mgr Manager) GetEmployeeSalary() int {
	return mgr.Salary
}

func main() {
	// Creating new Object of Type Manager
	newManager := Manager{Name: "Mayank", Age: 30, Designation: "Developer", Salary: 10}

	GetDetails(newManager)

	var employeeInterface Employee

	//Manager Object assigned to Interface type since Interface Contract is satisfied
	employeeInterface = newManager

	// Invoke Functions that belong to Interface "Employee"
	fmt.Println(employeeInterface.GetDetails())

}
