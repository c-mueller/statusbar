package bar

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
)

func TestParse_Valid(t *testing.T) {
	file, err := os.Open("testdata/config_valid.yml")
	defer file.Close()
	assert.NoError(t, err)

	content, err := ioutil.ReadAll(file)
	assert.NoError(t, err)

	sb, err := BuildFromConfig(content)
	assert.NoError(t, err)

	assert.Equal(t, 4, len(sb.components))
	assert.Equal(t, "hostname", sb.components[0].id)
}

func TestParse_Invalid_UnknownType(t *testing.T) {
	file, err := os.Open("testdata/config_invalid_unknown_type.yml")
	defer file.Close()
	assert.NoError(t, err)

	content, err := ioutil.ReadAll(file)
	assert.NoError(t, err)

	_, err = BuildFromConfig(content)
	assert.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "InvalidTestType"))
}

func TestParse_Invalid_SameID(t *testing.T) {
	file, err := os.Open("testdata/config_invalid_equal_id.yml")
	defer file.Close()
	assert.NoError(t, err)

	content, err := ioutil.ReadAll(file)
	assert.NoError(t, err)

	_, err = BuildFromConfig(content)
	assert.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "Invalid identifier name \"hostname\" is already in use"))
}
