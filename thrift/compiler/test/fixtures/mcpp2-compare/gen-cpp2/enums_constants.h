/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
#pragma once

#include "src/gen-cpp2/enums_types.h"
#include <thrift/lib/cpp2/protocol/Protocol.h>
namespace facebook { namespace ns { namespace qwerty {

struct enums_constants {

  static std::map<std::string,  ::facebook::ns::qwerty::AnEnumB> const& MapStringEnum();

  static std::map< ::facebook::ns::qwerty::AnEnumC, std::string> const& MapEnumString();

  static std::map< ::facebook::ns::qwerty::AnEnumA, std::set< ::facebook::ns::qwerty::AnEnumB>> const& ConstantMap1();

  static std::map< ::facebook::ns::qwerty::AnEnumC, std::map<int16_t, std::set<int16_t>>> const& ConstantMap2();

};

}}} // facebook::ns::qwerty
