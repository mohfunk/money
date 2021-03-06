package money

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/mohfunk/money/pkg/util"
)

// Config contains important data paths
type Config struct {
	DataDir    string
	DataFile   string
	TradeFile  string
	BudgetFile string
	LogFile    string
}

// NewConfig returns a pointer to an empty Config struct
func NewConfig() *Config {
	return &Config{}
}

// Configure initilizes the app's Config struct, creates the neccassary dirctories and files
func (c *Config) Configure() {
	usr, _ := user.Current()
	dir := usr.HomeDir
	ddir := filepath.Join(dir, ".money")
	assetsfile := filepath.Join(ddir, "assets.json")
	tradefile := filepath.Join(ddir, "trades.json")
	budgetfile := filepath.Join(ddir, "budget.json")
	baseAssetfile := filepath.Join(dir, "go/src/github.com/mohfunk/money/base/assets.json")
	baseTradefile := filepath.Join(dir, "go/src/github.com/mohfunk/money/base/trades.json")
	baseBudgetfile := filepath.Join(dir, "go/src/github.com/mohfunk/money/base/budget.json")
	logfile := filepath.Join(dir, "go/src/github.com/mohfunk/money/log.json")

	c.DataDir = ddir
	c.DataFile = assetsfile
	c.TradeFile = tradefile
	c.BudgetFile = budgetfile
	c.LogFile = logfile
	if _, err := os.Stat(c.DataDir); os.IsNotExist(err) {
		os.Mkdir(c.DataDir, 0700)
	}
	if _, err := os.Stat(c.DataFile); os.IsNotExist(err) {
		util.Copy(baseAssetfile, assetsfile)
	}
	if _, err := os.Stat(c.TradeFile); os.IsNotExist(err) {
		util.Copy(baseTradefile, tradefile)
	}
	if _, err := os.Stat(c.BudgetFile); os.IsNotExist(err) {
		util.Copy(baseBudgetfile, budgetfile)
	}
}
