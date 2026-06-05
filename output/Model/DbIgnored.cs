using System;

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
    
    public partial class DbIgnored  {

        public DbIgnored(){
        }
        public DbIgnored( string Name,
                 int Count) {
            this.Name = Name;
            this.Count = Count;
        }
        [XmlIgnore]
        public int ID { get; set; }
        
        [HashKey(1)]
        [DataMember(Name="Name")]
        [XmlAttribute(AttributeName="Name")]
        public string Name { get; set; }
        
        [Hash(2)]
        [DataMember(Name="Count")]
        [XmlAttribute(AttributeName="Count")]
        public int Count { get; set; }
        

        override public string ToString(){
            StringBuilder sb = new StringBuilder("DbIgnored - " + Environment.NewLine);
                sb.AppendFormat("{0}: {1}{2}", "ID", this.ID, Environment.NewLine);
                sb.AppendFormat("{0}: {1}{2}", "Name", this.Name, Environment.NewLine);
                sb.AppendFormat("{0}: {1}{2}", "Count", this.Count, Environment.NewLine);
            return sb.ToString();
        }

    }
}