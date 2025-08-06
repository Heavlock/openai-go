package responses_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/internal/testutil"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/responses"
	"github.com/openai/openai-go/shared"
)

func TestResponseWithBrowserAndCodeInterpreter(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := openai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Responses.New(context.TODO(), responses.ResponseNewParams{
		Model: shared.ResponsesModel("gpt-4o"),
		Input: responses.ResponseNewParamsInputUnion{
			OfString: openai.String("Search for a mountain and compute a value"),
		},
		Prompt: responses.ResponsePromptParam{
			ID:      "id",
			Version: openai.String("v1"),
			Variables: map[string]responses.ResponsePromptVariableUnionParam{
				"topic": {OfString: openai.String("mountain")},
			},
		},
		Tools: []responses.ToolUnionParam{
			responses.ToolParamOfWebSearchPreview(responses.WebSearchToolTypeWebSearchPreview),
			responses.ToolParamOfCodeInterpreter(responses.ToolCodeInterpreterContainerCodeInterpreterContainerAutoParam{}),
		},
	})
	if err != nil {
		var apierr *openai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
