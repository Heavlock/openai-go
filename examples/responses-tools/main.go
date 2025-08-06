package main

import (
	"context"
	"fmt"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/responses"
	"github.com/openai/openai-go/shared"
)

func main() {
	client := openai.NewClient()
	ctx := context.Background()

	question := "Find the world's tallest mountain and double its height using code"

	resp, err := client.Responses.New(ctx, responses.ResponseNewParams{
		Model: shared.ResponsesModel("gpt-4o-mini"),
		Input: responses.ResponseNewParamsInputUnion{
			OfString: openai.String(question),
		},
		Tools: []responses.ToolUnionParam{
			responses.ToolParamOfWebSearchPreview(responses.WebSearchToolTypeWebSearchPreview),
			responses.ToolParamOfCodeInterpreter(responses.ToolCodeInterpreterContainerCodeInterpreterContainerAutoParam{}),
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.OutputText())
}
