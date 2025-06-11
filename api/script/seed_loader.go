package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"utes-x-api/controller"

	"github.com/joho/godotenv"
)

var baseURL = ""

func init() {
	godotenv.Load("../.env")
	port := os.Getenv("API_PORT")
	if port == "" {
		panic("API_PORT is not set in .env")
	}
	baseURL = "http://localhost:" + port
}

var seedUsers = []controller.UserCreate{
	{
		Email: "alice@example.com", Username: "alice",
	},
	{
		Email: "bob@example.com", Username: "bob",
	},
	{
		Email: "carol@example.com", Username: "carol",
	},
}

// username => []PostCreate
var seedPosts = map[string][]controller.PostCreate{
	"alice": {
		{
			Content: "Hello, world!",
			UserId:  "",
		},
		{
			Content: "Just enjoying a cup of coffee.",
			UserId:  "",
		},
		{
			Content: "Sunny day today ☀️",
			UserId:  "",
		},
	},
	"bob": {
		{
			Content: "First post!",
			UserId:  "",
		},
		{
			Content: "Working on a new project.",
			UserId:  "",
		},
	},
	"carol": {
		{
			Content: "Good morning everyone.",
			UserId:  "",
		},
		{
			Content: "Had a great workout today.",
			UserId:  "",
		},
		{
			Content: "Reading an interesting book.",
			UserId:  "",
		},
		{
			Content: "Looking forward to the weekend!",
			UserId:  "",
		},
	},
}

// username => post index => []ReplyCreate
var seedReplies = map[string]map[int][]controller.ReplyCreate{
	"alice": {
		1: {
			{
				Content: "Thanks for sharing!",
				PostId:  "",
				UserId:  "bob", // あとでIDをセットする
			},
			{
				Content: "I love coffee too!",
				PostId:  "",
				UserId:  "carol",
			},
		},
		2: {
			{
				Content: "Enjoy your coffee ☕",
				PostId:  "",
				UserId:  "bob",
			},
		},
	},
	"bob": {
		1: {
			{
				Content: "Good luck with your project!",
				PostId:  "",
				UserId:  "alice",
			},
			{
				Content: "What kind of project is it?",
				PostId:  "",
				UserId:  "carol",
			},
		},
	},
	"carol": {
		1: {
			{
				Content: "Good morning!",
				PostId:  "",
				UserId:  "alice",
			},
		},
		3: {
			{
				Content: "Me too! Can't wait for the weekend 🎉",
				PostId:  "",
				UserId:  "bob",
			},
		},
	},
}

var client = &http.Client{}

func postUsers(user controller.UserCreate) (controller.User, error) {
	// JSONに変換
	body, err := json.Marshal(user)
	if err != nil {
		return controller.User{}, err
	}

	// POSTリクエスト作成
	req, err := http.NewRequest("POST", baseURL+"/users", bytes.NewReader(body))
	if err != nil {
		return controller.User{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	// 実行
	resp, err := client.Do(req)
	if err != nil {
		return controller.User{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return controller.User{}, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	// レスポンスをパース
	var createdUser controller.User
	if err := json.NewDecoder(resp.Body).Decode(&createdUser); err != nil {
		return controller.User{}, err
	}

	return createdUser, nil
}

func postPosts(post controller.PostCreate) (controller.Post, error) {
	body, err := json.Marshal(post)
	if err != nil {
		return controller.Post{}, err
	}

	req, err := http.NewRequest("POST", baseURL+"/posts", bytes.NewReader(body))
	if err != nil {
		return controller.Post{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return controller.Post{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return controller.Post{}, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var createdPost controller.Post
	if err := json.NewDecoder(resp.Body).Decode(&createdPost); err != nil {
		return controller.Post{}, err
	}

	return createdPost, nil
}

func postReplies(reply controller.ReplyCreate) (controller.Reply, error) {
	body, err := json.Marshal(reply)
	if err != nil {
		return controller.Reply{}, err
	}

	req, err := http.NewRequest("POST", baseURL+"/replies", bytes.NewReader(body))
	if err != nil {
		return controller.Reply{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return controller.Reply{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return controller.Reply{}, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var createdReply controller.Reply
	if err := json.NewDecoder(resp.Body).Decode(&createdReply); err != nil {
		return controller.Reply{}, err
	}

	return createdReply, nil
}

func loadSeedData() error {
	createdUsers := make(map[string]controller.User) // username => User
	for _, user := range seedUsers {
		createdUser, err := postUsers(user)
		if err != nil {
			fmt.Printf("Error creating user %s: %v\n", user.Username, err)
			return err
		}
		createdUsers[user.Username] = createdUser
	}
	for username, posts := range seedPosts {
		for i, post := range posts {
			user, ok := createdUsers[username]
			if !ok {
				return fmt.Errorf("user %s not found in created users", username)
			}
			post.UserId = user.Id // Set the UserId to the created user's ID
			createdPost, err := postPosts(post)
			if err != nil {
				fmt.Printf("Error creating post for user %s: %v\n", username, err)
				return err
			}
			seedPosts[username][i].UserId = createdPost.Id // Set the ID for later use

			// Set replies for this post
			if replies, ok := seedReplies[username][i]; ok {
				for _, reply := range replies {
					reply.PostId = createdPost.Id
					reply.UserId = createdUsers[reply.UserId].Id // Set the UserId to the created user's ID
					_, err := postReplies(reply)
					if err != nil {
						fmt.Printf("Error creating reply for post %d of user %s: %v\n", i+1, username, err)
						return err
					}
				}
			}
		}
	}
	return nil
}

func main() {
	if err := loadSeedData(); err != nil {
		fmt.Printf("Error loading seed data: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Seed data loaded successfully!")
}
