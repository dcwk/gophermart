## Event stroming накопительная система лояльности «Гофермарт»
![doc/img_1.png](doc/img_1.png)<br/>
https://miro.com/app/board/uXjVKJX4m7E=/?share_link_id=670000264537

## Таблицы бд
```mermaid
erDiagram
User ||--|| UserBalance: userId
User ||--o{ Order: userId
Order ||--|| Accrual: orderId
Order ||--o{ Withdrawal: orderId

User {
    int id PK
    string login UK
    string password
}

UserBalance {
    int userId PK
    double accrual
    double withdrawal
}

Order {
    int id PK
    int userId FK
    string number UK
    timestamp createdAt 
}

Accrual {
    int orderId PK
    enum status
    double value
    timestamp createdAt
    timestamp updatedAt
}

Withdrawal {
    int id PK
    int orderId FK
    double value
    timestamp createdAt
}
```
## Диаграмма последовательности для операции загрузки заказа
```mermaid
sequenceDiagram
    participant Client
    participant LoyaltySystem
    participant DB
    participant ScoringSystem
    
    Client->>LoyaltySystem: Загрузить заказ
    activate LoyaltySystem
    Note over LoyaltySystem,DB: Проверить существование пользователя
    LoyaltySystem->>DB: Открываем транзакцию. Создаем запись в таблице Order и в таблице Accrual в статусе Processing. 
    deactivate LoyaltySystem
       loop Пока не получим данные со статусом PROCESSED, INVALID, либо что заказ не зарегистрирован 
          LoyaltySystem->>ScoringSystem: Асинхронно делаем запрос по номеру заказа
          activate ScoringSystem
          alt Заказ найден
                  ScoringSystem->>LoyaltySystem: 200 возвращаем данные по заказу 
          else Заказ не найден
                  ScoringSystem->>LoyaltySystem: 204 заказ не зарегистрирован
          end
          deactivate ScoringSystem
       end
    activate LoyaltySystem
    alt Получили ответ в статусе PROCESSED
        LoyaltySystem->>DB: Блокируем баланс UserBalance, переводим Accrual в статус PROCESSED записываем в поле value значение из ответа сервиса расчета вознаграждений. Затем увеличиваем баланс и закрываем транзакцию. 
    else Получили ответ в статусе INVALID
        LoyaltySystem->>DB: Переводим Accrual в статус INVALID и закрываем транзакцию
    end
    deactivate LoyaltySystem
    LoyaltySystem->>Client: Возвращаем результат операции
```
## Диаграмма последовательности для операции вывода баллов
```mermaid
sequenceDiagram
    participant Client
    participant LoyaltySystem
    participant DB
    
    Client->>LoyaltySystem: Запрос на вывод баллов
    Note over LoyaltySystem,DB: Проверить существование пользователя + проверить что заказ принадлежит ему
    LoyaltySystem->>DB: Открываем транзакцию.Блокируем счет UserBalance получаем текущий баланс
    alt Если баллов на счету достаточно для вывода
        LoyaltySystem->>DB: Создаем запись в Withdrawal, уменьшаем в UserBalance поле accrual увеличиваем withdrawal
        LoyaltySystem->>Client: Возвращаем успешный ответ
    else Если баллов недостаточно
        LoyaltySystem->>DB: Откатываем транзакцию
        LoyaltySystem->>Client: Возвращаем ошибку
    end
```