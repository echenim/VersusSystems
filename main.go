package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type Friend struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
}
type PersonalData struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	UserName string   `json:"username"`
	Email    string   `json:"email"`
	Age      int64    `json:"age"`
	Power    string   `json:"power"`
	Friends  []Friend `json:"friends"`
}

func main() {
	p := GeneratePersonObject()
	fmt.Println("----- Before Tranformsation -----")
	fmt.Println(p)
	p.EmailTransformation()
	p.NameAndPasswordTransformation()

	fmt.Println("----- After Tranformsation write to disk -----")
	OutPutToFile(p)

	fmt.Println("----- After Tranformsation is applied -----")
	fmt.Println(p)
}

func (p *PersonalData) NameAndPasswordTransformation() *PersonalData {
	p.Name = Transformation(p.Name)
	p.UserName = Transformation(p.UserName)
	for i := 0; i < len(p.Friends); i++ {
		p.Friends[i].UserName = Transformation(p.Friends[i].UserName)
	}
	return p
}

func (p *PersonalData) EmailTransformation() *PersonalData {
	p.Email = EmailTransformation(p.Email)
	return p
}

func Transformation(s string) string {
	return strings.Repeat("*", len(s))
}

func EmailTransformation(s string) string {
	x := strings.Split(s, "@")
	x[0] = strings.Repeat("*", len(x[0]))
	return strings.Join(x, "@")
}

func GeneratePersonObject() PersonalData {
	p := PersonalData{}
	p.ID = 123
	p.Name = "Elsa"
	p.UserName = "xXfrozen_princessXx"
	p.Email = "elsa@arendelle.com"
	p.Age = 21
	p.Power = "ice ray"
	p.Friends = generateFriends()

	return p
}

func generateFriends() []Friend {
	f := []Friend{}
	f = append(f, Friend{ID: 234, UserName: "MagicSnowman32"})
	f = append(f, Friend{ID: 456, UserName: "call_me_anna"})
	return f
}

func OutPutToFile(p PersonalData) {
	file, err := json.MarshalIndent(p, "", "")
	if err != nil {
		fmt.Println("Error occured", err)
	}
	err = ioutil.WriteFile("output.json", file, 0o664)
	if err != nil {
		fmt.Println("Error occured while writing to disk", err)
	}
}
