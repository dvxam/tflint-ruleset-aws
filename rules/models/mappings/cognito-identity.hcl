import = "aws-sdk-go/models/apis/cognito-identity/2014-06-30/api-2.json"

mapping "aws_cognito_identity_pool" {
  identity_pool_name      = IdentityPoolName
  developer_provider_name = DeveloperProviderName
}

mapping "aws_cognito_identity_pool_roles_attachment" {
  identity_pool_id = IdentityPoolId
}

test "aws_cognito_identity_pool" "identity_pool_name" {
  valid   = ["identity pool"]
  invalid = ["identity:pool"]
}

test "aws_cognito_identity_pool_roles_attachment" "identity_pool_id" {
  valid   = ["us-east-1:0123456789"]
  invalid = ["0123456789"]
}
