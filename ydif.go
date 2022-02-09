package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func print_help()  {
	println("what file1 is uniq to file2:\t\tcat file1 | ydiff file2")
	println("what file2 is uniq to file1:\t\tcat file1 | ydiff file2 -")
	println("what file1 is uniq to file2:\t\tydiff 1 file1 file2")
	println("what file2 is uniq to file1:\t\tydiff 2 file1 file2")
	println("intersect of file1 and file2:\t\tydiff 3 file1 file2")
	println("union of file1 and file2:\t\tydiff 4 file1 file2")
	os.Exit(1)
}

func uniq(input []string) []string {
	//uniq sorted slice
	output := []string{}
	var prev string
	//output = append(output, input[0])
	for _, cur := range input {
		if cur != prev{
			output = append(output, cur)
		}
		prev = cur
	}
	return output
}

func suniq(input []string) []string {
	sort.Strings(input)
	return uniq(input)
}


func sorted_uniq_filecontent(filename string) []string {
	var finput io.Reader
	if filename == "-" {
		finput = os.Stdin
	} else {
		finput,_=os.Open(filename)
	}
	scanner := bufio.NewScanner(finput)
	content := []string{}
	for scanner.Scan() {
		//trim leading and trailing whitespaces; do not append empty line
		line:=strings.TrimSpace(scanner.Text())
		if len(line)!=0{
			content = append(content, line)
		}
	}
	return suniq(content)
}

func ydiff()  {
	//print help and exit if no 2/3 args passed in
	if len(os.Args)!=4 && len(os.Args)!=3 && len(os.Args)!=2 {
		print_help()
	}

	if  len(os.Args)==2 && ( os.Args[1]=="-h" || os.Args[1]=="--help"){
		print_help()
	}

	//parse args
	mode:="1"
	file1:=""
	file2:=""
	if len(os.Args)==2{
		mode="1"
		file1="-"
		file2=os.Args[1]
	}else if len(os.Args)==3{
		mode="1"
		file1=os.Args[1]
		file2=os.Args[2]
	}else{
		mode=os.Args[1]
		file1=os.Args[2]
		file2=os.Args[3]
	}

	//read two files
	file1_data:=sorted_uniq_filecontent(file1)
	file2_data:=sorted_uniq_filecontent(file2)

	//	diff and print
	if mode=="4"{
		//	union
		tmp:=[]string{}
		tmp=append(tmp, file1_data...)
		tmp=append(tmp, file2_data...)
		for _,line:=range suniq(tmp){
			fmt.Println(line)
		}
	}

	if mode=="3"{
	//	intersect
		for _, val1 := range file1_data {
			for _, val2 := range file2_data {
				if val1 == val2 { fmt.Println(val1);break }
			}
		}
	}


	if mode=="2"{file1_data,file2_data=file2_data,file1_data}
	if mode=="2" || mode=="1"{
		//what file1_data is uniq to file2_data
		for _, val1 := range file1_data {
			flag := true
			for _, val2 := range file2_data {
				if val1 == val2 { flag = false;break }
			}
			if flag { fmt.Println(val1) }
		}
	}
}

func main() {
	ydiff()
}
