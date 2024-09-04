//อ้างอิงจาก mikelobster 	https://docs.mikelopster.dev/c/goapi-essential/intro

package main //โครงสร้างพื้นฐาน 		//normal struct
//for import module
import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	// 		= 	คอมเมนต์แบบบรรทัดเดียว
	/*
		= 	คอมเมนต์แบบหลายบรรทัด
	*/

	/*
	   datatype
	   string เก็บข้อความ "a"
	   int เก็บเลขจำนวนเต็ม 1 2 3 213
	   float เก็บทศนิยม  //12.2 13.2
	   bool เก็บ "จริง" หรือ "เท็จ" //true // false
	   ประกาศตัวแปร
	   variable
	*/
	var a int = 4 //ประกาศแบบเต็มรูป
	var b = 4     //ประกาศแบบไม่ระบุ Datatype
	var c float32 //ประกาศแบบไม่ระบุข้อมูล
	d := "Hello"  //ประกาศแบบย่อ
	/*
	   Operator คือ symbol (สัญลักษณ์) พิเศษที่ใช้สำหรับดำเนินการ (operation) ระหว่าง variable (ตัวแปร) และ value (ค่าของตัวแปร)
	   + = บวก
	   - = ลบ
	   * =  คูณ
	   / = หาร
	   % = หารเอาเศษ
	   a := 10
	   b := 3
	   (a + b) // 13
	   (a - b) // 7
	   (a * b) // 30
	   (a / b) // 3
	   (a % b) // 1
	   Relational Operators คือ operator สำหรับการเปรียบเทียบ และคืนค่ากลับออกมาเป็น boolean
	   โดยปกติ จะประกอบไปด้วย == (เท่ากับ), != (ไม่เท่ากับ), < (น้อยกว่า), <= (น้อยกว่าหรือเท่ากับ), > (มากกว่า), >= (มากกว่าหรือเท่ากับ)
	   (a == b) // false
	   (a != b) // true
	   (a > b)  // true
	   (a < b)  // false
	   (a >= b) // true
	   (a <= b) // false
	   Logical Operators คือ operator ที่ใช้สำหรับรวม condition
	   โดยปกติ จะประกอบไปด้วย && (และ), || (หรือ), ! (ไม่ )
	   c := true
	   d := false
	   (c && d) // false
	   (c || d) // true
	   (!c)     // false

	   ** Datatype
	   //https://golang.google.cn/pkg/fmt/#:~:text=The%20verb%20%25v%20('v,comment%20for%20all%20the%20details. fmt document
	   int %d | string %s| float %f|bool %t
	*/
	//Print เป็นสั่งที่ใช้แสดงข้อความ ต้องมีmodule fmtเสมอ
	fmt.Print(d, "\n") //Print แสดงข้อความแบบไม่เคาะบรรทัดให้ (\n) เพื่อสร้างnewline หรือ เคาะบรรทัด
	fmt.Printf("This is%v\tAnd This is %v\n", a, b)
	//Printf แสดงข้อความแบบไม่เคาะบรรทัดให้แต่ต้องสร้าง "" เพื่อใส่ Datatype ให้ variable หรือ ตัวแปร เข้าไป  (\t) เพื่อTap
	fmt.Println(c) //Println แสงข้อความแบบเคาะบรรทัดให้

	//Scan เป็นสั่งที่ใช้Assign ค่า, ใส่ค่าเข้าไปในตัวแปร ต้องมีmodule fmtเสมอ
	fmt.Scan(&c) //Scanเป็นคำสั่งที่ใช้Assign ค่า ต้องใส่ & หน้าตัวแปรเสมอไม่งั้นคำสั่งจะไม่ทำงาน
	fmt.Scanf("%f", &c)
	//Scanf เป็นสั่งที่ใช้Assign ค่า ต้องใส่ & หน้าตัวแปรเสมอไม่งั้นคำสั่งจะไม่ทำงานแต่ต้องสร้าง "" เพื่อใส่ Datatype ให้ variable
	fmt.Scanln(&c) //Scanln เป็นคำสั่งที่ใช้Assign ค่า ต้องใส่ & หน้าตัวแปรเสมอไม่งั้นคำสั่งจะไม่ทำงาน
	fmt.Println("C = ", c)

	//Sprint เป็นคำสั่งที่ใช้เก็บข้อความไว้ ต้องสร้างตัวแปรมารับ   คล้ายคำสั่ง Print แค่ไม่ได้ใช้โดยตรง
	sprint := fmt.Sprint("This isd Sprint", c, "\n")
	sprintf := fmt.Sprintf("This isd Sprint %v", c)
	sprintln := fmt.Sprintln("This isd Sprint", c)

	fmt.Println(sprint, sprintf, sprintln)
	//Console InPut Advance
	Scan := bufio.NewScanner(os.Stdin)
	Scan.Scan()
	input := Scan.Text()
	fmt.Printf("This is Input op : %s\n", input)
	//Module Math
	fmt.Println(math.Ceil(3.00000001))    //4
	fmt.Println(math.Floor(3.00000001))   //3
	fmt.Println(math.Max(-2, 3.00000001)) //3.00000001   หาค่ามาก
	fmt.Println(math.Min(-2, 3.00000001)) //-2				หาค่าน้อย
	fmt.Println(math.Abs(-2))             //2				เป็นลบเป็นบวก
	fmt.Println(math.Sqrt(10))            //3.162276601683795   สแควร์รูท 10
	fmt.Println(math.Pow(10, 2))          //100       10 ยกกำลัง  2

	//String
	//https://www.cs.cmu.edu/~pattis/15-1XX/common/handouts/ascii.html  << อ่านASCII valueของคอมพิวเตอร์เพิ่ม
	H := "Hello Wolrd"
	fmt.Println(H)      //Hello wolrd
	fmt.Println(len(H)) //11
	//H[0] ใน[]คือลำดับของ H โดยลำดับของคอมพิวเตอร์จะเริ่มที่ 0 เสมอ
	//Ex H[0] = H , H[1] = e , H[2] = l
	fmt.Println(H[0])        //72        Conver Charactor H[0] to  Decimal Value
	fmt.Printf("%c\n", H[0]) //H			%c Conver Decimal value to Charator
	fmt.Println(H[0:6])      // Hello		H[ขอบเขต:ขอบเขต]
	fmt.Println(H[1:7])      // ello W

	H = H + "Hello" //Hell0 Wolrd + Hello = Hello WolrdHello
	H += "Hello"    //Hell0 Wolrd + Hello + Hello  = Hello WolrdHelloHello
	fmt.Println(H)  //Hello WolrdHelloHello

	fmt.Println("!Again \bboi") // \b == backspace  : !Againboi

	abc := "abc"
	l := []byte(abc) //conver abc to byte by l
	fmt.Printf("%s,%s\n", abc, l)
	abc2 := string(l) //converm l to string by abc 2
	fmt.Println(abc2)

	fmt.Println("KING,KONG")
	fmt.Println("King"[0])   //Conver Charactor "King"[0] to  Decimal Value
	fmt.Println(len("Kong")) //Conver How much charater in "Kong" is 4
	fmt.Println("King"[0:3]) //Kin

	// import module strings
	fmt.Println(strings.ToUpper("yut"))                                    //Using  command ToUpper to change Charater all To Upper 	"YUT"
	fmt.Println(strings.ToLower("YUT"))                                    //Using  command ToLower to change Charater all To Lower	"yut"
	fmt.Println(strings.HasPrefix("Hello Wolrd", "Hello"))                 //Using for check TexT  true
	fmt.Println(strings.HasSuffix("Hello Wolrd", "Wolrd"))                 // Using for check TexT  true
	fmt.Println(strings.HasSuffix("Hello Wolrd", "Hello"))                 // Using for check TexT  false
	fmt.Println(strings.Replace("Hello Golang Golang", "Golang", "HI", 1)) //Using for Replace Text
	fmt.Println(strings.Replace("Hello Golang Golang", "Golang", "HI", 2))

	var sb strings.Builder
	fmt.Println("This is String Builder :", sb.String()) // Hello
	fmt.Println(sb.Cap())                                // 0
	fmt.Println(sb.Len())                                // 0
	sb.WriteString("Hello")
	fmt.Println("This is String Builder :", sb.String()) // Hello Golang
	fmt.Println(sb.Cap())                                //8  capacity
	fmt.Println(sb.Len())                                //6	Length of Text
	sb.Grow(100)                                         // 100 + (cap * 2) = 116
	fmt.Println(sb.Cap())                                // 116
	fmt.Println(sb.Len())                                //6
	sb.Reset()                                           // RESET sb
	sb.Cap()                                             //Empty
	sb.Len()                                             //Empty
	fmt.Println(sb.String())
	sb.Write([]byte{65, 65, 65}) //อักษรASCII
	fmt.Println(sb.String())     //AAA

	//String Conver
	x := 12
	y := strconv.Itoa(x) //แปลงตัวเลขเป็นข้อความ
	fmt.Println(y)
	fmt.Printf("%T\n", y)     //string
	z, err := strconv.Atoi(y) //แปลงข้อความเป็นตัวเลข
	fmt.Println("This is Sting conv", z)
	if err != nil {
		fmt.Println("This Text have Any String in There")
	}

	//constants ค่าคงที่
	const const1 = 12
	fmt.Println("This is constants number : ", const1)
	//You can"t Chang const1 because const1 is constant if u make new assign that might be error

	//For loop
	sum := 1         //ตัวแปร
	for sum < 1000 { //เงื่อนไข คือ ถ้า sum น้อยกว่า 1000 ให้ทำต่อ
		sum += sum //คำสั่ง  sum = sum + sum
	}
	fmt.Println(sum)
	//ตัวแปร;เงื่อนไข;คำสั่ง
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//for len
	var forlen1 string = "hello"
	//for len is Lenght of forlen1
	for i := 0; i < len(forlen1); i++ {
		fmt.Println(i)
	}

	var forrange [3]int
	forrange[0] = 12
	forrange[1] = 20
	forrange[2] = 13
	//คล้ายfor len
	for leN, v := range forrange {
		fmt.Println(v)   // range v = แสดงค่าทั้งหมดใน array forrange
		fmt.Println(leN) // range Len = แสดงว่าในarray มีกี่ตัวนับตั้งแต่ 0
	}
	//Loop goto
	GOTO := 0
lable: //ป้าย
	if GOTO < 10 {
		GOTO++
		fmt.Println(GOTO)
		goto lable //ไปที่ป้าย lable
	}

	// Data structure
	// Array  เก็บข้อมูลหลายตัวไว้ในตัวแปรเดียวต้องระบุขนาดเสมอ เช่น var a [5]int
	// Slice	คล้ายArrayแต้สามารถเปลี่ยนขนาดได้ตามใจ เช่น var a []int
	// Map	ข้อมูล map ที่คล้ายๆ dictionary ที่สามารถเก็บ key คู่กับ value ตรงๆไว้ได้ (จากเดิมต้องระบุเป็นตำแหน่ง) เช่น map[string]int
	// Struct 	เก็บตัวแปรหลายตัวไว้ในStruct คล้ายๆเก็บ Object
	{
		//ARRAY
		//เลขของข้อมูลจะจาก 0 เสมอ
		var myArray [3]int // An array of 3 integers
		myArray[0] = 10    // Assign values
		//[]คือตำแหน่ง
		myArray[1] = 20
		myArray[2] = 30

		for i := 0; i < len(myArray); i++ {
			fmt.Println(myArray[i])
		}
		//หากต้องการเปลี่ยนข้อมูล
		myArray[0] = 100
		myArray[1] = 200
		myArray[2] = 300
		//ไม่สามารถเปลี่ยนแบบนี้ได้
		//myArray = [10, 20, 30]
	}
	//SLICE
	{
		mySlice := []int{10, 20, 30, 40, 50} // A slice of integers

		fmt.Println(mySlice)      // Output: [10 20 30 40 50]
		fmt.Println(len(mySlice)) // Length of the slice: 5
		fmt.Println(cap(mySlice)) // Capacity of the slice: 5

		// Slicing a slice
		subSlice := mySlice[1:3] // Slice from index 1 to 2
		fmt.Println(subSlice)    // Output: [20 30]
	}
	{
		myMap := make(map[string]int)

		// Add key-value pairs to the map
		myMap["apple"] = 5
		myMap["banana"] = 10
		myMap["orange"] = 8

		// Access and print a value for a key
		fmt.Println("Apples:", myMap["apple"])

		// Update the value for a key
		myMap["banana"] = 12

		// Delete a key-value pair
		delete(myMap, "orange")

		// Iterate over the map
		for key, value := range myMap {
			fmt.Printf("%s -> %d\n", key, value)
		}

		// Checking if a key exists
		val, ok := myMap["pear"]
		if ok {
			fmt.Println("Pear's value:", val)
		} else {
			fmt.Println("Pear not found inmap")
		}
	}


	//Struct
	//สร้าง type ชื่อ student
	type student struct {
		name string  //มีname เป็น string
		age  int    //มี age เป็น int
	}

	{
		//สร้างตัวแปร acc มี type เป็น student
		var acc student
		// เรียกใช้โดยใส่ชื่อตัวแปรที่เราสร้างแล้วดอทตามด้วยข้อมูลที่อยู่ในStruct
		acc.name = "Acc"
		acc.age = 12
		fmt.Println("Print acc : ",acc)
		fmt.Println("Print acc.name : ",acc.name)
		//สร้างตัวแปร acc มี ข้อมูลเป็นstudentที่ยังไม่กำหนดค่า
		var acc1 = student{}
		acc1.name = "Acc1"
		acc1.age = 15
		fmt.Println("Print acc1 : ",acc1)
		//สร้างตัวแปร acc มี ข้อมูลเป็นstudentที่ยังไม่กำหนดค่า แบบสั้น
		acc2 := student{}
		acc2.name = "Acc2"
		acc2.age = 14
		fmt.Println("Print acc2 : ",acc2)
		//กำหนดในตัวแปรเลย
		var acc3  = student{
			name: "Acc3",
			age: 15,
		}
		fmt.Println("Print acc3 : ",acc3)

		//เหมือนกันทุกอย่างต่างกันแค่วิธีใช้ อยู่ที่เราจะ handle เอง
	}
		

}