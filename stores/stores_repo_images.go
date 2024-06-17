package main

func RepoImageUrls() (urls []string, err error) {
	urls = []string{}
	err = sqlite.
		Model(&Image{}).
		Distinct().
		Pluck("url", &urls).Error
	return
}

func RepoImageByUrl(url string) (img Image, err error) {
	err = sqlite.Where(&Image{Url: url}).First(&img).Error
	return
}

func RepoImageID(url string) (int64, error) {
	img, err := RepoImageByUrl(url)
	return img.ID, err
}

func RepoImageSave(img *Image) (err error) {
	if x, err := RepoImageByUrl(img.Url); err == nil {
		img.ID = x.ID
		if img.ProductID == 0 {
			img.ProductID = x.ProductID
		}
		if img.Local == "" {
			img.Local = x.Local
		}
	}
	err = sqlite.Save(img).Error
	return
}

func RepoImageSaveAll(images []*Image) (err error) {
	for _, image := range images {
		if err := RepoImageSave(image); err != nil {
			return err
		}
	}
	return nil
}

func RepoImageUpdateProductID(url string, pid int64) (err error) {
	img, err := RepoImageByUrl(url)
	if err != nil {
		return
	}
	img.ProductID = pid
	err = RepoImageSave(&img)
	return
}

func RepoImageUpdateLocal(url, local string) (err error) {
	img, err := RepoImageByUrl(url)
	if err != nil {
		return
	}
	img.Local = local
	err = RepoImageSave(&img)
	return
}

func RepoImageUpdateLocalAndProductID(url, local string, pid int64) (err error) {
	img, err := RepoImageByUrl(url)
	if err != nil {
		return
	}
	img.ProductID = pid
	img.Local = local
	err = RepoImageSave(&img)
	return
}
