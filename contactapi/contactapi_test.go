package contactapi

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestRelation(t *testing.T) {
	client := &ContactApiClient{}
	client.Init("https://contact.app0.env0.luojm.com:9443/query")
	t1 := time.Now()
	relation, err := client.Relation(context.Background(), "LtXJjokUlmvzk5", "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODYzMjUwNjksInVzZXIiOnsiaWQiOiJ2OWNlMHBxNXZ2dnNrNCJ9fQ.PzylwCxuXo3USlQ-oPk-gbonB7fu0sjIFR4LA2myBZBcqcdOzXEbMQo6ZqPbid62mV4qIHeOPJ8SZgdh9-Q9dENoo7aosKPxlU9f_nvKNOh7cclgnqiJvMLqbHDy3r3wR8kS_fDFXAhqMbbchdhLvMhfnqUBKbb9bsxSjR65VXyFLLDdNIpSrnhZhtTkxjmST_vGgAhYvYe-DL1CTKHFabL5IOF692HnTZk5ozEljOFa5Wr8qyLtfAUXGePI7s-oum3UaLcwsb1E9OByjSC2f04BgQLjtJ2wghciatZX5baOMDiFB06Sb6Go0SUOGBmTby6QTleHj4aj5S6CgOyH9Q")
	fmt.Printf("relation(%v), err(%v)\n", relation, err)
	t2 := time.Now()
	diff := t2.Sub(t1)
	fmt.Println(diff)
}

// func TestContact(t *testing.T) {
// 	client := &ContactApiClient{}
// 	client.Init("https://contact.app0.env0.luojm.com:9443/query")
// 	t1 := time.Now()
// 	contact, err := client.Contact(context.Background(), "LtXJjokUlmvzk5", "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODYzMjUwNjksInVzZXIiOnsiaWQiOiJ2OWNlMHBxNXZ2dnNrNCJ9fQ.PzylwCxuXo3USlQ-oPk-gbonB7fu0sjIFR4LA2myBZBcqcdOzXEbMQo6ZqPbid62mV4qIHeOPJ8SZgdh9-Q9dENoo7aosKPxlU9f_nvKNOh7cclgnqiJvMLqbHDy3r3wR8kS_fDFXAhqMbbchdhLvMhfnqUBKbb9bsxSjR65VXyFLLDdNIpSrnhZhtTkxjmST_vGgAhYvYe-DL1CTKHFabL5IOF692HnTZk5ozEljOFa5Wr8qyLtfAUXGePI7s-oum3UaLcwsb1E9OByjSC2f04BgQLjtJ2wghciatZX5baOMDiFB06Sb6Go0SUOGBmTby6QTleHj4aj5S6CgOyH9Q")
// 	fmt.Printf("contact(%v), err(%v)\n", contact, err)
// 	t2 := time.Now()
// 	diff := t2.Sub(t1)
// 	fmt.Println(diff)
// }
