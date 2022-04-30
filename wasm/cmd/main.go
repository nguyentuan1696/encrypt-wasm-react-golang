package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"syscall/js"
	"time"
)

func encrypt(this js.Value, inputs []js.Value) interface{} {
	inputData := inputs[0].String()
	time := strconv.FormatInt(time.Now().Unix(), 10)
	encryptMD5 := md5.Sum([]byte("khonggica" + inputData + time))
	encryptAndroid := sha256.Sum256([]byte(string(encryptMD5[:]) + time))
	stringEncryptAndroid := hex.EncodeToString(encryptAndroid[:])

	mapResult := map[string]string{"data": inputData, "data_encrypt": stringEncryptAndroid, "time": time}

	byteResult, err := json.Marshal(mapResult)
	if err != nil {
		fmt.Println("Can not marshal json: \n", err)
	}

	stringResult := string(byteResult)

	return stringResult

}

func main() {
	c := make(chan string)
	js.Global().Set("encrypt", js.FuncOf(encrypt))

	<-c
}
