local imgByLocation = {"1672531741::fd46f602-8679-4a65-ad2c-08b6a9ea52a1.png",
"1672531861::8bc88b73-32df-4a28-867b-63413551a2cc.png",
"1672531921::b175838c-2c1e-4930-b849-140adf212561.png",
"1672531981::f946fe81-9df1-411f-8433-dc79515565f1.png",
"1672532041::dd64d611-9cb6-46cf-969f-69ba8f20d96a.png",
"1672532101::875d2e3c-b828-4f9a-869f-434bb8299772.png",
"1672532161::bb271c3c-1799-4dad-b2e8-201f35eb93ee.png",
"1672532221::385391e7-d746-4a0d-892c-1ee2d449ba80.png",
"1672532281::786513b1-16cd-4f4b-bbc3-c8242775d3fe.png",
"1672531201::1ac654dc-920f-43a2-a089-ee7ea94a4de8.png",
"1672531200::d918f72a-d329-4a25-ae91-8703cc5b39e8.png",
"1672531200::9c6f6eb7-844a-4c55-aec0-6423fc0b88a6.png",
"1674432000::56ee8d79-ab33-491f-97c4-f2658d21e066.png",
"1702339200::38907330-fe65-4a1f-83bb-f4ef1888a793.png",
"1702383132::15b51bde-25c5-452c-9caf-85032b5b32bd.png",
"1702422732::409f79b2-76d3-4c32-9947-deaac2cdeebb.png",
"1702417692::834cd083-70de-4da4-a574-6fc5b71faad9.png",
"1702417932::6f04d29e-d1e7-4648-883c-5537281c9634.png"}

local result = {}

for index, value in pairs(imgByLocation) do
	local time, name = string.match(value, "(.+)::(.+)")
    time = tonumber(time)

    if(time >= 1672532101 and time <=1702383132) then
        table.insert(result, name)
    end
    
end



 


for index, value in pairs(result) do
    print(value)
end