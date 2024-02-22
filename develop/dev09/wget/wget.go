package wget

import (
	"bytes"
	"dev09/parse"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

type environment struct {
	link      *url.URL
	outFile   string
	depth     int
	recursive bool
	resources bool
}

func (e *environment) run() error {
	stringLink := e.link.String()
	if e.recursive {
		directory := e.link.Host

		err := os.Mkdir(directory, os.ModePerm)
		if err != nil {
			return err
		}

		visited := make(map[string]struct{})
		s := Site{
			root:      stringLink,
			visited:   visited,
			directory: directory,
		}

		if err = s.downloadSite([]string{stringLink}, e.depth); err != nil {
			return err
		}
	}

	if err := downloadFile(stringLink, e.outFile); err != nil {
		return err
	}

	return nil
}

type Site struct {
	root      string
	visited   map[string]struct{}
	directory string
}

func (s *Site) downloadSite(queue []string, depth int) error {
	discoveredLinks := make([]string, 0)

	for _, v := range queue {
		resp, err := http.Get(v)
		if err != nil {
			return err
		}

		media, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			fmt.Printf("can't parse link type: %s", err.Error())
		}

		extension, err := mime.ExtensionsByType(media)
		if err != nil || len(extension) == 0 {
			fmt.Printf("can't parse link type: %s", err.Error())
			extension = append(extension, "")
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		r := bytes.NewReader(body)
		fileName := path.Join(s.directory, path.Base(resp.Request.URL.Path)+extension[0])
		file, err := os.Create(fileName)
		if err != nil {
			return err
		}

		size, err := io.Copy(file, r)
		if err != nil {
			return err
		}

		fmt.Printf("\nDownloaded a file %s with size %d bytes\n", fileName, size)

		if _, err = r.Seek(0, 0); err != nil {
			return err
		}

		l, err := s.parseLinks(r)
		if err != nil {
			return err
		}

		discoveredLinks = append(discoveredLinks, l...)

		resp.Body.Close()
		err = file.Close()
		if err != nil {
			return err
		}
	}

	if len(discoveredLinks) > 0 {
		depth--
		return s.downloadSite(discoveredLinks, depth)
	}

	return nil
}

func (s *Site) parseLinks(r io.Reader) ([]string, error) {
	res, err := parse.ParseHTML(r)
	if err != nil {
		return nil, err
	}

	links := make([]string, 0)

	for _, v := range res {
		href := v.Href.Host + v.Href.Path
		visited := true

		switch {
		case strings.HasPrefix(href, s.root):
			visited = s.isVisited(href)
		case strings.HasPrefix(href, "/"):
			href = s.root + href
			visited = s.isVisited(href)
		}

		if !visited {
			links = append(links, href)
		}
	}

	return links, nil
}

// isVisited checks if url visited
func (s *Site) isVisited(href string) (visited bool) {
	if _, visited = s.visited[href]; !visited {
		s.visited[href] = struct{}{}
	}

	return visited
}

func downloadFile(link, fileString string) error {
	file, err := os.Create(fileString)
	if err != nil {
		return err
	}

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			req.URL.Opaque = req.URL.Path
			return nil
		},
	}
	resp, err := client.Get(link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Printf("\nDownloaded a file %s\n", link)

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}

func (e *environment) fArgs(args []string) error {
	f := flag.NewFlagSet("wget", flag.ContinueOnError)
	f.StringVar(&e.outFile, "o", "", "Выходной файл")
	f.IntVar(&e.depth, "l", 0, "Максимальное количество ссылок, по которым необходимо перейти при создании загрузки сайта")
	f.BoolVar(&e.recursive, "r", false, "Рекурсивный проход по сайту")
	f.BoolVar(&e.resources, "p", false, "Загружать все файлы с сайта")

	if err := f.Parse(args); err != nil {
		f.Usage()
		return err
	}

	link, err := url.Parse(f.Arg(0))
	if err != nil {
		return err
	}

	e.link = link

	if e.outFile == "" {
		e.outFile = path.Base(e.link.Path)
	}

	return nil
}

func Run(args []string) int {
	var env environment
	if err := env.fArgs(args); err != nil {
		fmt.Println(err)
		return 2
	}

	if err := env.run(); err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}
