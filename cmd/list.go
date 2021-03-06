package cmd

import (
	"sort"
	"strings"

	"github.com/tuxagon/yata-cli/yata"
	"github.com/urfave/cli"
)

type listArgs struct {
	sort        string
	all         bool
	showTags    bool
	tag         string
	description string
	format      string
}

func (a *listArgs) Parse(ctx *cli.Context) {
	a.sort = ctx.String("sort")
	a.all = ctx.Bool("all")
	a.showTags = ctx.Bool("show-tags")
	a.tag = ctx.String("tag")
	a.description = ctx.String("description")
	a.format = ctx.String("format")
}

// List returns the list of tasks/todos that have been recorded
func List(ctx *cli.Context) error {
	args := &listArgs{}
	args.Parse(ctx)

	manager := yata.NewTaskManager()
	tasks, err := manager.GetAll()
	handleError(err)

	if args.showTags {
		return displayTags(tasks)
	}

	tasks = yata.FilterTasks(tasks, func(t yata.Task) bool {
		return (args.tag == "" || sliceContains(t.Tags, args.tag)) &&
			(args.description == "" || strings.Contains(t.Description, args.description)) &&
			(args.all || !t.Completed)
	})

	sortTasks(args.sort, &tasks)

	for _, v := range tasks {
		stringer := yata.NewTaskStringer(v, taskStringer(args.format))
		switch v.Priority {
		case yata.LowPriority:
			yata.PrintlnColor("cyan+h", stringer.String())
		case yata.HighPriority:
			yata.PrintlnColor("red+h", stringer.String())
		default:
			yata.Println(stringer.String())
		}
	}

	return nil
}

func sortTasks(sortField string, tasks *[]yata.Task) {
	switch {
	case sortField == "priority":
		sort.Sort(yata.ByPriority(*tasks))
	case sortField == "description":
		sort.Sort(yata.ByDescription(*tasks))
	case sortField == "timestamp":
		sort.Sort(yata.ByTimestamp(*tasks))
	default:
		sort.Sort(yata.ByID(*tasks))
	}
}

func displayTags(tasks []yata.Task) error {
	tagCounts := make(map[string]int)
	for _, v := range tasks {
		for _, t := range v.Tags {
			_, ok := tagCounts[t]
			if !ok {
				tagCounts[t] = 1
			} else {
				tagCounts[t] = tagCounts[t] + 1
			}
		}
	}

	var tags []string
	maxLength := 0
	for k := range tagCounts {
		tags = append(tags, k)
		if len(k) > maxLength {
			maxLength = len(k)
		}
	}
	sort.Strings(tags)

	if len(tags) > 0 {

		for _, k := range tags {
			yata.Printf("%-*s\t%d\n", maxLength, k, tagCounts[k])
		}
	}
	return nil
}

func sliceContains(arr []string, term string) bool {
	for _, v := range arr {
		if v == term {
			return true
		}
	}

	return false
}
