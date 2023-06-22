package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Post struct {
    Title string `json:"title"`
    Slug string `json:"slug"`
}

type ResponseData struct {
	Data struct {
		User struct {
			Publication struct {
                Posts []Post `json:"posts"`
			} `json:"publication"`
		} `json:"user"`
	} `json:"data"`
}

func FetchPosts() ([]*Post, error) {
    query := map[string]string {
        "query": `
            {
                user(username: "bpreston5393") {
                    publication {
                        posts(page: 0) {
                            title
                            slug
                        }
                    }
                }
            }
        `,
    }

    jsonData, err := json.Marshal(query)
    if err != nil {
        return nil, fmt.Errorf("JSON marshal error: %v", err)
    }

    newRequest, err := http.NewRequest("POST", "https://api.hashnode.com", bytes.NewBuffer(jsonData))
    if err != nil {
        return nil, fmt.Errorf("Error hitting api.hashnode.com: %v", err)
    }
    newRequest.Header.Set("Content-Type", "application/json")
    client := &http.Client{Timeout: time.Second * 5}
    res, err := client.Do(newRequest)
    if err != nil {
        return nil, fmt.Errorf("Error executing request: %v", err)
    }
    data, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return nil, fmt.Errorf("Error reading response body: %v", err)
    }

    var resp ResponseData
    err = json.Unmarshal(data, &resp)
    if err != nil {
        return nil, fmt.Errorf("Error unmarshalling data: %v", err)
    }
    var posts []*Post
    for _, p := range resp.Data.User.Publication.Posts {
        var post Post
        post.Title = p.Title
        post.Slug = p.Slug
        posts = append(posts, &post)
    }
    return posts, nil
}
