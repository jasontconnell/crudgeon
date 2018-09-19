# CRUDGEON

crudgeon is a multi-output CRUD generator written in Go for a very specific set of projects, but may be useful with modification to other projects. This project structure as is requires go 1.11 but it will compile in older versions if you get the dependencies elsehow.

## USAGE

go build

```crudgeon -file datafiles\example1.txt -path output -obj Business
crudgeon -file datafiles\example2.txt -path output -obj Employee
```

See genex.bat and sample input and output in datafiles and output folders.

## Documentation

Properties can be decorated with flags to control their output.

public [TYPE] FIELDNAME|PROPERTYNAME {get;set;} //FLAGS

ex.

public int years_old|Age {get;set;} //+sqlignore,+xmlignore

### TYPE:

As of now can be one of
* int
* short
* string
* decimal
* double
* long
* DateTime
* bool
* OR nullable versions of those with Nullable<TYPE>

### FIELDNAME:
This can be optional. If you're happy with the title cased version of the data source's name, you'll just need field name, not property name.
However, if your data source is coming back with underscores or words you find redundant (e.g. PlayerFirstName on a Player object), you can provide
a field name and a property name.  (e.g.  `public string PlayerFirstName|FirstName {get;set;}`)

### PROPERTYNAME:
See field name.

### FLAGS
Field Flags take the form //+flag,-flag after a field definition

* `+` Enables the flag
* `-` Disables the flag

Available Field Flags:

1. sqlignore
    - Don't generate sql properties for this, including column names and sql attributes.
2. jsonignore
    - Don't generate json attributes for this property
3. key
    - This field is part of the uniquely identifying set of fields for this object
4. nomap
    - Don't map this on the ORM call to get data from the database. Like sql ignore for "get" operations
5. xmlignore
    - Don't generate XML attributes for this object.
6. xmlwrapper  (array element name)
    - This field is an XML Array in C# parlance. It will generate [XmlArray] and [XmlArrayItem] attributes
    - The property name will be passed to [XmlArray] and the array element name will be passed to [XmlArrayItem]
        
In addition to fields, flags can be added to the entire object. For instance

//+xmlignore

Will have xml ignore on the entire file. This is a different type of flag though, which I called Generation Flags.

### Generation Flags

Generation Flags control Generation. In the config.json, these will tell what data each template needs in order to work. Some need more data, some need less (Delete proc just need keys, for instance). Here are the possible GenFlags.

1. id
    - include the "id" field which is implicit. If your object has a field named ID it has to be renamed with the ID|ObjectID field name|property name syntax.
2. fields
    - include the fields in the generation. This is just base type fields.
3. collections
    - include collections in generation. this will typically be used in code files as opposed to SQL files.
4. constructor
    - include fields for a sensible constructor (no collections, just base type fields)
5. concretes
    - also include concrete versions of interface classes
    - this one requires a bit of explanation. It uses convention. If you object is not a base type and it start with an I, it's an interface. Sorry.
6. keys
    - include keys as a separate object
7. sqlignore, csignore, jsonignore, xmlignore
    - these will usually be included in the data files, not in the config like 1-7
8. xmlroot (xml root element name)
    - this object is an xml root. you can use it like   //+xmlroot RootName

## Text Templates

The text template is passed a GenPackage object

```type GenPackage struct {
	Generate          bool
	Name              string
	Path              string
	OutputFile        string
	Namespace         string
	Fields            []GenField
	ConstructorFields []GenField
	KeyFields         []GenField
	TemplateFile      string
	Prefix            string
	Flags             GenFlags
}```

Really all you would care about are Name, Namespace, Fields, ConstructorFields, KeyFields, Prefix, and Flags.

You can look at the existing templates in the tmpl folder within this project.

To add a new one, add the template, then add the config to the config.json

## config.json

As this is a multi-output code generator, you can configure the specific outputs in the config.json. These use the generation flags.

The current "class" generator to generate classes:

`{ "file": "tmpl\\class.txt", "fileType": "cs", "outputPrefix": "", "folder": "Model", "flags": "+id,+collections,+concretes,+constructor,+fields","objdir":false}`

* file
    - use file as the template
* fileType
    - specify file type extension (cs in this case)
* outputPrefix
    - useful for generating multiple types of stored procedures, use "Get", "Update", "Delete", here for example
* folder
    - create this folder and put the output in there
* flags
    - generation flags. In this example, C# class requires pretty much everything.
* objdir
    - create a directory for this object's output.

With "folder" and "objdir", you can very easily control the file structure of the output.

You can use this project in conjunction with my [sqlrun project](https://github.com/jasontconnell/sqlrun) to get a database created in literally no time. Ok maybe a few seconds.

Happy Coding!
