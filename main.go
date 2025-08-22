package main

import (
	"fmt"
)

// Fungsi No 3
func swap(a, b *int){
    *a, *b = *b, *a
}

// Fungsi No 4
type Person struct {
    Name string
    Age  int
}

func safeStringLength(s *string) int {
    if s == nil {
        return 0
    }
    return len(*s)
}

func safeSliceLength(slice []int) int {
    if slice == nil {
        return 0
    }
    return len(slice)
}

func safeMapAccess(m map[string]int, key string) (int, bool) {
    // Return 0, false jika map adalah nil
    // Return value dan status dari map jika tidak nil
    if m == nil {
        return 0, false
    }
    value, ok := m[key]
    return value, ok
}

func safePrintPerson(p *Person) {
    if p == nil {
        fmt.Println("Person is nil")
    } else {
        fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
    }
}

func main() {

  // Nomor 1
  var varInt int = 100
  var varStr string = "Hello World"
  var varFloat float32 = 10.0
  var varBool bool = true
  var varSlice = []string{"Go", "Slices", "Are", "Powerful"}
  
  fmt.Println("=== SOAL NO 1 ===")
  fmt.Println("This is Integer = ", varInt)
  fmt.Println("This is String = ", varStr)
  fmt.Println("This is Float = ", varFloat)
  fmt.Println("This is Bool = ", varBool)
  fmt.Println("This is Slice = ", varSlice)

  // Nomor 2
  var dataMahasiswa = make(map[string]string)
  dataMahasiswa["Nama"] = "George Misael"
  dataMahasiswa["NIM"] = "434231115"
  dataMahasiswa["Email"] = "george@gmail.com"
  dataMahasiswa["Hp"] = "085175140670"

  var b []string 
  b = append(b, "Nama", "NIM", "Email", "Hp")
  
  fmt.Println("\n=== SOAL NO 2 ===")
  for _, element := range b {
    fmt.Printf("%v : %v \n", element, dataMahasiswa[element])
  }

  // Nomor 3
  x := 10
  y := 20

  swap(&x, &y)

  fmt.Println("\n=== SOAL NO 3 ===")
  fmt.Printf("x = %d, alamat memori: %p\n", x, &x)
  fmt.Printf("y = %d, alamat memori: %p\n", y, &y)

  // Nomor 4
  fmt.Println("\n=== SOAL NO 4 ===")
  var nilString *string
  validString := "Hello World"
  
  var nilSlice []int
  validSlice := []int{1, 2, 3}
  
  var nilMap map[string]int
  validMap := map[string]int{"a": 1, "b": 2}
  
  var nilPerson *Person
  validPerson := &Person{Name: "Alice", Age: 30}
  
  // safeStringLength
  fmt.Println("=== safeStringLength ===")
  fmt.Printf("Panjang dari nilString: %d\n", safeStringLength(nilString))
  fmt.Printf("Panjang dari validString: %d\n", safeStringLength(&validString))
  
  // safeSliceLength
  fmt.Println("\n=== safeSliceLength ===")
  fmt.Printf("Panjang dari nilSlice: %d\n", safeSliceLength(nilSlice))
  fmt.Printf("Panjang dari validSlice: %d\n", safeSliceLength(validSlice))
  
  // safeMapAccess
  fmt.Println("\n=== safeMapAccess ===")
  val1, ok1 := safeMapAccess(nilMap, "c")
  fmt.Printf("Akses map nil: value=%d, ok=%t\n", val1, ok1)
  
  val2, ok2 := safeMapAccess(validMap, "b")
  fmt.Printf("Akses map valid (key 'b'): value=%d, ok=%t\n", val2, ok2)
  
  val3, ok3 := safeMapAccess(validMap, "z")
  fmt.Printf("Akses map valid (key 'z' tidak ada): value=%d, ok=%t\n", val3, ok3)
  
  // safePrintPerson
  fmt.Println("\n=== safePrintPerson ===")
  safePrintPerson(nilPerson)
  safePrintPerson(validPerson)
}