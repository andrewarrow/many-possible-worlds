package server

type OpenGraph struct {
	Author      string
	Title       string
	Url         string
	Description string
	Image       string
}

func NewOpenGraph(author string) *OpenGraph {
	og := OpenGraph{}
	og.Author = author
	return &og
}
