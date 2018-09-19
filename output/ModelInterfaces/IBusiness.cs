using System;
using System.Collections.Generic;

namespace  Name.Space {
    public partial interface IBusiness {
        
        int ID { get; set; }
        string Name { get; set; }
        decimal Value { get; set; }
        decimal Revenue { get; set; }
        decimal Expenses { get; set; }
        List<Employee> Employees { get; set; }
    }
}