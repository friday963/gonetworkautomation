# Example Go template generator
## High level Objectives
1. Demonstrate ability to consume yaml files and parse into structured data.
2. Demonstrate understanding of Go templating language.  Including looping and conditional statements.
## Secondary discovery
1. More exposure to imports specifically, internal imports intra-project.
2. More exposure to structs and the hierarchical format required to consume yaml/json.  This was fairly difficult to wrap my head around and will require much more exposure before I'm comfortable using structs to pull out specific data from a dataset.
3. Yaml/Json formatting makes all the difference when creating structs.  If you design your file format poorly your structs may not offer the functionality you seek.
    - I faced sigifiant challenges figuring out how to iterate over certain data sets until I reformatted the yaml (and concequently the structs).  Once the Yaml was better written I was able to iterate over values better and the templates were easier to write.
4. Go is much more structured in the way it consumes data.  In python I could have consumed that yaml configuration file and passed all variables to a jinja template.  From there I would have put all the logic into the template.  Go forced me to parse all that data outside of the template (into structs) before pushing those into the template.   

**Closing thoughts**

*This is a rough example of using go to parse a yaml file as configuration and pushing to a text template.  Using the built in functions and methods to further parse and manipulate data before dumping it to standard out.  The overhead of doing something like this at scale could be more time consuming than python given the sheer number of structs required.  With the rigidity it also offers more safety so it could be a bonus in some circumstances.  This is certainly not a full fledged example, significantly more configuration and code changes are required to make this a more viable model for generating configuration but all-in-all it gave me the confidence that I could do this on a much larger scale if needed.*