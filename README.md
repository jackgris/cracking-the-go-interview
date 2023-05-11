# Cracking the Go Interview
In this repository, I will try to save normals questions and solved algorithms related to an interview about a Go developer position.

## Interview Questions

### 1. What do you need for two functions to be the same type?

```
 A. They should share the same signatures, including parameter types and return types.
 B. They should share the same parameter types but can return different types.
 C. All functions should be the same type.
 D. The functions should not be a first class type.
```

If we want two functions to be considered the same type in Go, they need to have the same function signature.

```go
type sigFunc func(a int, b float64) (bool, error)

func functionA(a int, b float64) (bool, error) {
  return true, nil
}

func functionB(a int, b float64) (bool, error) {
  return false, nil
}

func main() {
  var x sigFunc = functionA
  x = functionB
  print(x)
}
```

That just means they should have matching parameters (quantity, types) and return values.

### 2. What does the len() function return if passed a UTF-8 encoded string?

```
 A. the number of characters
 B. the number of bytes
 C. It does not accept string types.
 D. the number of code points
```

Here’s a cool little insight for you, in Go, strings are actually byte sequences under the hood. That means when you give a UTF-8 encoded string to the len() function, it’s counting bytes, not characters or runes:

```go
func main() {
  s := "世界"
  fmt.Println("Byte length:", len(s)) // 6 bytes
  fmt.Println("Rune count:", utf8.RuneCountInString(s)) // 2 letters
}
```

### 3. What is the value of Read?

```go
const (
  Write = iota
  Read
  Execute
)
```
```
A. 0
B. 1
C. 2
D. a random value
```

The iota keyword in Go is a fascinating feature — it represents an integer value that starts at 0 and increments by 1 for each item within a const block.

### 4. How do you tell go test to print out the tests it is running?

```
A. go test
B. go test -x
C. go test --verbose
D. go test -v
```

When employing the go test command, you can enable the -v ("verbose") flag for a more detailed output.

`go test -v`

By running go test with the -v flag, you'll see the name of each test, its outcome (PASS or FAIL), and the test's duration as it proceeds:
```
=== RUN   TestAdd
=== RUN   TestAdd/case_1_2_3
=== RUN   TestAdd/case_-1_-2
=== RUN   TestAdd/case_0
=== RUN   TestAdd/case_-1_2
--- PASS: TestAdd (0.00s)
    --- PASS: TestAdd/case_1_2_3 (0.00s)
    --- PASS: TestAdd/case_-1_-2 (0.00s)
    --- PASS: TestAdd/case_0 (0.00s)
    --- PASS: TestAdd/case_-1_2 (0.00s)
PASS
ok      command-line-arguments  0.523s
```

### 5. What does a sync.Mutex block while it is locked?

```
A. all goroutines
B. any other call to lock that Mutex
C. any reads or writes of the variable it is locking
D. any writes to the variable it is locking
```

In Go, a sync.Mutex serves as a mechanism for mutual exclusion, guaranteeing that only one goroutine can access a critical code section simultaneously.

The phrase “any reads or writes of the variable it is locking” aligns with our understanding, as we utilize a sync.Mutex to safeguard access to shared variables.

var mtx = sync.Mutex{}
```go
func Add() {
  mtx.Lock()
  defer mtx.Unlock()
  a++
}
```

### 6. What is an idiomatic way to pause execution of the current scope until an arbitrary number of goroutines have returned?
```
A. Pass an int and Mutex to each and count when they return.
B. Loop over a select statement.
C. Sleep for a safe amount of time.
D. sync.WaitGroup
```

```go
func main() {
  wg := sync.WaitGroup{}
  for i := 0; i < 500; i++ {
    wg.Add(1)
    go func() {
      defer wg.Done()
      Add()
    }()
  }

  wg.Wait()
  fmt.Println(a)
}
```

The sync.WaitGroup in Go is a handy tool that enables you to wait for a group of goroutines to complete before moving forward with the execution.

It works by using a simple counter and effectively blocks the current scope until all worker goroutines finish their tasks and the WaitGroup counter reaches zero.

### 7. What is a side effect of using time.After in a select statement?
```
A. It blocks the other channels.
B. It is meant to be used in select statements without side effects.
C. It blocks the select statement until the time has passed.
D. The goroutine does not end until the time passes.
```

If you’re not familiar with time.After, it's a function in Go's time package that returns a channel set to send the current time after a specified duration.

```go
func After(d Duration) <-chan Time
```

It’s typically used in select statements for implementing timeouts or delays. For example, imagine waiting for 3 seconds before printing something on the screen:
```go
func main() {
  timeout := 3 * time.Second
  start := time.Now()
  done := make(chan bool)

  select {
  case <-done:
      fmt.Println("Operation completed.")
      return
  case <-time.After(timeout):
    fmt.Printf("Timeout after %v\n", time.Since(start))
  }
}
```

Now, let’s talk about the side effect.

For short-lived time.After durations, it might not be a big deal, but consider a scenario where the timeout is set for 1 hour, and the work finishes before the timeout. In this situation, the timer still lingers in memory:
```go
func main() {
 done := make(chan bool)

 go func() {
  time.Sleep(500 * time.Millisecond)
  done <- true
 }()

  for {
    select {
    case <-done:
      fmt.Println("Operation completed.")
      return
    case <-time.After(time.Hour):
      fmt.Println("Still waiting...")
    }
  }
}
```

As a consequence, the goroutine created by time.After won't terminate until the full hour elapses, even if the operation concludes earlier than that.

### 8. What restriction is there on the type of var to compile this i := myVal.(int)?
```
A. myVal must be an integer type, such as int, int64, int32, etc.
B. myVal must be able to be asserted as an int.
C. myVal must be an interface.
D. myVal must be a numeric type, such as float64 or int64.
```

In the context of the type assertion i := myVal.(int), the variable myVal needs to be an interface type for the code to compile successfully.

However, using it in this manner can be risky and inefficient, as it may cause a panic at runtime if myVal is not of type int. To handle this situation more gracefully, it's better to use the two-value form of a type assertion, which provides a fallback mechanism:
```go
// BAD
i := myVal.(int)

// BETTER
i, ok := myVal.(int)
// ... doing something with ok
```

### 9. What is the correct way to pass this as a body of an HTTP POST request?
```go
data := "A group of Owls is called a parliament"
```
```
A. resp, err := http.Post("https://httpbin.org/post", "text/plain", []byte(data))
B. resp, err := http.Post("https://httpbin.org/post", "text/plain", data)
C. resp, err := http.Post("https://httpbin.org/post", "text/plain", strings.NewReader(data))
D. resp, err := http.Post("https://httpbin.org/post", "text/plain", &data)
```

In order to send the data as the body of an HTTP POST request, it’s important to know the content type. Since this is raw text, we’ll use the “text/plain” content type. The http.Post function requires an io.Reader as the body, rather than a string or bytes:
```go
Post(url string, contentType string, body io.Reader) (resp *http.Response, err error)
```

The Reader interface is defined as follows:

```go
type Reader interface {
  Read(p []byte) (n int, err error)
}
```

To fulfill the requirements, we’ll convert the body into a buffer that implements this interface:
```go
func main() {
  data := "A group of Owls is called a parliament"
  contentType := "text/plain"
  body := strings.NewReader(data)
  // or
  // body := bytes.NewBufferString(data)
  resp, err := http.Post("https://example.com", contentType, body)
  // ....
}
```

So the correct way to send the data as an HTTP POST request would be:
```go
resp, err := http.Post("https://httpbin.org/post", "text/plain", strings.NewReader(data))
```

### 10. What should the idiomatic name be for an interface with a single method and the signature Save() error?
```
A. Saveable
B. SaveInterface
C. ISave
D. Saver
```

Drawing from the Effective Go naming convention for interfaces, single-method interfaces are typically named by appending an “-er” suffix or a similar modification to the method name, resulting in agent nouns like Reader, Writer, Formatter, and CloseNotifier..

In this case, the idiomatic name for the interface would be Saver.

Oh, you may have also encountered the Stringer() interface, which is defined as follows:
```go
type Stringer interface {
  String() string
}
```

This is a trick when you pass a value implementing the Stringer interface to functions like fmt.Println() or fmt.Printf(), the String() method is automatically invoked to obtain the value's string representation.

### 11. What is the default case sensitivity of the JSON Unmarshal function?

```
A. The default behavior is case insensitive, but it can be overridden.
B. Fields are matched case sensitive.
C. Fields are matched case insensitive.
D. The default behavior is case sensitive, but it can be overridden.
```

The JSON Unmarshal function in Go is, by default, case-insensitive. However, this behavior can be overridden using JSON tags. Now, what happens when there are two fields such as "title" and "Title"?

Let's explore this through an example to understand why it's case-insensitive first:
```go
type Post struct {
  Title    string `json:"Title"`
  SubTitle string
}

func main() {
  p := Post{}
  json.Unmarshal([]byte(`{"title":"hello","subtitle":"world"}`), &p)
  fmt.Println(p)
}

// {hello world}
```

In this example, the Title field has an explicit field tag json:"Title", directing the json.Unmarshal function to match the JSON field "Title" (case-sensitive) to the struct field Title. The SubTitle field does not have a field tag, so it will use the default case-insensitive behavior to match the JSON field "subtitle" to the struct field SubTitle.

This is because the encoding/json package employs a fallback mechanism during unmarshaling:

It first looks for an exact match between JSON field names and struct field names or tags.
If it doesn’t find an exact match, it reverts to a case-insensitive linear search (which can be expensive for larger data sets).
```go
type Post struct {
  Title    string
  SubTitle string
  TitleSm  string `json:"title"`
}

func main() {
  p := Post{}
  json.Unmarshal([]byte(`{"title":"hello","subtitle":"world"}`), &p)
  fmt.Println(p)
}

// { world hello}
// p.Title is empty
```

### 12. Where is the built-in recover method useful?
```
A. in the main function
B. immediately after a line that might panic
C. inside a deferred function
D. at the beginning of a function that might panic
```

The built-in recover method is indeed useful within deferred functions, but it's not advised to call it directly with the defer keyword.
```go
// BAD
func main() {
  defer recover()
  panicCode() // <-- crash
}

// BETTER
func handlePanic() {
  if panicInfo := recover(); panicInfo != nil {
    fmt.Println(panicInfo)
  }
}

func main() {
  defer handlePanic()
  panicCode()
}
```

This example demonstrates the correct usage of the recover method.

### 13. What is the difference between the time package’s Time.Sub() and Time.Add() methods?
```
A. Time.Add() is for performing addition while Time.Sub() is for nesting timestamps.
B. Time.Add() always returns a later time while time.Sub always returns an earlier time.
C. They are opposites. Time.Add(x) is the equivalent of Time.Sub(-x).
D. Time.Add() accepts a Duration parameter and returns a Time while Time.Sub() accepts a Time parameter and returns a Duration.
```

The key difference between the Time.Add() and Time.Sub() methods in the time package lies in their parameters and return values. Time.Add() accepts a Duration parameter and returns a Time value, while Time.Sub() takes a Time parameter and returns a Duration.

`“Why don’t they both accept Duration and return Time?”`

The reason is that Time.Add() can handle negative arguments, effectively functioning as a subtraction operation. Consequently, it wouldn't make sense to have another Time.Sub() method that also accepts Duration.

The Time.Add() and Time.Sub() methods serve different purposes and have distinct signatures to address specific use cases:
```go
func main() {
	now := time.Now()
	newTime := now.Add(2 * time.Hour)
	fmt.Println("Time after 2 hours:", newTime)
	duration := newTime.Sub(now)
	newTime = now.Add(-2 * time.Hour)
	fmt.Println("Time before 2 hours:", newTime)
	fmt.Println("Duration newTime to now:", duration)
}
```
```
Time after 2 hours: 2009-11-11 01:00:00 +0000 UTC m=+7200.000000001
Time before 2 hours: 2009-11-10 21:00:00 +0000 UTC m=-7199.999999999
Duration newTime to now: 2h0m0s
```

As demonstrated in this example, Time.Add() is employed to add or subtract a duration from a time value, while Time.Sub() is utilized to compute the duration between two time values.

### 14. What is the risk of using multiple field tags in a single struct?

```
A. Every field must have all tags to compile.
B. It tightly couples different layers of your application.
C. Any tags after the first are ignored.
D. Missing tags panic at runtime.
```

The primary concern with using multiple field tags in a single struct is that it can lead to tight coupling between different layers of your application. To illustrate this concept, let’s look at an example:

But what is ‘couples different layers’?
```go
type Post struct {
  Title    string `json:"title" bson:"title"`
  SubTitle string `json:"subtitle" bson:"subtitle"`
}
```

In this example, the Post struct has two field tags for each field: json and bson. These tags may be used for different purposes, such as sending HTTP responses (using json) and handling MongoDB unmarshaling (using bson).

When multiple tags like these are used, the HTTP response layer (web server) and the storage layer (MongoDB) become tightly coupled. If you want to change the title to, for example, shortTitle, you'd need to update both the HTTP response (which may also impact clients handling the response) and the MongoDB storage.

It is essential to note that using multiple field tags in this manner is not inherently wrong.

### 15. If you iterate over a map in a for range loop, in which order will the key:value pairs be accessed?
```
A. in pseudo-random order that cannot be predicted
B. in reverse order of how they were added, last in first out
C. sorted by key in ascending order
D. in the order they were added, first in first out
```

When iterating over a map using a for range loop in Go, the order in which the key-value pairs are accessed is not guaranteed to follow any specific sequence, nor does it necessarily respect the insertion order:
```go
func main() {
  m := map[string]int{}
  m["a"] = 1
  m["b"] = 2
  m["c"] = 3

  for key, value := range m {
    fmt.Printf("Key: %s, Value: %d\n", key, value)
  }

  for key, value := range m {
    fmt.Printf("Key: %s, Value: %d\n", key, value)
  }
}
```
```
Key: a, Value: 1
Key: b, Value: 2
Key: c, Value: 3
Key: c, Value: 3
Key: a, Value: 1
Key: b, Value: 2
```

As a result, you might encounter different orders each time you iterate over the same map. This behavior stems from Go’s implementation of maps as hashmaps, leading to non-deterministic iteration order.


[Go Interview Questions](https://levelup.gitconnected.com/15-go-interview-questions-from-the-linkedin-assessment-detailed-explanations-4f0878c9ff05)
