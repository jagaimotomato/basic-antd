package upload

import (
	"basic-antd/pkg/jwt"
	"basic-antd/pkg/response"
	"basic-antd/tools"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FileResponse struct {
	Size     int64  `json:"size"`
	Path     string `json:"path"`
	FullPath string `json:"fullPath"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}

func UploadFile(c *gin.Context) {
	fileType, _ := c.GetPostForm("type")
	urlPerfix := fmt.Sprintf("http://%s/", c.Request.Host)
	if fileType == "" {
		response.Error(c, 200, errors.New(""), "缺少文件类型")
		return
	} else {
		switch fileType {
		case "1": // 单文件
			res, result := singleFile(c, urlPerfix)
			if !result {
				return
			}
			response.Success(c, res, "上传成功")
			return
		case "2":
			multipartFile := multipleFile(c, urlPerfix)
			response.Success(c, multipartFile, "上传成功")
			return
		case "3": // 头像
			res, result := avatar(c, urlPerfix)
			if !result {
				return
			}
			response.Success(c, res, "上传成功")
			return
		}
	}
}

func avatar(c *gin.Context, urlPerfix string) (res FileResponse, success bool) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Error(c, 200, errors.New(""), "图片不能为空")
		return res, false
	}
	id := uuid.New().String()
	fileName := id + ".jpg"
	path := "assets/uploads/avatar/" + jwt.GetUserName(c) + "/"
	notExist := tools.CheckExist(path)
	if notExist {
		err := tools.MkDir(path)
		if err != nil {
			fmt.Println(err)
		}
	}
	singleFile := path + fileName
	_ = c.SaveUploadedFile(file, singleFile)
	fileType, _ := tools.GetType(singleFile)
	res = FileResponse{
		Size:     tools.GetFileSize(singleFile),
		Path:     singleFile,
		FullPath: urlPerfix + singleFile,
		Name:     file.Filename,
		Type:     fileType,
	}
	return res, true
}

func singleFile(c *gin.Context, urlPerfix string) (res FileResponse, success bool) {
	file, err := c.FormFile("file")
	dir, _ := c.GetPostForm("dir")
	if err != nil {
		response.Error(c, 200, errors.New(""), "图片不能为空")
		return res, false
	}
	id := uuid.New().String()
	fileName := id + tools.GetExt(file.Filename)
	path := "assets/uploads/" + dir + "/"
	notExist := tools.CheckExist(path)
	if notExist {
		err := tools.MkDir(path)
		if err != nil {
			fmt.Println(err)
		}
	}
	singleFile := path + fileName
	_ = c.SaveUploadedFile(file, singleFile)
	fileType, _ := tools.GetType(singleFile)
	res = FileResponse{
		Size:     tools.GetFileSize(singleFile),
		Path:     singleFile,
		FullPath: urlPerfix + singleFile,
		Name:     file.Filename,
		Type:     fileType,
	}
	return res, true
}

func multipleFile(c *gin.Context, urlPerfix string) []FileResponse {
	files := c.Request.MultipartForm.File["file"]
	var multipartFile []FileResponse
	for _, f := range files {
		fileName := uuid.New().String() + tools.GetExt(f.Filename)
		path := "assets/uploads/"
		notExist := tools.CheckExist(path)
		if notExist {
			err := tools.MkDir(path)
			if err != nil {
				fmt.Println(err)
			}
		}
		multipartFileName := path + fileName
		err := c.SaveUploadedFile(f, multipartFileName)
		fileType, _ := tools.GetType(multipartFileName)
		if err != nil {
			response.Error(c, 200, errors.New(""), "上传失败")
		}
		fileRes := FileResponse{
			Size:     tools.GetFileSize(multipartFileName),
			Path:     multipartFileName,
			FullPath: urlPerfix + multipartFileName,
			Name:     f.Filename,
			Type:     fileType,
		}
		multipartFile = append(multipartFile, fileRes)
	}
	return multipartFile
}

/* 获取 oss 一些配置 */
/*func GetOssConfig(c *gin.Context) {
	var result = make(map[string]string)
	result["accessId"] = config.OssConfig.AccessId
	result["accessKey"] = config.OssConfig.AccessKey
	result["endpoint"] = config.OssConfig.Endpoint
	result["bucketName"] = config.OssConfig.BucketName
	response.Success(c, result, "")
}

func GetQiNiuToken(c *gin.Context) {
	bucket := config.QiniuConfig.Bucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(config.QiniuConfig.AccessKey, config.QiniuConfig.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	response.Success(c, upToken, "")
}
*/
