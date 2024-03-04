from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class Filename(_message.Message):
    __slots__ = ("id", "name")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    id: int
    name: str
    def __init__(self, id: _Optional[int] = ..., name: _Optional[str] = ...) -> None: ...

class Response(_message.Message):
    __slots__ = ("message",)
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    message: str
    def __init__(self, message: _Optional[str] = ...) -> None: ...

class ConversionData(_message.Message):
    __slots__ = ("id", "conversion", "amount")
    ID_FIELD_NUMBER: _ClassVar[int]
    CONVERSION_FIELD_NUMBER: _ClassVar[int]
    AMOUNT_FIELD_NUMBER: _ClassVar[int]
    id: int
    conversion: str
    amount: float
    def __init__(self, id: _Optional[int] = ..., conversion: _Optional[str] = ..., amount: _Optional[float] = ...) -> None: ...

class Response_2(_message.Message):
    __slots__ = ("amount_result",)
    AMOUNT_RESULT_FIELD_NUMBER: _ClassVar[int]
    amount_result: float
    def __init__(self, amount_result: _Optional[float] = ...) -> None: ...
