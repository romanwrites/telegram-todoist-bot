package todoist

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	"unicode/utf8"
)

const (
	maxTitleLength       = 500
	maxDescriptionLength = 16384
)

func GetProjectByName(name string, todoistToken string) (string, error) {
	request, err := http.NewRequest("GET", "https://api.todoist.com/rest/v2/projects", nil)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
		return err.Error(), err
	}

	request.Header.Add("Authorization", "Bearer "+todoistToken)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
		return err.Error(), err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error while closing the response body:", err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
		return err.Error(), err
	}

	var projects []Project
	err = json.Unmarshal(body, &projects)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
		return err.Error(), err
	}

	for project := range projects {
		if strings.EqualFold(projects[project].Name, name) {
			return projects[project].Id, nil
		}
	}

	err = errors.New("project " + name + "not found")
	return err.Error(), err
}

func AddNewTaskToProject(content string, projectId string, todoistToken string) {
	task := &Task{}
	task.ProjectId = projectId

	switch runeCountInString := utf8.RuneCountInString(content); {
	case runeCountInString > maxDescriptionLength:
		runesIdx := (utf8.RuneCountInString(content) - 1) / 2
		AddNewTaskToProject(string([]rune(content)[:runesIdx]), todoistToken, projectId)
		AddNewTaskToProject(string([]rune(content)[runesIdx:]), todoistToken, projectId)
	case runeCountInString > maxTitleLength:
		task.Content = string([]rune(content)[:maxTitleLength])
		task.Description = string([]rune(content)[maxTitleLength:])
	default:
		task.Content = content
	}

	taskJson, err := json.Marshal(task)
	if err != nil {
		log.Println("Error while marshalling the task:", err)
		return
	}

	request, err := http.NewRequest("POST", "https://api.todoist.com/rest/v2/tasks",
		bytes.NewBuffer(taskJson))
	request.Header.Add("Authorization", "Bearer "+todoistToken)
	request.Header.Add("Content-Type", "application/json")
	out, err := httputil.DumpRequestOut(request, true)
	if err != nil {
		return
	}
	log.Println("Request: " + string(out))

	client := &http.Client{}
	resp, err := client.Do(request)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println("Response: " + string(body))
}
