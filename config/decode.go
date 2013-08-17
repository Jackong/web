/**
 * User: Jackong
 * Date: 13-8-17
 * Time: 上午9:34
 */
package config

import (
	"io/ioutil"
	"fmt"
	"os"
)

func read(fileName string) []byte {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
		return nil
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
		return nil
	}
	return data
}
