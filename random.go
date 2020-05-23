package randomstring

import (
    time "time"
    big "math/big"
    mathRand "math/rand"
    cr "crypto/rand"
)

func GenString(length int) string {
    mathRand.Seed(time.Now().UnixNano())
    var result string
    for i := 0; i < length; i++ {
        newChar := rune(mathRand.Intn(26) + 'a')
        result += string(newChar)
    }
    return result
}

func GenCapString(length int) string {
    mathRand.Seed(time.Now().UnixNano())
    var result string
    for i := 0; i < length; i++ {
        newChar := rune(mathRand.Intn(26) + 'A')
        result += string(newChar)
    }
    return result
}

func GenStringCry(length int) (string, error) {
    var result string
    max := big.NewInt(26)
    for i := 0; i < length; i++ {
        newCharBigInt, err := cr.Int(cr.Reader, max)
        if err != nil {
            return "", err
        }
        newCharInt := newCharBigInt.Int64()
        newChar := rune(int(newCharInt) + 'a')
        result += string(newChar)
    }
    return result, nil
}

func GenCapStringCry(length int) (string, error) {
    var result string
    max := big.NewInt(26)
    for i := 0; i < length; i++{
        newCharBigInt, err := cr.Int(cr.Reader, max)
        if err != nil {
            return "", err
        }

        newCharInt := newCharBigInt.Int64()
        newChar := rune(int(newCharInt) + 'A')
        result += string(newChar)
    }
    return result, nil
}
