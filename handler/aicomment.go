package handler

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ollama/ollama/api"
	"github.com/sashabaranov/go-openai"
	"github.com/xieyuxuan109/homeworksystem/dao"
	"github.com/xieyuxuan109/homeworksystem/model"
	"github.com/xieyuxuan109/homeworksystem/pkg"
)

func buildPrompt(hw *model.Homework, sb *model.Submission) string {
	return fmt.Sprintf(`你是一位经验丰富的%s老师。请根据以下信息对学生的作业进行评价,并生成预估得分：
作业标题：%s
作业内容：%s
学生提交内容：%s

**评价要求**：
1. 先表扬优点，再指出不足。
2. 语气要亲切，像老师对学生说话。
3. 字数控制在 100 字以内。
4. 如果分数在 90 以上，给予特别表扬；如果低于 60，要鼓励加油。`,
		hw.Department, hw.Title, hw.Description, sb.Content)
}

func AIcomment(c *gin.Context) {
	var request model.AIcommentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		pkg.BadResponse(c, "参数错误", err)
		return
	}
	hw, _ := dao.SearchHomeworkById(request.HomeworkID)
	sub, _ := dao.SearchSubmissionById(request.SubmissionID)
	// 1. 设置环境变量 DASHSCOPE_API_KEY=你的API密钥
	apiKey := os.Getenv("DASHSCOPE_API_KEY")
	if apiKey == "" {
		log.Fatal("请设置环境变量 DASHSCOPE_API_KEY")
	}
	// 2. 配置客户端（关键：BaseURL 指向阿里云）
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = "https://dashscope.aliyuncs.com/compatible-mode/v1" // 阿里云兼容地址
	client := openai.NewClientWithConfig(config)
	// 要求 AI 返回 JSON 格式
	prompt := buildPrompt(hw, sub)
	// 3. 构造请求
	req := openai.ChatCompletionRequest{
		Model: "qwen-turbo", // 可选: qwen-plus, qwen-max
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleUser, Content: prompt},
		},
	}

	// 4. 发送请求
	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		log.Fatalf("调用失败: %v", err)
	}

	// 5. 输出结果
	pkg.GoodResponse(c, "调用成功", "AI 回答:"+resp.Choices[0].Message.Content)
}

func LocalAIcomment(c *gin.Context) {
	// 创建客户端
	os.Setenv("OLLAMA_HOST", "http://192.168.1.125:11434")
	client, err := api.ClientFromEnvironment()

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	stream := false
	// 生成请求
	req := &api.GenerateRequest{
		Model:  "deepseek-r1:7b",
		Prompt: "为什么天空是蓝色的？",
		Stream: &stream, // 可选：是否流式响应
	}

	// 生成响应
	respFunc := func(resp api.GenerateResponse) error {
		fmt.Print(resp.Response)
		return nil
	}

	err = client.Generate(ctx, req, respFunc)
	if err != nil {
		panic(err)
	}
}
