package main

func main() {
	n , length := FullName("Mar", "Mallo")

}

func FullName(f, l string) (string, int){
	full:= f + " " + l
	length := len(full)
	return full, length
}