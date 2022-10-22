package down

import (
	"fmt"
	"github.com/WHUPRJ/woj-server/pkg/utils"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Down(dest string, url string) error {
	dir := filepath.Dir(dest)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	tmp := fmt.Sprintf("%s.%s", dest, utils.RandomString(5))
	f, err := os.Create(tmp)
	if err != nil {
		return err
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		_ = os.Remove(tmp)
		return fmt.Errorf("bad status %s when accessing %s", resp.Status, url)
	}

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	err = os.Rename(tmp, dest)
	if err != nil {
		_ = os.Remove(dest)
		return err
	}

	return nil
}
