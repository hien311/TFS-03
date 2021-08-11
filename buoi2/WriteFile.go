package main

import (
	"os"
)

func main() {
	//data := []byte("hehehehehe")
	//ioutil.WriteFile("./data/write.txt", data, 1998) // tra ve loi

	f, _ := os.Create("./data/write.txt")
	defer f.Close()

	data2 := []byte("1,2,1,4,5,6,7,2,3,6,87,234")
	f.Write(data2) // f.write tra ve 2 gia tri, so byte da viet vao va loi
}
