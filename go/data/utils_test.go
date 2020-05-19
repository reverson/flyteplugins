package data

import (
	"context"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/lyft/flytestdlib/promutils"
	"github.com/lyft/flytestdlib/promutils/labeled"
	"github.com/lyft/flytestdlib/storage"
	"github.com/stretchr/testify/assert"
)

func TestIsFileReadable(t *testing.T) {
	tmpFolderLocation := ""
	tmpPrefix := "util_test"

	tmpDir, err := ioutil.TempDir(tmpFolderLocation, tmpPrefix)
	assert.NoError(t, err)
	defer func() {
		assert.NoError(t, os.RemoveAll(tmpDir))
	}()
	p := path.Join(tmpDir, "x")
	f, i, err := IsFileReadable(p, false)
	assert.Error(t, err)

	assert.NoError(t, ioutil.WriteFile(p, []byte("data"), os.ModePerm))
	f, i, err = IsFileReadable(p, false)
	assert.NoError(t, err)
	assert.Equal(t, p, f)
	assert.NotNil(t, i)

	noExt := path.Join(tmpDir, "y")
	p = path.Join(tmpDir, "y.png")
	f, i, err = IsFileReadable(noExt, false)
	assert.Error(t, err)

	assert.NoError(t, ioutil.WriteFile(p, []byte("data"), os.ModePerm))
	f, i, err = IsFileReadable(noExt, false)
	assert.Error(t, err)

	f, i, err = IsFileReadable(noExt, true)
	assert.NoError(t, err)
	assert.Equal(t, p, f)
	assert.NotNil(t, i)
}

func TestUploadFile(t *testing.T) {
	tmpFolderLocation := ""
	tmpPrefix := "util_test"

	tmpDir, err := ioutil.TempDir(tmpFolderLocation, tmpPrefix)
	assert.NoError(t, err)
	defer func() {
		assert.NoError(t, os.RemoveAll(tmpDir))
	}()

	exist := path.Join(tmpDir, "exist-file")
	data := []byte("data")
	l := int64(len(data))
	assert.NoError(t, ioutil.WriteFile(exist, data, os.ModePerm))
	nonExist := path.Join(tmpDir, "non-exist-file")

	store, err := storage.NewDataStore(&storage.Config{Type: storage.TypeMemory}, promutils.NewTestScope())
	assert.NoError(t, err)

	ctx := context.TODO()
	assert.NoError(t, UploadFile(ctx, exist, "exist", l, store))

	assert.Error(t, UploadFile(ctx, nonExist, "nonExist", l, store))
}

func TestDownloadFromHttp(t *testing.T) {
	loc := storage.DataReference("https://raw.githubusercontent.com/lyft/flyte/master/README.md")
	badLoc := storage.DataReference("https://no-exist")
	f, err := DownloadFromHttp(context.TODO(), loc)
	if assert.NoError(t, err) {
		if assert.NotNil(t, f) {
			f.Close()
		}
	}

	f, err = DownloadFromHttp(context.TODO(), badLoc)
	assert.Error(t, err)
}

func init() {
	labeled.SetMetricKeys("test")
}