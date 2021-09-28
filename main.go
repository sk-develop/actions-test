package main

func main() {
	client := connPsql()
	defer client.Close()
	if err := client.watch(); err != nil {
		panic(err)
	}
}
