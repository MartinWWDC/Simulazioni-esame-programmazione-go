# Domande Preparzione  Esame Prog I

**cos'è un array?:** un array è una  famiglia  di variabili   accessibili tramite indice di grandezza predefinitia  

esempio:

```go
var a [10]int
```

**cos'è una slice?:** una silce è una lista su un array allocata dinamicamente

esempio:

```go
var []t int
```

una  slice non contiene dati ma fa riferimento ad  un array dichirarato  sotto, indi per  ui i cambiamenti fatti su una slice  sono globabli.  NON E' un allocamento di memoria  ma  un riferimento ad  essa.

**Differnza tra lunghezza è capacità?:** 

The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

The length and capacity of a slice `s` can be obtained using the expressions `len(s)` and `cap(s)`.

```go
package main

import "fmt"

func main() {
    s := []int{2, 3, 5, 7, 11, 13}
    printSlice(s)

    // Slice the slice to give it zero length.
    s = s[:0]
    printSlice(s)

    // Extend its length.
    s = s[:4]
    printSlice(s)

    // Drop its first two values.
    s = s[2:]
    printSlice(s)
}

func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

output:

```
len=6 cap=6 [2 3 5 7 11 13]
len=0 cap=6 []
len=4 cap=6 [2 3 5 7]
len=2 cap=4 [5 7]
```

**Struttura slice:**

prendiamo un esempio

```go
var a [5]int
s:=a[1:3]
s2:=a[0:1]
```

una slice avrà tre puntatori per rappresentarla: ha un puntato alla parte iniziale  dell'array  e ha due contatori che ha come istanze la  grandezza e  la capacità della  slice.

**Funzionemento della  funzione  make**

Slices can be created with the built-in `make` function; this is how you create dynamically-sized arrays.

The `make` function allocates a zeroed array and returns a slice that refers to that array:

```go
a := make([]int, 5) // len(a)=5
```

To specify a capacity, pass a third argument to `make`:

```go
b := make([]int, 0, 5) // len(b)=0, cap(b)=5
b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:] // len(b)=4, cap(b)=4
```

**for copiano:**

```go
package main

import "fmt"

func main() {
	a := []int{0, 0, 0, 6}
	for i := range a {
		a[i] = i
	}
	for _, v := range a {
		v *= 2
	}
	fmt.Println(a)
}

```
