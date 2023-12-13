package main

import (
	"fmt"
	"net/http"
)

/*func readVariable() (string, string, string, string) {
	viper.SetConfigName("default")
	viper.SetConfigType("json")
	viper.AddConfigPath("/etc")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	cer := viper.GetString("server.certificate")
	key := viper.GetString("server.key")
	port := viper.GetString("server.port")
	addr := viper.GetString("server.address")
	return cer, key, port, addr
}*/

func main() {
	//cer, key, port, addr := readVariable()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello world")
		if err != nil {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
		fmt.Printf("Number of bytes written: %d", n)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
		return
	}

}
