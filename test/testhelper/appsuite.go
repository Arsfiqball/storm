package testhelper

import (
	"app/internal/system"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/Arsfiqball/codec/talker"
	"github.com/spf13/viper"
)

type AppSuite struct {
	app    *system.App
	cancel context.CancelFunc
}

func (a *AppSuite) Start(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	// Use current directory for config
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Search config in current directory with name ".config.yml" (without extension)
	viper.AddConfigPath(currentDir)
	viper.SetConfigName(".config.yml")

	// Read environment variables prefixed with APP_
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	// Set config mode for component testing
	viper.Set("mode", "component_test")

	app, err := system.New(ctx)
	if err != nil {
		t.Fatalf("failed to create app: %s", err.Error())
	}

	a.app = app
	a.cancel = cancel
}

func (a *AppSuite) Stop(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	defer a.cancel()

	// clean only necessary resources in test
	cleaner := talker.Sequential(
		talker.Parallel(
			a.app.Fiber.Clean,
			a.app.Watermill.Clean,
			a.app.Work.Stop,
		),
		a.app.Gorm.Close,
	)

	if err := cleaner(ctx); err != nil {
		t.Fatalf("failed to clean app: %s", err.Error())
	}
}

func (a *AppSuite) GetApp() *system.App {
	return a.app
}

func (a *AppSuite) ExecSQL(t *testing.T, sql string) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result := a.app.Gorm.DB().WithContext(ctx).Exec(sql)

	if result.Error != nil {
		t.Errorf("failed to execute sql: %s", result.Error.Error())
	}
}

func (a *AppSuite) ExecSQLTestFile(t *testing.T, filename string) {
	abspath, err := filepath.Abs("../testdata/" + filename)
	if err != nil {
		t.Errorf("failed to get %s path: %s", filename, err.Error())
	}

	b, err := os.ReadFile(abspath)
	if err != nil {
		t.Errorf("failed to load testdata: %s", err.Error())
		return
	}

	sql := string(b)
	a.ExecSQL(t, sql)
}

func JsonPrint(t *testing.T, data any) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		t.Errorf("failed to marshal json: %s", err.Error())
		return
	}

	fmt.Println(string(val))
}
