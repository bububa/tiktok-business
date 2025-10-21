package enum

// ProductImageStatus 商品图片的转存状态
type ProductImageStatus string

const (
	// ProductImageStatus_PROCESSING: 图片正在被转存。
	ProductImageStatus_PROCESSING ProductImageStatus = "PROCESSING"
	// ProductImageStatus_SUCCESS：图片转存成功。
	ProductImageStatus_SUCCESS ProductImageStatus = "SUCCESS"
	// ProductImageStatus_FAIL：图片转存失败。
	ProductImageStatus_FAIL ProductImageStatus = "FAIL"
	// ProductImageStatus_FILTERED：图片转存成功，但是因为尺寸不合尺寸要求（需要不小于500 x 500 像素)，无法使用。如果图片不满足要求，产品的审核状态会一直是processing。对于这种图片，你需要使用上传或更新产品接口上传一个满足条件的新图片。
	ProductImageStatus_FILTERED ProductImageStatus = "FILTERED"
	// ProductImageStatus_NOT_SUPPORTED：图片格式不支持。
	ProductImageStatus_NOT_SUPPORTED ProductImageStatus = "NOT_SUPPORTED"
	// ProductImageStatus_NO_FOUND：无法通过地址找到资源。请确保地址正确，文件可在网上公开访问。
	ProductImageStatus_NO_FOUND ProductImageStatus = "NO_FOUND"
)
