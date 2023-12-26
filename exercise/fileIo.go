package exercise

import (
	"fmt"
	"os"
)

func ReadFile() {
	f, err := os.Open("test.json")
	defer f.Close()

	if err != nil {
		panic(err)
	}

	fino, error := f.Stat()

	if error != nil {
		panic(error)
	}

	fmt.Println(fino.Size())

	contents := make([]byte, fino.Size())
	//var contents [10]byte

	//s := contents[:]

	for {
		n, er := f.Read(contents)
		if er != nil && er.Error() != "EOF" {
			panic(er)
		}

		if n == 0 {
			fmt.Println("No bytes read")
			fmt.Println(len(contents))
			break
		}
	}

	fmt.Println(string(contents))

}
