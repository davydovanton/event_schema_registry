# Event schema registry

This repository is an example of how to make event schema registry for JSON schema events using only github. The general idea - how to share schemas across different services plus how to validate data for specific events.

## Setup
### Ruby
Add this line into your Gemfile:

```
gem "schema_registry", git: "https://github.com/davydovanton/event_schema_registry.git"
```

## How to validate an event data by specific schema

### Ruby

For validating event data you need to use `SchemaRegistry#validate_event` method with following options:

* `data` - event data
* `name` - name of event which you will use for getting schema
* `version` - version of event data schema (default `1`)

Example:

```ruby
message = {
  # ...
}

# will try to search `schemas/Billing/CompliteCycle/1.json` file
result = SchemaRegistry.validate_event(data, 'Billing.CompliteCycle', version: 1)
# will try to search `schemas/billing/complite_cycle/1.json` file
result = SchemaRegistry.validate_event(data, 'billing.complite_cycle', version: 1)

# After you can work with result object
result.success?
result.failure?
result.failure
```

## How to use this library with producer
### Option one: with event object
```ruby
result = SchemaRegistry.validate_event(event, 'billing.refund', version: 1)

if result.success?
  kafka.produce('topic', event.to_json)
end
```

### Option two: with pure hash
```ruby
result = SchemaRegistry.validate_event(event, 'billing.refund', version: 1)

if result.success?
  kafka.produce('topic', event.to_json)
end
```
