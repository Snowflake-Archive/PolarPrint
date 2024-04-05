-- PolarPrint Agent
-- Erb3 GNU AGPL v3

settings.define("polarprint.host", {
  description =
  "The websocket host for the polarprint agent. Must include protocol and trailing slash. Example: wss://polar.snowflake.blue/",
  type = "string",
})

settings.define("polarprint.key", {
  description = "The secret key for this polarprint cluster",
  type = "string"
})

if not settings.get("polarprint.host") then
  error("Missing polarprint.host setting!")
end

if not settings.get("polarprint.key") then
  error("Missing polarprint.key setting!")
end

local ws = http.websocket(settings.get("polarprint.host"))
