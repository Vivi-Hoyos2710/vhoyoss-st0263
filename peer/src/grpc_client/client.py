import grpc
import time
from protobuf_files.filesystem_pb2_grpc import FileSystemStub
from protobuf_files.filesystem_pb2 import Filename
from protobuf_files.filesystem_pb2_grpc import CurrencyConverterStub
from protobuf_files.filesystem_pb2 import ConversionData

class Client_Remote:
    def _create_client(self,socket: str):
         channel: grpc.Channel = grpc.insecure_channel(socket)
         return  FileSystemStub(channel)
    
    def _create_client_2(self,socket: str):
         channel: grpc.Channel = grpc.insecure_channel(socket)
         return  CurrencyConverterStub(channel)
        
    def upload(self, socket, filenamestr) -> None:
        print(f"Intentando crear cliente con SOCKET={socket}", flush=True)
        client=self._create_client(socket)
        req=  Filename(id=1,name=filenamestr)
        response= client.Upload(req)
        return response

    
    def download(self, socket, filenamestr) -> None:
        print(f"Intentando crear cliente con SOCKET={socket}", flush=True)
        client = self._create_client(socket)
        req =  Filename(id=1,name=filenamestr)
        print(req)
        response = client.Download(req)
        return response    
    
    def currency_converter(self, socket, conversion, amount) -> None:
        print(f"Intentando crear cliente con SOCKET={socket}", flush=True)
        print(socket)
        client = self._create_client_2(socket)
        req =  ConversionData(id=1,conversion=conversion, amount=amount)
        print(req)
        response = client.Convert(req)
        return response