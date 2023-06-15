package ports

type RedisRepository interface {
	Get(key string) (*float64, error)
	Set(key string, data float64) error
	Pub(str *string)
	Sub()
}