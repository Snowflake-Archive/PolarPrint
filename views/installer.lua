local host, key, autopilot = arg[2], arg[1], nil
--{{AUTOPILOT_INJECT}}

print("Installing PolarPrint Agent" .. (autopilot and " with Autopilot" or ""))

if not host then
  print("Host URL: ")
  host = read()
end

if not key then
  print("Key: ")
  key = read()
end

print("Verifying...")

print("Verified connection! Downloading agent.lua")

if fs.exists("/agent.lua") then
  print("agent.lua file already exists. Do you wish to overwrite? (Y/n)")
  if read():lower() == "n" then
    return
  end
end

local response, sErr, errResponse = http.get(host .. "dl/agent", { Authorization = "Bearer " .. key })
if not response then
  error("Unable to download agent.lua from server! " .. sErr .. errResponse)
end

local sData = response.readAll()
response.close()

local file = fs.open("/agent.lua", 'w')
assert(file, "Unable to open /agent.lua. Is it write protected?")
file.write(sData)
file.close()

print("Successfully downloaded agent!\nAdding to startup.lua")
if fs.exists("/startup.lua") then
  print("Startup file already exists. Appending to the end. Current contents will be prioritized on boot.")
end

local file = fs.open("/startup.lua", "a")
assert(file, "Unable to open /startup.lua. Is it write protected?")
file.write(("shell.run \"/agent.lua\""))
file.close()

print("Agent installed!")
if autopilot then
  print("Restarting in 3 seconds...")
  sleep(3)
  os.reboot()
end
