using System;
using System.Collections.Generic;

namespace   {
    public partial interface IEmployee {
        
        int ID { get; set; }
        string FirstName { get; set; }
        string LastName { get; set; }
        string StartDateString { get; set; }
    }
}