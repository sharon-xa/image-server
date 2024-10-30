package main

import "os"

// create the dirctory if it's not already created
func (app *Application) initDir(path string) error {
	exists, err := exists(path)
	if err != nil {
		app.logger.Error(err.Error())
		return err
	}

	if exists {
		app.logger.Info("images directory exists at " + app.imgDirPath)
		return nil
	}

	if !exists {
		err = os.MkdirAll(path, 0777)
		if err != nil {
			app.logger.Error(err.Error())
			return err
		}
		app.logger.Info("directory created at " + app.imgDirPath)
	}

	return nil
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
