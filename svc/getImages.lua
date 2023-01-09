local radius = ARGV[1]
local unit = ARGV[2]
local lon = ARGV[3]
local lat = ARGV[4]
local timeStamp = ARGV[5]
local imgByLocation = redis.call('GEOSEARCH', KEYS[], 'FROMLONLAT', lon, lat, 'BYRADIUS', radius, unit)

local result = {}

for index, value in pairs(imgByLocation) do
	local time, name = string.match(value, "(.+)::(.+)")
    time = tonumber(time)

    if(time >= timeStamp-300 and time <=timeStamp+300) then
        table.insert(result, name)
    end
    
end
 
return result
