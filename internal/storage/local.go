package storage

import (
	"io"
	"os"
	"path/filepath"
	"sync"
)

type LocalCof struct {
	RootDir string `mapstructure:"root_dir" json:"root_dir" yaml:"root_dir"`
	Url     string `mapstructure:"url" json:"url" yaml:"url"`
}

type local struct {
	cof *LocalCof
}

var (
	l    *local
	once *sync.Once
)

func InitLocal(cof LocalCof) (Storage, error) {
	once = &sync.Once{}
	once.Do(func() {
		l = &local{
			cof: &cof,
		}
		Register(Local, l)
	})
	return l, nil
}

func (l *local) getPath(key string) string {
	key = NormalizeKey(key)
	return filepath.Join(l.cof.RootDir, key)
}

func (l *local) Put(key string, r io.Reader, dataLength int64) error {
	path := l.getPath(key)
	dir, _ := filepath.Split(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	fd, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		if os.IsPermission(err) {
			return FileNoPermissionErr
		}
		return err
	}
	defer fd.Close()

	_, err = io.Copy(fd, r)

	return err
}

func (l *local) PutFile(key string, localFile string) error {
	path := l.getPath(localFile)

	fd, fileInfo, err := OpenAsReadOnly(path)
	if err != nil {
		return err
	}
	defer fd.Close()

	return l.Put(key, fd, fileInfo.Size())
}

func (l *local) Get(key string) (io.ReadCloser, error) {
	path := l.getPath(key)

	fd, _, err := OpenAsReadOnly(path)
	if err != nil {
		return nil, err
	}

	return fd, nil
}

func (l *local) Rename(srcKey string, destKey string) error {
	srcPath := l.getPath(srcKey)
	ok, err := l.Exists(srcPath)
	if err != nil {
		return err
	}
	if !ok {
		return FileNotFoundErr
	}

	destPath := l.getPath(destKey)
	dir, _ := filepath.Split(destPath)
	if err = os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	return os.Rename(srcPath, destPath)
}

func (l *local) Copy(srcKey string, destKey string) error {
	srcPath := l.getPath(srcKey)
	srcFd, _, err := OpenAsReadOnly(srcPath)
	if err != nil {
		return err
	}
	defer srcFd.Close()

	destPath := l.getPath(destKey)
	dir, _ := filepath.Split(destPath)
	if err = os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	destFd, err := os.OpenFile(destPath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		if os.IsPermission(err) {
			return FileNoPermissionErr
		}
		return err
	}
	defer destFd.Close()

	_, err = io.Copy(destFd, srcFd)
	if err != nil {
		return err
	}

	return nil
}

func (l *local) Exists(key string) (bool, error) {
	path := l.getPath(key)
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		if os.IsPermission(err) {
			return false, FileNoPermissionErr
		}
		return false, err
	}

	return true, nil
}

func (l *local) Size(key string) (int64, error) {
	path := l.getPath(key)
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, FileNotFoundErr
		}
		if os.IsPermission(err) {
			return 0, FileNoPermissionErr
		}
		return 0, err
	}

	return fileInfo.Size(), nil
}

func (l *local) Delete(key string) error {
	path := l.getPath(key)
	err := os.Remove(path)
	if err != nil {
		if os.IsNotExist(err) {
			return FileNotFoundErr
		}
		if os.IsPermission(err) {
			return FileNoPermissionErr
		}
		return err
	}

	return nil
}

func (l *local) Url(key string) string {
	return l.cof.Url + "/" + NormalizeKey(key)
}
