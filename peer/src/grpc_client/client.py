import grpc
import time
from protobuf_files.filesystem_pb2_grpc import FileSystemStub
from protobuf_files.filesystem_pb2 import Filename
class Client_Remote:
    def _create_client(self,socket: str):
         channel: grpc.Channel = grpc.insecure_channel(socket)
         stub: FileSystemStub= FileSystemStub(channel)
         return stub
        
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
    
        
