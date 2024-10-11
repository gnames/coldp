package coldp

type Archive interface {
	Extract(Path string) error
	Load() error
}
