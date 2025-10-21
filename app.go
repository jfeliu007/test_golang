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

// Base struct for composition
type Entity struct {
	ID      int
	Name    string
	Version int     // NEW public field
	x       float64 // private field
	y       float64 // private field
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

// NEW: Public method
func (e *Entity) GetVersion() int {
	return e.Version
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
func (c *Circle) String() string {
	return fmt.Sprintf("Circle(Name: %s, Radius: %.2f)", c.Name, c.Radius)
}

// NEW: Public method (was private)
func (c *Circle) Scale(factor float64) {
	c.Radius *= factor
}

// Rectangle implements Shape interface
type Rectangle struct {
	Entity // composition
	Width  float64
	Height float64
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

// NEW: Public method
func (r *Rectangle) Diagonal() float64 {
	return r.diagonal()
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

// NEW: Public method (was private)
func (t *Triangle) IsValid() bool {
	return t.side1+t.side2 > t.side3 &&
		t.side1+t.side3 > t.side2 &&
		t.side2+t.side3 > t.side1
}

// private method
func (t *Triangle) isValid() bool {
	return t.IsValid()
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

// REMOVED: String() method

// NEW: Public method (was private)
func (p *Point) ManhattanDistance(other *Point) int {
	return p.manhattanDistance(other)
}

// private method
func (p *Point) manhattanDistance(other *Point) int {
	return abs(p.X-other.X) + abs(p.Y-other.Y)
}

// Container uses composition and doesn't implement Shape
type Container struct {
	Entity   // composition
	Capacity int // NEW public field
	shapes   []Shape
}

// Public method
func (c *Container) AddShape(s Shape) {
	c.shapes = append(c.shapes, s)
}

// Public method
func (c *Container) TotalArea() float64 {
	total := 0.0
	for _, s := range c.shapes {
		total += s.Area()
	}
	return total
}

// REMOVED: ListShapes() method

// NEW: Public method (was private)
func (c *Container) Count() int {
	return len(c.shapes)
}

// private method
func (c *Container) count() int {
	return c.Count()
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

// REMOVED: String() method

// NEW: Public method (was private)
func (p *Person) HasEmail() bool {
	return p.email != ""
}

// private method
func (p *Person) hasEmail() bool {
	return p.HasEmail()
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

// NEW: Public method (was private)
func (a *Account) CanWithdraw(amount float64) bool {
	return a.canWithdraw(amount)
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
		Entity: Entity{ID: 1, Name: "Red Circle", Version: 1},
		Radius: 5.0,
	}
	circle.SetPosition(10, 20)

	rect := &Rectangle{
		Entity: Entity{ID: 2, Name: "Blue Rectangle", Version: 1},
		Width:  4.0,
		Height: 6.0,
	}
	rect.SetPosition(5, 15)

	triangle := &Triangle{
		Entity: Entity{ID: 3, Name: "Green Triangle", Version: 1},
		Base:   3.0,
		Height: 4.0,
	}
	triangle.SetSides(3.0, 4.0, 5.0)

	// Use Shape interface
	shapes := []Shape{circle, rect, triangle}
	fmt.Println("Shapes:")
	for _, s := range shapes {
		fmt.Printf("%s - Area: %.2f, Perimeter: %.2f\n", s.String(), s.Area(), s.Perimeter())
	}

	// Create container with capacity
	container := &Container{
		Entity:   Entity{ID: 100, Name: "Main Container", Version: 1},
		Capacity: 10,
	}
	container.AddShape(circle)
	container.AddShape(rect)
	container.AddShape(triangle)
	fmt.Printf("\nTotal area in container: %.2f\n", container.TotalArea())
	fmt.Printf("Container has %d shapes\n", container.Count())

	// Create points
	p1 := &Point{X: 0, Y: 0}
	p2 := &Point{X: 3, Y: 4}
	fmt.Printf("\nEuclidean distance: %.2f\n", p1.Distance(p2))
	fmt.Printf("Manhattan distance: %d\n", p1.ManhattanDistance(p2))

	// Create person and account
	person := &Person{
		FirstName: "John",
		LastName:  "Doe",
	}
	person.SetAge(30)
	fmt.Printf("\nPerson: %s, Age: %d, Has email: %v\n", person.FullName(), person.GetAge(), person.HasEmail())

	account := &Account{
		Person:     *person,
		AccountNum: "ACC123456",
	}
	account.Deposit(1000.50)
	account.Deposit(250.00)
	fmt.Printf("\n%s\n", account.String())
	fmt.Printf("Can withdraw $500: %v\n", account.CanWithdraw(500.00))

	// Test new public methods
	fmt.Printf("\nCircle version: %d\n", circle.GetVersion())
	circle.Scale(2.0)
	fmt.Printf("Circle after scaling: Radius=%.2f, Area=%.2f\n", circle.Radius, circle.Area())
	fmt.Printf("Rectangle diagonal: %.2f\n", rect.Diagonal())
	fmt.Printf("Triangle is valid: %v\n", triangle.IsValid())