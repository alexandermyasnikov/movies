package parser

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"gitlab.com/amyasnikov/movies/common"
	"golang.org/x/net/html"
	"golang.org/x/net/html/charset"
)

var (
	host           = "https://imdb.com"
	DefaultOptions = Options{
		MediaIndexLimit:     200,
		Lang:                "en",
		TimeoutMilliSeconds: 100,
	}
)

type Movie = common.Movie

type Options struct {
	MediaIndexLimit     int
	Lang                string
	TimeoutMilliSeconds int
}

type Parser struct {
	opts Options
}

func NewParser(opts Options) Parser {
	return Parser{opts: opts}
}

func (p Parser) Ids(limit int) <-chan string {
	c := make(chan string, 1)
	go func() {
		for start, offset := 1, -1; offset != 0 && start < limit; start += offset {
			offset = p.search(start, c)
		}
		close(c)
	}()
	return c
}

func (p Parser) Movies(limit int) <-chan Movie {
	c := make(chan Movie, 1)
	go func() {
		for id := range p.Ids(limit) {
			movie := p.Movie(id)
			if movie != nil {
				c <- *movie
			}
			time.Sleep(time.Duration(p.opts.TimeoutMilliSeconds) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func (p Parser) Movie(id string) *Movie {
	url := host + "/title/" + id + "/"
	doc, err := p.getDoc(url)

	if err != nil {
		return nil
	}

	movie := &Movie{
		Id:      id,
		Name:    p.getXpath(doc, `//div[@id="ratingWidget"]/p/strong`),
		Genres:  p.getXpathList(doc, `//h4[contains(text(),'Genres:')]/following-sibling::a`),
		Similar: p.getXpathList(doc, `//div[@class='rec_view']//div[@class='rec_item']/@data-tconst`),
	}

	for i := 1; i <= p.opts.MediaIndexLimit; i++ {
		photos := p.mediaIndex(id, i)
		if len(photos) == 0 {
			break
		}
		movie.Photos = append(movie.Photos, photos...)
	}

	return movie
}

func (p Parser) search(start int, c chan string) int {
	url := host + "/search/title/?title_type=all&num_votes=100000,&view=simple&sort=num_votes,desc&start=" + strconv.Itoa(start)
	doc, err := p.getDoc(url)

	if err != nil {
		return 0
	}

	ids := p.getXpathList(doc, `//div[@class='lister-list']/div/div/a/img/@data-tconst`)

	for _, id := range ids {
		c <- id
	}

	return len(ids)
}

func (p Parser) mediaIndex(id string, page int) []string {
	url := host + "/title/" + id + "/mediaindex?refine=still_frame&page=" + strconv.Itoa(page)
	doc, err := p.getDoc(url)

	if err != nil {
		return []string{}
	}

	photos := p.getXpathList(doc, `//div[@class='media_index_thumb_list']//img/@src`)

	for i, photo := range photos {
		photos[i] = strings.TrimRight(photo, ".jpg")
	}

	return photos
}

func (p Parser) getDoc(url string) (*html.Node, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	if p.opts.Lang != "" {
		req.Header.Set("Accept-Language", p.opts.Lang)
	}

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

func (p Parser) getXpath(doc *html.Node, xpath string) string {
	values := p.getXpathList(doc, xpath)

	if len(values) == 0 {
		return ""
	}

	return values[0]
}

func (p Parser) getXpathList(doc *html.Node, xpath string) []string {
	ret := []string{}

	node := htmlquery.Find(doc, xpath)

	for _, n := range node {
		val := strings.TrimSpace(htmlquery.InnerText(n))
		ret = append(ret, val)
	}

	return ret
}
