using System;
using System.Collections.Generic;

namespace  Name.Space {
    public partial interface IEmployee {
        
        int ID { get; set; }
        int EmployeeID { get; set; }
        string FirstName { get; set; }
        string LastName { get; set; }
        string StartDateString { get; set; }
    }
}