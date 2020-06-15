package config

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

var dsn = "testDSN"
var yamlTest = []byte(fmt.Sprintf("dsn: %v", dsn))

func TestInitSuccessful(t *testing.T) {
	v := newViper()
	err, _ := Init(v)
	if err == nil {
		t.Errorf("init errored when it shouldn't have")
	}
}

func TestInitError(t *testing.T) {
	err, _ := Init(nil)
	if err != nil {
		t.Errorf("init did not error when it should have")
	}
}

func TestInitNoDSNError(t *testing.T) {
	v := viper.New()

	err, _ := Init(v)
	if err != nil {
		t.Errorf("init errored when it shouldn't have")
	}
}

func TestGetDSN(t *testing.T) {
	v := newViper()
	confDSN := GetDSN(v)

	if confDSN != dsn {
		t.Errorf("dsn does not match, got '%v' want '%v'", confDSN, dsn)
	}
}

func newViper() *viper.Viper {
	v := viper.New()
	v.SetConfigType("yaml")
	v.ReadConfig(bytes.NewBuffer(yamlTest))

	return v
}
