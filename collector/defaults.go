package collector

func RegisterDefaultCollectors() {
	svc := GetCollectorService()

	svc.Register(
		UsersCollector{})
}
