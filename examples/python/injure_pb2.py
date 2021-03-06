# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: injure.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='injure.proto',
  package='injure',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x0cinjure.proto\x12\x06injure\x1a\x1bgoogle/protobuf/empty.proto\"c\n\x07Request\x12\x0f\n\x07methods\x18\x01 \x03(\t\x12\r\n\x05\x65rrno\x18\x02 \x01(\r\x12\x0e\n\x06random\x18\x03 \x01(\x08\x12\x0b\n\x03pct\x18\x04 \x01(\r\x12\x0c\n\x04path\x18\x05 \x01(\t\x12\r\n\x05\x64\x65lay\x18\x06 \x01(\r\"\x1b\n\x08Response\x12\x0f\n\x07methods\x18\x01 \x03(\t2\xac\x02\n\x06Injure\x12\x35\n\x07Methods\x12\x16.google.protobuf.Empty\x1a\x10.injure.Response\"\x00\x12>\n\nRecoverAll\x12\x16.google.protobuf.Empty\x1a\x16.google.protobuf.Empty\"\x00\x12:\n\rRecoverMethod\x12\x0f.injure.Request\x1a\x16.google.protobuf.Empty\"\x00\x12\x35\n\x08SetFault\x12\x0f.injure.Request\x1a\x16.google.protobuf.Empty\"\x00\x12\x38\n\x0bSetFaultAll\x12\x0f.injure.Request\x1a\x16.google.protobuf.Empty\"\x00\x62\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_empty__pb2.DESCRIPTOR,])




_REQUEST = _descriptor.Descriptor(
  name='Request',
  full_name='injure.Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='methods', full_name='injure.Request.methods', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='errno', full_name='injure.Request.errno', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='random', full_name='injure.Request.random', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='pct', full_name='injure.Request.pct', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='path', full_name='injure.Request.path', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='delay', full_name='injure.Request.delay', index=5,
      number=6, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=53,
  serialized_end=152,
)


_RESPONSE = _descriptor.Descriptor(
  name='Response',
  full_name='injure.Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='methods', full_name='injure.Response.methods', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=154,
  serialized_end=181,
)

DESCRIPTOR.message_types_by_name['Request'] = _REQUEST
DESCRIPTOR.message_types_by_name['Response'] = _RESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Request = _reflection.GeneratedProtocolMessageType('Request', (_message.Message,), dict(
  DESCRIPTOR = _REQUEST,
  __module__ = 'injure_pb2'
  # @@protoc_insertion_point(class_scope:injure.Request)
  ))
_sym_db.RegisterMessage(Request)

Response = _reflection.GeneratedProtocolMessageType('Response', (_message.Message,), dict(
  DESCRIPTOR = _RESPONSE,
  __module__ = 'injure_pb2'
  # @@protoc_insertion_point(class_scope:injure.Response)
  ))
_sym_db.RegisterMessage(Response)



_INJURE = _descriptor.ServiceDescriptor(
  name='Injure',
  full_name='injure.Injure',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=184,
  serialized_end=484,
  methods=[
  _descriptor.MethodDescriptor(
    name='Methods',
    full_name='injure.Injure.Methods',
    index=0,
    containing_service=None,
    input_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    output_type=_RESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='RecoverAll',
    full_name='injure.Injure.RecoverAll',
    index=1,
    containing_service=None,
    input_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='RecoverMethod',
    full_name='injure.Injure.RecoverMethod',
    index=2,
    containing_service=None,
    input_type=_REQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='SetFault',
    full_name='injure.Injure.SetFault',
    index=3,
    containing_service=None,
    input_type=_REQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='SetFaultAll',
    full_name='injure.Injure.SetFaultAll',
    index=4,
    containing_service=None,
    input_type=_REQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_INJURE)

DESCRIPTOR.services_by_name['Injure'] = _INJURE

# @@protoc_insertion_point(module_scope)
