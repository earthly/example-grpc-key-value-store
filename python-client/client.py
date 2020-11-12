import sys
import grpc

import api_pb2
import api_pb2_grpc

addr = '127.0.0.1:50051'

if len(sys.argv) < 2:
    print('program requires arguments in the form key, or key=value')
    sys.exit(1)

channel = grpc.insecure_channel(addr)
stub = api_pb2_grpc.KeyValueStub(channel)

for arg in sys.argv[1:]:
    if '=' in arg:
        # send a value to the server
        key, value = arg.split('=')
        try:
            set_request = api_pb2.SetRequest(key=key, value=value)
            set_response = stub.Set(set_request)
        except grpc.RpcError as e:
            print(f'failed to send key to server: {e.details}')
        else:
            print(f'sent "{key}" to server')
    else:
        # get a value from the server
        key = arg
        try:
            get_request = api_pb2.GetRequest(key=key)
            get_response = stub.Get(get_request)
        except grpc.RpcError as e:
            print(f'failed to get key from server: {e.details}')
        else:
            value = get_response.value
            print(f'server returned value "{value}" for key "{key}"')
