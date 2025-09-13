package weather

type options struct {
	apkKey string
	id     string
}

type Option func(*options)

func WithAPIKey(apiKey string) Option {
	return func(opt *options) {
		opt.apkKey = apiKey
	}
}

func WithID(id string) Option {
	return func(opt *options) {
		opt.id = id
	}
}
