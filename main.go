package main

import (
	"github.com/joho/godotenv"
	"github.com/xanzy/go-gitlab"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	git, err := gitlab.NewClient(os.Getenv("GITLAB_TOKEN"), gitlab.WithBaseURL(os.Getenv("GITLAB_API_URL")))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	source := os.Getenv("GITLAB_COMPARE_SOURCE")
	target := os.Getenv("GITLAB_COMPARE_TARGET")

	opt := gitlab.ListProjectsOptions{}
	projects, _, err := git.Projects.ListUserProjects(4417440, &opt)
	for _, project := range projects {
		projectId := project.ID
		log.Println(project.Name)

		compares, _, _ := git.Repositories.Compare(projectId, &gitlab.CompareOptions{
			From:     gitlab.String(source),
			To:       gitlab.String(target),
			Straight: gitlab.Bool(true),
		})
		log.Println(compares)
	}
}
