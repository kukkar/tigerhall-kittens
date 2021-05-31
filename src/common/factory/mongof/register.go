package mongof

// func RegisterDependencyChecker() {
// 	depchecker.RegisterDependency(func() depchecker.Dependency {
// 		mongoDep := new(MongoChecker)
// 		return mongoDep
// 	}())
// }

// type MongoChecker struct{}

// func (this *MongoChecker) GetPinger() func() (bool, error) {
// 	return func() (bool, error) {
// 		mdb, err := GetPool(DEFAULT_KEY)
// 		if err != nil {
// 			return false, err
// 		}
// 		mdberr := mdb.Insert("test", map[string]interface{}{
// 			"testedAt": time.Now().Unix(),
// 		})
// 		if mdberr != nil {
// 			return false, fmt.Errorf(mdberr.DeveloperMessage)
// 		}
// 		return true, nil
// 	}
// }

// func (this *MongoChecker) GetName() string {
// 	return "mongo"
// }
