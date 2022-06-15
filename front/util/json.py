from datetime import datetime
from google.protobuf.message import Message
from google.protobuf.pyext._message import RepeatedCompositeContainer
from google.protobuf.json_format import MessageToDict
from json import JSONEncoder


class CustomEncoder(JSONEncoder):
    def default(self, o):
        """
        Serialize and de-serialize by using custom parser

        :param any o: custom object
        """
        if isinstance(o, datetime):
            return o.strftime('%Y-%m-%d %H:%M:%S%z')

        if isinstance(o, Message):
            return MessageToDict(o)

        if isinstance(o, RepeatedCompositeContainer):
            return list(o)

        return super(CustomEncoder, self).default(o)
