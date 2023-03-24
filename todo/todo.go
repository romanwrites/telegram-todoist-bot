package todo

import (
	"log"
	"net/http"
	"strings"
	"tg-function-go/auth"
	"tg-function-go/telegram"
	"tg-function-go/todoist"
	"unicode"
)

const InboxProjectName = "Inbox"

var IdeasMapping = map[string]bool{
	"idea":  true,
	"ideas": true,
}

var MaybeMapping = map[string]bool{
	"someday":       true,
	"maybe":         true,
	"somedaymaybe":  true,
	"someday-maybe": true,
	"someday/maybe": true,
}

func addTaskToProject(task string, projectId string, todoistToken string) {
	todoist.AddNewTaskToProject(task, projectId, todoistToken)
}

func HandleToDoBotRequest(w http.ResponseWriter, r *http.Request, validUserName string, todoistToken string) {
	update := telegram.GetUpdate(r)
	message, err := telegram.GetMessageOrEditedMessage(update)
	if err != nil {
		log.Println(err)
		return
	}

	if auth.IsAuthorisedUser(message, validUserName) {
		if message.Text == "" {
			log.Println("Empty message text")
			return
		}

		project, text := extractProjectFromText(message.Text)
		project = autoCorrectProject(project)
		text = strings.ReplaceAll(text, "\n", " ")
		projectId, err := todoist.GetProjectByName(project, todoistToken)
		if err != nil {
			log.Println(err)
			log.Println("Looking for default project ", InboxProjectName)
			projectId, err = todoist.GetProjectByName(InboxProjectName, todoistToken)
			if err != nil {
				log.Println(err)
				return
			}
			log.Println("Adding to default project ", InboxProjectName)
			addTaskToProject(project+" "+text, projectId, todoistToken)
		} else {
			addTaskToProject(text, projectId, todoistToken)
		}
	}
}

func autoCorrectProject(project string) string {
	if isMapContainsKey(IdeasMapping, project) {
		return "Ideas"
	}
	if isMapContainsKey(MaybeMapping, project) {
		return "Someday/Maybe"
	}
	return project
}

func isMapContainsKey(m map[string]bool, key string) bool {
	_, ok := m[strings.ToLower(key)]
	return ok
}

func extractProjectFromText(text string) (projectName string, textOfTask string) {
	if strings.HasPrefix(text, "#") || strings.HasPrefix(text, "/") {
		startIdx := 1
		isPrefix := true
		for i := 1; i < len([]rune(text)); i++ {
			r := []rune(text)[i]
			if isPrefix == false && unicode.IsSpace(r) {
				projectName := string([]rune(text)[startIdx:i])
				textOfTask := string([]rune(text)[i+1:])

				if projectName == "" || strings.HasPrefix(projectName, "#") || strings.HasPrefix(projectName, "/") {
					projectName = InboxProjectName
				}

				return projectName, textOfTask
			} else if isPrefix == true && (unicode.IsSpace(r) || r == '#' || r == '/') {
				startIdx++
			} else {
				isPrefix = false
			}
		}
	}

	return InboxProjectName, text
}
