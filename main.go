package main

import (
	"crypto/tls"

	"fmt"
	"io/ioutil"
	"os"
	//	"strconv"
	"strings"

	"gopkg.in/gomail"
)

// возвращает расширение файла из строки имени файла с расширением файла
func getExtInFileName(name string) string {
	res := ""
	tmpres := strings.Split(name, ".")
	if len(tmpres) >= 2 {
		res = tmpres[len(tmpres)-1]
	}
	return res
}

// получить все элементы из директории namedir
func getListFileNameFromDirectory(namedir string) []string {
	files, _ := ioutil.ReadDir(namedir) // "./")
	res := make([]string, len(files))
	index := 0
	for _, f := range files {
		res[index] = f.Name()
		index++
	}
	return res
}

// чтение файла с именем namefи возвращение содержимое файла, иначе текст ошибки
func readfilecsv(namef string) string {
	file, err := os.Open(namef)
	if err != nil {
		return "handle the error here"
	}
	defer file.Close()
	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return "error here"
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return "error here"
	}
	return string(bs)
}

//функция чтения конфиг файла для почты : адресСервера;порт;имяПользователя;пароль;
func readcfg(namef string) map[string]string {
	str := readfilecsv(namef)
	vv := strings.Split(str, ";")
	res := make(map[string]string)
	res["serverMail"] = vv[0]
	res["portMail"] = vv[1]
	res["userMail"] = vv[2]
	res["passwdMail"] = vv[3]
	return res
}

func main() {

	nameConfigMail := "mail-config.cfg"

	readcfg(nameConfigMail)

	maskFileForAttach := "xlsx"
	addressFrom := "client.service@kazan.2gis.ru"
	addressTo := "i.saifutdinov@kazan.2gis.ru"
	textSubject := "KAZAN LOG ZVONKOV"

	userMail := "client.service"
	passwdMail := "asdf1234!!"
	serverMail := "mx.2gis.local"
	portMail := 25

	readcfg(nameConfigMail)

	m := gomail.NewMessage()
	m.SetHeader("From", addressFrom)
	m.SetHeader("To", addressTo)
	m.SetHeader("Subject", textSubject)
	m.SetBody("text/html", "Hello !!!!")

	// фильтрация файлов по маске maskFileForAttach
	listFileForAttach := getListFileNameFromDirectory("./")
	filterFileForAttach := make(map[string]bool)
	for i := range listFileForAttach {
		if strings.Contains(getExtInFileName(listFileForAttach[i]), maskFileForAttach) {
			filterFileForAttach[listFileForAttach[i]] = true
		}
	}

	// добавляет вложения в письме
	for key, _ := range filterFileForAttach {
		m.Attach(key)
	}

	d := gomail.NewDialer(serverMail, portMail, userMail, passwdMail)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
