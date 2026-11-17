var (
    instance *MyService
    once     sync.Once
)

func GetInstance() *MyService {
    once.Do(func() {
        instance = &MyService{}
    })
    return instance
}
