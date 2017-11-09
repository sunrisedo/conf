package conf

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const middle = "========="

type Config struct {
	Mymap  map[string]string
	strcet string
}

func NewConfig(path string) *Config {
	c := new(Config)
	c.Mymap = make(map[string]string)
	if path == "" {
		return c
	}

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		s := strings.TrimSpace(string(b))
		//fmt.Println(s)
		if strings.Index(s, "#") == 0 {
			continue
		}

		n1 := strings.Index(s, "[")
		n2 := strings.LastIndex(s, "]")
		if n1 > -1 && n2 > -1 && n2 > n1+1 {
			c.strcet = strings.TrimSpace(s[n1+1 : n2])
			continue
		}

		if len(c.strcet) == 0 {
			continue
		}
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}

		frist := strings.TrimSpace(s[:index])
		if len(frist) == 0 {
			continue
		}
		second := strings.TrimSpace(s[index+1:])

		pos := strings.Index(second, "\t#")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " #")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, "\t//")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " //")
		if pos > -1 {
			second = second[0:pos]
		}

		if len(second) == 0 {
			continue
		}

		key := c.strcet + middle + frist
		c.Mymap[key] = strings.TrimSpace(second)
	}
	return c
}

func (c *Config) Read(node, key string) string {
	if v, ok := c.Mymap[node+middle+key]; ok {
		return v
	}
	return ""
}

func (c *Config) ReadFloat(node, key string) float64 {
	if v, ok := c.Mymap[node+middle+key]; ok {
		var value float64
		fmt.Sscanf(v, "%f", &value)
		return value
	}
	return 0
}
