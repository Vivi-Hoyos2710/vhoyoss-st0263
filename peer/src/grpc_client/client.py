import grpc
import time
from google.protobuf import empty_pb2
from protobuf.filesystem_pb2_grpc import FileSystemStub
from protobuf.filesystem_pb2 import Filename
class Client:
    def _create_client(self,socket: str):
         channel: grpc.Channel = grpc.insecure_channel(socket)
         return FileSystemStub(channel)
        
    def upload(self, socket: str) -> None:
        print(f"Intentando crear cliente con SOCKET={socket}", flush=True)
        client=self._create_client(socket)
        req=  Filename("test.txt")
        response= client.Upload(req)
        print(response)
    def download(self, socket: str) -> None:
        print(f"Intentando crear cliente con SOCKET={socket}", flush=True)
        client=self._create_client(socket)
        req=  Filename("test.txt")
        response= client.Download(req)
        print(response)    
    
        
