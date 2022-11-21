/*
 *  Copyright (C) 2020-2022 Embedded AMS B.V. - All Rights Reserved
 *
 *  This file is part of Embedded Proto.
 *
 *  Embedded Proto is open source software: you can redistribute it and/or
 *  modify it under the terms of the GNU General Public License as published
 *  by the Free Software Foundation, version 3 of the license.
 *
 *  Embedded Proto  is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 *  GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License
 *  along with Embedded Proto. If not, see <https://www.gnu.org/licenses/>.
 *
 *  For commercial and closed source application please visit:
 *  <https://EmbeddedProto.com/license/>.
 *
 *  Embedded AMS B.V.
 *  Info:
 *    info at EmbeddedProto dot com
 *
 *  Postal address:
 *    Johan Huizingalaan 763a
 *    1066 VH, Amsterdam
 *    the Netherlands
 */

// This file is generated. Please do not edit!
#ifndef EVENT_H
#define EVENT_H

#include <cstdint>
#include "MessageInterface.h"
#include "WireFormatter.h"
#include "Fields.h"
#include "MessageSizeCalculator.h"
#include "ReadBufferSection.h"
#include "RepeatedFieldFixedSize.h"
#include "FieldStringBytes.h"
#include "Errors.h"
#include "Defines.h"
#include <limits>

// Include external proto definitions

namespace main {

template<
    uint32_t Events_events_REP_LENGTH, 
    uint32_t Events_events_oneEvent_eventName_LENGTH
>
class Events final: public ::EmbeddedProto::MessageInterface
{
  public:
    Events() = default;
    Events(const Events& rhs )
    {
      set_events(rhs.get_events());
    }

    Events(const Events&& rhs ) noexcept
    {
      set_events(rhs.get_events());
    }

    ~Events() override = default;

    template<
        uint32_t oneEvent_eventName_LENGTH
    >
    class oneEvent final: public ::EmbeddedProto::MessageInterface
    {
      public:
        oneEvent() = default;
        oneEvent(const oneEvent& rhs )
        {
          set_id(rhs.get_id());
          set_eventName(rhs.get_eventName());
          set_eventTimeUnix(rhs.get_eventTimeUnix());
          set_createdAtUnix(rhs.get_createdAtUnix());
          set_userId(rhs.get_userId());
          set_isSilent(rhs.get_isSilent());
          set_snoozeDisabled(rhs.get_snoozeDisabled());
        }

        oneEvent(const oneEvent&& rhs ) noexcept
        {
          set_id(rhs.get_id());
          set_eventName(rhs.get_eventName());
          set_eventTimeUnix(rhs.get_eventTimeUnix());
          set_createdAtUnix(rhs.get_createdAtUnix());
          set_userId(rhs.get_userId());
          set_isSilent(rhs.get_isSilent());
          set_snoozeDisabled(rhs.get_snoozeDisabled());
        }

        ~oneEvent() override = default;

        enum class FieldNumber : uint32_t
        {
          NOT_SET = 0,
          ID = 1,
          EVENTNAME = 2,
          EVENTTIMEUNIX = 3,
          CREATEDATUNIX = 4,
          USERID = 5,
          ISSILENT = 6,
          SNOOZEDISABLED = 7
        };

        oneEvent& operator=(const oneEvent& rhs)
        {
          set_id(rhs.get_id());
          set_eventName(rhs.get_eventName());
          set_eventTimeUnix(rhs.get_eventTimeUnix());
          set_createdAtUnix(rhs.get_createdAtUnix());
          set_userId(rhs.get_userId());
          set_isSilent(rhs.get_isSilent());
          set_snoozeDisabled(rhs.get_snoozeDisabled());
          return *this;
        }

        oneEvent& operator=(const oneEvent&& rhs) noexcept
        {
          set_id(rhs.get_id());
          set_eventName(rhs.get_eventName());
          set_eventTimeUnix(rhs.get_eventTimeUnix());
          set_createdAtUnix(rhs.get_createdAtUnix());
          set_userId(rhs.get_userId());
          set_isSilent(rhs.get_isSilent());
          set_snoozeDisabled(rhs.get_snoozeDisabled());
          return *this;
        }

        static constexpr char const* ID_NAME = "id";
        inline void clear_id() { id_.clear(); }
        inline void set_id(const uint32_t& value) { id_ = value; }
        inline void set_id(const uint32_t&& value) { id_ = value; }
        inline uint32_t& mutable_id() { return id_.get(); }
        inline const uint32_t& get_id() const { return id_.get(); }
        inline uint32_t id() const { return id_.get(); }

        static constexpr char const* EVENTNAME_NAME = "eventName";
        inline void clear_eventName() { eventName_.clear(); }
        inline ::EmbeddedProto::FieldString<oneEvent_eventName_LENGTH>& mutable_eventName() { return eventName_; }
        inline void set_eventName(const ::EmbeddedProto::FieldString<oneEvent_eventName_LENGTH>& rhs) { eventName_.set(rhs); }
        inline const ::EmbeddedProto::FieldString<oneEvent_eventName_LENGTH>& get_eventName() const { return eventName_; }
        inline const char* eventName() const { return eventName_.get_const(); }

        static constexpr char const* EVENTTIMEUNIX_NAME = "eventTimeUnix";
        inline void clear_eventTimeUnix() { eventTimeUnix_.clear(); }
        inline void set_eventTimeUnix(const int64_t& value) { eventTimeUnix_ = value; }
        inline void set_eventTimeUnix(const int64_t&& value) { eventTimeUnix_ = value; }
        inline int64_t& mutable_eventTimeUnix() { return eventTimeUnix_.get(); }
        inline const int64_t& get_eventTimeUnix() const { return eventTimeUnix_.get(); }
        inline int64_t eventTimeUnix() const { return eventTimeUnix_.get(); }

        static constexpr char const* CREATEDATUNIX_NAME = "createdAtUnix";
        inline void clear_createdAtUnix() { createdAtUnix_.clear(); }
        inline void set_createdAtUnix(const int64_t& value) { createdAtUnix_ = value; }
        inline void set_createdAtUnix(const int64_t&& value) { createdAtUnix_ = value; }
        inline int64_t& mutable_createdAtUnix() { return createdAtUnix_.get(); }
        inline const int64_t& get_createdAtUnix() const { return createdAtUnix_.get(); }
        inline int64_t createdAtUnix() const { return createdAtUnix_.get(); }

        static constexpr char const* USERID_NAME = "userId";
        inline void clear_userId() { userId_.clear(); }
        inline void set_userId(const uint32_t& value) { userId_ = value; }
        inline void set_userId(const uint32_t&& value) { userId_ = value; }
        inline uint32_t& mutable_userId() { return userId_.get(); }
        inline const uint32_t& get_userId() const { return userId_.get(); }
        inline uint32_t userId() const { return userId_.get(); }

        static constexpr char const* ISSILENT_NAME = "isSilent";
        inline void clear_isSilent() { isSilent_.clear(); }
        inline void set_isSilent(const bool& value) { isSilent_ = value; }
        inline void set_isSilent(const bool&& value) { isSilent_ = value; }
        inline bool& mutable_isSilent() { return isSilent_.get(); }
        inline const bool& get_isSilent() const { return isSilent_.get(); }
        inline bool isSilent() const { return isSilent_.get(); }

        static constexpr char const* SNOOZEDISABLED_NAME = "snoozeDisabled";
        inline void clear_snoozeDisabled() { snoozeDisabled_.clear(); }
        inline void set_snoozeDisabled(const bool& value) { snoozeDisabled_ = value; }
        inline void set_snoozeDisabled(const bool&& value) { snoozeDisabled_ = value; }
        inline bool& mutable_snoozeDisabled() { return snoozeDisabled_.get(); }
        inline const bool& get_snoozeDisabled() const { return snoozeDisabled_.get(); }
        inline bool snoozeDisabled() const { return snoozeDisabled_.get(); }


        ::EmbeddedProto::Error serialize(::EmbeddedProto::WriteBufferInterface& buffer) const override
        {
          ::EmbeddedProto::Error return_value = ::EmbeddedProto::Error::NO_ERRORS;

          if((0U != id_.get()) && (::EmbeddedProto::Error::NO_ERRORS == return_value))
          {
            return_value = id_.serialize_with_id(static_cast<uint32_t>(FieldNumber::ID), buffer, false);
          }

          if(::EmbeddedProto::Error::NO_ERRORS == return_value)
          {
            return_value = eventName_.serialize_with_id(static_cast<uint32_t>(FieldNumber::EVENTNAME), buffer, false);
          }

          if((0 != eventTimeUnix_.get()) && (::EmbeddedProto::Error::NO_ERRORS == return_value))
          {
            return_value = eventTimeUnix_.serialize_with_id(static_cast<uint32_t>(FieldNumber::EVENTTIMEUNIX), buffer, false);
          }

          if((0 != createdAtUnix_.get()) && (::EmbeddedProto::Error::NO_ERRORS == return_value))
          {
            return_value = createdAtUnix_.serialize_with_id(static_cast<uint32_t>(FieldNumber::CREATEDATUNIX), buffer, false);
          }

          if((0U != userId_.get()) && (::EmbeddedProto::Error::NO_ERRORS == return_value))
          {
            return_value = userId_.serialize_with_id(static_cast<uint32_t>(FieldNumber::USERID), buffer, false);
          }

          if((false != isSilent_.get()) && (::EmbeddedProto::Error::NO_ERRORS == return_value))
          {
            return_value = isSilent_.serialize_with_id(static_cast<uint32_t>(FieldNumber::ISSILENT), buffer, false);
          }

          if((false != snoozeDisabled_.get()) && (::EmbeddedProto::Error::NO_ERRORS == return_value))
          {
            return_value = snoozeDisabled_.serialize_with_id(static_cast<uint32_t>(FieldNumber::SNOOZEDISABLED), buffer, false);
          }

          return return_value;
        };

        ::EmbeddedProto::Error deserialize(::EmbeddedProto::ReadBufferInterface& buffer) override
        {
          ::EmbeddedProto::Error return_value = ::EmbeddedProto::Error::NO_ERRORS;
          ::EmbeddedProto::WireFormatter::WireType wire_type = ::EmbeddedProto::WireFormatter::WireType::VARINT;
          uint32_t id_number = 0;
          FieldNumber id_tag = FieldNumber::NOT_SET;

          ::EmbeddedProto::Error tag_value = ::EmbeddedProto::WireFormatter::DeserializeTag(buffer, wire_type, id_number);
          while((::EmbeddedProto::Error::NO_ERRORS == return_value) && (::EmbeddedProto::Error::NO_ERRORS == tag_value))
          {
            id_tag = static_cast<FieldNumber>(id_number);
            switch(id_tag)
            {
              case FieldNumber::ID:
                return_value = id_.deserialize_check_type(buffer, wire_type);
                break;

              case FieldNumber::EVENTNAME:
                return_value = eventName_.deserialize_check_type(buffer, wire_type);
                break;

              case FieldNumber::EVENTTIMEUNIX:
                return_value = eventTimeUnix_.deserialize_check_type(buffer, wire_type);
                break;

              case FieldNumber::CREATEDATUNIX:
                return_value = createdAtUnix_.deserialize_check_type(buffer, wire_type);
                break;

              case FieldNumber::USERID:
                return_value = userId_.deserialize_check_type(buffer, wire_type);
                break;

              case FieldNumber::ISSILENT:
                return_value = isSilent_.deserialize_check_type(buffer, wire_type);
                break;

              case FieldNumber::SNOOZEDISABLED:
                return_value = snoozeDisabled_.deserialize_check_type(buffer, wire_type);
                break;

              case FieldNumber::NOT_SET:
                return_value = ::EmbeddedProto::Error::INVALID_FIELD_ID;
                break;

              default:
                return_value = skip_unknown_field(buffer, wire_type);
                break;
            }

            if(::EmbeddedProto::Error::NO_ERRORS == return_value)
            {
              // Read the next tag.
              tag_value = ::EmbeddedProto::WireFormatter::DeserializeTag(buffer, wire_type, id_number);
            }
          }

          // When an error was detect while reading the tag but no other errors where found, set it in the return value.
          if((::EmbeddedProto::Error::NO_ERRORS == return_value)
             && (::EmbeddedProto::Error::NO_ERRORS != tag_value)
             && (::EmbeddedProto::Error::END_OF_BUFFER != tag_value)) // The end of the buffer is not an array in this case.
          {
            return_value = tag_value;
          }

          return return_value;
        };

        void clear() override
        {
          clear_id();
          clear_eventName();
          clear_eventTimeUnix();
          clear_createdAtUnix();
          clear_userId();
          clear_isSilent();
          clear_snoozeDisabled();

        }

        static char const* field_number_to_name(const FieldNumber fieldNumber)
        {
          char const* name = nullptr;
          switch(fieldNumber)
          {
            case FieldNumber::ID:
              name = ID_NAME;
              break;
            case FieldNumber::EVENTNAME:
              name = EVENTNAME_NAME;
              break;
            case FieldNumber::EVENTTIMEUNIX:
              name = EVENTTIMEUNIX_NAME;
              break;
            case FieldNumber::CREATEDATUNIX:
              name = CREATEDATUNIX_NAME;
              break;
            case FieldNumber::USERID:
              name = USERID_NAME;
              break;
            case FieldNumber::ISSILENT:
              name = ISSILENT_NAME;
              break;
            case FieldNumber::SNOOZEDISABLED:
              name = SNOOZEDISABLED_NAME;
              break;
            default:
              name = "Invalid FieldNumber";
              break;
          }
          return name;
        }

    #ifdef MSG_TO_STRING

        ::EmbeddedProto::string_view to_string(::EmbeddedProto::string_view& str) const
        {
          return this->to_string(str, 0, nullptr, true);
        }

        ::EmbeddedProto::string_view to_string(::EmbeddedProto::string_view& str, const uint32_t indent_level, char const* name, const bool first_field) const override
        {
          ::EmbeddedProto::string_view left_chars = str;
          int32_t n_chars_used = 0;

          if(!first_field)
          {
            // Add a comma behind the previous field.
            n_chars_used = snprintf(left_chars.data, left_chars.size, ",\n");
            if(0 < n_chars_used)
            {
              // Update the character pointer and characters left in the array.
              left_chars.data += n_chars_used;
              left_chars.size -= n_chars_used;
            }
          }

          if(nullptr != name)
          {
            if( 0 == indent_level)
            {
              n_chars_used = snprintf(left_chars.data, left_chars.size, "\"%s\": {\n", name);
            }
            else
            {
              n_chars_used = snprintf(left_chars.data, left_chars.size, "%*s\"%s\": {\n", indent_level, " ", name);
            }
          }
          else
          {
            if( 0 == indent_level)
            {
              n_chars_used = snprintf(left_chars.data, left_chars.size, "{\n");
            }
            else
            {
              n_chars_used = snprintf(left_chars.data, left_chars.size, "%*s{\n", indent_level, " ");
            }
          }
          
          if(0 < n_chars_used)
          {
            left_chars.data += n_chars_used;
            left_chars.size -= n_chars_used;
          }

          left_chars = id_.to_string(left_chars, indent_level + 2, ID_NAME, true);
          left_chars = eventName_.to_string(left_chars, indent_level + 2, EVENTNAME_NAME, false);
          left_chars = eventTimeUnix_.to_string(left_chars, indent_level + 2, EVENTTIMEUNIX_NAME, false);
          left_chars = createdAtUnix_.to_string(left_chars, indent_level + 2, CREATEDATUNIX_NAME, false);
          left_chars = userId_.to_string(left_chars, indent_level + 2, USERID_NAME, false);
          left_chars = isSilent_.to_string(left_chars, indent_level + 2, ISSILENT_NAME, false);
          left_chars = snoozeDisabled_.to_string(left_chars, indent_level + 2, SNOOZEDISABLED_NAME, false);
      
          if( 0 == indent_level) 
          {
            n_chars_used = snprintf(left_chars.data, left_chars.size, "\n}");
          }
          else 
          {
            n_chars_used = snprintf(left_chars.data, left_chars.size, "\n%*s}", indent_level, " ");
          }

          if(0 < n_chars_used)
          {
            left_chars.data += n_chars_used;
            left_chars.size -= n_chars_used;
          }

          return left_chars;
        }

    #endif // End of MSG_TO_STRING

      private:


          EmbeddedProto::uint32 id_ = 0U;
          ::EmbeddedProto::FieldString<oneEvent_eventName_LENGTH> eventName_;
          EmbeddedProto::int64 eventTimeUnix_ = 0;
          EmbeddedProto::int64 createdAtUnix_ = 0;
          EmbeddedProto::uint32 userId_ = 0U;
          EmbeddedProto::boolean isSilent_ = false;
          EmbeddedProto::boolean snoozeDisabled_ = false;

    };

    enum class FieldNumber : uint32_t
    {
      NOT_SET = 0,
      EVENTS = 4
    };

    Events& operator=(const Events& rhs)
    {
      set_events(rhs.get_events());
      return *this;
    }

    Events& operator=(const Events&& rhs) noexcept
    {
      set_events(rhs.get_events());
      return *this;
    }

    static constexpr char const* EVENTS_NAME = "events";
    inline const oneEvent<Events_events_oneEvent_eventName_LENGTH>& events(uint32_t index) const { return events_[index]; }
    inline void clear_events() { events_.clear(); }
    inline void set_events(uint32_t index, const oneEvent<Events_events_oneEvent_eventName_LENGTH>& value) { events_.set(index, value); }
    inline void set_events(uint32_t index, const oneEvent<Events_events_oneEvent_eventName_LENGTH>&& value) { events_.set(index, value); }
    inline void set_events(const ::EmbeddedProto::RepeatedFieldFixedSize<oneEvent<Events_events_oneEvent_eventName_LENGTH>, Events_events_REP_LENGTH>& values) { events_ = values; }
    inline void add_events(const oneEvent<Events_events_oneEvent_eventName_LENGTH>& value) { events_.add(value); }
    inline ::EmbeddedProto::RepeatedFieldFixedSize<oneEvent<Events_events_oneEvent_eventName_LENGTH>, Events_events_REP_LENGTH>& mutable_events() { return events_; }
    inline oneEvent<Events_events_oneEvent_eventName_LENGTH>& mutable_events(uint32_t index) { return events_[index]; }
    inline const ::EmbeddedProto::RepeatedFieldFixedSize<oneEvent<Events_events_oneEvent_eventName_LENGTH>, Events_events_REP_LENGTH>& get_events() const { return events_; }
    inline const ::EmbeddedProto::RepeatedFieldFixedSize<oneEvent<Events_events_oneEvent_eventName_LENGTH>, Events_events_REP_LENGTH>& events() const { return events_; }


    ::EmbeddedProto::Error serialize(::EmbeddedProto::WriteBufferInterface& buffer) const override
    {
      ::EmbeddedProto::Error return_value = ::EmbeddedProto::Error::NO_ERRORS;

      if(::EmbeddedProto::Error::NO_ERRORS == return_value)
      {
        return_value = events_.serialize_with_id(static_cast<uint32_t>(FieldNumber::EVENTS), buffer, false);
      }

      return return_value;
    };

    ::EmbeddedProto::Error deserialize(::EmbeddedProto::ReadBufferInterface& buffer) override
    {
      ::EmbeddedProto::Error return_value = ::EmbeddedProto::Error::NO_ERRORS;
      ::EmbeddedProto::WireFormatter::WireType wire_type = ::EmbeddedProto::WireFormatter::WireType::VARINT;
      uint32_t id_number = 0;
      FieldNumber id_tag = FieldNumber::NOT_SET;

      ::EmbeddedProto::Error tag_value = ::EmbeddedProto::WireFormatter::DeserializeTag(buffer, wire_type, id_number);
      while((::EmbeddedProto::Error::NO_ERRORS == return_value) && (::EmbeddedProto::Error::NO_ERRORS == tag_value))
      {
        id_tag = static_cast<FieldNumber>(id_number);
        switch(id_tag)
        {
          case FieldNumber::EVENTS:
            return_value = events_.deserialize_check_type(buffer, wire_type);
            break;

          case FieldNumber::NOT_SET:
            return_value = ::EmbeddedProto::Error::INVALID_FIELD_ID;
            break;

          default:
            return_value = skip_unknown_field(buffer, wire_type);
            break;
        }

        if(::EmbeddedProto::Error::NO_ERRORS == return_value)
        {
          // Read the next tag.
          tag_value = ::EmbeddedProto::WireFormatter::DeserializeTag(buffer, wire_type, id_number);
        }
      }

      // When an error was detect while reading the tag but no other errors where found, set it in the return value.
      if((::EmbeddedProto::Error::NO_ERRORS == return_value)
         && (::EmbeddedProto::Error::NO_ERRORS != tag_value)
         && (::EmbeddedProto::Error::END_OF_BUFFER != tag_value)) // The end of the buffer is not an array in this case.
      {
        return_value = tag_value;
      }

      return return_value;
    };

    void clear() override
    {
      clear_events();

    }

    static char const* field_number_to_name(const FieldNumber fieldNumber)
    {
      char const* name = nullptr;
      switch(fieldNumber)
      {
        case FieldNumber::EVENTS:
          name = EVENTS_NAME;
          break;
        default:
          name = "Invalid FieldNumber";
          break;
      }
      return name;
    }

#ifdef MSG_TO_STRING

    ::EmbeddedProto::string_view to_string(::EmbeddedProto::string_view& str) const
    {
      return this->to_string(str, 0, nullptr, true);
    }

    ::EmbeddedProto::string_view to_string(::EmbeddedProto::string_view& str, const uint32_t indent_level, char const* name, const bool first_field) const override
    {
      ::EmbeddedProto::string_view left_chars = str;
      int32_t n_chars_used = 0;

      if(!first_field)
      {
        // Add a comma behind the previous field.
        n_chars_used = snprintf(left_chars.data, left_chars.size, ",\n");
        if(0 < n_chars_used)
        {
          // Update the character pointer and characters left in the array.
          left_chars.data += n_chars_used;
          left_chars.size -= n_chars_used;
        }
      }

      if(nullptr != name)
      {
        if( 0 == indent_level)
        {
          n_chars_used = snprintf(left_chars.data, left_chars.size, "\"%s\": {\n", name);
        }
        else
        {
          n_chars_used = snprintf(left_chars.data, left_chars.size, "%*s\"%s\": {\n", indent_level, " ", name);
        }
      }
      else
      {
        if( 0 == indent_level)
        {
          n_chars_used = snprintf(left_chars.data, left_chars.size, "{\n");
        }
        else
        {
          n_chars_used = snprintf(left_chars.data, left_chars.size, "%*s{\n", indent_level, " ");
        }
      }
      
      if(0 < n_chars_used)
      {
        left_chars.data += n_chars_used;
        left_chars.size -= n_chars_used;
      }

      left_chars = events_.to_string(left_chars, indent_level + 2, EVENTS_NAME, true);
  
      if( 0 == indent_level) 
      {
        n_chars_used = snprintf(left_chars.data, left_chars.size, "\n}");
      }
      else 
      {
        n_chars_used = snprintf(left_chars.data, left_chars.size, "\n%*s}", indent_level, " ");
      }

      if(0 < n_chars_used)
      {
        left_chars.data += n_chars_used;
        left_chars.size -= n_chars_used;
      }

      return left_chars;
    }

#endif // End of MSG_TO_STRING

  private:


      ::EmbeddedProto::RepeatedFieldFixedSize<oneEvent<Events_events_oneEvent_eventName_LENGTH>, Events_events_REP_LENGTH> events_;

};

} // End of namespace main
#endif // EVENT_H