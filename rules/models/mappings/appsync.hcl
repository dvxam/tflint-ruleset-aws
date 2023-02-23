import = "aws-sdk-go/models/apis/appsync/2017-07-25/api-2.json"

mapping "aws_appsync_datasource" {
  name = ResourceName
  type = DataSourceType
}

mapping "aws_appsync_graphql_api" {
  authentication_type = AuthenticationType
}

mapping "aws_appsync_resolver" {
  type              = ResourceName
  field             = ResourceName
  data_source       = ResourceName
  request_template  = MappingTemplate
  response_template = MappingTemplate
}

mapping "aws_appsync_function" {
  api_id                    = String
  data_source               = ResourceName
  name                      = ResourceName
  request_mapping_template  = MappingTemplate
  response_mapping_template = MappingTemplate
  description               = String
  function_version          = String
}

test "aws_appsync_datasource" "name" {
  valid   = ["tf_appsync_example"]
  invalid = ["01_tf_example"]
}

test "aws_appsync_datasource" "type" {
  valid   = ["AWS_LAMBDA"]
  invalid = ["AMAZON_SIMPLEDB"]
}

test "aws_appsync_graphql_api" "authentication_type" {
  valid   = ["API_KEY"]
  invalid = ["AWS_KEY"]
}
