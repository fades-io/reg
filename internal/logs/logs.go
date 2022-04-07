package logs

const (
	FailedHandleError   = "Не удалось обработать ошибку: %v"
	Error               = "Ошибка:"
	FileAccessFailed    = "Не удалось получить доступ к файлу '.env': %v"
	FileAccessSuccess   = "Значения из файла '.env' получены."
	NotFound            = "Не найдено"
	InternalServerError = "Внутренняя ошибка сервера"

	LoginRequired    = "Требуется логин"
	EmailRequired    = "Требуется почта"
	PasswordRequired = "Требуется пароль"

	InvalidRequestBody = "Не удалось считать тело запроса"
	JsonParsingError   = "Не удалось преобразовать JSON в модель"
	InvalidFormat      = "Неверный формат данных"

	LaunchServer         = "Запуск сервера на хосте"
	DatabaseAccessDenied = "Не могу подсоединиться к базе данных, используя драйвер %s"
	DatabaseConnection   = "База данных %s подключена\n"

	SignalReceived = "Получен сигнал: %v"
	Terminating    = "Завершение работы"
)
