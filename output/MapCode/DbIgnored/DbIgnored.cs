
var serviceObj = service.GetDbIgnored();
DbIgnored localObj = new DbIgnored(serviceObj.Name,
        serviceObj.Count
);
return localObj;