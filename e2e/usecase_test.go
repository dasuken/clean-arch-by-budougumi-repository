package e2e

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/noguchidaisuke/clean-arch-by-budougumi-repository/repository"
	"github.com/noguchidaisuke/clean-arch-by-budougumi-repository/usecase"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestUserCase_Find(t *testing.T) {
	okName := "budougumi"
	okEmail := "budougumi@gmail.com"
	type args struct {
		name, email string
	}
	okArgs := args {
		name: okName,
		email: okEmail,
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Correct",
			args: okArgs,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := os.Getenv("MYSQL_USER")
			db, err := sql.Open("mysql",fmt.Sprintf("%s:@(localhost:3306)/budougumi", u))
			if err != nil {
				log.Fatalln(err)
			}
			ctx := context.Background()
			repo := repository.NewRepo(db)
			uc := usecase.NewUserCase(repo)
			got, err := uc.FindUser(ctx, 1)

			if err != nil {
				t.Errorf("want not err, but has error: %v", err)
			}

			if got.ID != 1 {
				t.Errorf("got %v, want: 1", got.ID)
			}
		})
	}
}
