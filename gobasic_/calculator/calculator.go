package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
func getInput(promt string){
	fmt.Printf("%v", promt)
	input,_:=reader.ReadString('\n')
	value,err:=strconv.ParseFloat(strings.TrimSpace(input),64)
	if err != nil{
		message, _:=fmt.Scanf("%v must number",&promt)
		panic(message)
	}
}
func main() {
	
}