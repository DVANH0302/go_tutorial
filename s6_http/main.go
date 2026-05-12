package main

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	println(string(bs))
	return len(bs), nil
}

func main() {
	// resp, err := http.Get("https://google.com")
	// if err != nil {
	// 	log.Fatal("Error", err)
	// }

	// lw := logWriter{}

	// io.Copy(lw, resp.Body)

	// defer resp.Body.Close()

	s := square{
		sideLength: 2.0,
	}
	t := triangle{
		height: 2.0,
		base:   3.0,
	}
	printArea(s)
	printArea(t)
}
