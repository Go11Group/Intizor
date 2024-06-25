package main

import (
	"log"
	"net"
	"net/http"
)

var translation = map[string]string{
	"apple":        "olma",
	"banana":       "banan",
	"orange":       "apelsin",
	"grape":        "uzum",
	"watermelon":   "tarvuz",
	"strawberry":   "qulupnay",
	"cherry":       "gilos",
	"pear":         "nok",
	"peach":        "shaftoli",
	"plum":         "olxo'ri",
	"apricot":      "o'rik",
	"lemon":        "limon",
	"lime":         "laym",
	"pineapple":    "ananas",
	"mango":        "mango",
	"papaya":       "papayya",
	"kiwi":         "kivi",
	"pomegranate":  "anor",
	"blueberry":    "ko'kmeva",
	"raspberry":    "malina",
	"blackberry":   "qora smorodina",
	"cranberry":    "klyukva",
	"fig":          "anjir",
	"date":         "xurmo",
	"coconut":      "kokos",
	"melon":        "qovun",
	"guava":        "guava",
	"passionfruit": "marakuya",
	"dragonfruit":  "pitahaya",
	"jackfruit":    "jackfruit",
	"avocado":      "avokado",
	"carrot":       "sabzi",
	"potato":       "kartoshka",
	"tomato":       "pomidor",
	"cucumber":     "bodring",
	"pepper":       "qalampir",
	"onion":        "piyoz",
	"garlic":       "sarimsoq",
	"lettuce":      "salat",
	"spinach":      "ismaloq",
	"cabbage":      "karam",
	"cauliflower":  "gulkaram",
	"broccoli":     "brokkoli",
	"zucchini":     "kabachok",
	"eggplant":     "baqlajon",
	"pumpkin":      "qovoq",
	"corn":         "jo'xori",
	"peas":         "no'xat",
	"bean":         "loviya",
	"celery":       "selderey",
}

type Trans struct{}

func main() {

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	http.Serve(listener, nil)
}

func (t *Trans) Translate(words []string, resp *[]string) {

}
