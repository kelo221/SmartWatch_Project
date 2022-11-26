#include <nlohmann/json.hpp>

#pragma once

#include <optional>
#include <stdexcept>
#include <regex>

namespace eventSpace {
    using nlohmann::json;

    inline json get_untyped(const json & j, const char * property) {
        if (j.find(property) != j.end()) {
            return j.at(property).get<json>();
        }
        return {};
    }

    inline json get_untyped(const json & j, const std::string& property) {
        return get_untyped(j, property.data());
    }

    class event {
    public:
        event() = default;
        virtual ~event() = default;

    private:
        int64_t id{};
        std::string eventname;
        int64_t eventtime{};
        int64_t userid{};
        bool issilent{};
        bool snoozedisabled{};

    public:
        const int64_t & get_id() const { return id; }
        int64_t & get_mutable_id() { return id; }
        void set_id(const int64_t & value) { this->id = value; }

        const std::string & get_eventname() const { return eventname; }
        std::string & get_mutable_eventname() { return eventname; }
        void set_eventname(const std::string & value) { this->eventname = value; }

        const int64_t & get_eventtime() const { return eventtime; }
        int64_t & get_mutable_eventtime() { return eventtime; }
        void set_eventtime(const int64_t & value) { this->eventtime = value; }

        const int64_t & get_userid() const { return userid; }
        int64_t & get_mutable_userid() { return userid; }
        void set_userid(const int64_t & value) { this->userid = value; }

        const bool & get_issilent() const { return issilent; }
        bool & get_mutable_issilent() { return issilent; }
        void set_issilent(const bool & value) { this->issilent = value; }

        const bool & get_snoozedisabled() const { return snoozedisabled; }
        bool & get_mutable_snoozedisabled() { return snoozedisabled; }
        void set_snoozedisabled(const bool & value) { this->snoozedisabled = value; }
    };

    using events = std::vector<event>;
}

namespace eventSpace {
    using events = std::vector<event>;
}

namespace nlohmann {
    void from_json(const json & j, eventSpace::event & x);
    void to_json(json & j, const eventSpace::event & x);

    inline void from_json(const json & j, eventSpace::event& x) {
        x.set_id(j.at("id").get<int64_t>());
        x.set_eventname(j.at("eventname").get<std::string>());
        x.set_eventtime(j.at("eventtime").get<int64_t>());
        x.set_userid(j.at("userid").get<int64_t>());
        x.set_issilent(j.at("issilent").get<bool>());
        x.set_snoozedisabled(j.at("snoozedisabled").get<bool>());
    }

    inline void to_json(json & j, const eventSpace::event & x) {
        j = json::object();
        j["id"] = x.get_id();
        j["eventname"] = x.get_eventname();
        j["eventtime"] = x.get_eventtime();
        j["userid"] = x.get_userid();
        j["issilent"] = x.get_issilent();
        j["snoozedisabled"] = x.get_snoozedisabled();
    }
}

