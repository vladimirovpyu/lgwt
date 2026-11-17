package iteration

import "strings"

func Repeat(character string) string
{
    var builder strings.Builder
    // Мы заранее говорим системе, сколько байт нам нужно (например, 5 повторений)
    builder.Grow(5)

    for i := 0; i < 5; i++ {
        builder.WriteString(character)
    }

    return builder.String()
}
