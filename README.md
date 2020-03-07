# go-common


```go
import (
 "github.com/qq1060656096/go-common"
)

godotenv.Load(".env.example")
// 初始化数据连接
DbConnInit()
db, err := DbConnManager.Get("common").GetGormDB()

	
// 初始化redis连接
RedisConnInit()
redisClient := RedisConnManager.Get("default").GetRedisClient()

```

