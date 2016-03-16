package orm

var collectionList = []string{}

// CollectionExists ...
func CollectionExists(className string) bool {
	// 先在内存中查询
	for _, v := range collectionList {
		if v == className {
			return true
		}
	}
	// 内存中不存在，则去数据库中查询一次，更新到内存中
	collectionList = TomatoDB.getCollectionNames()
	for _, v := range collectionList {
		if v == className {
			return true
		}
	}
	return false
}

// Find ...
func Find(className string, where map[string]interface{}, options map[string]interface{}) []interface{} {
	return []interface{}{}
}

// Count ...
func Count(className string, where map[string]interface{}, options map[string]interface{}) int {
	return 0
}
