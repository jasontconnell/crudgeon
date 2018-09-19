using System;
using SqlMap.Attributes;
using System.Runtime.Serialization;

using System.Collections.Generic;
using System.Linq;
using System.Text;
using Hasher.Attributes;

// this code is generated, do not modify. use partial classes / interfaces to add stuff
namespace  {
    [DataContract]
    [StoredProc(Name="GetEmployee", ProcType=ProcType.Read)]
    [StoredProc(Name="SaveEmployee", ProcType=ProcType.Update)]
    [StoredProc(Name="DeleteEmployee", ProcType=ProcType.Delete)]
    
    public partial class Employee : IEmployee {

        public Employee(){
            
        }
        public Employee( string FirstName,
                 string LastName,
                 string StartDateString) {
            this.FirstName = FirstName;
            this.LastName = LastName;
            this.StartDateString = StartDateString;
        }
        public int ID { get; set;  }
        
        [Hash(1)]
        [Column]
        public string FirstName { get; set;  }
        
        [Hash(2)]
        [Column]
        public string LastName { get; set;  }
        
        [DataMember(Name="start_date")]
        [IgnoreCol]
        public string StartDateString { get; set;  }
        

        override public string ToString(){
            StringBuilder sb = new StringBuilder();
                sb.AppendFormat("{0}: {1}{2}", "ID", this.ID, Environment.NewLine);
                sb.AppendFormat("{0}: {1}{2}", "FirstName", this.FirstName, Environment.NewLine);
                sb.AppendFormat("{0}: {1}{2}", "LastName", this.LastName, Environment.NewLine);
                sb.AppendFormat("{0}: {1}{2}", "StartDateString", this.StartDateString, Environment.NewLine);
            return sb.ToString();
        }

    }
}