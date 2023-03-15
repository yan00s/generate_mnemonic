package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/1makarov/gen-prv-keys/file"
	"github.com/umahmood/mnemonic"
)

func create_mnemonic() string {
	m, err := mnemonic.New(mnemonic.DefaultConfig) // default 128 bits
	if err != nil {
		log.Fatal(err)
	}
	words, err := m.Words()
	if err != nil {
		log.Fatal(err)
	}
	result := strings.Join(words, " ")
	return result
}

func get_result_time(bef_time int64) string {
	result_time := time.Now().UnixMilli()
	result_second := (result_time - bef_time) / 1000
	result_milisec := (result_time - bef_time) % 1000
	result := fmt.Sprintf("result: %v.%v seconds", result_second, result_milisec)
	return result
}

func write_result(mnemonics string) {
	name_file := "result_mnemonics"
	file, err := file.New(name_file, os.O_APPEND|os.O_WRONLY|os.O_CREATE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	if err := file.Write(mnemonics); err != nil {
		log.Fatal(err)
	}

}

func main() {
	mnemonics := []string{}
	count_mnemon := 0

	fmt.Printf("How many mnemonic to create? ")
	fmt.Scanf("%v", &count_mnemon)
	bef_time := time.Now().UnixMilli()
	for i := count_mnemon; i > 0; i-- {
		mnemonics = append(mnemonics, create_mnemonic())
	}
	fmt.Println(get_result_time(bef_time))
	write_result(strings.Join(mnemonics, "\n"))
}

// при генерации просто приват кея 1.000.000 = 30сек на питоне или 60сек на го

//
// How many mnemonic to create? 1000000
// result: 3.439 seconds

// How many mnemonic to create? 10000000
// result: 23.230 seconds
