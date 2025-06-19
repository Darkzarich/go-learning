package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"slices"
	"sort"
	"time"
)

// Put your values here

const (
	KAITEN_API_TOKEN            = "<API_TOKEN>"                    // token for accessing the API, you can get it by following the link https://kaiten.tech/profile/api-key
	KAITEN_COLUMN_ID            = "7195"                           // kaiten column id for tasks in review
	KAITEN_BOARD_ID             = "212"                            // board id
	KAITEN_API_URL              = "https://kaiten.tech/api/latest" // kaiten api url
	KAITEN_BACKEND_COMPONENT_ID = 109
	KAITEN_CARD_URL             = "https://kaiten.tech/space/<space_id>/boards/card" // kaiten card url
)

type TaskProperties struct {
	MergeRequest  string `json:"id_294"`
	MergeRequests string `json:"id_352"`
	Component     []int  `json:"id_21"`
}

type Task struct {
	Id          int            `json:"id"`
	Title       string         `json:"title"`
	LaneId      int            `json:"lane_id"`
	Props       TaskProperties `json:"properties"`
	LastMovedAt time.Time      `json:"last_moved_at"`
}

type Lane struct {
	Id        int     `json:"id"`
	Title     string  `json:"title"`
	SortOrder float32 `json:"sort_order"`
}

func main() {
	tasks, err := fetchTasks()
	if err != nil {
		log.Fatalf("Error fetching tasks: %v\n", err)
	}

	lanes, err := fetchLanes()
	if err != nil {
		log.Fatalf("Error fetching lanes: %v\n", err)
	}

	fmt.Println("🗣️🔊 Kaiten tasks by lane priority (from top to bottom) 📅")
	fmt.Println("====================================")

	sortedTasks := sortTasksByLaneOrder(tasks, lanes)

	index := 0

	for _, task := range sortedTasks {
		timeRelative := formatRelativeTimeInRussian(task.LastMovedAt)

		if slices.Contains(task.Props.Component, KAITEN_BACKEND_COMPONENT_ID) {
			index++

			taskMarkdown := fmt.Sprintf("[%s](%s/%d)", task.Title, KAITEN_CARD_URL, task.Id)

			// A Merge Request field was used
			if task.Props.MergeRequest != "" {
				fmt.Printf("%d) %s [is already in review for %s] :\n* 🔀 %s\n", index, taskMarkdown, timeRelative, task.Props.MergeRequest)
			} else if task.Props.MergeRequests != "" { // A field with many links to merge requests was used
				re := regexp.MustCompile(`\[[^\]]*\]\(([^)]+)\)`) // RegExp for markdown links in the field
				matches := re.FindAllStringSubmatch(task.Props.MergeRequests, -1)

				fmt.Printf("%d) %s [is already in review for %s] :\n", index, taskMarkdown, timeRelative)

				for _, match := range matches {
					fmt.Printf("* 🔀 %s\n", match[1])
				}
			} else {
				fmt.Printf("%d) %s [is already in review for %s] :\n* 🔀 ⚠️ NO MERGE REQUEST ⚠️\n", index, taskMarkdown, timeRelative)
			}
		}
	}
}

func sortTasksByLaneOrder(tasks []Task, lanes []Lane) []Task {
	laneOrder := make(map[int]float32)
	for _, lane := range lanes {
		laneOrder[lane.Id] = lane.SortOrder
	}

	// Sort tasks by lane priority
	sort.Slice(tasks, func(i, j int) bool {
		iOrder := laneOrder[tasks[i].LaneId]
		jOrder := laneOrder[tasks[j].LaneId]

		if iOrder == jOrder {
			return tasks[i].Title < tasks[j].Title
		}

		return iOrder < jOrder
	})

	return tasks
}

func fetchTasks() ([]Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/cards?column_id=%s&archived=false", KAITEN_API_URL, KAITEN_COLUMN_ID), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+KAITEN_API_TOKEN)

	resp, err := processRequest(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var tasks []Task
	if err := json.NewDecoder(resp.Body).Decode(&tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func fetchLanes() ([]Lane, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/boards/%s/lanes", KAITEN_API_URL, KAITEN_BOARD_ID), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+KAITEN_API_TOKEN)

	resp, err := processRequest(req)

	defer resp.Body.Close()

	var lanes []Lane
	if err := json.NewDecoder(resp.Body).Decode(&lanes); err != nil {
		return nil, err
	}

	return lanes, nil
}

func processRequest(req *http.Request) (*http.Response, error) {
	client := new(http.Client)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("%s API returned status %d: %s", req.URL, resp.StatusCode, body)
	}

	return resp, nil
}

func formatRelativeTimeInRussian(t time.Time) string {
	duration := time.Since(t)

	switch {
	case duration < time.Minute:
		return "**меньше минуты** 🟢"
	case duration < time.Hour:
		minutes := int(duration.Minutes())
		return fmt.Sprintf("**%d %s** 🟢", minutes, pluralizeInRussian(minutes, "минуту", "минуты", "минут"))
	case duration < 24*time.Hour:
		hours := int(duration.Hours())
		return fmt.Sprintf("**%d %s** 🟢", hours, pluralizeInRussian(hours, "час", "часа", "часов"))
	case duration < 7*24*time.Hour:
		days := int(duration.Hours() / 24)
		if days == 1 {
			return "**1 день** 🟡"
		}
		return fmt.Sprintf("**%d %s** 🔥", days, pluralizeInRussian(days, "день", "дня", "дней"))
	case duration < 30*24*time.Hour:
		weeks := int(duration.Hours() / (24 * 7))
		return fmt.Sprintf("**%d %s** 🔥🔥", weeks, pluralizeInRussian(weeks, "неделю", "недели", "недель"))
	case duration < 365*24*time.Hour:
		months := int(duration.Hours() / (24 * 30))
		return fmt.Sprintf("**%d %s** 🔥🔥🔥", months, pluralizeInRussian(months, "месяц", "месяца", "месяцев"))
	default:
		years := int(duration.Hours() / (24 * 365))
		return fmt.Sprintf("**%d %s** 🔥🔥🔥🔥", years, pluralizeInRussian(years, "год", "года", "лет"))
	}
}

func pluralizeInRussian(n int, form1, form2, form5 string) string {
	n = n % 100
	if n >= 11 && n <= 19 {
		return form5
	}

	switch n % 10 {
	case 1:
		return form1
	case 2, 3, 4:
		return form2
	default:
		return form5
	}
}
