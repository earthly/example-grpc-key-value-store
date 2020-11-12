$LOAD_PATH.unshift '.'

require 'grpc'
require 'api_pb'
require 'api_services_pb'
 
stub = Simplekeyvalue::KeyValue::Stub.new(
  '127.0.0.1:50051', :this_channel_is_insecure
)
 
ARGV.map do |arg|
  request = Simplekeyvalue::GetRequest.new(key: arg)
  response = stub.get(request)
  puts response.value
end
