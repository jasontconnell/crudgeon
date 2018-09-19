using System;
using SqlMap.Attributes;
using System.Runtime.Serialization;
using System.Xml.Serialization;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using Hasher.Attributes;

// this code is generated, do not modify. use partial classes / interfaces to add stuff
namespace  {
    [DataContract]
    [StoredProc(Name="GetBusiness", ProcType=ProcType.Read)]
    [StoredProc(Name="SaveBusiness", ProcType=ProcType.Update)]
    [StoredProc(Name="DeleteBusiness", ProcType=ProcType.Delete)]
    
    public partial class Business : IBusiness {

        public Business(){
            
        }
        public Business( string Name,
                 decimal Value,
                 decimal Revenue,
                 decimal Expenses,
                 List<Employee> Employees) {
            this.Name = Name;
            this.Value = Value;
            this.Revenue = Revenue;
            this.Expenses = Expenses;
            this.Employees = Employees;
        }
        [XmlIgnore]
        public int ID { get; set;  }
        
        [HashKey(1)]
        [DataMember(Name="Name")]
        [XmlAttribute(AttributeName="Name")]
        [Column(Key=true)]
        public string Name { get; set;  }
        
        [Hash(2)]
        [DataMember(Name="Value")]
        [XmlAttribute(AttributeName="Value")]
        [Column]
        public decimal Value { get; set;  }
        
        [Hash(3)]
        [DataMember(Name="Revenue")]
        [XmlAttribute(AttributeName="Revenue")]
        [Column]
        public decimal Revenue { get; set;  }
        
        [Hash(4)]
        [DataMember(Name="Expenses")]
        [XmlAttribute(AttributeName="Expenses")]
        [Column]
        public decimal Expenses { get; set;  }
        
        [DataMember(Name="Employees")]
        [XmlElement("Employees")]
        [IgnoreCol]
        public List<Employee> Employees { get; set;  }
        

        override public string ToString(){
            StringBuilder sb = new StringBuilder();
                sb.AppendFormat("{0}: {1}{2}", "ID", this.ID, Environment.NewLine);
                sb.AppendFormat("{0}: {1}{2}", "Name", this.Name, Environment.NewLine);
                sb.AppendFormat("{0}: {1}{2}", "Value", this.Value, Environment.NewLine);
                sb.AppendFormat("{0}: {1}{2}", "Revenue", this.Revenue, Environment.NewLine);
                sb.AppendFormat("{0}: {1}{2}", "Expenses", this.Expenses, Environment.NewLine);
                sb.AppendFormat("{0}: {1}{2}", "Employees", this.Employees, Environment.NewLine);
            return sb.ToString();
        }

    }
}