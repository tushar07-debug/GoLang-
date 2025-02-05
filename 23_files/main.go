package main

import (
	"os"
)

func main() {
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err) //log the error
	// }
	// defer f.Close()
	// fileinfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(fileinfo.Name())
	// fmt.Println(fileinfo.Size())
	// fmt.Println(fileinfo.IsDir())
	// fmt.Println(fileinfo.Mode())
	// fmt.Println(fileinfo.ModTime())
	// fmt.Println(fileinfo.Sys())
	// fmt.Println()

	//readfile
	// f, err = os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// data := make([]byte, 10)
	// n, err := f.Read(data)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(n,data)
	// fmt.Println(string(data))

	// avoid this using ReadFile fro large rfile like gb and mb
	// f, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(f))

	// dir,err:=os.Open("../")
	// if err!=nil{
	//     panic(err)
	// }
	// defer dir.Close()
	// files,err:=dir.Readdir(-1)
	// // if err!=nil{
	// //     panic(err)
	// // }
	// // fmt.Println(files)
	// for _,f:=range files{
	//     fmt.Println(f.Name(), f.IsDir())
	// }

	//write file
	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// // n, err := f.WriteString("hello world")
	// // if err != nil {
	// // 	panic(err)
	// // }
	// // fmt.Println(n)
	// f.WriteString("hello ji kese ho sarre ")
	// f.WriteString("\n")
	// bytes := []byte("hello ji kese ho sarre ")
	// n, err := f.Write(bytes)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(n)

	//read and write to anothetr file (streaming fasion)
	// f1, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f1.Close()
	// f2, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f2.Close()
	// // io.Copy(f2, f1)
	// reader := bufio.NewReader(f1)
	// writer := bufio.NewWriter(f2)
	// for {
	// 	str, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// 	e := writer.WriteByte(str)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }
	// writer.Flush()

	//delete file
	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}

}
