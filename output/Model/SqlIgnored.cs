using System;
using SqlMap.Attributes;
using System.Runtime.Serialization;
using System.Xml.Serialization;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using Hasher.Attributes;

// this code is generated, do not modify. use partial classes / interfaces to add stuff
namespace Name.Space 
{
    [DataContract]
    [StoredProc(Name="GetSqlIgnored", ProcType=ProcType.Read)]
    [StoredProc(Name="SaveSqlIgnored", ProcType=ProcType.Update)]
    [StoredProc(Name="DeleteSqlIgnored", ProcType=ProcType.Delete)]
    
    public partial class SqlIgnored  {

        public SqlIgnored(){
        }
        public SqlIgnored( string Name,
                 int Count) {
            this.Name = Name;
            this.Count = Count;
        }
        [XmlIgnore]
        public int ID { get; set; }
        
        [HashKey(1)]
        [DataMember(Name="Name")]
        [XmlAttribute(AttributeName="Name")]
        [Column(Key=true)]
        public string Name { get; set; }
        
        [Hash(2)]
        [DataMember(Name="Count")]
        [XmlAttribute(AttributeName="Count")]
        [Column]
        public int Count { get; set; }
        

        override public string ToString(){
            StringBuilder sb = new StringBuilder("SqlIgnored - " + Environment.NewLine);
                sb.AppendFormat("{0}: {1}{2}", "ID", this.ID, Environment.NewLine);
                sb.AppendFormat("{0}: {1}{2}", "Name", this.Name, Environment.NewLine);
                sb.AppendFormat("{0}: {1}{2}", "Count", this.Count, Environment.NewLine);
            return sb.ToString();
        }

    }
}