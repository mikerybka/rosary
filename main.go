package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type Prayer struct {
	Verses []string `json:"verses"`
}

func (p *Prayer) WriteHTML(w io.Writer) error {
	tmpl := `<div style="">{{ range .Verses }}<div>{{ . }}</div>{{ end }}</div>`
	return template.Must(template.New("prayer").Parse(tmpl)).Execute(w, p)
}

var SignOfTheCross = Prayer{enSignOfTheCross}
var enSignOfTheCross = []string{
	"In the name of the Father,",
	"And of the Son,",
	"And of the Holy Spirit.",
	"Amen.",
}

var ApostlesCreed = Prayer{enApostlesCreed}
var enApostlesCreed = []string{
	"I believe in God,",
	"the Father almighty,",
	"Creator of heaven and earth,",
	"and in Jesus Christ, his only Son, our Lord,",
	"who was conceived by the Holy Spirit,",
	"born of the Virgin Mary,",
	"suffered under Pontius Pilate,",
	"was crucified, died and was buried;",
	"he descended into hell;",
	"on the third day he rose again from the dead;",
	"he ascended into heaven,",
	"and is seated at the right hand of God the Father almighty;",
	"from there he will come to judge the living and the dead.",
	"I believe in the Holy Spirit,",
	"the holy Catholic Church,",
	"the communion of saints,",
	"the forgiveness of sins,",
	"the resurrection of the body,",
	"and life everlasting.",
	"Amen.",
}

var HailMary = Prayer{enHailMary}
var enHailMary = []string{
	"Hail Mary, full of grace,",
	"The Lord is with thee.",
	"Blessed art thou amongst women,",
	"And blessed is the fruit of thy womb, Jesus.",
	"Holy Mary, Mother of God,",
	"Pray for us sinners,",
	"Now and at the hour of our death.",
	"Amen.",
}

var OurFather = Prayer{enOurFather}
var enOurFather = []string{
	"Our Father, who art in heaven,",
	"hallowed be thy name;",
	"thy kingdom come;",
	"thy will be done,",
	"on earth as it is in heaven.",
	"Give us this day our daily bread;",
	"and forgive us our trespasses,",
	"as we forgive those who trespass against us;",
	"and lead us not into temptation,",
	"but deliver us from evil.",
	"Amen.",
}

var GloryBe = Prayer{enGloryBe}
var enGloryBe = []string{
	"Glory be to Father,",
	"And to the Son,",
	"And to the Holy Spirit.",
	"As it was in the beginning,",
	"Is now, and ever shall be,",
	"World without end.",
	"Amen.",
}

var FatimaPrayer = Prayer{enFatimaPrayer}
var enFatimaPrayer = []string{
	"O my Jesus,",
	"Forgive us our sins,",
	"Save us from the fires of hell,",
	"Lead all souls to heaven,",
	"Especially those who are in most need of thy mercy.",
}

var HailHolyQueen = Prayer{enHailHolyQueen}
var enHailHolyQueen = []string{
	"Hail, Holy Queen,",
	"Mother of Mercy",
	"Our life, our sweetness and our hope,",
	"To thee do we cry,",
	"Poor banished children of Eve;",
	"To thee do we send up our sighs,",
	"Mourning and weeping in this vale of tears;",
	"Turn, then, most gracious Advocate,",
	"Thy eyes of mercy toward us,",
	"And after this, our exile,",
	"Show unto us the blessed fruit of thy womb, Jesus.",
	"O clement, O loving, O sweet Virgin Mary!",
	"Pray for us, O holy Mother of God,",
	"That we may be made worthy of the promises of Christ.",
	"Amen.",
}

var joyfulMysteries = []Mystery{
	{
		Name:       "The First Joyful Mystery: Annunciation",
		Verse:      "Luke 1:28",
		Dedication: "For the love of humility",
		Text:       "And when the angel had come to her, he said, \"Hail, full of grace, the Lord is with thee. Blexsed art thou among women.\"",
	},
	{
		Name:       "The Second Joyful Mystery: Visitation",
		Verse:      "Luke 1:41-42",
		Dedication: "For charity towards my neighbour",
		Text:       "Elizabeth was filled with the Holy Spirit, and cried out witha loud voice, \"Blessed art thou among women and blessed is the fruit of thy womb!",
	},
	{
		Name:       "The Third Joyful Mystery: Birth of Jesus",
		Verse:      "Luke 2:7",
		Dedication: "For the spirit of poverty",
		Text:       "She brought forth her firstborn son, and wrapped him in swaddling clothes, and laid him in a manger, because there was no room for them in the inn.",
	},
	{
		Name:       "The Fourth Joyful Mystery: Presentation",
		Verse:      "Luke 2:22-23",
		Dedication: "For the virtue of obedience",
		Text:       "Whem the days of her purification were fulfillled according to the Law of Moses, they took him up to Jerusalem to present him to the Lord.",
	},
	{
		Name:       "The Fifth Joyful Mystery: Finding the Child Jesus in the Temple",
		Verse:      "Luke 2:46",
		Dedication: "For the virtue of piety",
		Text:       "After three days ... they found him in the temple, sittin gin the midst of the teachers, listening to them and asking them questions.",
	},
}
var luminouslMysteries = []Mystery{
	{
		Name:       "The First Luminous Mystery: Baptism of Jesus",
		Verse:      "Matthew 3:16",
		Dedication: "For submission to God's will",
		Text:       "When Jesus had been baptized, he immediately came up from the water. And behold, the heavens were opened to him, and he saw the Spirit of God descending as a dove and coming upon him.",
	},
	{
		Name:       "The Second Luminous Mystery: Wedding Feast of Cana",
		Verse:      "John 2:3-5",
		Dedication: "For devotion to Mary",
		Text:       "And the wine having run short, the mother of Jesus said to him, \"They have no wine.\" ... His mother said to the attendants, \"Do whatever he tells you.\"",
	},
	{
		Name:       "The Third Luminous Mystery: Proclamation fo the Kingdom of God",
		Verse:      "Mark 1:14-15",
		Dedication: "For the grace of conversion",
		Text:       "Jesus cam ein to Galilee, preaching the gospel of the kingdom of God, \"The time is fulfilled, and the kingdom of God is at hand. Repend and believe in the gospel.\"",
	},
	{
		Name:       "The Fourth Luminous Mystery: Transfiguration",
		Verse:      "Matthew 17:1-2",
		Dedication: "For holy fear of God",
		Text:       "Jesus took Peter, James and his brother John, and led them up a high mountain by themselves, and he was transfigured before them. His face shone as the sun, and his garments became white as snow.",
	},
	{
		Name:       "The Fifth Luminous Mystery: Institution of the Holy Eucharist",
		Verse:      "Luke 22:19",
		Dedication: "For thanksgiving to God",
		Text:       "Having taken bread, he gave thanks and broke, and gave it to them, saying, \"This is my body, whis is being given for you; do this in rememberance of me.\"",
	},
}
var sorrowfulMysteries = []Mystery{
	{
		Name:       "The First Sorrowful Mystery: Agony in the Garden",
		Verse:      "Luke 22:43-44",
		Dedication: "For true contrition for sin",
		Text:       "Falling into an agony he prayed the more earnestly. And his sweat became as drops of blood running down upon the ground.",
	},
	{
		Name:       "The Second Sorrowful Mystery: Scourgin at the Pillar",
		Verse:      "John 19:1",
		Dedication: "For the virtue of purity",
		Text:       "Pilate, them, took Jesus and had him scourged.",
	},
	{
		Name:       "The Third Sorrowful Mystery: Crowning with Thorns",
		Verse:      "Matthew 27:28-29",
		Dedication: "For moral courage",
		Text:       "They stripped him and put on him a scarlet cloak; and plaiting a crown of throns, they put it upon his head, and a reed into his right hand.",
	},
	{
		Name:       "The Fourth Sorrowful Mystery: Carrying of the Cross",
		Verse:      "John 19:17",
		Dedication: "For the virtue of patience",
		Text:       "And bearing the cross for himself, we went forth to the place called the Skull, in Hebrew, Golgotha.",
	},
	{
		Name:       "The Fifth Sorrowful Mystery: Crucifixion",
		Verse:      "Like 23:46",
		Dedication: "For final perseverance",
		Text:       "Jesus cried out witha loud voice and said, \"Father, into thy hands I commend my spirit.\" And having said this, he expired.",
	},
}
var gloriousMysteries = []Mystery{
	{
		Name:       "The First Glorious Mystery: Resurrection",
		Verse:      "Mark 16:6",
		Dedication: "For the virtue of faith",
		Text:       "Do not be terrified. You are looking for Jesus of Nazareth, who was crucified. He has risen, he is not here. Behold the place where they laid him.",
	},
	{
		Name:       "The Second Glorious Mystery: Ascension",
		Verse:      "Mark 16:19",
		Dedication: "For the virtue of hope",
		Text:       "So then the Lord, after he had spoken to them, was taken up into heaven, and sits at the right hand of God.",
	},
	{
		Name:       "The Third Glorious Mystery: Descent of the Holy Spirit",
		Verse:      "Acts 2:4",
		Dedication: "For the virtue of love",
		Text:       "They were all filled with the Holy Spirit and began to speak in foreign tongues, even as the Holy Spirit prompted them to speak.",
	},
	{
		Name:       "The Fourth Glorious Mystery: Assumption of the blessed Virgin Mary",
		Verse:      "Revelation 11:19",
		Dedication: "For a happy death",
		Text:       "And the temple of God in heaven was opened, and there was seen the ark of his covenant.",
	},
	{
		Name:       "The Fifth Glorious Mystery: Coronation of the Blessed Virgin Mary",
		Verse:      "Revelation 12:1",
		Dedication: "For eternal salvation",
		Text:       "A great sign appeared in heaven: a woman clothed with the sun, and the moon was under her feet, and upon her head a crown of twelve stars.",
	},
}

func todaysMysteries() []Mystery {
	now := time.Now()
	switch now.Weekday() {
	case time.Monday, time.Saturday:
		return joyfulMysteries
	case time.Tuesday, time.Friday:
		return sorrowfulMysteries
	case time.Wednesday, time.Sunday:
		return gloriousMysteries
	case time.Thursday:
		return luminouslMysteries
	default:
		panic("unreachable")
	}
}

type Mystery struct {
	Name       string
	Verse      string
	Text       string
	Dedication string
}

func (m *Mystery) Prayer() *Prayer {
	return &Prayer{
		Verses: []string{
			m.Name,
			m.Dedication,
			m.Verse,
			m.Text,
		},
	}
}

func (m *Mystery) Print() {
	fmt.Println(m.Name)
	fmt.Println(m.Dedication)
	fmt.Println(m.Verse)
	fmt.Println(m.Text)
	fmt.Println()
}

func (m *Mystery) String() string {
	return strings.Join([]string{
		m.Name,
		m.Dedication,
		m.Verse,
		m.Text,
	}, "\n")
}

func p(lines []string) {
	for _, l := range lines {
		fmt.Println(l)
	}
	fmt.Println()
}

var of = strings.Join(enOurFather, "\n")
var soc = strings.Join(enSignOfTheCross, "\n")
var ac = strings.Join(enApostlesCreed, "\n")
var gb = strings.Join(enGloryBe, "\n")
var fp = strings.Join(enFatimaPrayer, "\n")
var hhq = strings.Join(enHailHolyQueen, "\n")
var hm = strings.Join(enHailMary, "\n")

func m(i int) *Prayer {
	return todaysMysteries()[i].Prayer()
}

func getPrayer(i int) *Prayer {
	switch i {
	case 0, 78:
		return &SignOfTheCross
	case 1:
		return &ApostlesCreed
	case 2, 8, 22, 36, 50, 64:
		return &OurFather
	case 6, 19, 33, 47, 61, 75:
		return &GloryBe
	case 7:
		return m(0)
	case 20, 34, 48, 62, 76:
		return &FatimaPrayer
	case 21:
		return m(1)
	case 35:
		return m(2)
	case 49:
		return m(3)
	case 63:
		return m(4)
	case 77:
		return &HailHolyQueen
	default:
		return &HailMary
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		fmt.Fprintf(w, `<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Rosary</title>
			<style>
				body {
					margin: 0;
					padding: 0;
					font-family: sans-serif;
					display: flex;
					justify-content: center;
					align-items: center;
					min-height: 100vh;
					text-align: center;
					background-color: #f9f9f9;
				}
		
				.container {
					padding: 1rem;
					max-width: 90%;
				}
		
				h1 {
					font-size: 2rem;
				}
		
				p {
					font-size: 1rem;
					color: #333;
				}
			</style>
		</head>
		<body>
		<div class="container"><a href="/0">Begin</a>
		</div>
	</body>
	</html>
	`)
	})
	http.HandleFunc("/{i}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		i, err := strconv.Atoi(r.PathValue("i"))
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintf(w, `<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Rosary</title>
			<style>
				body {
					margin: 0;
					padding: 0;
					font-family: sans-serif;
					display: flex;
					justify-content: center;
					align-items: center;
					min-height: 100vh;
					text-align: center;
					background-color: #f9f9f9;
				}
		
				.container {
					padding: 1rem;
					max-width: 90%;
				}
		
				h1 {
					font-size: 2rem;
				}
		
				p {
					font-size: 1rem;
					color: #333;
				}
			</style>
		</head>
		<body>
		<div class="container">`)
		getPrayer(i).WriteHTML(w)
		if i < 78 {
			fmt.Fprintf(w, "<p><a href=\"/%d\">Next</a> (%d/79)</p>", i+1, i+1)
		}
		fmt.Fprintf(w, `
		</div>
	</body>
	</html>
	`)
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.ListenAndServe(":"+port, nil)
}

var prayerHTML = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    <style>
        body {
            margin: 0;
            padding: 0;
            font-family: sans-serif;
            display: flex;
            justify-content: center;
            align-items: center;
            min-height: 100vh;
            text-align: center;
            background-color: #f9f9f9;
        }

        .container {
            padding: 1rem;
            max-width: 90%;
        }

        h1 {
            font-size: 2rem;
        }

        p {
            font-size: 1rem;
            color: #333;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>{{.Heading}}</h1>
        <p>{{.Message}}</p>
    </div>
</body>
</html>
`
