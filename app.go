package main

import (
	"fmt"
	"math"
)

// Common interface that some structs will implement
type Shape interface {
	Area() float64
	Perimeter() float64
	String() string
}

// Common interface that some structs will implement
type Shape2 interface {
	Area() float64
	Perimeter() float64
	Stringe() string
}

// Base struct for composition
type NewThing struct {
	ID   int
	Name string
	x    float64 // private field
	y    float64 // private field
}

// Base struct for composition
type Entity struct {
	ID   int
	Name string
	x    float64 // private field
	y    float64 // private field
}

// Public method
func (e *Entity) GetPosition() (float64, float64) {
	return e.x, e.y
}

// Public method
func (e *Entity) SetPosition(x, y float64) {
	e.x = x
	e.y = y
}

// private method
func (e *Entity) distance(other *Entity) float64 {
	dx := e.x - other.x
	dy := e.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

// Circle implements Shape interface
type Circle struct {
	Entity // composition
	Radius float64
}

// Public method - implements Shape interface
func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Public method - implements Shape interface
func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Public method - implements Shape interface
func (c *Circle) Stringe() string {
	return fmt.Sprintf("Circle(Name: %s, Radius: %.2f)", c.Name, c.Radius)
}

// private method
func (c *Circle) scale(factor float64) {
	c.Radius *= factor
}

// Rectangle implements Shape interface
type Rectangle struct {
	Entity     // composition
	Width      float64
	Height     float64
	AddedField float64
}

// Public method - implements Shape interface
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Public method - implements Shape interface
func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Public method - implements Shape interface
func (r *Rectangle) String() string {
	return fmt.Sprintf("Rectangle(Name: %s, Width: %.2f, Height: %.2f)", r.Name, r.Width, r.Height)
}

// Public method
func (r *Rectangle) IsSquare() bool {
	return r.Width == r.Height
}

// private method
func (r *Rectangle) diagonal() float64 {
	return math.Sqrt(r.Width*r.Width + r.Height*r.Height)
}

// Triangle implements Shape interface
type Triangle struct {
	Entity // composition
	Base   float64
	Height float64
	side1  float64 // private field
	side2  float64 // private field
	side3  float64 // private field
}

// Public method - implements Shape interface
func (t *Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Public method - implements Shape interface
func (t *Triangle) Perimeter() float64 {
	return t.side1 + t.side2 + t.side3
}

// Public method - implements Shape interface
func (t *Triangle) String() string {
	return fmt.Sprintf("Triangle(Name: %s, Base: %.2f, Height: %.2f)", t.Name, t.Base, t.Height)
}

// Public method
func (t *Triangle) SetSides(s1, s2, s3 float64) {
	t.side1 = s1
	t.side2 = s2
	t.side3 = s3
}

// private method
func (t *Triangle) isValid() bool {
	return t.side1+t.side2 > t.side3 &&
		t.side1+t.side3 > t.side2 &&
		t.side2+t.side3 > t.side1
}

// Point doesn't implement Shape interface
type Point struct {
	X int
	Y int
}

// Public method
func (p *Point) Distance(other *Point) float64 {
	dx := float64(p.X - other.X)
	dy := float64(p.Y - other.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

// Public method
func (p *Point) String() string {
	return fmt.Sprintf("Point(%d, %d)", p.X, p.Y)
}

// private method
func (p *Point) manhattanDistance(other *Point) int {
	return abs(p.X-other.X) + abs(p.Y-other.Y)
}

// Person doesn't use composition
type Person struct {
	FirstName string
	LastName  string
	age       int    // private field
	email     string // private field
}

// Public method
func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

// Public method
func (p *Person) GetAge() int {
	return p.age
}

// Public method
func (p *Person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	}
}

// Public method
func (p *Person) String() string {
	return fmt.Sprintf("Person(%s, Age: %d)", p.FullName(), p.age)
}

// private method
func (p *Person) hasEmail() bool {
	return p.email != ""
}

// Account uses composition with Person
type Account struct {
	Person     // composition
	AccountNum string
	balance    float64 // private field
}

// Public method
func (a *Account) Deposit(amount float64) {
	if amount > 0 {
		a.balance += amount
	}
}

// Public method
func (a *Account) GetBalance() float64 {
	return a.balance
}

// Public method
func (a *Account) String() string {
	return fmt.Sprintf("Account(%s, Num: %s, Balance: $%.2f)", a.FullName(), a.AccountNum, a.balance)
}

// private method
func (a *Account) canWithdraw(amount float64) bool {
	return amount > 0 && amount <= a.balance
}

// Helper function (private)
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Example usage
func main() {
	// Create shapes
	circle := &Circle{
		Entity: Entity{ID: 1, Name: "Red Circle"},
		Radius: 5.0,
	}
	circle.SetPosition(10, 20)

	rect := &Rectangle{
		Entity: Entity{ID: 2, Name: "Blue Rectangle"},
		Width:  4.0,
		Height: 6.0,
	}
	rect.SetPosition(5, 15)

	triangle := &Triangle{
		Entity: Entity{ID: 3, Name: "Green Triangle"},
		Base:   3.0,
		Height: 4.0,
	}
	triangle.SetSides(3.0, 4.0, 5.0)

	// Use Shape interface
	shapes := []Shape{rect, triangle}
	fmt.Println("Shapes:")
	for _, s := range shapes {
		fmt.Printf("%s - Area: %.2f, Perimeter: %.2f\n", s.String(), s.Area(), s.Perimeter())
	}

	// Create points
	p1 := &Point{X: 0, Y: 0}
	p2 := &Point{X: 3, Y: 4}
	fmt.Printf("\n%s to %s distance: %.2f\n", p1.String(), p2.String(), p1.Distance(p2))

	// Create person and account
	person := &Person{
		FirstName: "John",
		LastName:  "Doe",
	}
	person.SetAge(30)

	account := &Account{
		Person:     *person,
		AccountNum: "ACC123456",
	}
	account.Deposit(1000.50)
	account.Deposit(250.00)
	fmt.Printf("\n%s\n", account.String())
}
