package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	sites := []string{"https://eternallybored.org/misc/wget/", "https://habr.com/ru/articles/578744/", "https://itanddigital.ru/"}
	for i, site := range sites {
		count := strconv.Itoa(i)

		f, err := os.Create("site" + count + ".html")
		if err != nil {
			log.Println("os.CreateErr: ", err)
			return
		}
		defer f.Close()

		data, err := getPageContent(site)
		if err != nil {
			return
		}
		_, err = f.Write(data)
		if err != nil {
			log.Println("f.WriteErr: ", err)
			return
		}

	}
}
func getPageContent(url string) ([]byte, error) {
	cl := http.DefaultClient
	cl.Timeout = 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), cl.Timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Println("http.NewRequestWithContextErr: ", err)
		return nil, err
	}
	res, err := cl.Do(req)
	if err != nil {
		log.Println("cl.DoErr: ", err)
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("io.ReadAllErr: ", err)
		return nil, err
	}
	return data, nil
}
