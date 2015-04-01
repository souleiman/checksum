package checksum
import (
    "hash"
    "os"
    "math"
    "io"
)

const chunk = 8196 // 8KB

func compute_checksum(hash hash.Hash, filename string) ([]byte, error) {
    hash.Reset() // Clear any previous hashsum [Useful if you are re-using)

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

