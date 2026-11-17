<?php

// ==========================================
// 1. ЛЕКСИЧЕСКИЙ АНАЛИЗ (ТОКЕНИЗАТОР)
// ==========================================

enum TokenType {
    case NUMBER;
    case LBRACKET; // '['
    case RBRACKET; // ']'
    case STRING;
    case EOF;      // Конец строки
}

class Token {
    public function __construct(
        public TokenType $type,
        public string $value
    ) {}
}

class Scanner {
    private string $input;
    private int $pos = 0;
    private int $length;

    public function __construct(string $input) {
        $this->input = $input;
        $this->length = strlen($input);
    }

    /**
     * Поочередно возвращает лексемы из строки
     */
    public function nextToken(): Token {
        if ($this->pos >= $this->length) {
            return new Token(TokenType::EOF, '');
        }

        $char = $this->input[$this->pos];

        // 1. Обработка чисел (любой разрядности)
        if (is_numeric($char)) {
            $num = '';
            while ($this->pos < $this->length && is_numeric($this->input[$this->pos])) {
                $num .= $this->input[$this->pos];
                $this->pos++;
            }
            return new Token(TokenType::NUMBER, $num);
        }

        // 2. Обработка скобок
        if ($char === '[') {
            $this->pos++;
            return new Token(TokenType::LBRACKET, '[');
        }
        if ($char === ']') {
            $this->pos++;
            return new Token(TokenType::RBRACKET, ']');
        }

        // 3. Обработка буквенных строк
        $str = '';
        while (
            $this->pos < $this->length &&
            !is_numeric($this->input[$this->pos]) &&
            $this->input[$this->pos] !== '[' &&
            $this->input[$this->pos] !== ']'
        ) {
            $str .= $this->input[$this->pos];
            $this->pos++;
        }
        return new Token(TokenType::STRING, $str);
    }
}

// ==========================================
// 2. СИНТАКСИЧЕСКИЙ АНАЛИЗ (АВТОМАТ И РЕКУРСИЯ)
// ==========================================

class Parser {
    private Scanner $scanner;
    private Token $currentToken;

    public function __construct(Scanner $scanner) {
        $this->scanner = $scanner;
        // Инициализируем первую лексему
        $this->currentToken = $this->scanner->nextToken();
    }

    /**
     * Продвигает сканер к следующему токену
     */
    private function consume(TokenType $type): void {
        if ($this->currentToken->type === $type) {
            $this->currentToken = $this->scanner->nextToken();
        } else {
            throw new Exception("Неожиданный токен: " . $this->currentToken->value);
        }
    }

    /**
     * Главный управляющий автомат (Рекурсивный спуск)
     */
    public function parse(): string {
        $result = '';

        // Конечный автомат обрабатывает токены, пока не дойдет до конца или закрывающей скобки
        while (
            $this->currentToken->type !== TokenType::EOF &&
            $this->currentToken->type !== TokenType::RBRACKET
        ) {
            switch ($this->currentToken->type) {
                case TokenType::STRING:
                    // Состояние: Чтение обычной строки
                    $result .= $this->currentToken->value;
                    $this->consume(TokenType::STRING);
                    break;

                case TokenType::NUMBER:
                    // Состояние: Начало вложенного блока повторения k[...]
                    $k = (int)$this->currentToken->value;
                    $this->consume(TokenType::NUMBER);

                    // Обязана идти открывающая скобка
                    $this->consume(TokenType::LBRACKET);

                    // РЕКУРСИВНЫЙ ВЫЗОВ: автомат заходит внутрь скобок
                    $subResult = $this->parse();

                    // Обязана идти закрывающая скобка
                    $this->consume(TokenType::RBRACKET);

                    // Размножаем результат автомата и добавляем к строке
                    $result .= str_repeat($subResult, $k);
                    break;

                default:
                    throw new Exception("Синтаксическая ошибка на токене: " . $this->currentToken->value);
            }
        }

        return $result;
    }
}

// ==========================================
// ФУНКЦИЯ ДЛЯ ЗАПУСКА КЛИЕНТОМ
// ==========================================
function decodeString(string $input): string {
    $scanner = new Scanner($input);
    $parser = new Parser($scanner);
    return $parser->parse();
}

// --- Тесты наLeetCode примерах (включая сложные) ---
echo decodeString("3[a]2[bc]") . "\n";       // Вывод: aaabcbc
echo decodeString("3[a2[c]]") . "\n";        // Вывод: accaccacc
echo decodeString("2[abc]3[cd]ef") . "\n";   // Вывод: abcabccdcdcdef
echo decodeString("10[a]2[b]") . "\n";       // Вывод: aaaaaaaaaabb (корректно обрабатывает числа > 9)
