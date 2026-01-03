package govibeimpl

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func InvokeGenAI(prompt string) string {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Initialize the model
	model := client.GenerativeModel("gemini-3-flash-preview")

	// --- SET SYSTEM PROMPT HERE ---
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text("You are a coding assistant that only outputs Golang code. Return ONLY the raw code requested. Do not include markdown code blocks (```), do not include introductory text, and do not include any conversational filler."),
		},
	}

	// Make a standard request
	resp, err := model.GenerateContent(ctx,
		genai.Text("help me generate an implementation of the following golang interface (output of go doc). This generated file will be built with the provided interface file together. I have stubbed out the contract types as `any` for now, feel free to populate it with any custom struct needed for the implementation."),
		genai.Text(prompt),
	)
	if err != nil {
		log.Fatal(err)
	}

	return string(resp.Candidates[0].Content.Parts[0].(genai.Text))
}
