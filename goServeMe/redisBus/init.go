package redisBus

func init() {
	err := ConnectToServer(5)
	if err != nil {
		panic(err)
	}
}
