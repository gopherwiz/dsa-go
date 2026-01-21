package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"log"
	"os"
)

// The Secret ID can be the ARN or the friendly name of the secret.
// *** IMPORTANT: Change this to the name of your existing secret ***
const secretID = "value-finder-internal-auth-secret"

// The new secret value. For a simple API key, it's a string.
// If your secret is JSON, format it as a JSON string here (e.g., `{"user": "new_name", "pass": "new_pass"}`).
// *** IMPORTANT: Paste the new key generated with 'openssl rand -hex 32' here ***
const newSecretValue = "563506d1-e1c6-4e69-8913-32c490596a00" // 563506d1-e1c6-4e69-8913-32c490596a00(PROD) b29e34b3-b321-4c14-8517-eaa88bfe8432(STAGING)

func main() {
	//setSecret()

	getSecret()
}

func setSecret() {
	// The context for API calls, used to manage timeouts and cancellation.
	ctx := context.TODO()

	// 1. Load AWS Configuration
	// This automatically searches for credentials and region in the environment (e.g., AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION, or via assumed role).
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("Error loading AWS configuration: %v", err)
	}

	creds, err := cfg.Credentials.Retrieve(ctx)
	if err != nil {
		log.Fatalf("Error getting AWS creds: %v", err)
	}

	log.Printf("AWS creds: %#v\n", creds)

	// 2. Create Secrets Manager Client
	smClient := secretsmanager.NewFromConfig(cfg)

	// 3. Define the Input for PutSecretValue
	input := &secretsmanager.PutSecretValueInput{
		SecretId:     aws.String(secretID),
		SecretString: aws.String(newSecretValue),
	}

	// 4. Execute the API Call
	fmt.Printf("Attempting to update secret: %s...\n", secretID)
	result, err := smClient.PutSecretValue(ctx, input)
	if err != nil {
		log.Printf("Error: The AWS user/role may not have the required 'secretsmanager:PutSecretValue' permission or the Secret ID is incorrect.")
		log.Fatalf("AWS API Error: %v", err)
	}

	// 5. Success Confirmation
	fmt.Println("--- SUCCESS ---")
	fmt.Printf("Secret successfully updated: %s\n", *result.ARN)
	fmt.Printf("New version ID created: %s\n", *result.VersionId)
	fmt.Printf("The new value is now marked as the AWSCURRENT version.\n")

	// Optional: Exit the program with a successful status code
	os.Exit(0)
}

func getSecret() {
	ctx := context.TODO()

	// 1. Load AWS Configuration (using the same environment setup)
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("Error loading AWS configuration for get: %v", err)
	}

	creds, err := cfg.Credentials.Retrieve(ctx)
	if err != nil {
		log.Fatalf("Error getting AWS creds for get: %v", err)
	}
	log.Printf("AWS creds (Getter): %#v\n", creds)

	// 2. Create Secrets Manager Client
	smClient := secretsmanager.NewFromConfig(cfg)

	// 3. Define the Input for GetSecretValue
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretID),
		// Leaving VersionStage and VersionId unset defaults to retrieving the AWSCURRENT version.
	}

	// 4. Execute the API Call
	fmt.Printf("\n--- GET OPERATION ---\n")
	fmt.Printf("Attempting to retrieve secret: %s...\n", secretID)

	result, err := smClient.GetSecretValue(ctx, input)
	if err != nil {
		log.Printf("Error: The AWS user/role may not have the required 'secretsmanager:GetSecretValue' permission.")
		log.Fatalf("AWS API Error on Get: %v", err)
	}

	// 5. Display the retrieved value
	fmt.Println("RETRIEVAL SUCCESS")
	if result.SecretString != nil {
		// Log the actual value but only print a snippet to the console for security
		secretValue := *result.SecretString
		fmt.Printf("Secret Name: %s\n", secretID)
		fmt.Printf("Secret Version: %s\n", *result.VersionId)
		fmt.Printf("Secret Value (First 10 chars for verification): %s...\n", secretValue[:10])
		// Note: The full secret value is available in the 'secretValue' variable.
	} else {
		// Handle binary secrets (not used in this API key example but good practice)
		fmt.Printf("Secret Name: %s\n", secretID)
		fmt.Println("Secret value is stored as binary data.")
	}
}
