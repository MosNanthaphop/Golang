[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-22041afd0340ce965d47ae6ef1cefeee28c7c493a6346c4f15d667ab976d596c.svg)](https://classroom.github.com/a/kXKfjyWH)
# Mini Library Management System (6 ข้อ 100 คะแนน)

## 📝 ภาพรวม

ใน Assignment นี้ เราจะพัฒนา "ระบบจัดการหนังสือในห้องสมุดขนาดเล็ก" เพื่อฝึกทักษะในภาษา Go  โดยเน้นหลักการ OOP ครอบคลุมตั้งแต่ระดับพื้นฐาน (Pointers, Struct, Methods) จนถึงการใช้งาน Interface และ Error Handling

## 📚 รายละเอียดโจทย์

## โครงสร้าง Project (ไม่รวม File Unit Test)
```bash
mini_library/
├── models/
│   ├── book.go       # Package models สำหรับเก็บ Book struct และ Method (ex1-3)
│   ├── dvd.go        # Package models สำหรับเก็บ DVD struct และ Method (ex4-5)
│   └── ebook.go      # Package models สำหรับเก็บ EBook struct และ Method (ex6)
├── interfaces/
│   └── interfaces.go # Package interfaces สำหรับเก็บ LibraryItem และ Borrowable interface
├── go.mod            # Go module file
└── README.md         # คำอธิบาย Project
```
### เริ่มต้นโมดูล Go ด้วยคำสั่ง 
```go
go mod init mini_library
```

### 1. Basic Pointer Function ex1 : 10 คะแนน (นำ Code ไปวางไว้ที่ models/book.go)

**Scenario** เพิ่มจำนวนสต็อกหนังสือในห้องสมุดผ่าน Pointer

**โจทย์**
- สร้างฟังก์ชัน `increaseStock(stockPtr *int)` แบบ Private
- เพิ่มค่าที่ Pointer ชี้ไปขึ้นทีละ 1

```go
// IncreaseStock เพิ่มจำนวน stock ขึ้น 1 ผ่าน Pointer
func increaseStock(stockPtr *int) {
    // TODO: เติมโค้ดส่วนนี้
}
```

```go
func main() {
    stock := 5
    fmt.Println("Before:", stock) // 5
    increaseStock(&stock)
    fmt.Println("After:", stock)  // 6
}
```

### 2. Struct Basics ex2 : 15 คะแนน (นำ Code ไปวางไว้ที่ models/book.go)

**Scenario** สร้างโครงสร้างเก็บข้อมูลหนังสือ

**โจทย์**
1. ประกาศ struct Book ที่มี Private Field
   - title string
   - author string
   - stock int
2. สร้างฟังก์ชัน NewBook สำหรับสร้างหนังสือใหม่
3. สร้าง Methods สำหรับแสดงค่าแต่ละ field (Getter Methods)
   - GetTitle() string
   - GetAuthor() string
   - GetStock() int

```go
// Book แทนข้อมูลหนังสือพื้นฐาน
type Book struct {
    // TODO: กำหนด Field ที่ต้องการ
}

// NewBook เป็นเหมือน Constructor Function
func NewBook(title, author string, stock int) *Book {
    // TODO: สร้าง Book และ return Pointer
}

// Getter Methods
func (b *Book) GetTitle() string {
    // TODO: implement
}

func (b *Book) GetAuthor() string {
    // TODO: implement
}

func (b *Book) GetStock() int {
    // TODO: implement
}
```

```go
func main() {
    b := NewBook("Go 101", "Nuttachot Promrit", 3)
    fmt.Printf("Title: %s, Author: %s, Stock: %d\n", b.GetTitle(), b.GetAuthor(), b.GetStock())
    // Output: Title: Go 101, Author: Nuttachot Promrit, Stock: 3
}
```
**💡 Best Practice** ใช้ Constructor Function เพื่อให้การสร้างอ็อบเจ็กต์เป็นระบบ

### 3. Method Receiver กับ Struct ex3 : 15 คะแนน (นำ Code ไปวางไว้ที่ models/book.go)

**Scenario** เพิ่มความสามารถให้ Book ผ่าน Method

**โจทย์**
1. เพิ่ม Method `Details()` (Value Receiver) ให้ Book คืนข้อความ `<Title> by <Author> (Stock: <Stock>)`
2. เพิ่ม Method `AddStock(amount int)` (Pointer Receiver) ที่เรียกใช้ private method validateStock(amount int) error

```go
// validateStock ตรวจสอบความถูกต้องของจำนวน stock
func (b *Book) validateStock(amount int) error {
    if b.stock + amount < 0 {
        return errors.New("stock cannot be negative")
    }
    return nil
}

func (b Book) Details() string {
    // TODO: คืนข้อความ <Title> by <Author> (Stock: <Stock>)
}

func (b *Book) AddStock(amount int) error {
    // TODO: ตรวจสอบ stock ด้วย validateStock ก่อนเพิ่ม
}
```

```go
func main() {
    b := NewBook("Golang Concurrency", "Bob", 2)
    fmt.Println(b.Details()) // "Golang Concurrency by Bob (Stock: 2)"
    b.AddStock(5)
    fmt.Println(b.Details()) // "Golang Concurrency by Bob (Stock: 7)"
}
```

**💡 Best Practice** เลือกใช้ Value Receiver เมื่อไม่ต้องการแก้ไขค่าใน struct

### 4. Interface Implementation ex4 : 20 คะแนน

**Scenario** ห้องสมุดมีสื่อหลายประเภท (หนังสือ, DVD, CD) จึงต้องการ Interface กลาง

**โจทย์**
1. สร้าง Interface `LibraryItem` ที่มี Method `Info() string` (นำ Code ไปวางไว้ที่ interfaces/interfaces.go)
2. ให้ Book implement Interface นี้ โดยคืนค่า `<Title> by <Author>` (นำ Code ไปวางไว้ที่ models/book.go)
3. สร้าง struct DVD { Title string; Duration int; Stock int } ให้ Implement Info() โดยคืนค่า `DVD: <Title> (<Duration> mins)` (นำ Code ไปวางไว้ที่ models/dvd.go)
4. สร้างฟังก์ชัน NewDVD สำหรับสร้าง DVD ใหม่ (นำ Code ไปวางไว้ที่ models/dvd.go)
5. สร้างฟังก์ชัน PrintItemInfo สำหรับแสดงข้อมูล `item.Info()` (นำ Code ไปวางไว้ที่ interfaces/interfaces.go)

```go
type LibraryItem interface {
    Info() string
}

func (b Book) Info() string {
    // TODO: implement
}

type DVD struct {
    // TODO: implement
}

func (d DVD) Info() string {
    // TODO: implement
}
```

```go
func main() {
    b := NewBook("Go Patterns", "Sajjaporn", 5)
    d := NewDVD("Go Language Tour",120, 3)
    PrintItemInfo(b) // "Go Patterns by Sajjaporn"
    PrintItemInfo(d) // "DVD: Go Language Tour (120 mins)"
}
```
**💡 Best Practice** ใช้ชื่อที่สื่อความหมาย เช่น Info() เพื่อบอกว่า Function นี้ทำอะไร

### 5. Borrowable Interface ex5 : 25 คะแนน

**Scenario** ต้องการระบบยืม-คืนสำหรับหนังสือและ DVD

**โจทย์**
1. เพิ่ม Method AddStock(amount int) (Pointer Receiver) ให้ DVD (นำ Code ไปวางไว้ที่ models/dvd.go)
2. สร้าง Interface `Borrowable` (นำ Code ไปวางไว้ที่ interfaces/interfaces.go) ที่มี Method
   - `Borrow() error`
   - `Return() error`
3. ให้ Book และ DVD implement Interface นี้ (นำ Code ไปวางไว้ที่ models/book.go และ models/dvd.go ตามลำดับ)

   การยืม
   - ถ้า Stock > 0 ยืมได้ => เรียกใช้ Method `AddStock(-1)` แล้ว return nil
   - ถ้า Stock == 0 => return errors.New("out of stock")

   การคืน
   - เรียกใช้ Method `AddStock(1)` แล้ว return nil

```go
type Borrowable interface {
    // TODO: implement
}

func (b *Book) Borrow() error {
    // TODO: implement
}
func (b *Book) Return() error {
    // TODO: implement
}
func (d *DVD) Borrow() error {
    // TODO: implement
}
func (d *DVD) Return() error {
    // TODO: implement
}
```

```go
func main() {
    b := NewBook("Intro to Go", "Alice", 2)
    if err := b.Borrow(); err != nil {
        fmt.Println(err)
    }
    fmt.Println(b.GetStock()) // เหลือ 1
}
```

### 6. Struct Embedding ex6 : 15 คะแนน (นำ Code ไปวางไว้ที่ models/ebook.go)

**Scenario** สร้าง E-Book ที่ต่อยอดจาก Book

**โจทย์**
1. สร้าง struct EBook ที่ embed Book และมี private field `fileSizeMB` เป็น integer
2. สร้าง Function NewEBook สำหรับสร้าง EBook ใหม่ โดยให้เพิ่มเฉพาะ fileSizeMB โดยไม่ต้องเพิ่มจำนวน Stock (Stock = 0)
3. สร้าง Method GetFileSize() int 
4. Override Method Info() โดยให้คืนข้อความ `<Title> by <Author> (E-Book, <fileSizeMB>MB)`

```go
type EBook struct {
    // TODO: implement
}

func NewEBook(title, author string, fileSizeMB int) *EBook {
    // TODO: implement
}

func (eb *EBook) GetFileSize() int {
    // TODO: implement
}

func (eb *EBook) Info() string {
    // TODO: implement using getter methods
}
```

```go
func main() {
    eb := NewEBook("Go In Depth", "Nuttachot Promrit", 10) //การสร้าง EBook ให้เพิ่มเฉพาะ fileSizeMB โดยไม่ต้องเพิ่มจำนวน Stock (Stock = 0)
    fmt.Println(eb.Info()) 
    // "Go In Depth by Nuttachot Promrit (E-Book, 10MB)"
}
```

## 📤 การส่งงาน

1. Clone Repo จาก GitHub Classroom
2. สร้าง/แก้ไข File ตามโจทย์ที่กำหนด
3. ทดสอบด้วย `go test -v` สำหรับโจทย์แต่ละข้อดังนี้
   -  โจทย์ข้อ 1 `go test -v ./models -run ^TestEx1`
   -  โจทย์ข้อ 2 `go test -v ./models -run ^TestEx2`
   -  โจทย์ข้อ 3 `go test -v ./models -run ^TestEx3`
   -  โจทย์ข้อ 4 `go test -v ./models -run ^TestEx4`
   -  โจทย์ข้อ 5 `go test -v ./models -run ^TestEx5`
   -  โจทย์ข้อ 6 `go test -v ./models -run ^TestEx6`
4. Commit และ Push code
5. ตรวจสอบคะแนนใน Actions tab

**💡 Tips** Commit ด้วยข้อความที่ชัดเจน เช่น
```bash
git commit -m "Implement ex1-ex6"
```