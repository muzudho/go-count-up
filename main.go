package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	// Default working directory
	dwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GCU     | DefaultWorkingDirectory=%s\n", dwd)

	// コマンドライン引数
	wd := flag.String("WorkingDirectory", dwd, "Working directory path.")
	flag.Parse()
	fmt.Printf("GCU     | WorkingDirectory=%s\n", *wd)

	var number int

	fileName := filepath.Join(*wd, "count.log")
	fmt.Printf("GCU     | LogFile=%s\n", *wd)

	// ファイルの存在チェック
	_, err = os.Stat(fileName)
	if os.IsNotExist(err) {
		// 無ければ空ファイル作成
		file, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// 0開始
		number = 0
	} else {
		// ファイルを丸ごと読込
		bytes, err := ioutil.ReadFile(fileName)
		if err != nil {
			panic(err)
		}

		text := string(bytes)

		if text == "" {
			// 0開始
			number = 0
		} else {
			// 多分数字だろう
			num, err := strconv.Atoi(text)
			if err != nil {
				panic(err)
			}

			// 1を足す
			number = num + 1
		}
	}

	fmt.Println(fmt.Sprintf("GCU     | %d", number))

	writeString(fileName, strconv.Itoa(number))
}

func writeString(fileName string, contents string) {
	// 上書き用ファイル
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		// ログのファイル・オープン失敗
		panic(err)
	}
	defer file.Close()

	// 数字を書込
	file.WriteString(contents)
}
