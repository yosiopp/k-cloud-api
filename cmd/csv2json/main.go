package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("illegal parameter")
		return
	}
	data := reduce(readCsv(os.Args[1]))
	s, _ := json.MarshalIndent(data, "", "  ")
	fmt.Printf("%v", string(s))
}

func readCsv(fileName string) interface{} {
	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	data := make(map[string]interface{})
	reader := csv.NewReader(fp)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		ln := len(record)
		k := ""
		p := data
		for i := 0; i < ln; i++ {
			if len(record[i]) == 0 {
				break
			}

			k = record[i]
			if n := strings.Index(k, "["); n != -1 {
				idx := k[n+1 : len(k)-1]
				k = k[0:n] + "[]"
				if q, ok := p[k]; ok {
					p = q.(map[string]interface{})
				} else {
					p[k] = make(map[string]interface{})
					p = p[k].(map[string]interface{})
				}
				k = idx
			}

			if q, ok := p[k]; ok {
				p = q.(map[string]interface{})
			} else {
				if len(record[i+1]) != 0 {
					p[k] = make(map[string]interface{})
					p = p[k].(map[string]interface{})
				}
			}
		}
		p[k] = record[ln-1]
	}
	return data
}

func reduce(data interface{}) interface{} {
	_map, ok := data.(map[string]interface{})
	if ok {
		for k := range _map {
			if _child, ok := _map[k].(map[string]interface{}); ok {
				reduce(_child)
			}
			if strings.Contains(k, "[]") {
				if d, ok := _map[k].(map[string]interface{}); ok {
					ary := make([]map[string]interface{}, len(d))
					for i := 0; i < len(d); i++ {
						idx := strconv.Itoa(i)
						ary[i], _ = d[idx].(map[string]interface{})
					}
					_map[k[0:len(k)-2]] = ary
					delete(_map, k)
				}
			}
		}
	}
	return data
}
