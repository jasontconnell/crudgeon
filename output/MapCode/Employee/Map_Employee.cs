
var serviceObj = service.GetEmployee();
Employee localObj = new Employee(serviceObj.emp_first_name,
        serviceObj.emp_last_name,
        serviceObj.start_date
);
return localObj;