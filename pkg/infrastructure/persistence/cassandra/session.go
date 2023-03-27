package cassandra

// var sessionSingleton *gocql.Session

// func newSession() (*gocql.Session, error) {
// 	if sessionSingleton != nil {
// 		return sessionSingleton, nil
// 	}
// 	port, err := strconv.Atoi(os.Getenv("CASSANDRA_PORT"))
// 	if err != nil {
// 		return nil, err
// 	}
// 	cluster := gocql.NewCluster(os.Getenv("CASSANDRA_HOST"))
// 	cluster.ConnectTimeout = 10 * time.Second
// 	cluster.Keyspace = os.Getenv("CASSANDRA_KEYSPACE")
// 	cluster.Port = port
// 	cluster.ProtoVersion = 4
// 	cluster.Consistency = gocql.Quorum

// 	sessionSingleton, err := cluster.CreateSession()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return sessionSingleton, nil
// }
