package testhelper

import (
	"app/internal/system"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"
)

type AppSuite struct {
	app    *system.App
	cancel context.CancelFunc
}

func (a *AppSuite) Start(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	app, err := system.New(ctx)
	if err != nil {
		t.Error(err)
	}

	a.app = app
	a.cancel = cancel
}

func (a *AppSuite) Stop(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	defer a.cancel()

	err := a.app.Clean(ctx)
	if err != nil {
		t.Error(err)
	}
}

func (a *AppSuite) GetApp() *system.App {
	return a.app
}

func (a *AppSuite) ExecSQL(t *testing.T, sql string) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result := a.app.GormDb.WithContext(ctx).Exec(sql)

	if result.Error != nil {
		t.Error(result.Error)
	}
}

func (a *AppSuite) ExecSQLTestFile(t *testing.T, filename string) {
	abspath, err := filepath.Abs("../testdata/" + filename)
	if err != nil {
		t.Errorf("failed to get %s path", filename)
	}

	b, err := os.ReadFile(abspath)
	if err != nil {
		t.Error("failed to load testdata")
	}

	sql := string(b)
	a.ExecSQL(t, sql)
}

func JsonPrint(data interface{}) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatal("Failed to marshal value")
	}

	fmt.Println(string(val))
}
