package main

import "testing"

func TestSaveArticle(t *testing.T) {
	blog := New()
	blog.SaveArticle(Article{"Test Title", "Test Body"})
	if blog.Articles[0].Title != "Test Title" {
		t.Errorf("Item was not added")
	}
}

func TestFetchAll(t *testing.T) {
	blog := New()
	blog.SaveArticle(Article{"Test Title", "Test Body"})
	articles := blog.FetchAll()
	if len(articles) == 0 {
		t.Errorf("Fetch all fails")
	}
}
