
var serviceObj = service.GetBusiness();
Business localObj = new Business(serviceObj.Name,
        serviceObj.Value,
        serviceObj.Revenue,
        serviceObj.Expenses,
        serviceObj.Employees
);
return localObj;