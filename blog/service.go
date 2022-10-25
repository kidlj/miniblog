package blog

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"sort"
	"strings"
	"time"

	"github.com/gorilla/feeds"
	"github.com/kidlj/blog/templates"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

const (
	BLOG_LISTEN_ADDR = "127.0.0.1:8080"
	BLOG_PREFIX      = "http://127.0.0.1:8080/blog/"
	BLOG_FEED        = "http://127.0.0.1:8080/blog/feed"
)

//go:embed articles/*
var blogFS embed.FS

var markdown = goldmark.New(
	goldmark.WithExtensions(
		meta.Meta,
	),
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

type BlogEntry struct {
	Title   string
	Author  string
	Date    time.Time
	Excerpt string
	Content string
	URL     string
	Path    string
}

func (s *Service) list() ([]*BlogEntry, error) {
	pathName := "articles"
	entries, err := fs.ReadDir(blogFS, pathName)
	if err != nil {
		return nil, err
	}
	blogs := make([]*BlogEntry, 0, len(entries))

	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), "md") {
			blog, err := s.get(e.Name())
			if err != nil {
				return nil, err
			}
			blogs = append(blogs, blog)
		}
	}

	sort.Slice(blogs, func(i, j int) bool {
		return blogs[i].Date.After(blogs[j].Date)
	})

	return blogs, nil
}

func (s *Service) get(path string) (*BlogEntry, error) {
	pathName := fmt.Sprintf("articles/%s", path)
	f, err := fs.ReadFile(blogFS, pathName)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	context := parser.NewContext()
	if err := markdown.Convert([]byte(f), &buf, parser.WithContext(context)); err != nil {
		return nil, err
	}
	metaData := meta.Get(context)
	title := metaData["Title"].(string)
	author := metaData["Author"].(string)
	dateStr := metaData["Date"].(string)
	date, err := time.Parse(templates.DateTimeFormat, dateStr)
	if err != nil {
		return nil, err
	}
	content := buf.String()
	excerpt, _, _ := strings.Cut(content, "\n")
	seg, _, _ := strings.Cut(path, ".")
	url := fmt.Sprintf("/blog/%s", seg)

	blog := &BlogEntry{
		Title:   title,
		Author:  author,
		Date:    date,
		Excerpt: excerpt,
		Content: content,
		URL:     url,
		Path:    seg,
	}

	return blog, nil
}

func (s *Service) feed() (string, error) {
	entries, err := s.list()
	if err != nil {
		return "", err
	}

	now := time.Now()
	feed := &feeds.Feed{
		Title:       "The Blog",
		Link:        &feeds.Link{Href: BLOG_PREFIX},
		Description: "A Go static blog",
		Author:      &feeds.Author{Name: "Jian Li"},
		Created:     now,
	}
	for _, e := range entries {
		item := &feeds.Item{
			Title:       e.Title,
			Link:        &feeds.Link{Href: fmt.Sprintf("%s%s", BLOG_PREFIX, e.Path)},
			Description: e.Excerpt,
			Author:      &feeds.Author{Name: e.Author},
			Created:     e.Date,
		}
		feed.Items = append(feed.Items, item)
	}

	return feed.ToAtom()
}
