package wallet

type MysqlRepository interface{
	Create () 
	Update()
}

type RedisRepository interface{
	Get (key string) interface{} 
}