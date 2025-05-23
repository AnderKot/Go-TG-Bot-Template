# Go-TG-Bot-Template-V2

Добрый день, пройдёмся по файликам

BotService - сердце сервиса
  Тут реализоывано прослушивание сообщений от ТГ, алгоритм работы со "Страницами" и отправка ответа пользователь
  Алгоритм боработки включает следющие этампы:
  1. Определение ID чата и языка пользователя
  2. Передача "Странице" текста сообщения или нажатия кнопки
  3. Опрос "Страницы" на необходимость перехода к новой "Странице" или возвращение к родительской
  4. Получение текста ответа и набора кнопок от "Страницы"
  5. Отправка ответа

IPage - "Старница" основной строительный блок для бота
  Страницы хранят в себе логику бота
  Для чего каждый func
    // Common
  	GetName() string - Возвращаем уникальное имя своей страницы
  	// Input
  	OnProcessingMessage(text string) - Обрабатываем текстовые сообщения
  	OnProcessingKey(keyData string)  - Обрабатываем нажания кнопок
  	// Navigation
  	OnBackToParent() bool        - Отвечаем нужен ли переход на следующую страницу
  	OnGetNextPage() IConstructor - Возвращаем конструктор для следующего окна
  	// Print
  	GetMessageText() ITemplate - Возвращаем шаблон текстовоого сообщения 
  	GetKeyboard() IKeyboard    - Возвращаем набор кнопок

  Поскольку страницы могут быть достаточно большими структурами для правильной их инициализации следует использовать отдельные структуры реализующие IConstructor

IConstructor - аналог конструктора из ООП ЯП
  В структурах реализующих этот интерфейс следует определять все необхдимые поля для создания страницы
  New() IPage - возвращаем страницу, если для иницализации страницы нужно догрузить что-то из БД делаем это тут

IKeyboard - Контейнер для кнопок под сообщением, содержит строки IKeyRow, котрыйе в свою очередь содеержат кнопки IKey которая уже в свою очередь имеет (GetTemplate() ITemplate) шаблон для отображаемого на кнопке текста и (GetData() string) строку которая вернётся в страницу 

ITemplate - Шаблон текста
  isTranslated() - Если нужно подкрузить текст под язяк пользователя из БД возвращаем true
  GetTemplateCode() string - Код шаблона текста в БД (Если такого нет в БД создатутся строки на основе GetTemplateText() string)
  GetTemplateText() string - Базовый текст сообщения, используется если нужно ввывести текст "как есть" и как дефолтный текст если не будет найден шаблон под язык пользователя

Примеры
  Шалон для текста страницы меню
  type PageTemplate struct {
    text string
    code string
  }
  func (pt PageTemplate) isTranslated() bool { return true }
  func (pt PageTemplate) GetTemplateText() string { return pt.text }
  func (pt PageTemplate) GetTemplateCode() string { return pt.code }

  Стандартный шаблон для кнопки возврата
  const (
	  onBackToParent string = "onBackToParent" - для использования в боработчиках
  )
  type onBackToParentTemplate struct{}
  func (obpt onBackToParentTemplate) isTranslated() bool { return true }
  func (obpt onBackToParentTemplate) GetTemplateText() string { return "Back" }
  func (obpt onBackToParentTemplate) GetTemplateCode() string { return onBackToParent }

Сейчас на основе этих интерфейсов реализовано 2 базовые страницы
  1. PageMenu - проставя менюшка, позвооляет строить базовую навигацию
    Создается через PageMenuConstructor где:
  	  name        string      - Имя
  	  template    ITemplate   - Шаблон текста сообщения
  	  items       []MenuItem  - Кнопки меню, содержат (Name ITemplate) шаблон для текскта на кнопке и (Constructor IConstructor) конструктор следуюшей страници открываемой по нажатию
  	  isHasParent bool        - Нужна ли возможность вернуться назад (добавляет кнопку возврата ниже кнопок меню)
    Для примера главное меню (буду дополнять)
    func CreateMainMenu() *IPage {
    	constructor := PageMenuConstructor{
    		name:        "mainMenu",
    		template:    PageTemplate{code: "mainMenuPage"},
    		items:       make([]MenuItem, 0),
    		isHasParent: false,
    	}
    	page := constructor.New()
    	return &(page)
    }

  2. List - простой список с возможностью постраничного отображения
    Буду переделывать, текушая реализация требует создания конструктора под каждый элемент в списке что достаточно расточительно по ресурсам, смотрю в сторону дженериков

  TO DO:
  1. Возможность передачи данных от дочерней страницы к родительской
  2. Реализация универсальной страницы для анкет, опросов, голосований и т.п.
  3. Мультистраничность, для отображения нескольких страниц со своими функциональностями
  4. Пример простого бота с использованием каждого из стандартных реализаций страниц
