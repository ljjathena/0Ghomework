package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strconv"

	// 假设引入了 0g-storage-client 的 transfer 包
	"github.com/0glabs/0g-storage-client/transfer"
	"github.com/0glabs/0g-storage-client/core"
)

func main() {
	largeFileName := "large_data.bin" // 假设存在的4GB文件
	chunkSize := 400 * 1024 * 1024    // 400MB

	// 1. 打开源文件
	file, err := os.Open(largeFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 2. 切分并上传
	for i := 0; i < 10; i++ {
		// 创建分片文件名
		partFileName := "part_" + strconv.Itoa(i)
		partFile, err := os.Create(partFileName)
		if err != nil {
			panic(err)
		}

		// 读取 400MB 数据写入分片文件
		// 使用 io.LimitReader 限制读取大小
		written, err := io.Copy(partFile, io.LimitReader(file, int64(chunkSize)))
		if err != nil && err != io.EOF {
			panic(err)
		}
		partFile.Close()
		fmt.Printf("Created chunk %d: %s (%d bytes)\n", i, partFileName, written)

		// 3. 使用 0g-storage-client 上传
		// 注意：这里演示核心逻辑，实际需配置 client、privateKey 等
		upload(partFileName)
	}
}

func upload(filename string) {
	ctx := context.Background()
	
	// 初始化 Client (伪代码，需填入实际 RPC 和 Key)
	// client := transfer.NewUploader(...) 

	// 配置上传选项，这里参考题目要求的 fragment size
	// Fragment Size 控制传输层的分块大小，例如设置为 1MB
	ops := []transfer.UploadOption{
		transfer.WithFragmentSize(1024 * 1024), // 1MB fragment size
		transfer.WithTaskSize(400 * 1024 * 1024), // 任务大小
	}

	// 执行上传
	// rootHash, err := client.Upload(ctx, filename, ops...)
	
	fmt.Printf("Uploading %s with fragment size settings...\n", filename)
	// 处理 err 和 rootHash...
}
