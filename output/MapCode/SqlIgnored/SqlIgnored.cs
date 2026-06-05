
var serviceObj = service.GetSqlIgnored();
SqlIgnored localObj = new SqlIgnored(serviceObj.Name,
        serviceObj.Count
);
return localObj;