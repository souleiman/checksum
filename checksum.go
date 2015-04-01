package checksum
import (
    "hash"
    "os"
    "math"
    "io"
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
)

const chunk = 8196 // 8KB

var HashMap map[string]hash.Hash = map[string]hash.Hash { // No pun intended :P
    "md5sum": md5.New(),
    "sha1sum": sha1.New(),
    "sha224sum": sha256.New224(),
    "sha256sum": sha256.New(),
    "sha384sum": sha512.New384(),
    "sha512sum": sha512.New(),
}

func Compute(hash hash.Hash, filename string) ([]byte, error) {
    hash.Reset() // Clear any previous hashsum [Useful if you are re-using object)

    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    stat, _ := file.Stat()
    size := stat.Size()

    blocks := uint64(math.Ceil(float64(size) / chunk))
    for i := uint64(0); i < blocks; i++ {
        block_size := int(math.Min(chunk, float64(size - int64(i * chunk))))
        buffer := make([]byte, block_size)

        file.Read(buffer)
        io.WriteString(hash, string(buffer))
    }

    return hash.Sum(nil), nil
}