# How to contribute to flow


## Contributing modules
Modules templates can be created by running `create_module <module_name>` from
the command line. Every module has a run method that does all the processing. The run method pulls from the in port, processes the pulled data, ad returns the
processed data. Every module will have to import the processing and constants.
