package seed

import (
	"context"
	"database/sql"
	"os"

	db "github.com/brkss/gogql/db/sqlc"
	"github.com/google/uuid"
)

type Awareness struct {
	Title 		string
	Content 	string
	Image		string
	FileName 	string
}

var data = []Awareness{
	{
		Title: "Anexiety",
		Content: "",
		Image: "",
		FileName: "./awareness/anxiety.html",
	},
	/*
	{
		Title: "Depression",
		Content: "",
		Image: "",
		FileName: "./awareness/depression.html",
	},
	{
		Title: "Mood Swing",
		Content: "",
		Image: "",
		FileName: "./awareness/mood_swing.html",
	},
	*/
}

func gatheData() ([]*db.CreateAwarenessContentParams){
	
	var results []*db.CreateAwarenessContentParams;

	for _, aw := range data {
		f, err := os.ReadFile(aw.FileName);
		if err != nil {
			panic(err)
		}
		results = append(results, &db.CreateAwarenessContentParams{
			ID: uuid.New().String(),
			Title: aw.Title,
			Content: string(f),
			SurveyID: sql.NullString{
				Valid: false,
				String: "",
			},
		})	
	}

	return results 
}


func SeedAwareness(store db.Store){
	results := gatheData();
	for _, res := range results {
		_, err := store.CreateAwarenessContent(context.Background(), *res);
		if err != nil {
			panic(err);
		}
	}
	return;
}
