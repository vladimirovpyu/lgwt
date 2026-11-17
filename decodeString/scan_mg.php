<?php

// Использование перечисления токенов из предыдущего шага
enum TokenType {
    case NUMBER;
    case LBRACKET; // '['
    case RBRACKET; // ']'
    case STRING;
    case EOF;
}

class Token {
    public function __construct(public TokenType $type, public string $value) {}
}

// Сканер (токенизатор) остается прежним
class Scanner {
    private string $input;
    private int $pos = 0;
    private int $length;

    public function __construct(string $input) {
        $this->input = $input;
        $this->length = strlen($input);
    }

    public function nextToken(): Token {
        if ($this->pos >= $this->length) return new Token(TokenType::EOF, '');
        $char = $this->input[$this->pos];

        if (is_numeric($char)) {
            $num = '';
            while ($this->pos < $this->length && is_numeric($this->input[$this->pos])) {
                $num .= $this->input[$this->pos];
                $this->pos++;
            }
            return new Token(TokenType::NUMBER, $num);
        }
        if ($char === '[') { $this->pos++; return new Token(TokenType::LBRACKET, '['); }
        if ($char === ']') { $this->pos++; return new Token(TokenType::RBRACKET, ']'); }

        $str = '';
        while ($this->pos < $this->length && !is_numeric($this->input[$this->pos]) && $this->input[$this->pos] !== '[' && $this->input[$this->pos] !== ']') {
            $str .= $this->input[$this->pos];
            $this->pos++;
        }
        return new Token(TokenType::STRING, $str);
    }
}

// Состояния автомата
enum State {
    case INIT;
    case NUM;
    case OPEN;
    case CLOSE;
}

class PushdownAutomaton {
    /**
     * ТАБЛИЦА ПЕРЕХОДОВ
     * Формат: Матрица [Текущее_Состояние][Тип_Токена] -> Следующее_Состояние
     */
    private array $transitionTable;

    public function __construct() {
        $this->transitionTable = [
            State::INIT->name => [
                TokenType::STRING->name   => State::INIT,
                TokenType::NUMBER->name   => State::NUM,
                TokenType::RBRACKET->name => State::CLOSE,
                TokenType::EOF->name      => State::INIT
            ],
            State::NUM->name => [
                TokenType::LBRACKET->name => State::OPEN
            ],
            State::OPEN->name => [
                TokenType::STRING->name   => State::INIT,
                TokenType::NUMBER->name   => State::NUM,
                TokenType::RBRACKET->name => State::CLOSE
            ],
            State::CLOSE->name => [
                TokenType::STRING->name   => State::INIT,
                TokenType::NUMBER->name   => State::NUM,
                TokenType::RBRACKET->name => State::CLOSE,
                TokenType::EOF->name      => State::INIT
            ]
        ];
    }

    public function decode(string $input): string {
        $scanner = new Scanner($input);

        // Магазин (Стек) для хранения контекстов строк и коэффициентов повторения
        $pushdownStore = [];

        // Рабочие регистры автомата
        $currentString = '';
        $currentK = 0;

        // Стартовое состояние
        $state = State::INIT;

        while (true) {
            $token = $scanner->nextToken();

            // Если дошли до конца и автомат в базовом состоянии — завершаем работу
            if ($token->type === TokenType::EOF && $state !== State::NUM && empty($pushdownStore)) {
                break;
            }

            // Проверяем валидность перехода по таблице
            if (!isset($this->transitionTable[$state->name][$token->type->name])) {
                throw new Exception("Синтаксическая ошибка: некорректный токен {$token->value} в состоянии {$state->name}");
            }

            // Меняем состояние автомата согласно таблице переходов
            $state = $this->transitionTable[$state->name][$token->type->name];

            // ДЕЙСТВИЯ АВТОМАТА (Семантика переходов)
            switch ($state) {
                case State::INIT:
                    if ($token->type === TokenType::STRING) {
                        $currentString .= $token->value;
                    }
                    break;

                case State::NUM:
                    $currentK = (int)$token->value;
                    break;

                case State::OPEN:
                    // Заталкиваем в магазин текущую накопленную строку и множитель перед '['
                    array_push($pushdownStore, [$currentString, $currentK]);
                    // Очищаем регистры для сбора подстроки внутри скобок
                    $currentString = '';
                    $currentK = 0;
                    break;

                case State::CLOSE:
                    // Извлекаем (выталкиваем из магазина) предыдущий контекст
                    if (empty($pushdownStore)) {
                        throw new Exception("Несбалансированная закрывающая скобка");
                    }
                    [$previousString, $multiplier] = array_pop($pushdownStore);

                    // Выполняем циклическое повторение текущей подстроки и склеиваем с предыдущей
                    $currentString = $previousString . str_repeat($currentString, $multiplier);
                    break;
            }
        }

        return $currentString;
    }
}

// --- Тестирование ---
$pda = new PushdownAutomaton();

echo $pda->decode("3[a]2[bc]") . "\n";       // Вывод: aaabcbc
echo $pda->decode("3[a2[c]]") . "\n";        // Вывод: accaccacc
echo $pda->decode("2[abc]3[cd]ef") . "\n";   // Вывод: abcabccdcdcdef
echo $pda->decode("12[x]") . "\n";           // Вывод: xxxxxxxxxxxx
