package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	//body在退出时需要关闭
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//fmt.Println("Error, status code is ", resp.StatusCode)
		return nil, fmt.Errorf("error, status code is %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}
