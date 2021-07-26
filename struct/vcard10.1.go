package main

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

type Address struct {
	Country string
	Province string
	City string
	Village string
	Street           string
	HouseNumber      int
	POBox            string
	ZipCode          string
}

type VCard struct {
	FirstName string
	LastName  string
	NickName  string
	BirtDate  time.Time
	Photo     string
	Addresses []*Address
}

func main() {
	//addr1 := &Address{"Elfenstraat", 12, "", "", "2600", "Mechelen", "België"}
	//addr2 := &Address{"Heideland", 28, "", "", "2640", "Mortsel", "België"}
	addr1 := &Address{"中国", "陕西省", "宝鸡市", "扶风县", "安上村", 80, "", "722200"}
	addr2 := &Address{"China", "Shanxi", "Baoji", "Fufeng", "AnShang", 90, "", "700000"}
	// make 一个长度为5的*Address的切片
	addrs := make([]*Address, 5)
	//addrs := map[string]*Address{}
	addrs[0] = addr1
	addrs[1] = addr2
	birthdt := time.Date(1996, 1, 5, 0, 0, 0, 0, time.Local)
	photo := "MyDocuments/MyPhotos/photo1.jpg"
	vcard := &VCard{"瑞峰", "安", "", birthdt, photo, addrs}
	// vcard 是一个结构体指针
	fmt.Printf("Here is the full VCard: %v\n", vcard)
	fmt.Println(reflect.TypeOf(vcard), unsafe.Sizeof(vcard))
	fmt.Printf("My Addresses are:\n %v\n %v \n", addr1, addr2)
	fmt.Println(reflect.TypeOf(addr1), reflect.TypeOf(addr2))
}

/* Output:
Here is the full VCard: &{Ivo Balbaert  Sun Jan 17 15:04:05 +0000 1956 MyDocuments/MyPhotos/photo1.jpg map[now:0x126d57c0 youth:0x126d5500]}
My Addresses are:
 &{Elfenstraat 12   2600 Mechelen België}
 &{Heideland 28   2640 Mortsel België}
*/
