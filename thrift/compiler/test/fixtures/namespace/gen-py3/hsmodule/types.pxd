#
# Autogenerated by Thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

from libcpp.string cimport string
from libcpp cimport bool as cbool, nullptr, nullptr_t
from cpython cimport bool as pbool
from libc.stdint cimport int8_t, int16_t, int32_t, int64_t
from libcpp.memory cimport shared_ptr, unique_ptr
from libcpp.vector cimport vector
from libcpp.set cimport set as cset
from libcpp.map cimport map as cmap, pair as cpair
from thrift.py3.exceptions cimport cTException
cimport folly.iobuf as __iobuf
cimport thrift.py3.exceptions
cimport thrift.py3.types
from thrift.py3.types cimport bstring, move, optional_field_ref
from folly.optional cimport cOptional





cdef extern from "gen-cpp2/hsmodule_types_custom_protocol.h" namespace "::cpp2":
    # Forward Declaration
    cdef cppclass cHsFoo "::cpp2::HsFoo"

cdef extern from "gen-cpp2/hsmodule_types.h" namespace "::cpp2":
    cdef cppclass cHsFoo__isset "::cpp2::HsFoo::__isset":
        bint MyInt

    cdef cppclass cHsFoo "::cpp2::HsFoo":
        cHsFoo() except +
        cHsFoo(const cHsFoo&) except +
        bint operator==(cHsFoo&)
        bint operator!=(cHsFoo&)
        bint operator<(cHsFoo&)
        bint operator>(cHsFoo&)
        bint operator<=(cHsFoo&)
        bint operator>=(cHsFoo&)
        int64_t MyInt
        cHsFoo__isset __isset


cdef extern from "<utility>" namespace "std" nogil:
    cdef shared_ptr[cHsFoo] move(unique_ptr[cHsFoo])
    cdef shared_ptr[cHsFoo] move_shared "std::move"(shared_ptr[cHsFoo])
    cdef unique_ptr[cHsFoo] move_unique "std::move"(unique_ptr[cHsFoo])

cdef extern from "<memory>" namespace "std" nogil:
    cdef shared_ptr[const cHsFoo] const_pointer_cast "std::const_pointer_cast<const ::cpp2::HsFoo>"(shared_ptr[cHsFoo])

# Forward Definition of the cython struct
cdef class HsFoo(thrift.py3.types.Struct)


cdef class HsFoo(thrift.py3.types.Struct):
    cdef object __hash
    cdef object __weakref__
    cdef shared_ptr[cHsFoo] _cpp_obj

    @staticmethod
    cdef unique_ptr[cHsFoo] _make_instance(
        cHsFoo* base_instance,
        bint* __isNOTSET,
        object MyInt
    ) except *

    @staticmethod
    cdef create(shared_ptr[cHsFoo])





