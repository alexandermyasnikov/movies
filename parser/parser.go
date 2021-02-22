package parser

import (
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"golang.org/x/net/html/charset"
)

var (
	host            = "https://imdb.com"
	mediaIndexLimit = 0
	lang            = "en"
)

type Movie struct {
	Id      string
	Name    string
	Genres  []string
	Similar []string
	Photos  []string
}

type Parser struct {
}

func (parser *Parser) Search(limit int) chan Movie {
	c := make(chan Movie, 1)
	var wg sync.WaitGroup
	go func() {
		for i := 1; i <= limit; i++ {
			wg.Add(1)
			go func(i int) {
				movie := parser.searchOne(i)
				if movie != nil {
					c <- *movie
				}
				wg.Done()
			}(i)
		}
		wg.Wait()
		close(c)
	}()
	return c
}

func (parser *Parser) Movie(id string) *Movie {
	url := host + "/title/" + id + "/"
	doc, err := parser.getDoc(url)

	if err != nil {
		return nil
	}

	movie := &Movie{
		Id:      id,
		Name:    parser.getXpath(doc, `//div[@id="ratingWidget"]/p/strong`),
		Genres:  parser.getXpathList(doc, `//h4[contains(text(),'Genres:')]/following-sibling::a`),
		Similar: parser.getXpathList(doc, `//div[@class='rec_view']//div[@class='rec_item']/@data-tconst`),
	}

	for i := 1; i <= mediaIndexLimit; i++ {
		photos := parser.mediaIndex(id, i)
		if len(photos) == 0 {
			break
		}
		movie.Photos = append(movie.Photos, photos...)
	}

	return movie
}

func (parser *Parser) searchOne(start int) *Movie {
	url := host + "/search/title/?title_type=all&num_votes=100000,&view=simple&sort=num_votes,desc&start=" + strconv.Itoa(start)
	doc, err := parser.getDoc(url)

	if err != nil {
		return nil
	}

	xpath := "//div[@class='lister-list']/div[1]/div/a/img/@data-tconst"
	node := htmlquery.FindOne(doc, xpath)

	if node == nil {
		return nil
	}

	id := htmlquery.InnerText(node)

	return parser.Movie(id)
}

func (parser *Parser) mediaIndex(id string, page int) []string {
	url := host + "/title/" + id + "/mediaindex?refine=still_frame&page=" + strconv.Itoa(page)
	doc, err := parser.getDoc(url)

	if err != nil {
		return []string{}
	}

	photos := parser.getXpathList(doc, `//div[@class='media_index_thumb_list']//img/@src`)

	for i, photo := range photos {
		photos[i] = strings.TrimRight(photo, ".jpg")
	}

	return photos
}

func (parser *Parser) getDoc(url string) (*html.Node, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept-Language", lang)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	r, err := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}

	return html.Parse(r)
}

func (parser *Parser) getXpath(doc *html.Node, xpath string) string {
	values := parser.getXpathList(doc, xpath)

	if len(values) == 0 {
		return ""
	}

	return values[0]
}

func (parser *Parser) getXpathList(doc *html.Node, xpath string) []string {
	ret := []string{}

	node := htmlquery.Find(doc, xpath)

	for _, n := range node {
		val := strings.TrimSpace(htmlquery.InnerText(n))
		ret = append(ret, val)
	}

	return ret
}
