package tools

type (
	Mode string
)

const (
	ModeDev  Mode = "dev"
	ModeTest Mode = "test" // 测试模式
	ModeProd Mode = "prod" // 生产模式
)