// Golang provides packages for important RFCs like HTTP, URI, HML, etc.

import "net/http" //web communication protocol
import "net" //TCP/IP socket programming

//JSON Marshalling - transpform golang projects into JSon objects
//marshall() returns JSON representation as []byte.

//p1 := Person (name:"lucas", endereço: "rua2", fone:"+55093281")
//barr, err := json.Marshall(p1)
//Json.Unmarshall() converts a JSON []byte into a golang object

package main
import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}


type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

type Person struct {
	name int
	addr []string
	phone []string
}

func main() {

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))


	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

res1D := &response1{
Page:   1,
Fruits: []string{"apple", "peach", "pear"}}
res1B, _ := json.Marshal(res1D)
fmt.Println(res1B)
fmt.Println(string(res1B))

res2D := &response2{
Page:   1,
Fruits: []string{"apple", "peach", "pear"}}
res2B, _ := json.Marshal(res2D)
fmt.Println(res2D)
fmt.Println(string(res2B))

res := &Person{
	name:   1,
	addr: []string{"Rua2","Rua3","Rua4"},
	phone: []string{"+55873218","837287"}}
res1, _ := json.Marshal(true)
fmt.Println(res)
fmt.Println(string(res1))


byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

var dat map[string]interface{}


if err := json.Unmarshal(byt, &dat); err != nil {
panic(err)
}
fmt.Println(dat)


num := dat["num"].(float64)
fmt.Println(num)


strs := dat["strs"].([]interface{})
str1 := strs[0].(string)
fmt.Println(str1)

str := `{"page": 1, "fruits": ["apple", "peach"]}`
res4 := response2{}
json.Unmarshal([]byte(str), &res4)
fmt.Println(res4)
fmt.Println(res4.Fruits[0])


enc := json.NewEncoder(os.Stdout)
d := map[string]int{"apple": 5, "lettuce": 7}
enc.Encode(d)
data,e := ioutil.ReadFile("main/test.txt")
if e != nil {
	log.Fatal(e)
}
fmt.Println(string(data))
//ioutil.ReaFile does not need open/close structures.
//ioutil.ReadFile can be a problem when you're dealing with LARGE files because of the realtionship between disk/RAM memory
data2 := "Hello World"
ioutil.WriteFile("main/outuput.txt",[]byte(data2), 0777) // 0777 is unix-style permission bytes

}

