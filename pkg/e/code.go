package e

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400

	// User Error
	ErrorExistUser         = 3001
	ErrorFailEncryption    = 3002
	ErrorExistUserNotFound = 3003
	ErrorNotCompare        = 3004
	ErrorAuthToken         = 3005
	ErrorAuthTokenTimeout  = 3006
	ErrorUploadFail        = 3007
	ErrorSendEmail         = 3008

	// product Error
	ErrorProductImgUpload = 4001

	// favorite Error
	ErrorFavoriteExist = 5001
)
